package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	t.Run("channel is unbuffered", func(t *testing.T) {
		ch := generate(t.Context())
		assert.Equal(t, 0, cap(ch))
	})

	t.Run("generates values", func(t *testing.T) {
		ch := generate(t.Context())

		got := make(map[int]struct{})

		for range 10 {
			val, ok := <-ch
			require.True(t, ok)
			got[val] = struct{}{}
		}

		assert.Greaterf(t, len(got), 1, "expected at least 2 different values, got %d", len(got))
	})

	t.Run("stops on context done", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		ch := generate(ctx)
		_, ok := <-ch
		assert.True(t, ok)

		cancel()
		_, ok = <-ch
		assert.False(t, ok)
	})
}
