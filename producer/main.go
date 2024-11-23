package main

import (
	"log"
	"time"

	"producer/zmq"
)

func main() {
	
	serverAddress := "tcp://localhost:5555"
	delay := 4 * time.Second

	log.Println("Starting ZeroMQ Producer...")
	zmq.StartZeroMQProducer(serverAddress, delay)
}
