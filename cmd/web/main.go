package main

import (
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	logger *slog.Logger
}

func main() {
	app := App{}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app.logger = logger

	server := http.Server{
		Addr:    ":3000",
		Handler: app.Routes(),
	}
	app.logger.Info("Starting server on port :3000")
	err := server.ListenAndServe()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}

}
