package _package

import (
	"github.com/tcnksm/go-httpstat"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHttpstat(t *testing.T) {
	req, err := http.NewRequest("GET", "http://deeeet.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create go-httpstat powered context and pass it to http.Request
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	result.End(time.Now())

	// Show results
	log.Printf("%+v", result)
}

type key1 struct{}
type key2 struct{}

func TestNullStructKey(t *testing.T) {
	m := map[interface{}]int{
		struct{}{}:  3,
		key1{}:      1,
		key2{}:      2,
		&struct{}{}: 3,
		&key1{}:     1,
		&key2{}:     2,
	}
	t.Log(m[key1{}], m[key2{}], m[struct{}{}])    // 1 2 3
	t.Log(m[&key1{}], m[&key2{}], m[&struct{}{}]) // 0 0 0
}
