package main

import (
	"log"
	"net"

	cf "learning-service/config"
	"learning-service/storage"

	pb "learning-service/genprotos"
	service "learning-service/service"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()
	db, err := storage.NewPostgresStorage(config)
	em.CheckErr(err)
	defer db.PgClient.Close()

	listener, err := net.Listen("tcp", config.LEARNING_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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

	q, err := ch.QueueDeclare(
		"progress_updates_queue", // name
		true,                     // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	vocabularyService := service.NewVocabularyService(db)

	go func() {
		for d := range msgs {
			vocabularyService.ProcessMessage(d.Body)
		}
	}()

	s := grpc.NewServer()
	pb.RegisterLessonServiceServer(s, service.NewLessonService(db))
	pb.RegisterExerciseServiceServer(s, service.NewExerciseService(db))
	pb.RegisterVocabularyServiceServer(s, vocabularyService)

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
