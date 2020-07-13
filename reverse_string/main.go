package main

import "fmt"

/*
Reverse a string - example justin -> nitsuj

138


Interpreted string literals are character sequences between double quotes "" using the (possibly multi-byte) UTF-8 encoding of individual characters. In UTF-8, ASCII characters are single-byte corresponding to the first 128 Unicode characters. Strings behave like slices of bytes. A rune is an integer value identifying a Unicode code point.
https://golangbyexample.com/swap-characters-string-golang/
*/

func reverseString(s string) string {
	r := []rune(s)
	l := len(r)
	m := l / 2

	for x := 0; x < m; x++ {
		e := l - 1 - x
		r[x], r[e] = r[e], r[x]
	}

	return string(r)
}

func main() {
	fmt.Println(reverseString("justin"))
}
