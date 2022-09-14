package globalpresence

import (
	"hackattic_solutions/modules/common"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func callEndPoint(requestUrl string, proxy string) string {

	// urenrl, err := url.Parse(presenceEndPoint)
	// common.HandleError(err)

	// proxyUrl, err := url.Parse("http://" + proxy)
	// common.HandleError(err)

	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyUrl),
	// }
	// client := &http.Client{
	// 	Transport: transport,
	// }

	// request, err := http.NewRequest("GET", presenceEndPoint, nil)
	// common.HandleError(err)

	// response, err := client.Do(request)
	// common.HandleError(err)

	// data, err := ioutil.ReadAll(response.Body)
	// common.HandleError(err)

	// println(string(data))

	// defer response.Body.Close()

	// urenrl, err := url.Parse(requestUrl)
	// if err != nil {
	// 	log.Println(err)
	// }

	var client *http.Client

	proxyURL, err := url.Parse("http://" + proxy)
	if err != nil {
		log.Println(err)
	}

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client = &http.Client{
		Transport: &transport,
	}

	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println(err)
	}
	request.Header.Add("Accept-Encoding", "identity")
	request.Close = true


	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	log.Println(string(data))

	return string(data)

}

func GlobalPresence() {
	problemUrl := "https://hackattic.com/challenges/a_global_presence/problem?access_token=8e80fec0cbe25049"

	tokenDto := common.GetResponse[TokenDto](problemUrl)

	common.PrintDto(tokenDto)

	// solutionEndPoint := "https://hackattic.com/challenges/a_global_presence/solve?access_token=8e80fec0cbe25049"

	presenceEndPoint := "https://hackattic.com/_/presence/" + tokenDto.Token
		
	callEndPoint(presenceEndPoint, "188.0.147.102")

}
