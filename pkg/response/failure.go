package response

import (
	"encoding/json"
	"net/http"
)

type FailureResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r FailureResponse) ToJSON(w http.ResponseWriter) {
	if r.Code == 0 {
		r.Code = http.StatusInternalServerError
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
