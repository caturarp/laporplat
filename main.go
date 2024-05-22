package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mailer "github.com/caturarp/laporplat/external/config"
	"github.com/caturarp/laporplat/external/mail"
	"github.com/caturarp/laporplat/handler"
	"github.com/caturarp/laporplat/logger"
	"github.com/caturarp/laporplat/repository"
	"github.com/caturarp/laporplat/server"
	"github.com/caturarp/laporplat/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
	defaultLog "github.com/rs/zerolog/log"
)

func main() {
	var configFileName string
	flag.StringVar(&configFileName, "c", "config.yml", "Config file name")

	flag.Parse()
	cfg := defaultConfig()
	cfg.loadFromEnv()

	if len(configFileName) > 0 {
		err := loadConfigFromFile(configFileName, &cfg)
		if err != nil {
			defaultLog.Warn().Str("file", configFileName).Err(err).Msg("cannot load config file, use defaults")
		}
	}
	log := logger.NewLogger()
	logger.SetLogger(log)

	pool, err := pgxpool.New(context.Background(), cfg.DBConfig.ConnectionString())
	if err != nil {
		defaultLog.Error().Err(err).Msg("unable to connect to database")
	}

	smtpConfig := mailer.GetSMTPConfig()
	sr := mail.NewSMTP(smtpConfig)

	ur := repository.NewUserRepository(pool)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	uur := repository.NewUnverifiedUserRepository(pool)

	au := usecase.NewAuthUsecase(ur, uur, sr, pool)
	ah := handler.NewAuthHandler(au)

	opts := server.RouterOpts{
		UserHandler: uh,
		AuthHandler: ah,
	}

	r := server.NewRouter(opts)
	appPort := cfg.Listen.Address()

	srv := http.Server{
		Addr:    appPort,
		Handler: r,
	}
	logger.Log.Info("Server listening on port: ", appPort)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Info("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Errorf("Server forced to shutdown:", err)
	}

	logger.Log.Info("Server exiting")
}
