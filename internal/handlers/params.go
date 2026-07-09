package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func idDeURL(r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "id"))
}
