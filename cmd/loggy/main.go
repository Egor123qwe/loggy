package main

import (
	"log"

	logger "github.com/Egor123qwe/loggy/pkg/service"
	"github.com/google/uuid"
)

func init() {
	logOpts := logger.Options{
		Level:  logger.DEBUG,
		Module: "loggy",

		ToStderr: true,

		Server: []logger.Server{
			{
				URL:  "http://localhost:8082",
				URLs: []string{"localhost:19092"},

				Credentials: &logger.Credentials{Username: "admin", Password: "admin"},
			},
		},

		File: []logger.File{
			{
				Name:      "test.log",
				MaxSizeMb: 100,
				MaxFiles:  10,
			},
		},
	}

	if err := logger.Init(logOpts); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer logger.Close()

	traceID := uuid.New().String()

	log := logger.Log(traceID)

	// tests
	log.Infof("Hello, world!")
	log.Criticalf("Hello, world!")
}
