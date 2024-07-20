package main

import (
	"fmt"
	"gateway-service/api"
	cf "gateway-service/config"
	"log"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()

	LearningConn, err := grpc.NewClient(fmt.Sprintf("learning_service%s", config.LEARNING_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer LearningConn.Close()

	ProgressConn, err := grpc.NewClient(fmt.Sprintf("progress_service%s", config.PROGRESS_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer ProgressConn.Close()

	conn, err := amqp.Dial(config.RABBITMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"progress_updates", // name
		"fanout",           // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	r := api.NewRouter(LearningConn, ProgressConn, ch)

	fmt.Printf("Server started on port %s\n", config.HTTP_PORT)
	if err := r.Run(config.HTTP_PORT); err != nil {
		panic(err)
	}
}
