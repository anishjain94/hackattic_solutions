package jottingjwts

import (
	"bytes"
	"encoding/json"
	"hackattic_solutions/modules/common"
	"io"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func VerifyJwts() {

	problemUrl := "https://hackattic.com/challenges/jotting_jwts/problem?access_token=8e80fec0cbe25049"

	secretDto := common.GetResponse[jwtSecretDto](problemUrl)

	go sendEndPoint()

	webServer(secretDto.Secret)
}

func verifyToken(jwtString string, secret string) (*string, bool) {
	jwtToken, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		common.HandleError(err)
		return nil, false
	}
	common.PrintDto(jwtToken.Claims)

	var result string
	if err != nil && jwtToken.Valid {
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && claims["append"] != nil {
			result = claims["append"].(string)
			return &result, true
		} else {
			temp := "Time to finish"
			return &temp, false
		}
	}
	return &result, false
}

func webServer(secrets string) {

	var answer strings.Builder

	requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqToken, err := io.ReadAll(r.Body)
		common.HandleError(err)
		print(string(reqToken) + "\n")
		stringToAppend, isValid := verifyToken(string(reqToken), secrets)

		if isValid {
			answer.WriteString(*stringToAppend)
		} else if stringToAppend != nil && *stringToAppend == "Time to finish" {
			solution := answer.String()
			finalSolution := FinalSolution{
				Solution: solution,
			}

			finalSolutionJson, err := json.Marshal(&finalSolution)
			common.HandleError(err)

			print(string(finalSolutionJson))
			w.Header().Set("Content-Type", "application/json")
			w.Write(finalSolutionJson)

		}
	})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/app", requestHandler).Methods(http.MethodPost)
	http.ListenAndServe(":3000", router)
}

func sendEndPoint() {
	solution := AppUrlDto{
		AppUrl: "https://3c58-2405-201-d033-a844-7966-5681-2648-75c9.ngrok.io/app",
	}

	marshalled, _ := json.Marshal(solution)
	bytesReader := bytes.NewReader(marshalled)
	print(string(marshalled) + "\n")

	resp, err := http.Post("https://hackattic.com/challenges/jotting_jwts/solve?access_token=8e80fec0cbe25049", "application/json", bytesReader)
	common.HandleError(err)

	defer resp.Body.Close()
	common.PrintReadClosure(resp.Body)
}
