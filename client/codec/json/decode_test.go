package json

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Alma-media/eop09/client/codec"
	"github.com/Alma-media/eop09/proto"
)

var _ codec.DecoderFunc = Decode

var testJSON = `{"MYTBA":{"code":"55700","name":"Tanjong Baran","city":"Tanjong Baran","province":"Sarawak","country":"Malaysia","timezone":"Asia/Kuala_Lumpur","coordinates":[113.9769444,4.593333299999999],"unlocs":["MYTBA"]},"NOBVK":{"code":"40313","name":"Brevik","city":"Brevik","province":"Telemark","country":"Norway","timezone":"Europe/Oslo","coordinates":[9.7,59.05],"unlocs":["NOBVK"]},"TRBDM":{"code":"48963","name":"Bandirma","city":"Bandirma","province":"Balikesir","country":"Turkey","timezone":"Europe/Istanbul","coordinates":[27.97,40.35],"unlocs":["TRBDM"]}}`

func Test_Decode(t *testing.T) {
	t.Run("test if JSON input is ptoperly decoded to the receiver", func(t *testing.T) {
		var (
			reader = strings.NewReader(testJSON)
			stream = make(chan *proto.Payload)
		)

		go func() {
			if err := Decode(reader, stream); err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			close(stream)
		}()

		var index int

		for decoded := range stream {
			expected := testPorts[index]

			if decoded.Id != expected.Id {
				t.Errorf("id does not match: got %q, expected %q", decoded.Id, expected.Id)
			}

			if !reflect.DeepEqual(decoded.Port, expected.Port) {
				t.Errorf("port does not match:\n%#v\nexpected:\n%#v\n", decoded.Port, expected.Port)
			}

			index++
		}

		if expected := len(testPorts); index != expected {
			t.Errorf("unexpected index: %d, expected %d", index, expected)
		}
	})
}
