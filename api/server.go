package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"levelup/api/controllers"
	"levelup/api/seed"
)

var server = controllers.Server{}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DBDriver"), os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBPort"), os.Getenv("DBHost"), os.Getenv("DBName"))

	seed.Load(server.DB)

	server.Run(":8080")

}