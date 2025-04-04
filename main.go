package main

import (
	"fmt"
	"log"

	"github.com/VictorHRRios/blog_aggregator/internal/config"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	configFile.SetUser("VictorHugo")

	configFile, err = config.Read()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("configFile: %v\n", configFile)
}
