package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadParams(r *http.Request, params interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(params)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}
