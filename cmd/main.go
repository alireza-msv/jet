package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alireza-msv/jet/internal/app"
	"github.com/alireza-msv/jet/internal/config"
	"github.com/alireza-msv/jet/internal/storage"
	"github.com/go-co-op/gocron/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal("Error on creating Scheduler", err)
	}

	strg := NewStorage()
	a := app.NewApp(cfg, strg)

	j, err := s.NewJob(
		gocron.CronJob(cfg.Schedule, false),
		gocron.NewTask(func(ap *app.App) {
			ap.Start()
		}, a),
	)
	if err != nil {
		log.Fatal("Error on creating a new Jon", err)
	}

	fmt.Printf("Job with ID %s created", j.ID())

	s.Start()

	select {
	case <-time.After(time.Minute):
	}
}

// The function is just a simple demonstration the power of using interfaces
// over single purpose structs
func NewStorage() storage.Storage {
	// Based on the config a new instance of Storage (e.g. S3, Local file or Google Cloud Storage)
	// will be created and returned
	return storage.NewLocalStorage()
}
