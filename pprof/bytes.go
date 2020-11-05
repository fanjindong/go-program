package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
)

// ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("bytes.Buffer: too large")

// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []byte {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {
			fmt.Println(ErrTooLarge)
		}
	}()
	return make([]byte, n)
}

//func TestMakeSlice(t *testing.T) {
//	makeSlice(math.MaxInt64)
//}

func main() {
	f, _ := os.Create("./pprof.prof")
	defer f.Close()
	makeSlice(math.MaxInt64)
	pprof.WriteHeapProfile(f)
}
