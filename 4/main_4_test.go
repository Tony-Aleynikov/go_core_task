package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diff(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "both empty",
			s1:   []string{},
			s2:   []string{},
			want: nil,
		},
		{
			name: "first empty",
			s1:   []string{},
			s2:   []string{"a"},
			want: nil,
		},
		{
			name: "second empty",
			s1:   []string{"a"},
			s2:   []string{},
			want: []string{"a"},
		},
		{
			name: "both nil",
			s1:   nil,
			s2:   nil,
			want: nil,
		},
		{
			name: "first nil",
			s1:   nil,
			s2:   []string{"a"},
			want: nil,
		},
		{
			name: "second nil",
			s1:   []string{"a"},
			s2:   nil,
			want: []string{"a"},
		},
		{
			name: "both same",
			s1:   []string{"a", "b"},
			s2:   []string{"a", "b"},
			want: nil,
		},
		{
			name: "both same but different order",
			s1:   []string{"a", "b"},
			s2:   []string{"b", "a"},
			want: nil,
		},
		{
			name: "different",
			s1:   []string{"a", "b"},
			s2:   []string{"c", "d"},
			want: []string{"a", "b"},
		},
		{
			name: "one more element in first",
			s1:   []string{"a", "b", "c"},
			s2:   []string{"a", "b"},
			want: []string{"c"},
		},
		{
			name: "one more element in second",
			s1:   []string{"a", "b"},
			s2:   []string{"a", "b", "c"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, diff(tt.s1, tt.s2))
		})
	}
}
