package caller

import (
	"context"

	"github.com/Alma-media/eop09/proto"
)

// Caller interface describes a caller for remote Port API.
type Caller interface {
	UploadStream(context.Context, <-chan *proto.Payload) error
	DownloadStream(context.Context, chan<- *proto.Payload) error
}
