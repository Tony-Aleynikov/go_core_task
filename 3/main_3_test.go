package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	tests := []struct {
		name     string
		data     map[string]int
		key      string
		value    int
		wantData map[string]int
	}{
		{
			name:     "add to nil data",
			data:     nil,
			key:      "a",
			value:    1,
			wantData: map[string]int{"a": 1},
		},
		{
			name:     "add to empty data",
			data:     map[string]int{},
			key:      "a",
			value:    1,
			wantData: map[string]int{"a": 1},
		},
		{
			name:     "add to existing data",
			data:     map[string]int{"a": 1},
			key:      "b",
			value:    2,
			wantData: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "overwrite key",
			data:     map[string]int{"a": 1, "b": 2},
			key:      "a",
			value:    10,
			wantData: map[string]int{"a": 10, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringIntMap{data: tt.data}
			s.Add(tt.key, tt.value)
			assert.Equal(t, tt.wantData, s.data)
		})
	}
}

func Test_Remove(t *testing.T) {
	tests := []struct {
		name     string
		data     map[string]int
		key      string
		wantData map[string]int
	}{
		{
			name:     "remove existing key",
			data:     map[string]int{"a": 1, "b": 2, "c": 3},
			key:      "b",
			wantData: map[string]int{"a": 1, "c": 3},
		},
		{
			name:     "remove non-existing key",
			data:     map[string]int{"a": 1, "b": 2},
			key:      "c",
			wantData: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "remove from nil data",
			data:     nil,
			key:      "a",
			wantData: nil,
		},
		{
			name:     "remove from empty data",
			data:     map[string]int{},
			key:      "a",
			wantData: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringIntMap{data: tt.data}
			s.Remove(tt.key)
			assert.Equal(t, tt.wantData, s.data)
		})
	}
}

func Test_Copy(t *testing.T) {
	tests := []struct {
		name     string
		data     map[string]int
		wantData map[string]int
	}{
		{
			name:     "copy non-empty data",
			data:     map[string]int{"a": 1, "b": 2, "c": 3},
			wantData: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name:     "copy nil data",
			data:     nil,
			wantData: map[string]int{},
		},
		{
			name:     "copy empty data",
			data:     map[string]int{},
			wantData: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringIntMap{data: tt.data}
			copied := s.Copy()
			assert.Equal(t, tt.wantData, copied.data)
		})
	}
}

func Test_Copy_TrueCopy(t *testing.T) {
	original := StringIntMap{data: map[string]int{"a": 1, "b": 2}}
	copied := original.Copy()

	copied.Add("c", 3)
	copied.Remove("a")

	assert.Equal(t, map[string]int{"a": 1, "b": 2}, original.data)
	assert.Equal(t, map[string]int{"b": 2, "c": 3}, copied.data)
}

func Test_Exists(t *testing.T) {
	tests := []struct {
		name string
		data map[string]int
		key  string
		want bool
	}{
		{
			name: "existing key",
			data: map[string]int{"a": 1, "b": 2},
			key:  "a",
			want: true,
		},
		{
			name: "non-existing key",
			data: map[string]int{"a": 1, "b": 2},
			key:  "c",
			want: false,
		},
		{
			name: "nil data",
			data: nil,
			key:  "a",
			want: false,
		},
		{
			name: "empty data",
			data: map[string]int{},
			key:  "a",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringIntMap{data: tt.data}
			assert.Equal(t, tt.want, s.Exists(tt.key))
		})
	}
}

func Test_Get(t *testing.T) {
	tests := []struct {
		name      string
		data      map[string]int
		key       string
		wantValue int
		wantOk    bool
	}{
		{
			name:      "existing key",
			data:      map[string]int{"a": 1, "b": 2},
			key:       "a",
			wantValue: 1,
			wantOk:    true,
		},
		{
			name:      "non-existing key",
			data:      map[string]int{"a": 1, "b": 2},
			key:       "c",
			wantValue: 0,
			wantOk:    false,
		},
		{
			name:      "nil data",
			data:      nil,
			key:       "a",
			wantValue: 0,
			wantOk:    false,
		},
		{
			name:      "empty data",
			data:      map[string]int{},
			key:       "a",
			wantValue: 0,
			wantOk:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringIntMap{data: tt.data}
			val, ok := s.Get(tt.key)
			assert.Equal(t, tt.wantValue, val)
			assert.Equal(t, tt.wantOk, ok)
		})
	}
}
