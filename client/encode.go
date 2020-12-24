package client

import (
	"encoding/json"
	"io"

	"github.com/Alma-media/eop09/proto"
)

// Encode entities from the stream channel writing to the provided io.Writer.
func Encode(writer io.Writer, stream <-chan *proto.Payload) error {
	var (
		encoder = json.NewEncoder(writer)
		isFirst = true
	)

	if _, err := writer.Write([]byte{'{'}); err != nil {
		return err
	}

	for payload := range stream {
		if isFirst {
			isFirst = false
		} else if _, err := writer.Write([]byte{','}); err != nil {
			return err
		}

		if err := encoder.Encode(payload.Id); err != nil {
			return err
		}

		if _, err := writer.Write([]byte{':'}); err != nil {
			return err
		}

		if err := encoder.Encode(payload.Port); err != nil {
			return err
		}
	}

	_, err := writer.Write([]byte{'}'})

	return err
}
