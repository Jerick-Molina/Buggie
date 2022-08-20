package api

import (
	"net/http"
)

type ReturnResponse struct {
	code    int
	message string
}

func Response(code int) (int, string) {

	switch code {

	case http.StatusOK:

	case http.StatusBadRequest:

	default:

	}
	return http.StatusOK, "ss"
}
