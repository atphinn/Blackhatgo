package smb

import (
	"github.com/blackhat-go/bhg/ch-6/smb/gss"
)

type NegotiateRes struct {
	Header
	StructureSize        uint16
	SecurityMode         uint16
	DialectRevision      uint16
	Reserved             uint16
	ServerGuid           []byte `smb:"fixed:16"`
	Capabilities         uint32
	MaxTransactSize      uint32
	MaxReadSize          uint32
	MaxWriteSize         uint32
	SystemTime           uint64
	ServerStartTime      uint64
	SecurityBufferOffset uint16 `smb:"offset:SecurityBlob"`
	SecurityBufferLength uint16 `smb:"len:SecurityBlob"`
	Reserved2            uint32
	SecurityBlob         *gss.NegTokenInit
}
