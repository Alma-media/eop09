package service

import (
	"errors"
	"io"

	"github.com/Alma-media/eop09/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PortStorage is a port repository providing basic data operation.
type PortStorage interface {
	Save(key string, port *proto.Port)
	Each(func(key string, port *proto.Port) error) error
}

// PortServer is the server that provides operations on Port(s).
type PortServer struct {
	proto.UnimplementedStorageServer

	PortStorage
}

// NewPortServer returns a new PortServer.
func NewPortServer(storage PortStorage) *PortServer {
	return &PortServer{
		PortStorage: storage,
	}
}

// Load is a server-streaming RPC to return all available ports.
// TODO:
// - logging
func (server *PortServer) Load(_ *empty.Empty, stream proto.Storage_LoadServer) error {
	ctx := stream.Context()

	if err := server.PortStorage.Each(
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
			if errors.Is(err, io.EOF) {
				break
			}

			if err != nil {
				return err
			}

			server.PortStorage.Save(pair.Id, pair.Port)
		}

		return nil
	}(); err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return stream.SendAndClose(new(empty.Empty))
}
