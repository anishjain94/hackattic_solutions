package websocketchitchat

import (
	"bytes"
	"encoding/json"
	"hackattic_solutions/modules/common"
	"net/http"

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

	problemUrl := "https://hackattic.com/challenges/websocket_chit_chat/problem?access_token=8e80fec0cbe25049"

	secretDto := websockerSecret{
		Secret: secret,
	}

	marshalled, _ := json.Marshal(secretDto)
	bytesReader := bytes.NewReader(marshalled)

	resp, err := http.Post(problemUrl, "application/json", bytesReader)
	common.HandleError(err)

	defer resp.Body.Close()
	common.PrintReadClosure(resp.Body)

}

func Chat() {

	token := getToken()

	// token := "ae4d1c55.5b00.429b.953f.084d163a93d7"
	url := "wss://hackattic.com/_/ws/" + token
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	common.HandleError(err)
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		common.HandleError(err)

		msg := string(p)
		print(msg)

		if msg == "!ping" {
			err = conn.WriteMessage(messageType, p)
			common.HandleError(err)
		} else if msg == "!good" {
			print("doing correctly")
		} else {
			sendSecretMsg(msg)
		}

	}

}
