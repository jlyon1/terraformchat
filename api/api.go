package api

import (
	"encoding/json"
	"github.com/jlyon1/appcache/database"
	"net/http"
	"time"
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

type Chat struct {
	Text string `json:Text`
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func (api *API) GetRoomName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(api.RoomName))
}

func (api *API) GetChatLog(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, api.Log)
}

func (api *API) AppendChatLog(w http.ResponseWriter, r *http.Request) {
	chat := Chat{}
	err := json.NewDecoder(r.Body).Decode(&chat)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		http.Error(w, "Invalid format", 500)
		return
	}
	api.Log.Log = append(api.Log.Log, "Anon: "+chat.Text)
}

func (api *API) Save() {
	for {
		api.DB.Set(api.RoomName, api.Log)
		<-time.After(5 * time.Second)
	}
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}
