package filler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFaulty(t *testing.T) {
	filler := NewFaulty()

	completed, err := filler.Fill(nil)
	require.Error(t, err)
	require.False(t, completed)
}
