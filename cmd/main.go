package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Raziur306/kon-dol/internal/router"
	"github.com/Raziur306/kon-dol/internal/scheduler"
	"github.com/joho/godotenv"
)

// first load
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	port := os.Getenv("PORT")
	r := router.NewRouter()

	//start scheduler
	scheduler.StartScheduler()

	fmt.Println("🚀 Server is running at http://localhost:8080")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
