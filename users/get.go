/*
 * @File: users.get.go
 * @Description: handle our GET operation on our API
 * @LastModifiedTime: 2021-01-20 08:37:47
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package users creates or fetchs user data from Cassandra
package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Akshit8/go-cassandra/db"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// Get -- handles GET request to /users/ to fetch all users
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params (unused here)
func Get(w http.ResponseWriter, r *http.Request) {
	var userList []User
	m := map[string]interface{}{}

	query := "SELECT id, age, firstname, lastname, city, email FROM users"
	iterable := db.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		userList = append(userList, User{
			ID:        m["id"].(gocql.UUID),
			Age:       m["age"].(int),
			FirstName: m["firstname"].(string),
			LastName:  m["lastname"].(string),
			Email:     m["email"].(string),
			City:      m["city"].(string),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(AllUsersResponse{Users: userList})
}

// GetOne -- handles GET request to /users/{user_uuid} to fetch one user
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func GetOne(w http.ResponseWriter, r *http.Request) {
	var user User
	var errs []string
	var found bool = false

	vars := mux.Vars(r)
	id := vars["user_uuid"]

	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		m := map[string]interface{}{}
		query := "SELECT id,age,firstname,lastname,city,email FROM users WHERE id=? LIMIT 1"
		iterable := db.Session.Query(query, uuid).Consistency(gocql.One).Iter()
		for iterable.MapScan(m) {
			found = true
			user = User{
				ID:        m["id"].(gocql.UUID),
				Age:       m["age"].(int),
				FirstName: m["firstname"].(string),
				LastName:  m["lastname"].(string),
				Email:     m["email"].(string),
				City:      m["city"].(string),
			}
		}
		if !found {
			errs = append(errs, "User not found")
		}
	}

	if found {
		json.NewEncoder(w).Encode(GetUserResponse{User: user})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

// Enrich -- turns an array of user UUIDs into a map of {uuid: "firstname lastname"}
// params:
// uuids - array of user UUIDs to fetch
// returns:
// a map[string]string of {uuid: "firstname lastname"}
func Enrich(uuids []gocql.UUID) map[string]string {
	if len(uuids) > 0 {
		log.Println("----\nfetching names", uuids)
		names := map[string]string{}
		m := map[string]interface{}{}

		query := "SELECT id, firstname, lastname FROM users WHERE id IN ?"
		iterable := db.Session.Query(query, uuids).Iter()
		for iterable.MapScan(m) {
			log.Println("m", m)
			userID := m["id"].(gocql.UUID)
			log.Println("userID", userID.String())
			names[userID.String()] = fmt.Sprintf("%s %s", m["firstname"].(string), m["lastname"].(string))
			m = map[string]interface{}{}
		}
		log.Println("names", names)
		return names
	}
	return map[string]string{}
}
