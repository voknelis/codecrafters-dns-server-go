package bitutils

// Returns true if the bit(s) specified by mask are set in b.
//
// Example:
//   GetBit(0b10100000, 0b10000000) => true (bit 7 is set)
//   GetBit(0b00100000, 0b10000000) => false (bit 7 is clear)
//
func GetBit(b byte, mask byte) bool {
	value := b & mask
	return value != 0
}

// Sets or clears the bit(s) specified by mask in b based on value.
//
// Example:
//   b := byte(0b00000000)
//   SetBit(b, 0b00000010, true)  => 0b00000010
//   SetBit(b, 0b00000010, false) => 0b00000000
//
func SetBit(b byte, mask byte, value bool) byte {
	if value {
		b |= mask // set bit
	} else {
		b &^= mask // clear bit
	}
	return b
}

// Extracts a multi-bit field from b using the mask and shift.
// The shift parameter defines how far right the bits need to be shifted after masking.
//
// Example:
//   GetBits(0b01111000, 0b01111000, 3) => 0b00001111
//
func GetBits(b byte, mask byte, shift uint8) byte {
	value := (b & mask) >> shift
	return value
}

// Clears the masked bits in b, then sets them to value shifted left by shift.
//
// Example:
//   b := byte(0b00000000)
//   SetBits(b, 0b01111000, 3, 0b00001011) => b == 0b01011000
//
func SetBits(b byte, mask byte, shift uint8, value byte) byte {
	b &^= mask // clear bits
	b |= byte(value << shift)
	return b
}
