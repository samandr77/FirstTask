package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MessageRequestBody struct {
	Message string `json:"message"`
}

var storedMessage string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", storedMessage)
}

func HelloPostHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody MessageRequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Невозможно декодировать JSON", http.StatusBadRequest)
		return
	}

	storedMessage = requestBody.Message

	fmt.Fprintf(w, "Получено сообщение: %s\n", storedMessage)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", HelloPostHandler).Methods("POST")
	fmt.Println("YES")
	http.ListenAndServe(":8080", router)
}
