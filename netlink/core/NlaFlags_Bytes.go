package core

import (
	"github.com/sam-caldwell/convert"
)

// Bytes - Convert NlaFlags (uint16) to bytes
func (m *NlaFlags) Bytes() []byte {

	return convert.Uint16ToBytes(uint16(*m))

}
