package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("dev.env")
}

func main() {
	//get command line args
	args := os.Args[1:]
	var server Server
	if len(args) == 0 {
		server = NewServer(os.Getenv("DB_URL"))
	} else if args[0] == "-local" {
		fmt.Println("Running local mock server")
		server = NewMockServer()
	} else {
		log.Fatal("Invalid command line args")
	}
	//log to a file
	log.Fatal(server.Run(os.Getenv("PORT")))

}
