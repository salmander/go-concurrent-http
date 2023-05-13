package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const URL = "https://api.open-meteo.com/v1/forecast?latitude=51.51&longitude=-0.13&current_weather=true"

func main() {
	http.HandleFunc("/", MyHandler)

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
	log.Println("Making a http request to get current weather information", time.Now().Format(time.RFC822))
	time.Sleep(5 * time.Second)
	// Make an HTTP GET request to the API
	resp, err := http.Get(URL)
	if err != nil {
		log.Println("failed to make request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body:", err)
		return
	}

	// Unmarshal the response JSON into the Response struct
	var response struct {
		CurrentWeather struct {
			Temperature float64 `json:"temperature"`
		} `json:"current_weather"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to unmarshal response:", err)
		return
	}

	// Log the temperature
	log.Println("Temperature:", response.CurrentWeather.Temperature)
}
