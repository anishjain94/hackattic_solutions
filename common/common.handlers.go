package common

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGet[outputDto any](callingFunc func(ctx *context.Context) *outputDto) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		queries := r.URL.Query()
		ctx = context.WithValue(ctx, CTX_QUERIES, queries)

		params := mux.Vars(r)
		ctx = context.WithValue(ctx, CTX_PARAMS, params)

		ctx = context.WithValue(ctx, CTX_HEADERS, r.Header)

		response := callingFunc(&ctx)

		w.Header().Set(string("content-type"), "application/json")
		msg := "Request Successfull"
		json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Message: &msg,
			},
			Data: response,
		})

	}
}

func HandlePost[InputDtoType any, OutputDtoType any](callingFunc func(ctx *context.Context, inputDto *InputDtoType) *OutputDtoType) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		queries := r.URL.Query()
		ctx = context.WithValue(ctx, CTX_QUERIES, queries)

		params := mux.Vars(r)
		ctx = context.WithValue(ctx, CTX_PARAMS, params)

		ctx = context.WithValue(ctx, CTX_HEADERS, r.Header)

		ctx = context.WithValue(ctx, CTX_REQUEST, r)

		ctx = context.WithValue(ctx, CTX_WRITER, w)

		var dto InputDtoType
		_ = json.NewDecoder(r.Body).Decode(&dto)

		response := callingFunc(&ctx, &dto)

		w.Header().Set(string("content-type"), "application/json")
		msg := "Request Successfull"
		json.NewEncoder(w).Encode(SuccessDto{
			Meta: AckDto{
				Success: true,
				Message: &msg,
			},
			Data: response,
		})

	}
}
