package gob

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestUint(t *testing.T) {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(uint(256))
	t.Log(buf.Bytes())
}
