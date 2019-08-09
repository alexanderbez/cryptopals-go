// Set 1: Challenge 1
//
// Convert hex to base64
//
// https://cryptopals.com/sets/1/challenges/1
package set_1

var (
	base64EncodingStr      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	padChar           rune = '='  // standard base64 padding character
	hexMask           uint = 0x3F // mask to get 6 bit chunk
)

// HexToBase64 converts hex (base16) encoded bytes to base64 encoded bytes. Each
// base64 digit represents 6-bits from the hex encoded byte padding as necessary.
func HexToBase64(input []byte) []byte {
	var (
		result  []byte // TODO: Allocate fixed size slice for encoding
		scanIdx int
	)

	// loop over and construct three octet chunks
	chunks := (len(input) / 3) * 3
	for scanIdx < chunks {
		chunk := uint(input[scanIdx+0])<<16 | uint(input[scanIdx+1])<<8 | uint(input[scanIdx+2])

		result = append(result, base64EncodingStr[chunk>>18&hexMask])
		result = append(result, base64EncodingStr[chunk>>12&hexMask])
		result = append(result, base64EncodingStr[chunk>>6&hexMask])
		result = append(result, base64EncodingStr[chunk&hexMask])

		// scan to next three octets
		scanIdx += 3
	}

	remaining := len(input) - scanIdx
	if remaining == 0 {
		return result
	}

	// Construct remaining chunk which should only have two or one significant
	// octets. Regardless, the remaining chunk can be constructed as three bytes.
	// Depending on the remainin number of input octets, a padding character will
	// be added.
	chunk := uint(input[scanIdx]) << 16

	if remaining == 2 {
		chunk |= uint(input[scanIdx:][1]) << 8
	}

	result = append(result, base64EncodingStr[chunk>>18&hexMask])
	result = append(result, base64EncodingStr[chunk>>12&hexMask])

	if remaining == 2 {
		result = append(result, base64EncodingStr[chunk>>6&hexMask], byte(padChar))
	} else {
		result = append(result, byte(padChar), byte(padChar))
	}

	return result
}
