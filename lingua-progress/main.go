package main

import (
	"fmt"
	"log"
	"net"

	cf "progress-service/config"
	"progress-service/storage"

	pb "progress-service/genprotos"
	service "progress-service/service"

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

	LearningConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.LEARNING_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer LearningConn.Close()

	db, err := storage.NewPostgresStorage(config, LearningConn)
	em.CheckErr(err)
	defer db.PgClient.Close()

	listener, err := net.Listen("tcp", config.PROGRESS_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// ###### RabbitMQ ##### //
	rabbitConn, err := amqp.Dial(config.RABBITMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer rabbitConn.Close()

	ch, err := rabbitConn.Channel()
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

	q, err := ch.QueueDeclare(
		"progress_updates", // name
		true,               // durable
		false,              // delete when unused
		true,               // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,             // queue name
		"",                 // routing key
		"progress_updates", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	// msgs, err := ch.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto-ack
	// 	false,  // exclusive
	// 	false,  // no-local
	// 	false,  // no-wait
	// 	nil,    // args
	// )
	failOnError(err, "Failed to register a consumer")

	quizService := service.NewQuizService(db, LearningConn, rabbitConn)

	s := grpc.NewServer()
	pb.RegisterUserLessonServiceServer(s, service.NewUserLessonService(db))
	pb.RegisterUserDataServiceServer(s, service.NewUserDataService(db))
	pb.RegisterQuizServiceServer(s, quizService)

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
