package response

import "net/http"

type Response interface {
	ToJSON(w http.ResponseWriter)
}

func JSON(w http.ResponseWriter, r Response) {
	r.ToJSON(w)
}
