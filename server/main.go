package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Alma-media/eop09/proto"
	"github.com/Alma-media/eop09/server/service"
	"github.com/Alma-media/eop09/server/storage"
	"google.golang.org/grpc"
)

var port int

func init() {
	flag.IntVar(&port, "port", 5050, "the server port")
	flag.Parse()
}

func main() {
	var (
		repository = storage.NewInMemory()
		portServer = service.NewPortServer(repository)
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("cannot start the server: %s", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterStorageServer(grpcServer, portServer)

	log.Printf("starting GRPC server on %s", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to start the server: %s", err)
	}
}
