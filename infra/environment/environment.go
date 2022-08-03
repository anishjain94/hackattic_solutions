package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
	ENV  string
)

func InitializeEnvs(variant ...string) {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/hackattic_solutions/config/dev.env"))
	if err != nil {
		log.Fatal("unable to load " + ENV + ".env file")
	}

	ENV = os.Getenv("ENV")

	PORT = os.Getenv("APP_PORT")

}
