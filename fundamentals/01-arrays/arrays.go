/*
	Arrays, strings, slices | are the data types related to each other,
							| have the same memory structure

		size: 4
		value: 2 | 5 | 4 | 0
		index: 0   1   2   3
*/

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"reflect"

	"github.com/hoangtk0100/go-advanced/utils"
)

func main() {
	utils.PrintTitle("defineMultipleWays")
	defineMultipleWays()

	utils.PrintTitle("defineMultipleTypes")
	defineMultipleTypes()

	utils.PrintTitle("getType")
	getType()

	utils.PrintTitle("arrayPointer")
	arrayPointer()

	utils.PrintTitle("getSize")
	getSize()

	utils.PrintTitle("loops")
	loops()

	utils.PrintTitle("emptyArray")
	emptyArray()
}

func defineMultipleWays() {
	// Ways to define an array

	// Define array with 3 elements, all initial value is 0
	var a [3]int
	fmt.Printf("a: %v - size: %d\n", a, len(a))

	// Define with initial values: 1 2 3, so size is the number of arrays's elements
	var b = [...]int{1, 2, 3}
	fmt.Printf("b: %v - size: %d\n", b, len(b))

	// Array values: 0 0 4 0 3
	var c = [...]int{2: 4, 4: 3}
	fmt.Printf("c: %v - size: %d\n", c, len(c))

	// Array values: 1 2 0 0 5 7
	var d = [...]int{1, 2, 4: 5, 7}
	fmt.Printf("d: %v - size: %d\n", d, len(d))

	// Array values: 2 1 6 5
	e := [...]int{2, 1, 6, 5}
	fmt.Printf("e: %v - size: %d\n", e, len(e))

	// Array values: 0 0
	f := make([]int, 2)
	fmt.Printf("f: %v - size: %d\n", f, len(f))
}

func defineMultipleTypes() {
	// String array
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"Hello!", "World"}
	var s3 = [...]string{1: "Hello", 0: "World"}
	fmt.Printf("\n[%v] - [%v] - [%v]\n", s1, s2, s3)

	// Struct array
	var line1 = [2]image.Point{}
	var line2 = [...]image.Point{image.Point{X: 1, Y: 2}, image.Point{X: 3, Y: 4}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Printf("\n[%v] - [%v] - [%v]\n", line1, line2, line3)

	// Image decoder array
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	fmt.Printf("\n[%v] - [%v]\n", decoder1, decoder2)

	// Interface array
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{[]int{1, 2, 3}, "Hello"}
	fmt.Printf("\n[%v] - [%v]\n", unknown1, unknown2)

	// Channel array
	var chanList1 [2]chan int
	var chanList2 = [...]chan int{make(chan int), make(chan int)}
	fmt.Printf("\n[%v] - [%v]\n", chanList1, chanList2)
}

func getType() {
	a := [...]int{1, 2, 3}
	fmt.Printf("%v - Type: %T\n", a, a)
	fmt.Printf("%v - Type: %v\n", a, reflect.TypeOf(a))
	fmt.Printf("%v - Type: %v\n", a, reflect.ValueOf(a).Kind())
}

func arrayPointer() {
	/*
	   When the array varies or transmitted, the entire Array will be copied.
	   If the size of the Array is large, the Array assignment will suffer a large cost.
	   To avoid overhead (costs) in copying array, you can transmit the Array's cursor.
	*/
	var a = [...]int{1, 2, 3}

	// Declare a pointer to an array
	var b = &a

	// Print the values of the array using the pointer is the same as the original array
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for index, value := range b {
		b[index] += 1
		fmt.Println(index, value)
	}

	// The values of array a are changed, but the pointer is still pointing to the original array
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func getSize() {
	// Len: returns the number of elements in the array
	// Cap: returns the capacity of the array | The capacity is the number of elements the array can hold without reallocating
	// Because the array is fixed size, so len = cap
	var a = [...]int{1, 2, 3}
	fmt.Printf("a: %v - size: %d - cap: %d\n", a, len(a), cap(a))
}

func loops() {
	var a = [...]int{1, 2, 3}

	// For loop in range is the best way to iterate through an array
	// Because of not having to know the size of the array, ensuring the access is always within the bounds of the array
	for index := range a {
		fmt.Printf("a[%d]: %d\n", index, a[index])
	}

	for index, value := range a {
		fmt.Printf("a[%d]: %d\n", index, value)
	}

	for index := 0; index < len(a); index++ {
		fmt.Printf("a[%d]: %d\n", index, a[index])
	}
}

func emptyArray() {
	var a [0]int
	var b = [0]int{}
	var c = [...]int{}
	fmt.Printf("%v - %v - %v\n", a, b, c)
}
