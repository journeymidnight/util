package util

import (
	"testing"
)

func TestMsgPackMarshalEmpty(t *testing.T) {
	_, err := MsgPackMarshal(nil)
	if err != nil {
		t.Error("msgpack marshal empty failed:", err)
	}
}
