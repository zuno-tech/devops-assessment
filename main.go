package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var serverStartTime uint

func main() {

	serverRouter := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      serverRouter,
	}

	fmt.Println("Starting the assessment webserver...")
	serverStartTime = uint(time.Now().Unix())
	serverRouter.HandleFunc("/", Uptime)
	serverRouter.HandleFunc("/external-check", ExternalCheck)
	log.Fatal(server.ListenAndServe())
}

func Uptime(response http.ResponseWriter, request *http.Request) {
	diceRoll := rand.Float32()
	if diceRoll < 0.9 {
		io.WriteString(response, fmt.Sprintf("Uptime: %d", uint(time.Now().Unix())-serverStartTime))
	} else {
		response.WriteHeader(500)
		io.WriteString(response, "We're down!")
	}
}

func ExternalCheck(response http.ResponseWriter, request *http.Request) {
	externalResponse, err := http.Get("https://example.com")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, fmt.Sprintf("External check failed: %s", err))
		return
	}
	defer externalResponse.Body.Close()
	if externalResponse.StatusCode != http.StatusOK {
		response.WriteHeader(http.StatusBadGateway)
		io.WriteString(response, fmt.Sprintf("External check returned: %s", externalResponse.Status))
		return
	}
	io.WriteString(response, "External check successful")
}
