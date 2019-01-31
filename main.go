package main

import (
	"go-docker-example/app/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", controllers.Index)

	log.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println(err)
	}

}
