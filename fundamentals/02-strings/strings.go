/*
	String: an array of bytes | every element is immutable
			an reflect.StringHeader struct | StringHeader{Data, Len} | (pointer, length)
*/

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s1 = "hello world"
	// The same as
	var s2 = [...]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	fmt.Printf("s1: %s - len: %d\n", s1, len(s1))
	fmt.Printf("s2: %s - len: %d\n", s2, len(s2))

	// String is not a slice, but supports slicing
	// Get elements from index 0 -> 4
	hello := s1[:5] // s[0:5]

	// Get elements from index 5 -> end
	world := s1[6:]

	// Directly slicing
	s3 := "hello world"[:5]
	s4 := "hello world"[6:]

	fmt.Printf("%s %s\n", hello, world)
	fmt.Printf("%s %s\n", s3, s4)

	// Get length of string
	fmt.Printf("len(s) - len: %d\n", len(s1))
	fmt.Printf("(*reflect.StringHeader)(unsafe.Pointer(&s)).Len - len: %d\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len)
}
