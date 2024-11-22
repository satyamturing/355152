package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	websocketServerURL = "ws://localhost:8080/ws" // Replace with your WebSocket server URL
	restAPIURL         = "http://localhost:8080/api" // Replace with your REST API endpoint URL
)

func main() {
	// Connect to the WebSocket server
	wsDialer := websocket.DefaultDialer
	wsConn, _, err := wsDialer.Dial(websocketServerURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer wsConn.Close()

	// Send a message to the WebSocket server
	err = wsConn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket Server!"))
	if err != nil {
		log.Fatal("Error sending message to WebSocket server:", err)
	}
	fmt.Println("Message sent to WebSocket server.")

	// Read a message from the WebSocket server
	_, message, err := wsConn.ReadMessage()
	if err != nil {
		log.Fatal("Error reading message from WebSocket server:", err)
	}
	fmt.Println("Received message from WebSocket server:", string(message))

	// Make a request to the REST API endpoint
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(restAPIURL)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status: %s", response.Status)
	}

	fmt.Println("REST API response status:", response.Status)

	// Read the REST API response body
	body, err := io.ReadAll(response.Body) // Use io.ReadAll
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	fmt.Println("REST API response body:", string(body))
}
