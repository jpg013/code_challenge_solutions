package main

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
*/

import "fmt"

var symbolValues map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func getValue(s string) int {
	val, ok := symbolValues[s]

	if !ok {
		return 0
	}

	return val
}

func getCharAtIndex(s string, idx int) string {
	return s[idx : idx+1]
}

func getSubtraction(pre string, post string) (i int) {
	fmt.Println(pre, post)
	if pre == "I" && post == "V" {
		return 4
	}
	if pre == "I" && post == "X" {
		return 9
	}
	if pre == "X" && post == "L" {
		return 40
	}
	if pre == "X" && post == "C" {
		return 90
	}
	if pre == "C" && post == "D" {
		return 400
	}
	if pre == "C" && post == "M" {
		return 900
	}
	return i
}

func romanToInt(s string) int {
	var fn func(num int, idx int) int

	fn = func(num int, idx int) int {
		if idx >= len(s) {
			return num
		}
		if idx == 0 {
			num = getValue(getCharAtIndex(s, idx))
		} else {
			ch := getCharAtIndex(s, idx)
			prevCh := getCharAtIndex(s, idx-1)

			sub := getSubtraction(prevCh, ch)

			if sub > 0 {
				num = num - getValue(prevCh) + sub
			} else {
				num = num + getValue(ch)
			}
		}

		return fn(num, idx+1)
	}

	return fn(0, 0)
}

func main() {
	fmt.Println(romanToInt("IV"))
}
