package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	slc := arr[:]
	slc = append(slc, 4)

	fmt.Println("Array - Len: ", len(arr), "Cap: ", cap(arr))
	fmt.Println("Slice - Len: ", len(slc), "Cap: ", cap(slc))
}
