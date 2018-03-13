package api

import (
	"github.com/jlyon1/appcache/database"
	"net/http"
)

type API struct {
	DB *database.Redis
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Setting this up later"))
}
