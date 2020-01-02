package requests__test

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)
import (
	"github.com/fanjindong/go-requests"
)

func TestRequestsGet(t *testing.T) {
	resp, _ := requests.Get("http://dpp.boluome.com/ping")
	assert.True(t, resp.StatusCode == 200)

	data := make(map[string]interface{})
	_ = resp.Json(&data)

	assert.Equal(t, data["code"], float64(0))
	t.Log(data)
}

func TestRequestsPost(t *testing.T) {
	name := strconv.Itoa(rand.Intn(10))
	resp, _ := requests.Put("http://192.168.3.3:40010/v0/user/14076", requests.Json{"name": name})
	assert.Equal(t, resp.StatusCode, 200)

	resp, _ = requests.Get("http://192.168.3.3:40010/v0/user/getInfo?ownerUid=14076&publicQaAmount=1")
	assert.Equal(t, resp.StatusCode, 200)
	respData := make(map[string]interface{})
	_ = resp.Json(&respData)
	assert.Equal(t, respData["data"].(map[string]interface{})["name"], name)
}
