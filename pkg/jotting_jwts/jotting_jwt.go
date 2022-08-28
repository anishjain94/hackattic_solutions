package jottingjwts

import (
	"bytes"
	"encoding/json"
	"hackattic_solutions/pkg/common"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type jwtSecretDto struct {
	Secret string `json:"jwt_secret"`
}

type SolutionDto struct {
	Solution string `json:"app_url"`
}

type FinalSolution struct {
	Solution string `json:"solution"`
}

func VerifyJwts() {

	problemUrl := "https://hackattic.com/challenges/jotting_jwts/problem?access_token=8e80fec0cbe25049"

	secretDto := common.GetResponse[jwtSecretDto](problemUrl)
	common.PrintDto(secretDto)

	go webserve(secretDto.Secret)

	solution := SolutionDto{
		Solution: "http://127.0.0.1:3000/app",
	}

	marshalled, _ := json.Marshal(solution)
	body := bytes.NewReader(marshalled)
	resp, err := http.Post("https://hackattic.com/challenges/reading_qr/solve?access_token=8e80fec0cbe25049", "application/json", body)
	common.HandleError(err)

	common.PrintDto(resp.Body)
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

func webserve(secrets string) {

	var answer strings.Builder

	requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqToken, _ := io.ReadAll(r.Body)
		print(reqToken)

		stringToAppend, isValid := verifyToken(string(reqToken), secrets)

		if isValid {
			answer.WriteString(*stringToAppend)
		} else if *stringToAppend == "Time to finish" {
			solution := answer.String()
			finalSolution := FinalSolution{
				Solution: solution,
			}

			finalSolutionJson, err := json.Marshal(&finalSolution)
			common.HandleError(err)

			w.Header().Set("Content-Type", "application/json")
			w.Write(finalSolutionJson)
		}
	})

	router := mux.NewRouter().StrictSlash(true)
	// http.ListenAndServe(":3000", router)
	http.ListenAndServe(":3000", handlers.CombinedLoggingHandler(os.Stdout, router))

	router.HandleFunc("/app", requestHandler).Methods(http.MethodPost)
}
