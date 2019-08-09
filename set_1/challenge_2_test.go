package set_1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFixedXOR(t *testing.T) {
	testCases := []struct {
		input1, input2 string
		expectedErr    bool
		expectedResult string
	}{
		{
			"1c0111001f010100061a024b535350",
			"686974207468652062756c6c277320657965",
			true,
			"",
		},
		{
			"1c0111001f010100061a024b53535009181c",
			"686974207468652062756c6c277320657965",
			false,
			"746865206b696420646f6e277420706c6179",
		},
	}

	for i, tc := range testCases {
		res, err := FixedXOR(tc.input1, tc.input2)
		require.Equal(t, err != nil, tc.expectedErr)
		require.Equal(t, tc.expectedResult, res, "test case #%d failed", i)
	}
}
