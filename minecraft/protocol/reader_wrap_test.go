package protocol

import (
	"bytes"
	"testing"
)

func TestNewReaderWrapsAnyByteReader(t *testing.T) {
	r := NewReader(bytes.NewReader([]byte{7, 9}), 0, true)

	var first, second uint8
	r.Uint8(&first)
	r.Uint8(&second)
	if first != 7 || second != 9 {
		t.Fatalf("read %d, %d, want 7, 9", first, second)
	}
}
