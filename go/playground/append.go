package main

import "fmt"

// see https://stackoverflow.com/questions/35920534/why-does-append-modify-passed-slice
func GenerateSlices(fullArray []int) [][]int {
	var slices [][]int
    for i := range fullArray {
        fmt.Println("before", fullArray[:i], fullArray[i:], fullArray)
        newSlice := append(fullArray[:i], fullArray[i+1:]...)
        fmt.Println("after", fullArray[:i], fullArray[i:], fullArray)
		slices = append(slices, newSlice)
	}
	return slices
}

// func main() {
//     GenerateSlices([]int{1,2})
// }
