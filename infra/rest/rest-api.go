package rest

import (
	"encoding/json"
	"fmt"
	"hackattic_solutions/infra/environment"
	"hackattic_solutions/modules/hackattic"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func InitializeApiRestServer() {

	router := mux.NewRouter().StrictSlash(true)

	initializeApiRoutes(router)

	zap.L().Info("HTTP Api Server starting : " + fmt.Sprint(environment.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", environment.PORT), router))
}

func initializeApiRoutes(router *mux.Router) {

	initHealthRoutes(router.PathPrefix("/health").Subrouter())
	hackattic.InitUserRoutes(router.PathPrefix("/").Subrouter())
}

func initHealthRoutes(router *mux.Router) {
	router.StrictSlash(true)
	router.HandleFunc("", healthCheck).Methods(http.MethodGet)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	msg := "Healthy"
	json.NewEncoder(w).Encode(msg)
}
