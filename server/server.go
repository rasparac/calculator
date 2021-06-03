package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/rasparac/calculator/calculator"
)

type (
	Server struct {
		config Config
	}
	Config struct {
		Host string
		Port string
	}
)

func New(c Config) *Server {
	return &Server{
		config: c,
	}
}

func (s *Server) Run(ctx context.Context) {
	log.Println("starting server")

	m := calculator.New()

	ser := http.Server{
		Addr:        s.config.Host + ":" + s.config.Port,
		Handler:     m,
		ReadTimeout: time.Duration(30 * time.Second),
		IdleTimeout: time.Duration(30 * time.Second),
	}
	go func() {
		if err := ser.ListenAndServe(); err != nil {
			log.Printf("event:ListenAndServe status: error, error: %v", err)
		}
	}()

	log.Println("server started on addres " + s.config.Host + " and port " + s.config.Port)

	<-ctx.Done()

	log.Println("graceful shutdown")

	if err := ser.Shutdown(ctx); err != nil {
		log.Printf("event:Shutdown status: error, error: %v", err)
	}
}
