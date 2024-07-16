package main

import (
	"log"
	"net"

	cf "learning-service/config"
	"learning-service/storage"

	pb "learning-service/genprotos"
	service "learning-service/service"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()
	db, err := storage.NewPostgresStorage(config)
	em.CheckErr(err)
	defer db.DB.Close()

	listener, err := net.Listen("tcp", config.LEARNING_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterLessonServiceServer(s, service.NewLessonService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
