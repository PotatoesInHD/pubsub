package main

import (
    amqp "github.com/rabbitmq/amqp091-go"
    "fmt"
    "log"
    "os"
    "os/signal"
)

func main() {
	fmt.Println("Starting Peril server...")

	addr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Fatalf("Error Dialing rabbitmq: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connection successful!")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Shutting down server!")
}
