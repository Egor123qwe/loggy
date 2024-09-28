package main

import (
	"log"

	loggerModel "github.com/Egor123qwe/loggy/pkg/model/log"
	logger "github.com/Egor123qwe/loggy/pkg/service"
	"github.com/google/uuid"
)

func init() {
	logOpts := logger.Options{
		Level:  loggerModel.DEBUG,
		Module: "loggy",

		ToStderr: true,

		Server: &logger.Server{
			URLs:     []string{"localhost:19092"},
			WithWait: true,

			Credentials: nil,
		},

		File: &logger.File{
			Name:      "test.log",
			MaxSizeMb: 100,
			MaxFiles:  10,
		},
	}

	if err := logger.Init(logOpts); err != nil {
		log.Fatal(err)
	}
}

func main() {
	traceID := uuid.New().String()

	log := logger.Log(traceID)

	// tests
	log.Infof("Hello, world!")
	log.Criticalf("Hello, world!")
}
