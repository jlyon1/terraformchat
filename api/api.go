package api

import (
	"github.com/jlyon1/appcache/database"
	"net/http"
)

type ChatLog struct {
	Name string   `json:Name`
	Log  []string `json:Log`
}

type API struct {
	DB       *database.Redis
	RoomName string
	Log      ChatLog
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Setting this up later"))
}

func (api *API) GetRoomName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(api.RoomName))
}
