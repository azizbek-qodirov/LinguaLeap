package main

import (
	"log"
	"net"

	cf "progress-service/config"
	"progress-service/storage"

	pb "progress-service/genprotos"
	service "progress-service/service"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()
	db, err := storage.NewPostgresStorage(config)
	em.CheckErr(err)
	defer db.DB.Close()

	listener, err := net.Listen("tcp", config.FORUM_SERVICE_PORT)
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
