package hackattic

import (
	"hackattic_solutions/common"
	"net/http"

	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.StrictSlash(true)

	initRoutes(router)
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/signup", common.HandleGet(signUp)).Methods(http.MethodPost)
}
