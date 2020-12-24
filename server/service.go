package server

import (
	"io"

	"github.com/Alma-media/eop09/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PortServer is the server that provides operations on Port(s).
type PortServer struct {
	*Storage

	proto.UnimplementedStorageServer
}

// NewPortServer returns a new PortServer.
func NewPortServer(storage *Storage) *PortServer {
	return &PortServer{Storage: storage}
}

// Load is a server-streaming RPC to return all available ports.
// TODO:
// - advanced logging
// - avoid logging to stderr/stdout (performance issue)
func (server *PortServer) Load(_ *empty.Empty, stream proto.Storage_LoadServer) error {
	ctx := stream.Context()

	if err := server.Storage.Each(
		func(id string, port *proto.Port) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if err := stream.Send(&proto.Payload{
				Id:   id,
				Port: port,
			}); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
}

// Save is a server-streaming RPC to store ports to the repository.
func (server *PortServer) Save(stream proto.Storage_SaveServer) error {
	ctx := stream.Context()

	if err := func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			pair, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			server.Storage.Save(pair.Id, pair.Port)
		}

		return nil
	}(); err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return stream.SendAndClose(new(empty.Empty))
}
