package bitutils_test

import (
	"testing"

	"github.com/codecrafters-io/dns-server-starter-go/app/utils/bitutils"
	"github.com/stretchr/testify/assert"
)

func TestGetBit(t *testing.T) {
	var b byte = 0b10100000

	result := bitutils.GetBit(b, 0b10000000)
	assert.Equal(t, true, result)

	result = bitutils.GetBit(b, 0b01000000)
	assert.Equal(t, false, result)
}

func TestSetBit(t *testing.T) {
	var b byte = 0b00000000

	// set 0 value to 1
	result := bitutils.SetBit(b, 0b10000000, true)
	assert.Equal(t, uint8(0b10000000), result)

	// set 1 value to 0
	b = 0b01000000
	result = bitutils.SetBit(b, 0b01000000, false)
	assert.Equal(t, uint8(0b00000000), result)

	// set 1 value to 1
	b = 0b00100000
	result = bitutils.SetBit(b, 0b00100000, true)
	assert.Equal(t, uint8(0b00100000), result)

	// set 0 value to 0
	b = 0b00000000
	result = bitutils.SetBit(b, 0b00010000, false)
	assert.Equal(t, uint8(0b00000000), result)
}

func TestGetBits(t *testing.T) {
	var b byte = 0b10111100

	result := bitutils.GetBits(b, 0b00111100, 2)
	assert.Equal(t, uint8(0b00001111), result)

	result = bitutils.GetBits(b, 0b10100000, 5)
	assert.Equal(t, uint8(0b00000101), result)
}

func TestSetBits(t *testing.T) {
	var b byte = 0b10000000

	result := bitutils.SetBits(b, 0b00111100, 2, 0b00001111)
	assert.Equal(t, uint8(0b10111100), result)

	b = 0b00001101
	result = bitutils.SetBits(b, 0b00000110, 1, 0b00000011)
	assert.Equal(t, uint8(0b00001111), result)
}
