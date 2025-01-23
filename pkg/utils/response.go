package utils

import (
	"net/http"
)

type HandlerReturn struct {
	Code    int
	Payload interface{}
	Error   error
}

func SendResponse(service func(w http.ResponseWriter, r *http.Request) (hr *HandlerReturn)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hr := service(w, r)

		if hr.Error != nil {
			respondWithError(w, hr.Code, hr.Error.Error())
			return
		}

		sendJSON(w, hr.Code, hr.Payload)
	}
}
