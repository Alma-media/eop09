package grpc

import (
	"context"
	"errors"
	"io"

	"github.com/Alma-media/eop09/client/caller"
	"github.com/Alma-media/eop09/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var _ caller.Caller = (*PortCaller)(nil)

// PortCaller is an RPC service caller.
type PortCaller struct {
	proto.StorageClient
	stopChan <-chan struct{}
}

// NewPortCaller returns a new caller for RPC service.
func NewPortCaller(stop <-chan struct{}, cc *grpc.ClientConn) *PortCaller {
	return &PortCaller{
		StorageClient: proto.NewStorageClient(cc),
		stopChan:      stop,
	}
}

// UploadStream calls streaming RPC to upload the data.
func (caller *PortCaller) UploadStream(ctx context.Context, stream <-chan *proto.Payload) error {
	client, err := caller.StorageClient.Save(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case <-caller.stopChan:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		case payload, ok := <-stream:
			if !ok {
				_, err = client.CloseAndRecv()

				return err
			}

			if err := client.Send(payload); err != nil {
				return err
			}
		}
	}
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

		select {
		case <-caller.stopChan:
			return nil
		case stream <- payload:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
