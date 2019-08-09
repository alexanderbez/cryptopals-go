package set_1

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexToBase64(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			"",
			"",
		},
		{
			"0a0a0a",
			"CgoK",
		},
		{
			"2a4b6c8d",
			"KktsjQ==",
		},
		{
			"1e4b6c8d0f",
			"HktsjQ8=",
		},
		{
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}

	for i, tc := range testCases {
		inputBz, err := hex.DecodeString(tc.input)
		require.NoError(t, err, "unexpected error decoding input string")

		output := HexToBase64(inputBz)
		require.Equal(t, tc.expected, string(output), "test case #%d failed", i)
	}
}
