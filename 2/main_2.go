package main

import (
	"fmt"
)

func main() {
	//originalSlice := make([]int, 10)
	//for i := range originalSlice {
	//	originalSlice[i] = rand.Int()
	//}
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(sliceExample(originalSlice))
	fmt.Println(addElements(originalSlice, 5))
	fmt.Println(copySlice(originalSlice))
	fmt.Println(removeElement(originalSlice, 3))
}

func sliceExample(slice []int) []int {
	var newS []int
	for _, v := range slice {
		if v%2 == 0 {
			newS = append(newS, v)
		}
	}
	return newS
}

func addElements(slice []int, v int) []int {
	s := make([]int, 0, len(slice)+1)
	s = append(s, slice...)
	s = append(s, v)
	return s
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		// вернуть копию, чтобы не мутировать исходный, если присланный индекс не валидный
		result := make([]int, len(slice))
		copy(result, slice)
		return result
	}

	s := make([]int, 0, len(slice)-1)
	s = append(s, slice[:i]...)
	s = append(s, slice[i+1:]...)
	return s
}
