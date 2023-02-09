package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("dev.env")
}

func main() {
	server := NewServer(os.Getenv("DB_URL"))
	//log to a file
	log.Fatal(server.Run(os.Getenv("PORT")))

}
