package main

import (
	"docker_go/internal/http/handlers"
	"docker_go/internal/repo"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Go REST API Tutorial")
	if err := Run(); err != nil {
		log.Println(err)
	}

}

func Run() error {
	fmt.Println("starting up the application...")

	// create a database instance variable
	store, err := repo.NewDatabase()
	if err != nil {
		log.Println("Database Connection Failure")
		return err
	}

	// initialize the migrations functionality on the new database
	if err := store.MigrateDB(); err != nil {
		log.Println("failed to setup store migrations")
		return err
	}

	// set the database instance as the store for the user service implementation
	userService := repo.NewUserRepo(store)

	//// initialize a new handler with the user service
	handler := handlers.NewHandler(*userService)
	//
	// call the serve function to start the server
	if err := handler.Serve(); err != nil {
		log.Println("failed to gracefully serve our application")
		return err
	}

	return nil

}
