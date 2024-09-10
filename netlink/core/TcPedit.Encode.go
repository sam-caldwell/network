package core

import (
	"bytes"
	"github.com/sam-caldwell/convert"
	"unsafe"
)

// Encode serializes the TcPedit structure into netlink attributes and adds it to the parent RtAttr.
// It encodes the `pedit` action, including the `TcPeditSel`, `TcPeditKey`, and `TcPeditKeyEx` fields.
//
// Parameters:
//
//	parent - The parent RtAttr to which the encoded `TcPedit` structure will be added.
func (p *TcPedit) Encode(parent *RtAttr) {
	// Add the "kind" attribute for the pedit action.
	parent.AddRtAttr(uint16(TcaActKind), convert.GoStringToNullTerminated("pedit"))

	// Add the options attribute for the pedit action.
	actOpts := parent.AddRtAttr(uint16(TcaActOptions), nil)

	// Create a buffer to serialize TcPeditSel and TcPeditKey.
	// The size of the buffer is determined by the size of TcPeditSel and TcPeditKey.
	bbuf := bytes.NewBuffer(make([]byte, 0, int(unsafe.Sizeof(p.Sel)+unsafe.Sizeof(p.Keys))))

	// Serialize TcPeditSel structure and write it into the buffer.
	bbuf.Write((*(*[SizeOfTcPeditSel]byte)(unsafe.Pointer(&p.Sel)))[:])

	// Serialize each TcPeditKey and write it into the buffer.
	for i := uint8(0); i < p.Sel.NKeys; i++ {
		bbuf.Write((*(*[SizeOfTcPeditKey]byte)(unsafe.Pointer(&p.Keys[i])))[:])
	}

	// Add the serialized TcPeditSel and TcPeditKey data to the act options.
	actOpts.AddRtAttr(uint16(TcaPeditParmsEx), bbuf.Bytes())

	// Add extended keys (TcPeditKeyEx) to the act options.
	exAttrs := actOpts.AddRtAttr(uint16(TcaPeditKeysEx)|uint16(NlaFNested), nil)
	for i := uint8(0); i < p.Sel.NKeys; i++ {
		// For each extended key, add a nested attribute.
		keyAttr := exAttrs.AddRtAttr(uint16(TcaPeditKeyEx)|uint16(NlaFNested), nil)

		// Create buffers to serialize the HeaderType and Command (Cmd) for the extended key.
		htypeBuf := make([]byte, 2)
		cmdBuf := make([]byte, 2)

		// Write the HeaderType and Command values into the buffers in native endian format.
		NativeEndian.PutUint16(htypeBuf, uint16(p.KeysEx[i].HeaderType))
		NativeEndian.PutUint16(cmdBuf, uint16(p.KeysEx[i].Cmd))

		// Add the serialized HeaderType and Command to the key attributes.
		keyAttr.AddRtAttr(uint16(TcaPeditKeyExHtype), htypeBuf)
		keyAttr.AddRtAttr(uint16(TcaPeditKeyExCmd), cmdBuf)
	}
}
