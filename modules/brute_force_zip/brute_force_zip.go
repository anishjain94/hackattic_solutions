package bruteforcezip

import (
	"bytes"
	"hackattic_solutions/modules/common"
	"io"
	"log"
	"os"
	"os/exec"
)

// type zipDto struct {
// 	ZipUrl string `json:"zip_url"`
// }

func BruteForceZip() {

	// problemUrl := "https://hackattic.com/challenges/brute_force_zip/problem?access_token=8e80fec0cbe25049"

	// zipObj := common.GetResponse[zipDto](problemUrl)

	// respZip, err := http.Get(zipObj.ZipUrl)
	// common.HandleError(err)

	// defer respZip.Body.Close()

	// zipFile, err := os.Create("../modules/brute_force_zip/hackattic.zip")
	// common.HandleError(err)

	// _, err = io.Copy(zipFile, respZip.Body)
	// common.HandleError(err)

	// read, err := zip.OpenReader("../modules/brute_force_zip/hackattic.zip")
	// common.HandleError(err)

	// Used john the reaper for breaking password

	// lucky featherbold boatshy bonusbillowing smoke
	runCommand("/opt/homebrew/Cellar/john-jumbo/1.9.0/share/john/zip2john hackattic.zip > secure.hashes", "", "")

	_, err := exec.Command("/opt/homebrew/Cellar/john-jumbo/1.9.0/share/john/zip2john hackattic.zip > secure.hashes").CombinedOutput()
	common.HandleError(err)

	passwordBytes, err := exec.Command("sh", "-c", "john secure.hashes").CombinedOutput()
	common.HandleError(err)

	println(string(passwordBytes))

	// for _, file := range read.File {
	// 	if file.Name == "secret.txt" {
	// 		password := ""
	// 		file.SetPassword(password)
	// 		_, err := file.Open()
	// 		if err != nil {
	// 			println(err.Error())
	// 			continue
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
}

func runCommand(cmd string, args string, dir string) {
	command := exec.Command(cmd, args)
	command.Dir = dir

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	command.Stdout = mw
	command.Stderr = mw

	// Execute the command
	if err := command.Run(); err != nil {
		log.Panic(err)
	}

	log.Println(stdBuffer.String())

}
