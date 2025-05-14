package dns

import "github.com/codecrafters-io/dns-server-starter-go/app/utils/bitutils"

// Here's the diagram of all flag fields:
//
// 			byte 1	       |		byte 2
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//   7  6  5  4  3  2  1  0  7  6  5  4  3  2  1  0
// +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
//
// more: https://datatracker.ietf.org/doc/html/rfc1035#section-4.1

type OPCode uint8

const (
	OP_QUERY    OPCode = 0 // a standard query
	OP_IQUERY   OPCode = 1 // an inverse query
	OP_STATUS   OPCode = 2 // a server status request
	OP_RESERVED OPCode = 3 // 3-15 reserved fo future use
)

type RCode uint8

const (
	RCodeNoErr          RCode = 0 // no error
	RCodeFormarErr      RCode = 1 // format error
	RCodeServerFail     RCode = 2 // server failure
	RCodeNameErr        RCode = 3 // name error
	RCodeNotImplemented RCode = 4 // not implemented error
	RCodeRefused        RCode = 5 // server refused
	RCodeReserved       RCode = 6 //  6-15 reserved fo future use
)

const (
	mask_QR     = 0x80 // 0b10000000
	mask_OPCode = 0x78 // 0b01111000
	mask_AA     = 0x04 // 0b00000100
	mask_TC     = 0x02 // 0b00000010
	mask_RD     = 0x01 // 0b00000001
	mask_RA     = 0x80 // 0b10000000
	mask_Z      = 0x70 // 0b01110000
	mask_RCode  = 0x0F // 0b00001111

	shift_OPCode uint8 = 3 // as OPCode uses bits 6-3
	shift_Z      uint8 = 4 // as OPCode uses bits 6-4
	shift_RCode  uint8 = 0 // as OPCode uses bits 3-0
)

type HeaderFlags struct {
	flagBytes [2]byte
}

func (f *HeaderFlags) GetQR() bool {
	return bitutils.GetBit(f.flagBytes[0], mask_QR)
}

func (f *HeaderFlags) SetQR(value bool) {
	f.flagBytes[0] = bitutils.SetBit(f.flagBytes[0], mask_QR, value)
}

func (f *HeaderFlags) GetOPCode() OPCode {
	value := bitutils.GetBits(f.flagBytes[0], mask_OPCode, shift_OPCode)
	return OPCode(value)
}

func (f *HeaderFlags) SetOPCode(value OPCode) {
	f.flagBytes[0] = bitutils.SetBits(f.flagBytes[0], mask_OPCode, shift_OPCode, byte(value))
}

func (f *HeaderFlags) GetAA() bool {
	return bitutils.GetBit(f.flagBytes[0], mask_AA)
}

func (f *HeaderFlags) SetAA(value bool) {
	f.flagBytes[0] = bitutils.SetBit(f.flagBytes[0], mask_AA, value)
}

func (f *HeaderFlags) GetTC() bool {
	return bitutils.GetBit(f.flagBytes[0], mask_TC)
}

func (f *HeaderFlags) SetTC(value bool) {
	f.flagBytes[0] = bitutils.SetBit(f.flagBytes[0], mask_TC, value)
}

func (f *HeaderFlags) GetRD() bool {
	return bitutils.GetBit(f.flagBytes[0], mask_RD)
}

func (f *HeaderFlags) SetRA(value bool) {
	f.flagBytes[0] = bitutils.SetBit(f.flagBytes[0], mask_RD, value)
}

func (f *HeaderFlags) GetZ() byte {
	return bitutils.GetBits(f.flagBytes[0], mask_Z, shift_Z)

}

func (f *HeaderFlags) SetZ(value byte) {
	f.flagBytes[0] = bitutils.SetBits(f.flagBytes[0], mask_Z, shift_Z, value)
}

func (f *HeaderFlags) GetRCode() RCode {
	value := bitutils.GetBits(f.flagBytes[0], mask_RCode, shift_RCode)
	return RCode(value)
}

func (f *HeaderFlags) SetRCode(value RCode) {
	f.flagBytes[0] = bitutils.SetBits(f.flagBytes[0], mask_RCode, shift_RCode, byte(value))
}

func NewHeaderFlags(b []byte) *HeaderFlags {
	flagBytes := [2]byte{b[0], b[1]}
	return &HeaderFlags{flagBytes}
}
