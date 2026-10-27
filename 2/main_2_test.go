package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sliceExample(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "example from main",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "nil slice",
			s:    nil,
			want: nil,
		},
		{
			name: "empty slice",
			s:    []int{},
			want: nil,
		},
		{
			name: "only odd numbers",
			s:    []int{1, 3, 5, 7, 9},
			want: nil,
		},
		{
			name: "only even numbers",
			s:    []int{2, 4, 6, 8, 10},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "mixed negative and positive numbers with zero",
			s:    []int{-4, -3, -2, -1, 0, 1, 2, 3, 4},
			want: []int{-4, -2, 0, 2, 4},
		},
		{
			name: "single odd element",
			s:    []int{7},
			want: nil,
		},
		{
			name: "single even element",
			s:    []int{8},
			want: []int{8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sliceExample(tt.s))
		})
	}
}

func Test_addElement(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{
			name: "example from main",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:    5,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5},
		},
		{
			name: "nil slice",
			s:    nil,
			n:    42,
			want: []int{42},
		},
		{
			name: "empty slice",
			s:    []int{},
			n:    0,
			want: []int{0},
		},
		{
			name: "add negative number",
			s:    []int{1, 2, 3},
			n:    -5,
			want: []int{1, 2, 3, -5},
		},
		{
			name: "single element slice",
			s:    []int{10},
			n:    20,
			want: []int{10, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addElements(tt.s, tt.n))
		})
	}
}

func Test_copySlice(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "example from main",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "nil slice",
			s:    nil,
			want: []int{},
		},
		{
			name: "empty slice",
			s:    []int{},
			want: []int{},
		},
		{
			name: "single element",
			s:    []int{42},
			want: []int{42},
		},
		{
			name: "slice with negative numbers and zero",
			s:    []int{-1, 0, 1},
			want: []int{-1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, copySlice(tt.s))
		})
	}
}

func Test_copySlice_OriginalNotModified(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := copySlice(original)

	copied[0] = 999
	assert.Equal(t, []int{1, 2, 3, 4, 5}, original)
	assert.Equal(t, 999, copied[0])
}

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{
			name: "example from main",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			n:    3,
			want: []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "remove first element",
			s:    []int{1, 2, 3, 4},
			n:    0,
			want: []int{2, 3, 4},
		},
		{
			name: "remove last element",
			s:    []int{1, 2, 3, 4},
			n:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "remove middle element",
			s:    []int{10, 20, 30, 40},
			n:    2,
			want: []int{10, 20, 40},
		},
		{
			name: "single element slice removing index 0",
			s:    []int{100},
			n:    0,
			want: []int{},
		},
		{
			name: "negative index out of bounds",
			s:    []int{1, 2, 3},
			n:    -1,
			want: []int{1, 2, 3},
		},
		{
			name: "index equal to slice length out of bounds",
			s:    []int{1, 2, 3},
			n:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "index greater than slice length out of bounds",
			s:    []int{1, 2, 3},
			n:    10,
			want: []int{1, 2, 3},
		},
		{
			name: "nil slice with index 0",
			s:    nil,
			n:    0,
			want: []int{},
		},
		{
			name: "empty slice with index 0",
			s:    []int{},
			n:    0,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeElement(tt.s, tt.n))
		})
	}
}

func Test_removeElement_OriginalNotModified(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	removed := removeElement(original, 2)

	removed[0] = 999
	assert.Equal(t, []int{1, 2, 3, 4, 5}, original)
	assert.Equal(t, 999, removed[0])
}
