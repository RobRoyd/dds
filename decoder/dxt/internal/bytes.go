package internal

import "math"

// Extract some bits form a byte array, being able to go across byte boundaries. The slice is interpreted
// in ascending order. The length is interpreted as slicing the slice into fixed length items and the
// zero-based index refers to the item of index n.
//
// The maximum length allowed is 7. As 8 bit would be simply accessing the byte slice and more than 8 byte
// is not supported.
//
// It needs to be verified before that (index+1)*length <= len(bytes) either by convention or by testing.
//
// Example: bytes: [ 0b1010_1000, 0b0001_0011 ], i: 1, l: 5 => 0b01110
func Extract(bytes []byte, index, length byte) byte {
	return extractSimple(bytes, index*length, length)
}

// extractSimple does the actual shift and length calculation described in Extract without any security checks
func extractSimple(bytes []byte, shift, length byte) byte {
	var (
		byteStart = shift / 8
		bitStart  = shift % 8
	)

	// get the bits at bs and shift them to b0
	result := bytes[byteStart] >> bitStart

	// check if the next byte needs to be accessed
	if bitStart+length > 8 {
		// get the bits and shift them to (8-bs) and add it to the result
		result += bytes[byteStart+1] << (8 - bitStart)
	}
	return result % (1 << length) // drop all the bits > length
}

// Weighted is a simple weighting function vor two values to create a weighted median.
func Weighted(w0 float64, v0 byte, w1 float64, v1 byte) byte {
	return byte(math.Round((w0*float64(v0) + w1*float64(v1)) / (w0 + w1)))
}
