/*
 * @File: main.go
 * @Description: impls main package of application
 * @LastModifiedTime: 2021-01-19 16:56:12
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Akshit8/go-cassandra/config"
	"github.com/Akshit8/go-cassandra/db"
	"github.com/Akshit8/go-cassandra/messages"
	"github.com/Akshit8/go-cassandra/stream"
	"github.com/Akshit8/go-cassandra/users"
	"github.com/gorilla/mux"
)

func main() {
	// load config with viper
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", err)
	}

	// connect to cassandra cluster
	db.CassandraConnect(
		config.CassandraHost,
		config.CassandraPort,
		config.CassandraKeyspace,
	)

	// defer closing our Cassandra connection:
	defer db.Session.Close()

	// int stream api client
	err = stream.Connect(
		config.StreamAPIKey,
		config.StreamAPISecret,
		config.StreamAPIRegion,
	)
	if err != nil {
		log.Fatal("Could not connect to Stream: ", err)
	}

	// create API router using mux
	listeningAddress := fmt.Sprintf(":%s", config.AppPort)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/health", health)

	router.HandleFunc("/users", users.Get)
	router.HandleFunc("/users/new", users.Post)
	router.HandleFunc("/users/{user_uuid}", users.GetOne)

	router.HandleFunc("/messages", messages.Get)
	router.HandleFunc("/messages/new", messages.Post)
	router.HandleFunc("/messages/{message_uuid}", messages.GetOne)

	log.Fatal(http.ListenAndServe(listeningAddress, router))
}

type healthResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func health(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(healthResponse{Code: 200, Status: "OK"})
}
