package patricia

import (
	"encoding/binary"
)

const _leftmost32Bit = uint32(1 << 31)

// IPv4Address is a representation of an IPv4 address and CIDR
type IPv4Address struct {
	Address uint32
	Length  uint
}

// NewIPv4Address creates an address from the input IPv4 address
func NewIPv4Address(address uint32, length uint) IPv4Address {
	return IPv4Address{
		Address: address,
		Length:  length,
	}
}

// NewIPv4AddressFromBytes creates an address from the input IPv4 address bytes
// - address must be 4 bytes
func NewIPv4AddressFromBytes(address []byte, length uint) IPv4Address {
	return IPv4Address{
		Address: binary.BigEndian.Uint32(address),
		Length:  length,
	}
}

// ShiftLeft shifts the address to the left
func (i *IPv4Address) ShiftLeft(shiftCount uint) {
	i.Address <<= shiftCount
	i.Length -= shiftCount
}

// IsLeftBitSet returns whether the leftmost bit is set
func (i *IPv4Address) IsLeftBitSet() bool {
	return i.Address >= _leftmost32Bit
}
