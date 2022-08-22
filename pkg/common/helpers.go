package common

import (
	"encoding/json"
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

	return &respDto
}
