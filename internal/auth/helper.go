package auth

import (
	"errors"
	"net/http"
	"strings"
)

func extractApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no Authorization header provided")
	}

	head := strings.Split(authHeader, " ")
	if len(head) != 2 {
		return "", errors.New("header has been malformed")
	}

	name, token := head[0], head[1]

	if name != "ApiKey" {
		return "", errors.New("header has been malformed")
	}

	return token, nil
}
