package filler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFaulty(t *testing.T) {
	filler := NewFaulty[int8]()

	completed, err := filler.Fill(nil, nil)
	require.Error(t, err)
	require.False(t, completed)
}
