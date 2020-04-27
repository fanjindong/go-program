package demo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"testing"
)

func TestDemo(t *testing.T) {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}
	// Output: 1, 2, 3
}

func TestTwoSum(t *testing.T) {
	r := twoSum([]int{2, 7, 9, 12}, 9)
	t.Log(r)

	isPalindrome(1000030001)
	t.Log(romanToInt("IV"))
}
func longestCommonPrefix(strs []string) string {
	var result string
	var n int

	for n <= len(strs[0]) {
		s := strs[0][n]
		for i := 1; i < len(strs); i++ {
			if n >= len(strs[i]) || strs[i-1][n] != strs[i][n] {
				return result
			}
		}
		result += string(s)
		n++
	}
	return result
}
func romanToInt(s string) int {
	symbolValueMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	type RomanNumeral struct {
		Symbol string
		Value  int
	}

	rn := make([]*RomanNumeral, len(s))

	for i, str := range s {
		rn[i] = &RomanNumeral{
			Symbol: string(str),
			Value:  symbolValueMap[string(str)],
		}
	}

	result := rn[len(s)-1].Value
	for j := 0; j < len(s)-1; j++ {
		value := rn[j].Value
		if value < rn[j+1].Value {
			value *= -1
		}

		result += value
	}
	return result
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for index, num := range nums {
		result[0] = index
		fmt.Println(result)
		for otherIndex, otherNum := range nums[index+1:] {
			if num+otherNum == target {
				result[1] = otherIndex + index + 1
				fmt.Println(result)
				return result
			}
		}
	}
	return result
}

func isPalindrome(x int) bool {
	if x < 0 || x == 10 {
		return false
	}

	a := make([]int, 0)

	for x != 0 {
		a = append(a, x%10)
		x /= 10
	}

	k := len(a) - 1
	for i := 0; i <= (k / 2); i++ {
		fmt.Println(i, k-i, a[i], a[k-i], a)
		if a[i] != a[k-i] {
			return false
		}
	}
	return true
}

func Append(a []int) {
	a = append(a, 0)
}

func Set(a []int) {
	a[0] = 0
}

func TestSlice(t *testing.T) {
	a := []int{1, 2}
	Append(a)      // Change a[0].
	fmt.Println(a) // Output: [8 2]
	a = append(a, 3)
	fmt.Println(a)

	Set(a)
	fmt.Println(a)
}

func TestJson(t *testing.T) {
	str := "[]"
	type D struct {
		A string
	}
	var a []*D
	err := json.Unmarshal([]byte(str), &a)
	t.Log(a, err)
}

var consumerSeq uint64

const consumerTagLengthMax = 0xFF // see writeShortstr

func uniqueConsumerTag() string {
	return commandNameBasedUniqueConsumerTag(os.Args[0])
}

func commandNameBasedUniqueConsumerTag(commandName string) string {
	tagPrefix := "ctag-"
	tagInfix := commandName
	tagSuffix := "-" + strconv.FormatUint(atomic.AddUint64(&consumerSeq, 1), 10)

	if len(tagPrefix)+len(tagInfix)+len(tagSuffix) > consumerTagLengthMax {
		tagInfix = "streadway/amqp"
	}

	return tagPrefix + tagInfix + tagSuffix
}

func TestAmqp(t *testing.T) {
	t.Log(uniqueConsumerTag())
}
