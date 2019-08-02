package main

import (
	"time"

	"github.com/spf13/viper"
)

type imageScalerConfig struct {
	hostname              string
	port                  int
	username              string
	password              string
	filename              string
	imageExchange         string
	imageUpdateQueue      string
	imageUpdateRoutingKey string
	timeout               time.Duration
	minioURL              string
	minioAccessKey        string
	minioSecret           string
	minioBucketName       string
	minioSecure           bool
}

func readConfig() imageScalerConfig {
	viper.SetConfigFile("config.properties")
	viper.SetConfigType("properties")

	//default values suitable for vanilla rabbitmq docker container
	viper.SetDefault("rabbitmq.hostname", "localhost")
	viper.SetDefault("rabbitmq.port", "5672")
	viper.SetDefault("rabbitmq.username", "guest")
	viper.SetDefault("rabbitmq.password", "guest")
	viper.SetDefault("rabbitmq.timeout", "5s")
	viper.SetDefault("rabbitmq.image.exchange", "user.event")
	viper.SetDefault("rabbitmq.image.update.queue", "user.image.url.updated.dev")
	viper.SetDefault("rabbitmq.image.update.routingkey", "user.image.url.updated.#")

	//default values suitable for min.io docker container
	viper.SetDefault("minio.url", "localhost:9000")
	viper.SetDefault("minio.accesskey", "admin")
	viper.SetDefault("minio.secret", "password")
	viper.SetDefault("minio.bucketname", "testbucket")
	viper.SetDefault("minio.secure", false)

	//load config
	confErr := viper.ReadInConfig()
	logOnError(confErr, "No configuration file loaded - using defaults")

	return imageScalerConfig{
		hostname:              viper.GetString("rabbitmq.hostname"),
		port:                  viper.GetInt("rabbitmq.port"),
		username:              viper.GetString("rabbitmq.username"),
		password:              viper.GetString("rabbitmq.password"),
		timeout:               viper.GetDuration("rabbitmq.timeout"),
		filename:              viper.GetString("filename"),
		imageExchange:         viper.GetString("rabbitmq.image.exchange"),
		imageUpdateQueue:      viper.GetString("rabbitmq.image.update.queue"),
		imageUpdateRoutingKey: viper.GetString("rabbitmq.image.update.routingkey"),
		minioURL:              viper.GetString("minio.url"),
		minioAccessKey:        viper.GetString("minio.accesskey"),
		minioSecret:           viper.GetString("minio.secret"),
		minioBucketName:       viper.GetString("minio.bucketname"),
		minioSecure:           viper.GetBool("minio.secure"),
	}
}
