package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}
func length(s string) {
	fmt.Printf("length of %s is %d,%d\n", s, len(s), utf8.RuneCountInString(s))
}
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func main() {
	name := "Hello World"
	printBytes(name)
	printChars("Hello World")

	//Constructing string from slice of bytes
	//byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str := string(byteSlice)
	fmt.Println(str)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	//Rune
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	//String is immutable so s[0] ='a' will lead to error but when runes of string are passed, u can change bcoz runes are slices
	h := "hello"
	//fmt.Println(mutate(h))
	fmt.Println(mutate([]rune(h)))
}
