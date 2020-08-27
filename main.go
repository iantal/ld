package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/internal/files"
	"github.com/iantal/ld/internal/server"
	protos "github.com/iantal/ld/protos/ld"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	viper.AutomaticEnv()
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	bp := fmt.Sprintf("%v", viper.Get("BASE_PATH"))
	rkHost := fmt.Sprintf("%v", viper.Get("RK_HOST"))

	stor, err := files.NewLocal(bp, 1024*1000*1000*5)
	if err != nil {
		log.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	c := server.NewLinguist(log, bp, rkHost, stor)

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
