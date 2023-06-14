package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sijibomii/cryptopay/config"
	"github.com/sijibomii/cryptopay/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pg_url := os.Getenv("POSTGRES_URL")

	port := os.Getenv("PORT")
	num, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic("invalid port")
	}

	host := os.Getenv("HOST")

	sc := config.ServerConfig{
		Host: host,
		Port: num,
	}

	c := config.Config{
		Postgres: pg_url,
		Server:   sc,
	}

	server.Run(c)

}
