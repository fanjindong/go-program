package compress

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"log"
	"testing"
)

func TestZlib(t *testing.T) {
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write([]byte(data))
	w.Close()
	t.Log(len(data), len(b.Bytes()))
	// 1558606 23643
	buf := bytes.NewReader(b.Bytes())
	r, err := zlib.NewReader(buf)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	got, _ := io.ReadAll(r)
	t.Log(data == string(got))
}

func TestGzip(t *testing.T) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	t.Log(len(data), len(buf.Bytes()))
	// 1558606 23655
}
