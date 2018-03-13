package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/appcache/database"
	"github.com/jlyon1/terraformchat/api"
	"github.com/jlyon1/terraformchat/config"

	"net/http"
)

type Config struct {
	Sites []string `json:strings`
}

var db database.Redis

func connectDB(db *database.Redis) {
	for db.Connect() == false {
		fmt.Printf("Trying to connect\n")
	}
	fmt.Printf("Connected\n")

}

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}
	db.IP = cfg.Redis
	db.Port = cfg.RedisPort
	db.DB = 0
	db.Password = cfg.Password

	connectDB(&db)

	r := mux.NewRouter()
	api := api.API{
		&db,
	}
	r.HandleFunc("/", api.IndexHandler).Methods("GET")
	http.ListenAndServe("0.0.0.0:8081", r)
}
