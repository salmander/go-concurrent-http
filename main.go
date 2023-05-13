package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/my-endpoint", MyHandler)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	go MyConcurrentFunction()

	// Return a 204 No Content response immediately
	w.WriteHeader(http.StatusNoContent)
}

func MyConcurrentFunction() {
	// Perform background tasks or computations here
	fmt.Println("concurrent function doing task 1", time.Now().Format(time.RFC822))
	time.Sleep(1 * time.Second)

	fmt.Println("concurrent function doing task 2", time.Now().Format(time.RFC822))
	time.Sleep(5 * time.Second)
	fmt.Println("concurrent function all done", time.Now().Format(time.RFC822))
}
