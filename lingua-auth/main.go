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

	pgsql, mongo, err := postgresql.ConnectDB(&cf)
	em.CheckErr(err)
	defer pgsql.Close()

	us := service.NewUserService(pgsql, mongo)
	handler := handlers.NewHandler(us)

	roter := api.NewRouter(handler)
	fmt.Println("Server is running on port ", cf.AUTH_PORT)
	if err := roter.Run(cf.AUTH_PORT); err != nil {
		panic(err)
	}
}
