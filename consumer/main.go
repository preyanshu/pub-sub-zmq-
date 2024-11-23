package main

import (
	"consumer/database"
	"consumer/zmq"
	"context"
)

func main() {
	
	mongoClient, collection := database.InitMongoDB("mongodb://localhost:27017", "zmqdb", "processed_messages")
	defer mongoClient.Disconnect(context.TODO())

	
	zmq.StartZeroMQServer( collection)
}
