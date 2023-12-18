package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"sbs-be/app"
	"sbs-be/config"
)

func main() {

	mysqlConn, errMySQL := config.ConnectMySQL()
	if errMySQL != nil {
		fmt.Sprintf("error postgresql connection: ", errMySQL)
	}

	router := app.InitRouter(mysqlConn)
	fmt.Printf("routes Initialized")

	port := config.CONFIG["PORT"]
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Server Initialized")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Sprintf("Server Shutdown:", err)
	}
	fmt.Printf("Server exiting")
}
