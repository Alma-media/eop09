package json

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/Alma-media/eop09/proto"
)

var errInvalidInput = errors.New("invalid input format")

// Decode JSON input from io.Reader sending entities one by one to the stream channel.
func Decode(reader io.Reader, stream chan<- *proto.Payload) error {
	const (
		stateWaitForObjectOpen int = iota
		stateWaitForObjectKey
		stateWaitForObjectClose
	)

	var (
		decoder = json.NewDecoder(reader)
		state   int
	)

	for {
		token, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		switch state {
		case stateWaitForObjectOpen:
			if t, ok := token.(json.Delim); !ok || t != '{' {
				return errInvalidInput
			}

			state = stateWaitForObjectKey

			continue
		case stateWaitForObjectKey:
			switch t := token.(type) {
			case json.Delim:
				if t != '}' {
					return errInvalidInput
				}

				state = stateWaitForObjectClose

				continue
			case string:
				port := new(proto.Port)

				// TODO: implement wrapper over proto.Entity providing
				// custom JSON unmarshaler to avoid using reflect and
				// speed up the decoder
				if err := decoder.Decode(port); err != nil {
					return err
				}

				// blocking operation: parser should not be faster than
				// storage since we are going to use straming API
				stream <- &proto.Payload{
					Id:   t,
					Port: port,
				}
			}

			continue
		case stateWaitForObjectClose:
			// consume the last token
		default:
			// decoder is broken
			return errInvalidInput
		}
	}
}
