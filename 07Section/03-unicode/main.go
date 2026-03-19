package main

import (
	"fmt"
)

func main() {
	str := []rune{'早', '安'}
	fmt.Println(string(str))

	// length of str in golang doesn't make any sence. because in go strings are encoded in unicode.
	// size of single char goes upto 4 bytes depending on language.
	fmt.Println("length of str(bytes): ", len(str))
	for _, chr := range str {
		fmt.Println(string(chr))
	}
}
