/*
 * @File: messages.post.go
 * @Description: handle our POST operation on our API
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
	"github.com/Akshit8/go-cassandra/stream"
	getstream "github.com/GetStream/stream-go2/v5"
	"github.com/gocql/gocql"
)

// Post -- handles POST request to /messages/new to create a new message
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var errStr, userIDStr, message string

	if userIDStr, errStr = processFormField(r, "userID"); len(errStr) != 0 {
		errs = append(errs, errStr)
	}

	userID, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		errs = append(errs, "parameter 'userID' not a UUID")
	}

	if message, errStr = processFormField(r, "message"); len(errStr) != 0 {
		errs = append(errs, errStr)
	}

	gocqlUUID := gocql.TimeUUID()

	var created bool = false
	if len(errs) == 0 {
		query := "INSERT INTO messages (id, user_id, message) VALUES (?, ?, ?)"
		err := db.Session.Query(
			query,
			gocqlUUID, userID, message,
		).Exec()
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}

	if created {
		// send message to stream
		globalMessages, err := stream.Client.FlatFeed("akshit", "global")
		log.Print("stream error", err)
		if err == nil {
			_, err := globalMessages.AddActivity(getstream.Activity{
				Actor:  userID.String(),
				Verb:   "post",
				Object: gocqlUUID.String(),
			})
			log.Print("stream error 2 ", err)
		}
		json.NewEncoder(w).Encode(NewMessageResponse{ID: gocqlUUID})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
