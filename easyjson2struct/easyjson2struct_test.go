package easyjson2struct

import (
	"testing"
)

// easyjson is to lower, 要解析的json字符串里面的key必须首字母大写？？黑人问号
func TestEasyjson2Struct(t *testing.T) {
	p := &People{}
	p.UnmarshalJSON([]byte(`{"Name": "fjd", "sex": 1, "phone": "10086", "Id": 123}`))
	t.Log(p)
}
