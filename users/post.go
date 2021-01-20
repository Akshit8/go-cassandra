/*
 * @File: users.post.go
 * @Description: handle our POST operation on our API
 * @LastModifiedTime: 2021-01-20 08:37:47
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package users creates or fetchs user data from Cassandra
package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Akshit8/go-cassandra/db"
	"github.com/gocql/gocql"
)

// Post -- handles POST request to /users/new to create new user
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUUID gocql.UUID

	user, errs := FormToUser(r)

	var created bool = false
	if len(errs) == 0 {
		log.Println("creating a new user")
		gocqlUUID = gocql.TimeUUID()
		query := "INSERT INTO users (id, firstname, lastname, email, city, age) VALUES (?, ?, ?, ?, ?, ?)"
		err := db.Session.Query(
			query,
			gocqlUUID, user.FirstName, user.LastName, user.Email, user.City, user.Age,
		).Exec()
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}

	if created {
		log.Println("user_id", gocqlUUID)
		json.NewEncoder(w).Encode(NewUserResponse{ID: gocqlUUID})
	} else {
		log.Println("error", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
