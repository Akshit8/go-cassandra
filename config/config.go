/*
 * @File: config.config.go
 * @Description: describes app configurations
 * @LastModifiedTime: 2021-01-19 08:30:18
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package config impls method to load and use configuration
package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are ready by viper from a config file or environment variable
type Config struct {
	CassandraHost string `mapstructure:"CASSANDRA_HOST"`
	CassandraPort int    `mapstructure:"CASSANDRA_PORT"`
}

// LoadConfigFromEnv loads env variables to config object
func LoadConfigFromEnv() (config Config, err error) {
	viper.AutomaticEnv()
	config.CassandraHost = viper.GetString("CASSANDRA_HOST")
	config.CassandraPort = viper.GetInt("CASSANDRA_PORT")
	return
}
