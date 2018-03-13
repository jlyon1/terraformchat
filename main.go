package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jlyon1/appcache/database"
	"github.com/jlyon1/terraformchat/api"
	"github.com/jlyon1/terraformchat/config"
	"math/rand"
	"net/http"
	"os"
	"time"
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

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}
	db.IP = cfg.Redis
	db.Port = cfg.RedisPort
	db.DB = 0
	db.Password = cfg.Password

	connectDB(&db)

	animals, err := readLines("./animals.txt")
	roomName := animals[rand.Intn(len(animals))]
	r := mux.NewRouter()
	chatLog := api.ChatLog{
		roomName,
		[]string{},
	}
	api := api.API{
		&db,
		roomName,
		chatLog,
	}

	if err != nil {
		panic(err.Error())
	}
	go api.Save()

	r.HandleFunc("/", api.IndexHandler).Methods("GET")
	r.HandleFunc("/name", api.GetRoomName).Methods("GET")
	r.HandleFunc("/chat", api.GetChatLog).Methods("GET")
	r.HandleFunc("/send", api.AppendChatLog).Methods("Post")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.ListenAndServe("0.0.0.0:"+cfg.Port, r)
}
