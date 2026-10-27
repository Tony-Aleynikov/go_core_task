package main

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getType(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  string
	}{
		{name: "decimal int", value: 1, want: "int"},
		{name: "octal int", value: 02, want: "int"},
		{name: "hexadecimal int", value: 0x3, want: "int"},
		{name: "float64", value: 4.0, want: "float64"},
		{name: "float32", value: float32(4.0), want: "float32"},
		{name: "string", value: "5", want: "string"},
		{name: "bool true", value: true, want: "bool"},
		{name: "bool false", value: false, want: "bool"},
		{name: "complex64", value: complex64(6i), want: "complex64"},
		{name: "complex128", value: 6i, want: "complex128"},
		{name: "rune", value: 'a', want: "int32"},
		{name: "byte", value: byte(1), want: "uint8"},
		{name: "slice of runes", value: []rune{'g', 'o'}, want: "[]int32"},
		{name: "slice of bytes", value: []byte{1, 2, 3}, want: "[]uint8"},
		{name: "nil", value: nil, want: "<nil>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getType(tt.value))
		})
	}
}

func Test_merge(t *testing.T) {
	tests := []struct {
		name   string
		values []any
		want   string
	}{
		{
			name:   "example variables from main",
			values: []any{1, 02, 0x3, 4.0, "5", true, complex64(6i)},
			want:   "12345true(0+6i)",
		},
		{
			name:   "empty arguments",
			values: []any{},
			want:   "",
		},
		{
			name:   "empty string",
			values: []any{""},
			want:   "",
		},
		{
			name:   "single integer",
			values: []any{42},
			want:   "42",
		},
		{
			name:   "mixed strings and numbers",
			values: []any{"hello", " ", "world", 123},
			want:   "hello world123",
		},
		{
			name:   "booleans and floats",
			values: []any{false, 3.1415, true},
			want:   "false3.1415true",
		},
		{
			name:   "slices and runes",
			values: []any{[]rune{'a', 'b'}, []byte{1, 2}},
			want:   "[97 98][1 2]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, merge(tt.values...))
		})
	}
}

func Test_hash(t *testing.T) {
	tests := []struct {
		name string
		r    []rune
		want [32]byte
	}{
		{
			name: "empty string",
			r:    []rune(""),
			want: sha256.Sum256([]byte("go-2024")),
		},
		{
			name: "single character",
			r:    []rune{'a'},
			want: sha256.Sum256([]byte("go-2024a")),
		},
		{
			name: "two characters",
			r:    []rune{'g', 'o'},
			want: sha256.Sum256([]byte("ggo-2024o")),
		},
		{
			name: "word golang",
			r:    []rune("golang"),
			want: sha256.Sum256([]byte("golgo-2024ang")),
		},
		{
			name: "main string merged",
			r:    []rune("12345true(0+6i)"),
			want: sha256.Sum256([]byte("12345trgo-2024ue(0+6i)")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hash(tt.r))
		})
	}
}
