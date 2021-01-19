package main

import (
	"log"

	"github.com/Akshit8/go-cassandra/config"
	"github.com/gocql/gocql"
)

func main() {
	// load config from env
	var err error
	config, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	// connect to cassandra cluster
	cluster := gocql.NewCluster(config.CassandraHost)
	cluster.Port = config.CassandraPort
	cluster.Keyspace = "akshit"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("failed to connect to cassandra: ", err)
	}

	defer session.Close()

	log.Print("connection successful")
}
