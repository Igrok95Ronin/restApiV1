package handlers

import "net/http"

type Ihandler interface {
	Register(mux *http.ServeMux)
}
