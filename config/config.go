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
	AppPort       string `mapstructure:"APP_PORT"`
	CassandraHost string `mapstructure:"CASSANDRA_HOST"`
	CassandraPort int    `mapstructure:"CASSANDRA_PORT"`
}

// LoadConfig loads env variables to config object
func LoadConfig() (config Config, err error) {
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
