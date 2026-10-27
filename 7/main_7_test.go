package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name    string
		merging func() []<-chan int
		want    []int
	}{
		{
			name:    "no channels",
			merging: func() []<-chan int { return []<-chan int{} },
			want:    nil,
		},
		{
			name: "opened channels",
			merging: func() []<-chan int {
				a := make(chan int)
				b := make(chan int)
				c := make(chan int)

				go func() { a <- 1; a <- 2; close(a) }()
				go func() { b <- 3; b <- 4; close(b) }()
				go func() { c <- 5; c <- 6; close(c) }()

				return []<-chan int{a, b, c}
			},
			want: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name: "opened and closed channels",
			merging: func() []<-chan int {
				a := make(chan int)
				b := make(chan int)
				c := make(chan int)

				go func() { a <- 1; a <- 2; close(a) }()
				go func() { close(b) }()
				go func() { c <- 5; c <- 6; close(c) }()

				return []<-chan int{a, b, c}
			},
			want: []int{6, 5, 2, 1},
		},
		{
			name: "with nil channels",
			merging: func() []<-chan int {
				a := make(chan int)
				b := make(chan int)
				var c <-chan int

				go func() { a <- 1; a <- 2; close(a) }()
				go func() { close(b) }()

				return []<-chan int{a, b, c}
			},
			want: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := merge(tt.merging()...)

			var got []int
			timeout := time.After(2 * time.Second) // даем 2 секунды

		loop:
			for {
				select {
				case v, ok := <-out:
					if !ok {
						break loop
					}
					got = append(got, v)
				case <-timeout:
					t.Fatal("тест завис: канал не закрылся за 2 секунды")
				}
			}

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
