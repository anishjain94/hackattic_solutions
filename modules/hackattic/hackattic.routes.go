package hackattic

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.StrictSlash(true)

	initRoutes(router)
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/readQr", readingQr).Methods(http.MethodPost)
}
