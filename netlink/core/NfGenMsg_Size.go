package core

import "unsafe"

// NfGenMsgSize - Track the message sizes for the correct serialization/deserialization
const NfGenMsgSize = int(unsafe.Sizeof(NfGenMsg{}))
