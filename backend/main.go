package main

import (
	"backend/db"
	"backend/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	logfile, err := os.OpenFile("backend.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer logfile.Close()
	//log.SetOutput(logfile)

	log.SetPrefix("Quantum - Backend: ")

	log.Print("Connecting to database")
	db.Init()
	log.Println("Server started on port 8080")
	err = http.ListenAndServe("127.0.0.1:8080", routers.GetRouter())
	if err != nil {
		log.Print(err.Error())
	}

}
