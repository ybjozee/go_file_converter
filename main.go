package main

import (
	"app/database"
	"app/handler"
	"fmt"
	"net/http"
)

func main() {
	database.Connect()
	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080", handler.GetRouter())
}
