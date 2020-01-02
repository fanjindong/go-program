package http__test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

var PingUrl = "http://dpp.boluome.com/ping"

func TestNetHttpGet(t *testing.T) {
	resp, _ := http.Get(PingUrl)
	defer resp.Body.Close()
	//resp: &{Status:200 OK StatusCode:200 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Cache-Control:[private] Content-Type:[application/json] Set-Cookie:[QINGCLOUDELB=d9a2454c187d2875afb6701eb80e9c8761ebcf3b54797eae61b25b90f71273ea; path=/; HttpOnly]] Body:0xc00016e060 ContentLength:-1 TransferEncoding:[chunked] Close:true Uncompressed:true Trailer:map[] Request:0xc000114200 TLS:<nil>}
	t.Logf("resp: %+v", resp)

	client := &http.Client{}
	resp, _ = client.Get(PingUrl)
	defer resp.Body.Close()
	t.Logf("resp: %+v", resp)

	req, _ := http.NewRequest("GET", PingUrl, nil)
	req.Header.Set("name", "fjd")
	req.Header.Add("platform", "mac")
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	t.Logf("resp: %+v", resp)

	respBodyRead, _ := ioutil.ReadAll(resp.Body)
	t.Logf("respBodyRead: %+v", respBodyRead)
	respMap := make(map[string]interface{})
	_ = json.Unmarshal(respBodyRead, &respMap)
	t.Logf("respMap: %+v", respMap)
}
