package json

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Alma-media/eop09/client/codec"
	"github.com/Alma-media/eop09/proto"
)

var _ codec.EncoderFunc = Encode

var testPorts = []*proto.Payload{
	{
		Id: "MYTBA",
		Port: &proto.Port{
			Code:        "55700",
			Name:        "Tanjong Baran",
			City:        "Tanjong Baran",
			Province:    "Sarawak",
			Country:     "Malaysia",
			Timezone:    "Asia/Kuala_Lumpur",
			Coordinates: []float64{113.9769444, 4.593333299999999},
			Unlocs:      []string{"MYTBA"},
		},
	},
	{
		Id: "NOBVK",
		Port: &proto.Port{
			Code:        "40313",
			Name:        "Brevik",
			City:        "Brevik",
			Province:    "Telemark",
			Country:     "Norway",
			Timezone:    "Europe/Oslo",
			Coordinates: []float64{9.7, 59.05},
			Unlocs:      []string{"NOBVK"},
		},
	},
	{
		Id: "TRBDM",
		Port: &proto.Port{
			Code:        "48963",
			Name:        "Bandirma",
			City:        "Bandirma",
			Province:    "Balikesir",
			Country:     "Turkey",
			Timezone:    "Europe/Istanbul",
			Coordinates: []float64{27.97, 40.35},
			Unlocs:      []string{"TRBDM"},
		},
	},
}

func Test_Encode(t *testing.T) {
	var (
		buff   bytes.Buffer
		stream = make(chan *proto.Payload)
	)

	go func() {
		for _, payload := range testPorts {
			stream <- payload
		}

		close(stream)
	}()

	if err := Encode(&buff, stream); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if actual := strings.Replace(buff.String(), "\n", "", -1); actual != testJSON {
		t.Errorf("output:\n%s\nwas expected to be:\n%s\n", actual, testJSON)
	}
}
