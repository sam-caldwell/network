package core

// AtmCell defines constants related to ATM (Asynchronous Transfer Mode) cell structure.
//
// ATM is a telecommunications technology that uses fixed-size cells to transfer data.
// The constants below define the size of the payload and the overall ATM cell size,
// according to the ATM cell structure.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/atm.h
type AtmCell uint8

const (
	// ATMCellPayload - ATM_CELL_PAYLOAD - defines the payload size of an ATM cell, which is 48 bytes.
	ATMCellPayload AtmCell = 48

	// ATMCellSize - ATM_CELL_SIZE - defines the total size of an ATM cell, which is 53 bytes (48 bytes of payload + 5 bytes of header).
	ATMCellSize AtmCell = 53
)
