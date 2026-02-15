package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer response.Body.Close()

	// Check if status code is OK
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d", response.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse JSON
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Print the data
	fmt.Println("Todo Item:")
	fmt.Printf("User ID: %d\n", todo.UserId)
	fmt.Printf("ID: %d\n", todo.Id)
	fmt.Printf("Title: %s\n", todo.Title)
	fmt.Printf("Completed: %t\n", todo.Completed)
}
