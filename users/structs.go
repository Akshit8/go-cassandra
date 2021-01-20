/*
 * @File: users.structs.go
 * @Description: contain all structs for user package
 * @LastModifiedTime: 2021-01-20 08:29:21
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package users creates or fetchs user data from Cassandra
package users

import "github.com/gocql/gocql"

// User struct to hold profile data for our user
type User struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	City      string     `json:"city"`
}

// GetUserResponse to form payload returning a single User struct
type GetUserResponse struct {
	User User `json:"user"`
}

// AllUsersResponse to form payload of an array of User structs
type AllUsersResponse struct {
	Users []User `json:"users"`
}

// NewUserResponse builds a payload of new user resource ID
type NewUserResponse struct {
	ID gocql.UUID `json:"id"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
