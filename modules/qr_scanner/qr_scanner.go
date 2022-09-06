package main

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/liyue201/goqr"
)

type qr struct {
	Image_url string `json:"image_url"`
}

type qrPost struct {
	Code string `json:"code"`
}

func ScanQR() {

	resp, err := http.Get("https://hackattic.com/challenges/reading_qr/problem?access_token=8e80fec0cbe25049")

	var qrDto qr

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
		println(resp.Body)
		temp, _ := json.Marshal(resp.Body)
		println(string(temp))

	}
}
