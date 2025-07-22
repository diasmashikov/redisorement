package main

import (
	"log"
	"net/http"
	"redisorement/internal/handler"
)


func main() {
	http.HandleFunc("/", handler.HelloHandler)
	http.HandleFunc("/set", handler.SetHandler)
	http.HandleFunc("/get", handler.GetHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
