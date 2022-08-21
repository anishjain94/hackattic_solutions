package bruteforcezip

import (
	"archive/zip"
	"encoding/json"
	"hackattic_solutions/pkg/common"
	"io"
	"log"
	"net/http"
	"os"
)

type zipDto struct {
	ZipUrl string `json:"zip_url"`
}

func BruteForceZip() {

	resp, err := http.Get("https://hackattic.com/challenges/brute_force_zip/problem?access_token=8e80fec0cbe25049")
	var zipDto zipDto

	if err == nil {
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&zipDto); err != nil {
			log.Fatal("ooopsss! an error occurred, please try again")
		}
	}

	respZip, err := http.Get(zipDto.ZipUrl)
	common.HandleError(err)

	defer respZip.Body.Close()

	zipFile, err := os.Create("../pkg/brute_force_zip/hackattic.zip")
	common.HandleError(err)

	_, err = io.Copy(zipFile, respZip.Body)
	common.HandleError(err)

	zipReadCloser, err := zip.OpenReader("../pkg/brute_force_zip/hackattic.zip")
	common.HandleError(err)

	for _, file := range zipReadCloser.File {
		print(file)
	}

}
