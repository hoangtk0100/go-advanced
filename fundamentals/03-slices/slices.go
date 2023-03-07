/*
	Slice is an reflect.SliceHeader struct | SliceHeader{Data, Len, Cap} | (pointer, length, capacity)
			 fat pointer (https://nullprogram.com/blog/2019/06/30)

	https://go.dev/blog/slices
*/

package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/hoangtk0100/go-advanced/utils"
)

func main() {
	utils.PrintTitle("defineMultipleWays")
	defineMultipleWays()

	utils.PrintTitle("loops")
	loops()

	utils.PrintTitle("addElement")
	addElements()

	utils.PrintTitle("removeElements")
	removeElements()

	utils.PrintTitle("manageMemoryInSlice")
	manageMemoryInSlice()
}

func defineMultipleWays() {
	// nil slice - Values: [] - len: 0 - cap: 0
	var a []int
	printSlice("a", a)
	if a == nil {
		fmt.Println("a - nil!")
	}

	// empty slice, difference from nil slice
	// Values: [] - len: 0 - cap: 0
	var b = []int{}
	printSlice("b", b)

	// Values: [2 3 5 7 11] - len: 5 - cap: 5
	var c = []int{2, 3, 5, 7, 11}
	printSlice("c", c)

	// Values: [3 5] - len: 2 - cap: 4
	// Create a new slice this way, the slice will not copy the data of the original slice c to d
	// it creates a new slice value that points to the original array.
	var d = c[1:3]
	printSlice("d", d)

	// Values: [1 2 3] - len: 3 - cap: 3
	var e = []int{1, 2, 3}
	printSlice("e", e)

	// Values: [1 2] - len: 2 - cap: 3
	var f = e[:2]
	printSlice("f", f)

	// Values: [1 2] - len: 2 - cap: 3
	var g = e[0:2:cap(e)]
	printSlice("g", g)

	// Values: [] - len: 0 - cap: 3
	var h []int = e[:0]
	printSlice("h", h)

	// Values: [] - len: 0 - cap: 0
	// make([]T, length, capacity) - (length, capacity) : optional
	var i = make([]int, 0)
	printSlice("i", i)

	// Values: [0 0 0] - len: 3 - cap: 3
	var j = make([]int, 3)
	printSlice("j", j)

	// Values: [0 0] - len: 2 - cap: 3
	var k = make([]int, 2, 3)

	// Values: [0 0 4 5 6 7 8] - len: 7 - cap: 8
	k = append(k, 4, 5, 6, 7, 8)
	printSlice("k", k)

	// Values: [] - len: 0 - cap: 3
	var l = make([]int, 0, 3)
	printSlice("l", l)
}

func loops() {
	var a = []int{1, 2, 3}

	for index := range a {
		fmt.Printf("a[%d] = %d\n", index, a[index])
	}

	for index, value := range a {
		fmt.Printf("a[%d] = %d\n", index, value)
	}

	for index := 0; index < len(a); index++ {
		fmt.Printf("a[%d] = %d\n", index, a[index])
	}
}

func addElements() {
	/*
		==== Add element to the end of the slice ====
	*/
	// append(slice, value) - append the value to the end of the slice

	utils.PrintSubTitle("Add element to the end of the slice")
	var a []int
	a = append(a, 1)
	a = append(a, 2, 3, 4)
	a = append(a, []int{5, 6, 7, 8}...)
	printSlice("a", a)

	/*
		In case the original slice does not have enough space when adding an element, the append function will perform a memory allocation of the size:
			- If old size (cap) < 1024: allocate double (x2) old memory.
			- If old size >= 1024: allocate 1.25x old memory.

		After that, old data will be copied to the new memory area and the pointer will point to the new memory area.
		Details in the following link: https://go.dev/src/runtime/slice.go#L66
	*/
	utils.PrintSubTitle("Scale the capacity of the slice")
	sl := make([]int, 1)
	printSlice("sl", sl)
	for index := 0; index < 5; index++ {
		sl = myAppend(sl, index)
	}

	/*
		==== Add element to the beginning of the slice ====
	*/
	/*
		Adding an element to the beginning of the slice causes a reallocation of memory
		and causes the existing elements in the slice to be copied again.
		Therefore, the performance of adding the element to the beginning of the slice will be lower than adding the element to the end of the slice.
	*/
	utils.PrintSubTitle("Add elements to the beginning of the slice")
	var b = []int{1, 2, 3, 4}
	b = append([]int{-2, -1}, b...)
	printSlice("b", b)

	/*
		==== Add an element at position i of the slice ====
	*/
	// Using append function | initialize temporary slice to concatenate with the original slice
	utils.PrintSubTitle("Add an element at position i of the slice using append function")
	i := 3
	b = append(
		b[:i],
		// Create a temporary slice to concatenate with b[:i]
		append([]int{0}, b[i:]...,
		)...,
	)
	printSlice("b", b)

	// Using copy and append function | avoid initializing temporary slices
	// How copy function works (https://www.educative.io/answers/how-to-use-the-copy-function-in-go)
	utils.PrintSubTitle("Add an element at position i of the slice using copy function")
	var c = []int{4, 5, 6, 1, 2, 3}
	printSlice("c", c)

	// Expand the space of the slice c with 1 element
	c = append(c, 0)
	printSlice("c", c)
	printSlice("c[i]", c[i:])
	printSlice("c[i+1]", c[i+1:])
	numberCopied := copy(c[i+1:], c[i:])
	c[i] = 0
	printSlice("c", c)
	fmt.Println("Number of elements copied: ", numberCopied)

	/*
		==== Add elements at position i of the slice ====
	*/
	utils.PrintSubTitle("Add elements at position i of the slice using append function")
	var d = []int{7, 8, 9, 10}
	printSlice("d", d)

	d = append(d[:i], append([]int{1, 2, 3}, d[i:]...)...)
	printSlice("d", d)

	utils.PrintSubTitle("Add elements at position i of the slice using copy function")
	var e = []int{55, 66, 77, 88, 99}
	var x = []int{11, 22}
	printSlice("e", e)
	printSlice("x", x)

	// Expand the space of the slice e with the slice x
	e = append(e, x...)
	numberCopied = copy(e[i+len(x):], e[i:])
	copy(e[i:], x)
	fmt.Println("Number of elements copied: ", numberCopied)
	printSlice("e", e)
}

func removeElements() {
	/*
		==== Remove elements at the end of the slice is fastest way ====
	*/
	utils.PrintSubTitle("Remove elements at the end of the slice")
	a := []int{1, 2, 3, 4}
	printSlice("a", a)

	// Remove 1 element at the end
	a = a[:len(a)-1]
	printSlice("a", a)

	// Remove N elements at the end
	N := 2
	a = a[:len(a)-N]
	printSlice("a", a)

	/*
		==== Remove elements at the beginning of the slice ====
	*/
	// Removing the element at the beginning is actually moving the data pointer back
	utils.PrintSubTitle("Remove elements at the beginning of the slice")
	b := []int{1, 2, 3, 4}
	printSlice("b", b)

	// Remove first element
	b = b[1:]
	printSlice("b", b)

	// Remove first N elements
	b = b[N:]
	printSlice("b", b)

	/*
		==== Remove elements in the middle ====
	*/
	// Removing the elements in the middle, you need to move the elements in the back to the front
	utils.PrintSubTitle("Remove elements at the position i")
	i := 2
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice("c", c)

	// Remove 1 element
	c = append(c[:i], c[i+1:]...)
	printSlice("c", c)

	// Remove N elements
	c = append(c[:i], c[i+N:]...)
	printSlice("c", c)

	// Remove 1 element using copy
	c = c[:i+copy(c[i:], c[i+1:])]
	printSlice("c", c)

	// Remove N elements using copy
	c = c[:i+copy(c[i:], c[i+N:])]
	printSlice("c", c)
}

func manageMemoryInSlice() {
	/*
		Limit the need to reallocate memory

		*** append(): not reallocate memory until the maximum capacity cap has been reached.
		=> Keeping append not reach the slice's capacity is to reduce the number of allocations
		and reduce the size of the allocation at all times.
	*/
	utils.PrintSubTitle("Limit the need to reallocate memory")
	s := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	trimmedSpaceSlice := filter(s, func(x byte) bool {
		if x == ' ' {
			return true
		}

		return false
	})

	printSlice("trimmedSpaceSlice", trimmedSpaceSlice)

	/*
		Avoid causing memory leaks

		Remove the memory allocated to the slice/removed elements in slice when the slice/removed elements is no longer needed.
		So that the garbage collector can free the memory allocated to the slice/removed elements.
	*/
	utils.PrintSubTitle("Avoid causing memory leaks")
	fmt.Println("+++ Assign the last element to nil before removing it +++")
	x, y := 78, 88
	a := []*int{&x, &y}

	a[len(a)-1] = nil
	a = a[:len(a)-1]
	printSlice("a", a)

	fmt.Println("+++ Assign the result to a new variable +++")
	phoneNumbers := findPhoneNumber("phone_numbers.txt")
	fmt.Println("Phone numbers - ", string(phoneNumbers))
}

func findPhoneNumber(filename string) []byte {
	/*
		Returns an array of bytes pointing to the entire file.
		Because slice refers to the entire original array,
		automatic garbage collection cannot free the space below the array in the meantime.
		A small result request, but has to store all the data for a long time.

		=> Copy the needed data to a new slice
	*/
	b, _ := os.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

func filter(s []byte, fn func(x byte) bool) []byte {
	// Define slice b with s's capacity, reduce the number of allocations
	b := s[:0]
	for _, val := range s {
		if !fn(val) {
			b = append(b, val)
		}
	}

	return b
}

func myAppend(sl []int, val int) []int {
	sl = append(sl, val)
	printSlice("sl", sl)
	return sl
}

func printSlice[T comparable](s string, sl []T) {
	fmt.Printf("%s - %v - len: %d - cap: %d\n", s, sl, len(sl), cap(sl))
}
