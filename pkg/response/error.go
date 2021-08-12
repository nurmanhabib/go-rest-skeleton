package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int           `json:"code"`
	Errors  []interface{} `json:"errors"`
	Message string        `json:"message"`
}

func (r ErrorResponse) ToJSON(w http.ResponseWriter) {
	if r.Code == 0 {
		r.Code = http.StatusBadRequest
	}

	if r.Message == "" {
		r.Message = http.StatusText(r.Code)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		panic(err)
	}
}
