package benchmark

import (
	"encoding/json"
	"testing"
)

func BenchmarkJson2Struct(b *testing.B) {
	var Str = `{"name": "fjd", "sex": 1, "phone": "10086"}`
	type People struct {
		Name  string `json:"name"`
		Sex   int    `json:"sex"`
		Phone string `json:"phone"`
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := &People{}
		// b.Log(Str)
		json.Unmarshal([]byte(Str), p)
		// b.Log(p)
	}
}
