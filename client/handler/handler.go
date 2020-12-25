package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Alma-media/eop09/client/codec"
	"github.com/Alma-media/eop09/proto"
)

// Downloader interface.
type Downloader interface {
	DownloadStream(ctx context.Context, stream chan<- *proto.Payload) error
}

// CreateGetAllHandler creates an HTTP handler to return the entire port list.
// TODO: proper error handling.
func CreateGetAllHandler(caller Downloader, encode codec.EncoderFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stream := make(chan *proto.Payload)

		go func() {
			defer close(stream)

			if err := caller.DownloadStream(r.Context(), stream); err != nil {
				log.Printf("failed to encode stream: %s", err)
			}
		}()

		if err := encode(w, stream); err != nil {
			log.Printf("failed to encode stream: %s", err)
		}
	}
}
