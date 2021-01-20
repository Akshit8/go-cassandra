/*
 * @File: messages.get.go
 * @Description: handle our GET operation on our API
 * @LastModifiedTime: 2021-01-20 08:37:47
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package messages creates or fetchs message data from Cassandra
package messages

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Akshit8/go-cassandra/db"
	"github.com/Akshit8/go-cassandra/users"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

// Get -- handles GET request to /messages/ to fetch all messages
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params (unused here)
func Get(w http.ResponseWriter, r *http.Request) {
	var messageList []Message
	var enrichedMessages []Message
	var userList []gocql.UUID
	// var err error
	m := map[string]interface{}{}

	// globalMessages, err := stream.Client.FlatFeed("messages", "global")
	// // fetch from Stream
	// if err == nil {
	// 	activities, err := globalMessages.Activities(nil)
	// 	if err == nil {
	// 		log.Println("Fetching activities from Stream")
	// 		for _, activity := range activities.Activities {
	// 			log.Println(activity)
	// 			userID, _ := gocql.ParseUUID(activity.Actor)
	// 			messageID, _ := gocql.ParseUUID(activity.Object)
	// 			messageList = append(messageList, Message{
	// 				ID:      messageID,
	// 				UserID:  userID,
	// 				Message: activity.MetaData["message"],
	// 			})
	// 			userList = append(userList, userID)
	// 		}
	// 	}
	// } else {
	// if Stream fails, pull from database instead
	log.Println("fetching activities from db")
	query := "SELECT id, user_id, message FROM messages"
	iterable := db.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		userID := m["user_id"].(gocql.UUID)
		messageList = append(messageList, Message{
			ID:      m["id"].(gocql.UUID),
			UserID:  userID,
			Message: m["message"].(string),
		})
		userList = append(userList, userID)
		m = map[string]interface{}{}
	}
	// }

	names := users.Enrich(userList)
	for _, message := range messageList {
		message.UserFullName = names[message.UserID.String()]
		enrichedMessages = append(enrichedMessages, message)
	}

	log.Println("message list after enrichment", enrichedMessages)

	json.NewEncoder(w).Encode(AllMessagesResponse{Messages: enrichedMessages})
}

// GetOne -- handles GET request to /messages/{message_uuid} to fetch one message
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func GetOne(w http.ResponseWriter, r *http.Request) {
	var message Message
	var errs []string
	var found bool = false

	vars := mux.Vars(r)
	id := vars["message_uuid"]

	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		m := map[string]interface{}{}
		query := "SELECT id, user_id, message FROM messages WHERE id=? LIMIT 1"
		iterable := db.Session.Query(query, uuid).Consistency(gocql.One).Iter()
		for iterable.MapScan(m) {
			found = true
			userID := m["user_id"].(gocql.UUID)
			names := users.Enrich([]gocql.UUID{userID})
			log.Println("names", names)
			message = Message{
				ID:           userID,
				UserID:       m["user_id"].(gocql.UUID),
				UserFullName: names[userID.String()],
				Message:      m["message"].(string),
			}
		}
		if !found {
			errs = append(errs, "Message not found")
		}
	}

	if found {
		json.NewEncoder(w).Encode(GetMessageResponse{Message: message})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
