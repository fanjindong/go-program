package codec

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Params struct{}

func (Params) Marshal(v interface{}) ([]byte, error) {
	fmt.Printf("Marshal: %+v\n", v)
	//return json.Marshal(v)
	return []byte("name=aaa"), nil
}

func (Params) Unmarshal(data []byte, v interface{}) error {
	// data is a=1&b=2
	fmt.Println("data: ", string(data))
	kvs := strings.Split(string(data), "&")
	m := make(map[string]interface{}, len(kvs))
	for _, kv := range kvs {
		item := strings.Split(kv, "=")
		m[item[0]] = item[1]
	}
	pData, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(pData, v)
}

func (Params) String() string {
	return "params"
}

func (Params) Name() string {
	return "params"
}
