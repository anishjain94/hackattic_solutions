package hackattic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/liyue201/goqr"
)

func recognizeFile(path string) string {
	fmt.Printf("recognize file: %v\n", path)
	imgFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgFile))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
	}
	return string(qrCodes[0].Payload)
}

func readingQr(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://hackattic.com/challenges/reading_qr/problem?access_token=8e80fec0cbe25049")

	var qrDto qr
	// recognizeFile("outimage.png")

	if err == nil {
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&qrDto); err != nil {
			log.Fatal("ooopsss! an error occurred, please try again")
		}

		respImg, _ := http.Get(qrDto.Image_url)

		file, _ := os.Create("outimage.png")
		if err != nil {
			// Handle error
		}
		img, _, err := image.Decode(respImg.Body)

		err = png.Encode(file, img)
		if err != nil {
			// Handle error
		}
		// file.Seek(0, 0)
		qrCodes, err := goqr.Recognize(img)

		qr := qrPost{}
		qr.Code = string(qrCodes[0].Payload)

		marshalled, _ := json.Marshal(qr)
		body := bytes.NewReader(marshalled)
		resp, err = http.Post("https://hackattic.com/challenges/reading_qr/solve?access_token=8e80fec0cbe25049", "application/json", body)
		if err != nil {
			// Handle error
		}
		print(resp.Body)
		temp, _ := json.Marshal(resp.Body)
		print(string(temp))

	}

}
