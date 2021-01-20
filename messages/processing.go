/*
 * @File: messages.processing.go
 * @Description: contain form data methods
 * @LastModifiedTime: 2021-01-20 08:29:21
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package messages creates or fetchs message data from Cassandra
package messages

import "net/http"

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
