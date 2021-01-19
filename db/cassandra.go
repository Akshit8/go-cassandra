/*
 * @File: db.cassandra.go
 * @Description: file includes cassandra db utilities
 * @LastModifiedTime: 2021-01-19 16:57:09
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package db impls database access and connection functionalities
package db

import (
	"log"

	"github.com/gocql/gocql"
)

// Session is singleton cassandra tcp client for the service
var Session *gocql.Session

// CassandraConnect connects to cassandra db and inits Session var
func CassandraConnect(cassandraHost string, cassandraPort int, cassandraKeyspace string) {
	var err error
	cluster := gocql.NewCluster(cassandraHost)
	cluster.Port = cassandraPort
	cluster.Keyspace = cassandraKeyspace
	cluster.Consistency = gocql.Quorum
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("failed to connect to cassandra: ", err)
	}
	log.Print("connected to cassandra db")
}
