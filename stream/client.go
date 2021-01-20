/*
 * @File: stream.api.go
 * @Description: impl http client for Stream API
 * @LastModifiedTime: 2021-01-20 10:17:16
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package stream defines http client for Stream API
package stream

import (
	"errors"

	getstream "github.com/GetStream/stream-go"
)

// Client holds our connection to Stream
var Client *getstream.Client

// Connect -- connect to Stream, set our Client variable or report error
// params:
// apiKey    - string, Stream API key
// apiSecret - string, Stream API Secret
// apiRegion - string, Stream region (ie, "us-east", "eu-central")
func Connect(apiKey, apiSecret, apiRegion string) error {
	var err error
	if apiKey == "" || apiSecret == "" || apiRegion == "" {
		return errors.New("missing api credentials")
	}

	Client, err = getstream.New(&getstream.Config{
		APIKey:    apiKey,
		APISecret: apiSecret,
		Location:  apiRegion,
	})
	return err
}
