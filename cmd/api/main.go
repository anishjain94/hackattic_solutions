package main

import (
	"hackattic_solutions/infra/environment"
	"hackattic_solutions/infra/rest"
)

func main() {
	environment.InitializeEnvs()

	rest.InitializeApiRestServer()
}
