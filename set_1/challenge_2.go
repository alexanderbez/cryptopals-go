// Set 1: Challenge 2
//
// Fixed XOR:
//
// Write a function that takes two equal-length buffers and produces their XOR
// combination.
//
// https://cryptopals.com/sets/1/challenges/2
package set_1

import "fmt"

func hexCharToByte(c byte) byte {
	switch {
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10

	case c >= 'A' && c <= 'F':
		return c - 'A' + 10

	case c >= '0' && c <= '9':
		return c - '0'

	default:
		panic(fmt.Sprintf("invalid hex char: %d", c))
	}
}

// FixedXOR accepts two equal-length buffers and produces their XOR combination.
// Each input string is hex decoded and each nibble is XOR'd. An error is
// returned if the string buffers are not equal in length.
func FixedXOR(in1, in2 string) (string, error) {
	if len(in1) != len(in2) {
		return "", fmt.Errorf("incompatible inputs %d and %d", len(in1), len(in2))
	}

	buf1 := make([]byte, len(in1)*2)
	buf2 := make([]byte, len(in1)*2)

	// Convert input string into hex bytes where each nibble is decoded into a
	// single byte.
	var i int
	for i < len(in1) {
		buf1[i] = hexCharToByte(in1[i])
		buf2[i] = hexCharToByte(in2[i])
		i++
	}

	var (
		j      int
		result string
	)

	// XOR each byte adding the result to a string as a hex nibble
	for j < len(in1) {
		x := buf1[j] ^ buf2[j]
		y := buf1[j+1] ^ buf2[j+1]

		result += fmt.Sprintf("%x", (x<<4)|y) // TOOD: replace fmt.Sprinf with byteToHexChar
		j += 2
	}

	return result, nil
}
