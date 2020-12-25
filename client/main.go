package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpcCaller "github.com/Alma-media/eop09/client/caller/grpc"
	"github.com/Alma-media/eop09/client/codec/json"
	"github.com/Alma-media/eop09/client/handler"
	"github.com/Alma-media/eop09/proto"
	"google.golang.org/grpc"
)

var (
	fileName string
	httpPort int
	grpcAddr string
)

func init() {
	flag.StringVar(&fileName, "file", "", "*path to the file (required)")
	flag.IntVar(&httpPort, "http-port", 8080, "HTTP port")
	flag.StringVar(&grpcAddr, "grpc-addr", "localhost:5050", "GRPC address")
	flag.Parse()

	if fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// TODO:
// - https
// - use advanced routing
// - http middleware (jwt, auth, timeout etc)
// - choose one of available codecs according to the content type
// - TLS for grpc
// - grpc interceptors
// - configure with ENV variables
func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	defer file.Close()

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %s", err)
	}

	var (
		httpPort    = fmt.Sprintf(":%d", httpPort)
		stream      = make(chan *proto.Payload)
		ctx, cancel = context.WithCancel(context.Background())
		caller      = grpcCaller.NewPortCaller(ctx.Done(), conn)
		errs        = make(chan error)
	)

	go func() {
		defer close(stream)

		if err := json.Decode(file, stream); err != nil {
			errs <- fmt.Errorf("cannot decode stream: %w", err)
		}
	}()

	go func() {
		if err := caller.UploadStream(ctx, stream); err != nil {
			errs <- fmt.Errorf("failed to bootstrap the service: %w", err)
		}
	}()

	router := http.NewServeMux()
	router.Handle("/json", handler.CreateGetAllHandler(caller, json.Encode))

	server := &http.Server{
		Addr:    httpPort,
		Handler: router,
	}

	go func() {
		log.Printf("starting HTTP server on port %s", httpPort)
		if err := server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			errs <- fmt.Errorf("failed to start HTTP server: %w", err)
		}
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-term:
		cancel()

		// TODO: pass a proper context
		if err := server.Shutdown(context.Background()); err != nil {
			errs <- fmt.Errorf("failed to stop HTTP server: %w", err)
		}

		log.Printf("client gracefully stopped")
	case err = <-errs:
		log.Fatal(err)
	}
}
