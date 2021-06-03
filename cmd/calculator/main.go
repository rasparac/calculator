package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/rasparac/calculator/server"
)

func main() {

	serverConf := server.Config{}

	flag.StringVar(&serverConf.Host, "HOST", "0.0.0.0", "use this variable to set server host")
	flag.StringVar(&serverConf.Port, "PORT", "9999", "use this variable to set server port")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		<-sigCh
		cancel()
	}()

	s := server.New(serverConf)

	s.Run(ctx)
}
