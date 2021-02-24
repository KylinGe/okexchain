package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBloomKey(t *testing.T) {
	expectedHeight, expectedBloomKey := int64(1), []byte{0, 0, 0, 0, 0, 0, 0, 1}
	require.True(t, bytes.Equal(BloomKey(expectedHeight), expectedBloomKey))
}
