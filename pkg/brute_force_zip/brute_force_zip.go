package bruteforcezip

import (
	"archive/zip"
	"hackattic_solutions/pkg/common"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type zipDto struct {
	ZipUrl string `json:"zip_url"`
}

func listFile(file *zip.File) error {

	fileRead, err := file.Open()
	common.HandleError(err)

	defer fileRead.Close()

	str, err := ioutil.ReadAll(fileRead)
	common.HandleError(err)

	print(string(str))

	return nil
}

func BruteForceZip() {

	problemUrl := "https://hackattic.com/challenges/brute_force_zip/problem?access_token=8e80fec0cbe25049"

	zipObj := common.GetResponse[zipDto](problemUrl)

	respZip, err := http.Get(zipObj.ZipUrl)
	common.HandleError(err)

	defer respZip.Body.Close()

	zipFile, err := os.Create("../pkg/brute_force_zip/hackattic.zip")
	common.HandleError(err)

	_, err = io.Copy(zipFile, respZip.Body)
	common.HandleError(err)

	read, err := zip.OpenReader("../pkg/brute_force_zip/hackattic.zip")
	common.HandleError(err)

	for _, file := range read.File {
		if file.Name == "secret.txt" {
			if err := listFile(file); err != nil {
				log.Fatal("unable to open file")
			}
		}
	}

}
