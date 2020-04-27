package requests__test

import (
	"fmt"
	"github.com/fanjindong/go-requests"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
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

func TestRequestsFiles(t *testing.T) {
	file, err := requests.FileFromPath("./a.text")
	assert.NoError(t, err)
	resp, err := requests.Post(
		"https://email.neoclub.cn/send_email",
		requests.Files{"subject": "go-requests", "text": "nice", "toAddr": "jd.fan@neoclub.cn", "file1": file},
	)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var r struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	err = resp.Json(&r)
	assert.NoError(t, err)
	assert.Equal(t, 0, r.Code)
}

func TestRequests(t *testing.T) {
	resp, err := requests.Get("http://dpp.boluome.com/ping", requests.Params{"name": "fjd"})
	if err != nil {
		fmt.Printf("Get err: %v", err)
	}
	fmt.Println(resp.Request.URL)
	fmt.Println(resp.Header)
	fmt.Println(resp.Text)
	var rMap map[string]interface{}
	var rStruct struct {
		Code int `json:"code"`
	}

	err = resp.Json(&rMap)
	if err != nil {
		fmt.Printf("resp.Json to map err: %v \n", err)
	} else {
		fmt.Printf("resp.Json to map: %v \n", rMap)
	}

	err = resp.Json(&rStruct)
	if err != nil {
		fmt.Printf("resp.Json to struct err: %v \n", err)
	} else {
		fmt.Printf("resp.Json to struct: %+v \n", rStruct)
	}

}
