package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes今天要放假过年了！"
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println("===============")
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
