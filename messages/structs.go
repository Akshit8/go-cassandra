/*
 * @File: messages.structs.go
 * @Description: contain all structs for message package
 * @LastModifiedTime: 2021-01-20 08:29:21
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package messages creates or fetchs message data from Cassandra
package messages

import "github.com/gocql/gocql"

// Message struct for preparing JSON payload
type Message struct {
	ID           gocql.UUID `json:"id"`
	UserID       gocql.UUID `json:"user_id"`
	UserFullName string     `json:"user_full_name"`
	Message      string     `json:"lastname"`
}

// GetMessageResponse struct for embedding a single message
type GetMessageResponse struct {
	Message Message `json:"message"`
}

// AllMessagesResponse struct for an array of Message structs
type AllMessagesResponse struct {
	Messages []Message `json:"messages"`
}

// NewMessageResponse struct for returning ID of message in payload
type NewMessageResponse struct {
	ID gocql.UUID `json:"id"`
}

// ErrorResponse for sending back a potential array of error strings
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
