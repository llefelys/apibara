package codes

import "net/http"

type (
	StatusOK               struct{}
	StatusBadRequest       struct{}
	StatusNotFound         struct{}
	StatusMethodNotAllowed struct{}
)

func (StatusOK) Code() int               { return http.StatusOK }
func (StatusBadRequest) Code() int       { return http.StatusBadRequest }
func (StatusNotFound) Code() int         { return http.StatusNotFound }
func (StatusMethodNotAllowed) Code() int { return http.StatusMethodNotAllowed }
