package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/alwindoss/fragmos"
	"github.com/alwindoss/fragmos/internal/engine"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetPrefix("FRAGMOS: ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := fragmos.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	err = engine.Run(&cfg)
	if err != nil {
		log.Fatal(err)
	}

}
