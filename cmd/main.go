package main

import (
	"fmt"
	"log"

	"github.com/alireza-msv/jet/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
