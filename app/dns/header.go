package dns

import (
	"encoding/binary"
	"fmt"
)

// Here's the diagram of all flag fields:
//
//   0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |                      ID                       |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |                    QDCOUNT                    |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |                    ANCOUNT                    |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |                    NSCOUNT                    |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |                    ARCOUNT                    |
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//
// more: https://datatracker.ietf.org/doc/html/rfc1035#section-4.1

type Header struct {
	ID      uint16
	Flags   *HeaderFlags
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func (h *Header) Unmarshall(buf []byte) error {
	if len(buf) < 12 {
		return fmt.Errorf("invalid header format")
	}

	h.ID = binary.BigEndian.Uint16(buf[:2])
	h.Flags = NewHeaderFlags(buf[2:4])
	h.QDCount = binary.BigEndian.Uint16(buf[4:6])
	h.ANCount = binary.BigEndian.Uint16(buf[6:8])
	h.NSCount = binary.BigEndian.Uint16(buf[8:10])
	h.ARCount = binary.BigEndian.Uint16(buf[10:12])

	return nil
}

func (h *Header) Marshall() []byte {
	buf := make([]byte, 12)

	binary.BigEndian.PutUint16(buf[:2], h.ID)
	copy(buf[2:4], h.Flags.flagBytes[:])
	binary.BigEndian.PutUint16(buf[4:6], h.QDCount)
	binary.BigEndian.PutUint16(buf[6:8], h.ANCount)
	binary.BigEndian.PutUint16(buf[8:10], h.NSCount)
	binary.BigEndian.PutUint16(buf[10:12], h.ARCount)

	return buf
}

func NewHeader() *Header {
	return &Header{
		Flags: &HeaderFlags{},
	}
}
