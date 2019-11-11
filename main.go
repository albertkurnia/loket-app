package main

import (
	"flag"
	"os"
	"sync"

	config "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := config.Load(".env")
	if err != nil {
		log.Debug(err)
	}
	service := MakeHandler()

	// read specific flag
	HTTPPort := flag.String("http-port", os.Getenv("PORT"), "HTTP port")
	// parse all command flags
	flag.Parse()

	// override default port from .env file
	os.Setenv("PORT", *HTTPPort)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		service.HTTPServeMain()
	}()

	wg.Wait()
}
