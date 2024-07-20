package main

import (
	"fmt"
	"gateway-service/api"
	cf "gateway-service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()

	LearningConn, err := grpc.NewClient(fmt.Sprintf("learning_service%s", config.LEARNING_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer LearningConn.Close()

	ProgressConn, err := grpc.NewClient(fmt.Sprintf("progress_service%s", config.PROGRESS_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer LearningConn.Close()

	r := api.NewRouter(LearningConn, ProgressConn)

	fmt.Printf("Server started on port %s\n", config.HTTP_PORT)
	if r.Run(config.HTTP_PORT); err != nil {
		panic(err)
	}
}
