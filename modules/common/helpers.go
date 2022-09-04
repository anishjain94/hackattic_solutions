package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		print(err.Error())
	}
}

func GetResponse[dto any](url string) *dto {

	resp, err := http.Get(url)
	HandleError(err)

	defer resp.Body.Close()

	var respDto dto

	err = json.NewDecoder(resp.Body).Decode(&respDto)
	HandleError(err)
	PrintDto(respDto)

	return &respDto
}

func PrintDto[dto any](dtoObj dto) {
	fmt.Printf("%#v \n", dtoObj)
}

func PrintReadClosure(resp io.ReadCloser) {
	responseStr, _ := io.ReadAll(resp)
	print(string(responseStr) + "\n")
}
