package jottingjwts

import (
	"fmt"
	"hackattic_solutions/pkg/common"

	"github.com/golang-jwt/jwt/v4"
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

var globalString = ""

func VerifyJwts() {

	// problemUrl := "https://hackattic.com/challenges/jotting_jwts/problem?access_token=8e80fec0cbe25049"

	// secretDto := common.GetResponse[jwtSecretDto](problemUrl)

	secretDto := jwtSecretDto{
		Secret: "1KMe&tcv&3!L57QR",
	}
	common.PrintDto(secretDto)

	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	common.HandleError(err)
	common.PrintDto(jwtToken.Claims.(jwt.MapClaims)["append"])
	claims := jwtToken.Claims.(jwt.MapClaims)

	globalString = globalString + fmt.Sprint(claims["append"])
}
