package main

import (
	"context"
	"fmt"
	"insight/pkg/application"
	"insight/pkg/infrastructure/persistance"
	"insight/pkg/infrastructure/server"
	"insight/pkg/views"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	repoStore, repoErr := persistance.NewRepoStore()
	if repoErr != nil {
		log.Println(repoErr.Error())
		os.Exit(1)
	}

	migrationErr := repoStore.MigrateModels()
	if migrationErr != nil {
		log.Println(migrationErr.Error())
		os.Exit(1)
	}

	apps := application.NewApplication(*repoStore)
	viewStore := views.NewViewStore(apps)
	server := server.NewServer("", 8000, viewStore)
	server.Serve()
	httpServer := server.Run()

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	log.Println(fmt.Sprintf("Server Shutting Down... %s ", sig))

	//gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if shutDownError := httpServer.Shutdown(ctx); shutDownError != nil {
		log.Println(shutDownError.Error())
	}
	log.Println("Server Shut Down.")
}
