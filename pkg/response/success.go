package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r SuccessResponse) ToJSON(w http.ResponseWriter) {
	if r.Code == 0 {
		r.Code = http.StatusOK
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
