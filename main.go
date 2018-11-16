package main

import (
	"github.com/caarlos0/env"
	"github.com/caarlos0/env/parsers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type server struct {
	router *http.ServeMux
	config struct {
		LISTEN_ADDR  string `env:"LISTEN_ADDR" envDefault:":8000"`
		PONG_MESSAGE string `env:"PONG_MESSAGE" envDefault:"pong!"`
	}
}

func (s *server) setupRoutes() {
	s.router.HandleFunc("/ping", s.handlePing())
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	app := server{
		router: http.NewServeMux(),
	}
	app.setupRoutes()

	err := env.ParseWithFuncs(
		&app.config,
		env.CustomParsers{
			parsers.URLType: parsers.URLFunc,
		})

	if err != nil {
		os.Exit(1)
	}

	http.ListenAndServe(app.config.LISTEN_ADDR, app.router)
}
