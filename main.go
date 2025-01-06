package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/adamwgriffin/go-microservice/api"
	db "github.com/adamwgriffin/go-microservice/db/sqlc"
	"github.com/adamwgriffin/go-microservice/lib"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := lib.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// This context is used to gracefully shutdown the app when any of the
	// specified OS interrupt signals are received
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()
	connPool, err := pgxpool.New(ctx, config.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to database")
	}
	store := db.NewStore(connPool)

	runGinServer(config, store)
}

func runGinServer(config lib.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start server")
	}
}
