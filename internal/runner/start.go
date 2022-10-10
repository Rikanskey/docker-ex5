package runner

import (
	"docker_ex5/internal/app"
	"docker_ex5/internal/config"
	"docker_ex5/internal/port/http"
	"docker_ex5/internal/server"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

func Start(configDir string) {
	cfg := newConfig(configDir)
	application := newApplication()
	startServer(cfg, application)
}

func newConfig(configDir string) *config.Config {
	cfg, err := config.New(configDir)
	if err != nil {
		log.Panicln(err)
	}

	return cfg
}

func newApplication() app.Application {

	return app.Application{}
}

func startServer(cfg *config.Config, application app.Application) {
	logrus.Info(fmt.Sprintf("Starting HTTP server on address: %s", cfg.HTTP.Port))

	httpServer := server.New(cfg, http.NewHandler(application))

	err := httpServer.Run()

	logrus.WithError(err).Fatal("HTTP server stopped")
}
