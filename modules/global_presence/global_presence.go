package globalpresence

import (
	"hackattic_solutions/modules/common"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func callEndPoint(requestUrl string, proxy string) string {
	var client *http.Client

	proxyURL, err := url.Parse("http://" + proxy)
	common.HandleError(err)

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client = &http.Client{
		Transport: &transport,
	}

	request, err := http.NewRequest("GET", requestUrl, nil)
	common.HandleError(err)
	request.Header.Add("Accept-Encoding", "identity")
	request.Close = true

	request.Body = ioutil.NopCloser(strings.NewReader("{}"))

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

	callEndPoint(presenceEndPoint, "137.184.197.190:80")

}
