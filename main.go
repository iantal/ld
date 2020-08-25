package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/internal/server"
	protos "github.com/iantal/ld/protos/ld"
	"github.com/nicholasjackson/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
	basePath = env.String("BASE_PATH", false, "./repos", "Base path to save uploads")
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewLinguist(log)

	// register the currency server
	protos.RegisterUsedLanguagesServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8003))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	log.Info("Starting server", "bind_address", l.Addr().String())
	// listen for requests
	gs.Serve(l)
}
