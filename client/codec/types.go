package codec

import (
	"io"

	"github.com/Alma-media/eop09/proto"
)

// EncoderFunc is a codec agnostic encoder.
type EncoderFunc func(io.Writer, <-chan *proto.Payload) error

// DecoderFunc is a codec agnostic decoder.
type DecoderFunc func(reader io.Reader, stream chan<- *proto.Payload) error
