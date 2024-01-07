package mem

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEngine_Del(t *testing.T) {
	t.Parallel()

	var (
		engine = NewEngine()
		ctx    = context.Background()
		key    = "key"
		value  = "value"
	)

	engine.Set(ctx, key, value)

	result, ok := engine.Get(ctx, key)
	require.Equal(t, value, result)
	require.True(t, ok)

	engine.Del(ctx, key)

	result, ok = engine.Get(ctx, key)
	require.Empty(t, result)
	require.False(t, ok)
}

func TestEngine_Get(t *testing.T) {
	t.Parallel()

	var (
		engine = NewEngine()
		ctx    = context.Background()
		key    = "key"
		value  = "value"
	)

	engine.Set(ctx, key, value)

	result, ok := engine.Get(ctx, key)
	require.Equal(t, value, result)
	require.True(t, ok)
}

func TestEngine_Set(t *testing.T) {
	t.Parallel()

	var (
		engine = NewEngine()
		ctx    = context.Background()
		key    = "key"
		value  = "value"
	)

	result, ok := engine.Get(ctx, key)
	require.Empty(t, result)
	require.False(t, ok)

	engine.Set(ctx, key, value)

	result, ok = engine.Get(ctx, key)
	require.Equal(t, value, result)
	require.True(t, ok)
}
