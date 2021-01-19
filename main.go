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
	"github.com/gorilla/mux"
)

type healthResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

func main() {
	// load config with viper
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config", err)
	}

	// connect to cassandra cluster
	db.CassandraConnect(config.CassandraHost, config.CassandraPort)

	// create API router using mux
	listeningAddress := fmt.Sprintf(":%s", config.AppPort)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(listeningAddress, router))
}

func health(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(healthResponse{Code: 200, Status: "OK"})
}
