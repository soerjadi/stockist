package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/soerjadi/stockist/internal/config"
	"github.com/soerjadi/stockist/internal/delivery/rest"
	strHndl "github.com/soerjadi/stockist/internal/delivery/rest/store"
	userHndl "github.com/soerjadi/stockist/internal/delivery/rest/user"
	"github.com/soerjadi/stockist/internal/pkg/log"
	"github.com/soerjadi/stockist/internal/pkg/log/logger"
	"github.com/soerjadi/stockist/internal/repository/store"
	"github.com/soerjadi/stockist/internal/repository/user"
	strUcs "github.com/soerjadi/stockist/internal/usecase/store"
	userUcs "github.com/soerjadi/stockist/internal/usecase/user"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Errorw("[Config] error reading config from file.", logger.KV{
			"err": err,
		})
		return
	}

	// initialize log
	log.InitLog(cfg.Server.LogPath, cfg.Server.Name)

	// open database connection
	dataSource := fmt.Sprintf("user=%s password=%s	host=%s port=%s dbname=%s sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sqlx.Open(cfg.Database.Driver, dataSource)
	if err != nil {
		log.Errorw("cannot connect to db.", logger.KV{
			"err": err,
		})
		return
	}

	handlers, err := initiateHandler(cfg, db)
	if err != nil {
		log.Errorw("unable to initiate handler.", logger.KV{
			"err": err,
		})
		return
	}

	r := mux.NewRouter()
	rest.RegisterHandlers(r, handlers...)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Infow("server running in ", logger.KV{
		"port": cfg.Server.Port,
	})

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Errorw("error running apps.", logger.KV{
				"err": err,
			})
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait  for.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.WaitTimeout())
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down")
	os.Exit(0)
}

func initiateHandler(cfg *config.Config, db *sqlx.DB) ([]rest.API, error) {
	validate := validator.New()

	userRepository, err := user.GetRepository(db)
	if err != nil {
		log.Errorw("[initiateHandler] failed initiate userRepository", logger.KV{
			"err": err,
		})
		return nil, err
	}
	storeRepository, err := store.GetRepository(db)

	userUsecase := userUcs.GetUsecase(userRepository, cfg)
	storeUsecase := strUcs.GetUsecase(storeRepository)

	userHandler := userHndl.NewHandler(userUsecase, validate)
	storeHandler := strHndl.NewHandler(storeUsecase, userUsecase, validate, cfg)

	return []rest.API{
		userHandler,
		storeHandler,
	}, nil
}
