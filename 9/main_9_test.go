package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pipeline(t *testing.T) {
	tests := []struct {
		name   string
		values []uint8
		want   []float64
	}{
		{
			name:   "cubes values in order",
			values: []uint8{2, 3, 5},
			want:   []float64{8, 27, 125},
		},
		{
			name:   "handles the largest uint8 value",
			values: []uint8{255},
			want:   []float64{16581375},
		},
		{
			name: "handles empty input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(chan uint8, len(tt.values))
			for _, value := range tt.values {
				input <- value
			}
			close(input)

			output := make(chan float64)
			go pipeline(input, output)

			var got []float64
			for value := range output {
				got = append(got, value)
			}

			assert.Equal(t, tt.want, got)

			_, ok := <-output
			assert.Falsef(t, ok, "output channel should be closed")
		})
	}
}
