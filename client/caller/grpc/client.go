package grpc

import (
	"context"
	"errors"
	"io"

	"github.com/Alma-media/eop09/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// PortCaller is an RPC service caller.
type PortCaller struct {
	proto.StorageClient
}

// NewPortCaller returns a new caller for RPC service.
func NewPortCaller(cc *grpc.ClientConn) *PortCaller {
	service := proto.NewStorageClient(cc)
	return &PortCaller{service}
}

// UploadStream calls streaming RPC to upload the data.
func (caller *PortCaller) UploadStream(ctx context.Context, stream <-chan *proto.Payload) error {
	client, err := caller.StorageClient.Save(ctx)
	if err != nil {
		return err
	}

	for payload := range stream {
		if err := client.Send(payload); err != nil {
			return err
		}
	}

	_, err = client.CloseAndRecv()

	return err
}

// DownloadStream calls streaming RPC to fetch the data.
func (caller *PortCaller) DownloadStream(ctx context.Context, stream chan<- *proto.Payload) error {
	client, err := caller.StorageClient.Load(ctx, &empty.Empty{})
	if err != nil {
		return err
	}

	for {
		payload, err := client.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		stream <- payload
	}

	return nil
}
