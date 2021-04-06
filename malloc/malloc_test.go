package malloc

import (
	"fmt"
	"testing"
)

type smallobj struct {
	arr [1 << 10]byte
}

type largeobj struct {
	arr [1 << 26]byte
}

func tiny() {
	x := 100000
	y := 100000
	fmt.Println(x, y)
}

func large() {
	large := largeobj{}
	println(&large)
}

func small() {
	small := smallobj{}
	print(&small)
}

func TestTiny(t *testing.T) {
	tiny()
}
