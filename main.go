package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to cassandra cluster
	cluster := gocql.NewCluster("35.229.149.49")
	cluster.Port = 3005
	cluster.Keyspace = "akshit"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	defer session.Close()

	if err != nil {
		log.Fatal("failed to connect to cassandra: ", err)
	}
	log.Print("connection successful");
}