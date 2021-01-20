/*
 * @File: config.config.go
 * @Description: describes app configurations
 * @LastModifiedTime: 2021-01-19 08:30:18
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package config impls method to load and use configuration
package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are ready by viper from a config file or environment variable
type Config struct {
	AppPort           string `mapstructure:"APP_PORT"`
	CassandraHost     string `mapstructure:"CASSANDRA_HOST"`
	CassandraPort     int    `mapstructure:"CASSANDRA_PORT"`
	CassandraKeyspace string `mapstructure:"CASSANDRA_KEYSPACE"`
	StreamAPIKey      string `mapstructure:"STREAM_API_KEY"`
	StreamAPISecret   string `mapstructure:"STREAM_API_SECRET"`
	StreamAPIRegion   string `mapstructure:"STREAM_API_REGION"`
}

// LoadConfig loads env variables to config object
func LoadConfig() (config Config, err error) {
	// config missing jugaad
	_, err = os.Open("config/config.yml")
	if err != nil {
		log.Print("copying config file template...")
		os.Link("config/sample.config.yml", "config/config.yml")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
