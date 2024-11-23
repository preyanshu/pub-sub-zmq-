package zmq

import (
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
)

// StartZeroMQProducer initializes a ZeroMQ PUSH socket and sends messages in a loop.
func StartZeroMQProducer(serverAddress string, delay time.Duration) {
	// Create a new ZeroMQ PUSH socket
	socket, err := zmq4.NewSocket(zmq4.PUSH)
	if err != nil {
		log.Fatalf("Failed to create ZeroMQ socket: %v", err)
	}
	defer socket.Close()

	// Connect the socket to the ZeroMQ server address
	if err := socket.Connect(serverAddress); err != nil {
		log.Fatalf("Failed to connect to ZeroMQ server at %s: %v", serverAddress, err)
	}
	log.Printf("Connected to ZeroMQ server at %s", serverAddress)

	// Loop to send messages
	messageID := 1
	for {
		// Generate a message with an incrementing ID
		message := fmt.Sprintf("Message %d - Sent at: %s", messageID, time.Now().Format(time.RFC3339))

		// Send the message to the server
		_, err := socket.Send(message, 0)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}

		log.Printf("Message sent: %s", message)

		// Increment message ID and add a delay between messages
		messageID++
		time.Sleep(delay) // Use the delay duration provided
	}
}
