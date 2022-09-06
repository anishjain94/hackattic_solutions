package websocketchitchat

import (
	"bytes"
	"encoding/json"
	"hackattic_solutions/modules/common"
	"math"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getToken() string {

	problemUrl := "https://hackattic.com/challenges/websocket_chit_chat/problem?access_token=8e80fec0cbe25049"

	tokenDto := common.GetResponse[websocketToken](problemUrl)

	return tokenDto.Token

}

func sendSecretMsg(secret string) {

	solution := websockerSecret{
		Secret: secret,
	}

	marshalled, _ := json.Marshal(solution)
	bytesReader := bytes.NewReader(marshalled)

	resp, err := http.Post("https://hackattic.com/challenges/websocket_chit_chat/solve?access_token=8e80fec0cbe25049", "application/json", bytesReader)
	common.HandleError(err)

	defer resp.Body.Close()
	common.PrintReadClosure(resp.Body)

}

func findBestTime(timeTaken int64) int64 {
	intervals := [6]int64{700, 1000, 1500, 2000, 2500, 3000}
	floatTimeTaken := float64(timeTaken)

	minimumDistance := math.Inf(1)
	var answer int64

	for _, interval := range intervals {
		floatInterval := float64(interval)

		curDistance := math.Abs(floatInterval - floatTimeTaken)

		if minimumDistance > curDistance {
			minimumDistance = curDistance
			answer = interval
		}
	}
	return answer
}

func findDiff(stopChan chan struct{}, timeValueChan chan time.Duration) {
	startTime := time.Now()
	for {
		select {
		case <-stopChan:
			endTime := time.Now()
			timeElapsed := endTime.Sub(startTime)
			timeValueChan <- timeElapsed
			return
		}
	}
}

func Chat() {

	done := make(chan struct{})
	stopChan := make(chan struct{})
	timeValueChan := make(chan time.Duration)
	sendMessageChan := make(chan int64)

	token := getToken()

	// TODO:
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "wss://hackattic.com/_/ws/" + token
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	common.HandleError(err)
	defer conn.Close()

	go findDiff(stopChan, timeValueChan)

	go func() {
		defer close(done)
		for {
			_, p, err := conn.ReadMessage()
			common.HandleError(err)

			msg := string(p)
			println(msg)

			if strings.Contains(msg, "hello") {
				println("first message")
			} else if msg == "ping!" {
				close(stopChan)

				timeTaken := <-timeValueChan
				close(timeValueChan)

				stopChan = make(chan struct{})
				timeValueChan = make(chan time.Duration)

				go findDiff(stopChan, timeValueChan)
				answer := findBestTime(int64(timeTaken / time.Millisecond))
				println(answer)

				sendMessageChan <- answer

			} else if strings.Contains(msg, "congratulations") {
				var rgx = regexp.MustCompile(`\"(.*?)\"`)
				answer := rgx.FindStringSubmatch(msg)[1]
				sendSecretMsg(answer)
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case calculatedDuration := <-sendMessageChan:
			err := conn.WriteMessage(1, []byte(string(calculatedDuration)))
			common.HandleError(err)
		case <-interrupt:
			println("Interrupt signal received")
			// close the connection cleanly by sending a message and waiting for a server to close it
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				println("Closed writing channel:", err)
				return
			}
			return
		}

	}

}
