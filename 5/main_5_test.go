package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersection(t *testing.T) {
	tests := []struct {
		name      string
		s1        []int
		s2        []int
		wantOk    bool
		wantSlice []int
	}{
		{
			name:      "both empty",
			s1:        []int{},
			s2:        []int{},
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "first empty",
			s1:        []int{},
			s2:        []int{1},
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "second empty",
			s1:        []int{1},
			s2:        []int{},
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "both nil",
			s1:        nil,
			s2:        nil,
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "first nil",
			s1:        nil,
			s2:        []int{1},
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "second nil",
			s1:        []int{1},
			s2:        nil,
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "both same",
			s1:        []int{1, 2},
			s2:        []int{1, 2},
			wantOk:    true,
			wantSlice: []int{1, 2},
		},
		{
			name:      "both same but different order",
			s1:        []int{1, 2},
			s2:        []int{2, 1},
			wantOk:    true,
			wantSlice: []int{2, 1},
		},
		{
			name:      "different",
			s1:        []int{1, 2},
			s2:        []int{3, 4},
			wantOk:    false,
			wantSlice: nil,
		},
		{
			name:      "one element intersection",
			s1:        []int{1, 2, 3},
			s2:        []int{3, 4, 5},
			wantOk:    true,
			wantSlice: []int{3},
		},
		{
			name:      "multiple elements intersection",
			s1:        []int{10, 20, 30, 40},
			s2:        []int{30, 50, 10},
			wantOk:    true,
			wantSlice: []int{30, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotSlice := intersection(tt.s1, tt.s2)
			assert.Equal(t, tt.wantOk, gotOk)
			assert.Equal(t, tt.wantSlice, gotSlice)
		})
	}
}
