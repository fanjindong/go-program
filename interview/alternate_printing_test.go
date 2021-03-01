package interview

import (
	"fmt"
	"testing"
)

// 问题描述
//
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//
// 0A1B2C3D4E5F6G7H8I9J10K11L12M13N14O15P16Q17R18S19T20U21V22W23X24Y25Z

func TestAlternatePrinting(t *testing.T) {
	numReadyChan := make(chan bool)
	strReadyChan := make(chan bool)
	mainChan := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			<-numReadyChan
			fmt.Print(i)
			strReadyChan <- true
		}
	}()
	go func() {
		for s := 'A'; s <= 'Z'; s++ {
			<-strReadyChan
			fmt.Print(string(s))
			numReadyChan <- true
		}
		close(mainChan)
	}()
	numReadyChan <- true
	<-mainChan
}
