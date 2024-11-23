package zmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func StartZeroMQServer(collection *mongo.Collection) {
	
	socket, err := zmq4.NewSocket(zmq4.PULL)
	if err != nil {
		log.Fatalf("Failed to create ZeroMQ socket: %v", err)
	}
	defer socket.Close()

	
	address := "tcp://*:5555"
	if err := socket.Bind(address); err != nil {
		log.Fatalf("Failed to bind ZeroMQ socket to address %s: %v", address, err)
	}
	log.Printf("ZeroMQ server started and bound to %s", address)

	// Infinite loop to receive and process messages
	for {
		msg, err := socket.Recv(0)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}
		if err := processMessage(msg, collection); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}

// processMessage processes the received message and saves it to MongoDB.
func processMessage(msg string, collection *mongo.Collection) error {
	// Log the received message and prepare a processed version
	log.Printf("Received message: %s", msg)
	processedMsg := fmt.Sprintf("%s - Processed", msg)
	log.Printf("Processed message: %s", processedMsg)

	// Prepare the document for MongoDB insertion
	document := bson.M{
		"original_message":  msg,
		"processed_message": processedMsg,
		"timestamp":         time.Now(),
	}

	// Use a context with a timeout to ensure MongoDB operations don't hang indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert the document into the MongoDB collection
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return fmt.Errorf("failed to insert document into MongoDB: %w", err)
	}

	log.Println("Message successfully saved to MongoDB.")
	return nil
}
