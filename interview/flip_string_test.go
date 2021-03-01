package interview

import (
	"testing"
)

// 问题描述
//
// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
//
// 给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
func TestFlipString(t *testing.T) {
	s := "12"
	sr := []rune(s)
	sr[0], sr[1] = sr[1], sr[0]
	t.Log(s, string(sr))
	// f := func(s string) {
	// 	b := []rune(s)
	//
	// }
}
