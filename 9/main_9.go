package main

import "fmt"

func pipeline(in <-chan uint8, out chan<- float64) {
	defer close(out)
	for value := range in {
		number := float64(value)
		out <- number * number * number
	}
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)
	go pipeline(in, out)

	go func() {
		defer close(in)
		for i := range 10 {
			in <- uint8(i)
		}
	}()

	for value := range out {
		fmt.Println(value)
	}
}
