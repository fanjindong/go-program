package main

func main() {
	r := &Reader{data: []byte("我三等奖逢狼时刻")}
	read := make([]byte, 8)
	r.ReadNotEscape(read)
	r.ReadEscape(8)

}

type Reader struct {
	data []byte
}

//ReadNotEscape good
func (r Reader) ReadNotEscape(p []byte) (n int, err error) {
	copy(p, r.data)
	return len(p), nil
}

//ReadEscape bad
func (r Reader) ReadEscape(n int) (p []byte, err error) {
	p = make([]byte, n)
	copy(p, r.data)
	return p, nil
}
