package main

import (
	"auth-service/api"
	"auth-service/api/handlers"
	"auth-service/config"
	"auth-service/postgresql"
	"auth-service/service"
	"fmt"
)

func main() {
	cf := config.Load()
	em := config.NewErrorManager()

	conn, err := postgresql.ConnectDB(&cf)
	em.CheckErr(err)
	defer conn.Close()

	us := service.NewUserService(conn)
	handler := handlers.NewHandler(us)

	roter := api.NewRouter(handler)
	fmt.Println("Server is running on port ", cf.AUTH_PORT)
	if err := roter.Run(cf.AUTH_PORT); err != nil {
		panic(err)
	}
}
