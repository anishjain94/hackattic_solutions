package bruteforcezip

import (
	// "archive/zip"
	"hackattic_solutions/pkg/common"
	"io"
	"log"

	"github.com/alexmullins/zip"
)

// type zipDto struct {
// 	ZipUrl string `json:"zip_url"`
// }

func listFile(file *zip.File) error {

	// brute force password

	var fileRead io.ReadCloser
	var err error

	for i := 1000; i <= 999999; i++ {
		file.FileHeader.SetPassword(string(rune(i)))
		fileRead, err = file.Open()
		if err != nil {
			continue
		} else {
			print("password is ", i)
			break
		}
	}

	// defer fileRead.Close()

	str, err := io.ReadAll(fileRead)
	common.HandleError(err)

	print(string(str))

	return nil
}

func BruteForceZip() {

	// problemUrl := "https://hackattic.com/challenges/brute_force_zip/problem?access_token=8e80fec0cbe25049"

	// zipObj := common.GetResponse[zipDto](problemUrl)

	// respZip, err := http.Get(zipObj.ZipUrl)
	// common.HandleError(err)

	// defer respZip.Body.Close()

	// zipFile, err := os.Create("../pkg/brute_force_zip/hackattic.zip")
	// common.HandleError(err)

	// _, err = io.Copy(zipFile, respZip.Body)
	// common.HandleError(err)

	read, err := zip.OpenReader("../pkg/brute_force_zip/hackattic.zip")
	common.HandleError(err)

	for _, file := range read.File {
		print(file.Name)
		print(file.FileHeader.IsEncrypted())

		if file.Name == "secret.txt" {
			if err := listFile(file); err != nil {
				log.Fatal("unable to open file")
			}
		}
	}

}


// m = list('abcdefghijklmnopqrstuvwxyz0123456789')
// def to_base_lm(n):
//     result = ''
//     base = len(m)
//     while n > 0:
//         index = n % base
//         result = m[index] + result
//         n = n // base
//     return result
// def from_base_lm(s):
//     s = s[::-1]
//     result = 0
//     n = len(s)
//     base = len(m)
//     for i in range(n):
//         number = m.index(s[i])
//         result += number * (base ** i)
//     return result
// for i in range(from_base_lm('aaaa'), from_base_lm('99999')+1):
//     print(to_base_lm(i).rjust(4,'a'))