package core

import (
	"encoding/binary"
	"fmt"
)

// Print - print the values of the given attribute
func (attr *Attribute) Print() {
	nested := attr.Type&NlaFNested != 0
	fmt.Printf("type=%d nested=%v len=%v %v\n",
		attr.Type&NlaTypeMask, nested, len(attr.Value), attr.Value)
}

// PrintRawAttributes - Print the netlink attributes raw attribute bytes
func PrintRawAttributes(rawAttributes []byte) error {
	return printRawAttributes(rawAttributes, 0)
}

// printRawAttributes - Given a byte slice and starting levels to print, write the attributes to stdout
func printRawAttributes(rawAttributes []byte, startingLevel int) error {
	offset := 0
	for offset < len(rawAttributes) {
		// Ensure there's enough data for the header
		if len(rawAttributes[offset:]) < 4 {
			return fmt.Errorf("insufficient data at offset %d for attribute header", offset)
		}

		// Read the attribute length and type
		nlaLen := binary.LittleEndian.Uint16(rawAttributes[offset : offset+2])
		nlaType := binary.LittleEndian.Uint16(rawAttributes[offset+2 : offset+4])

		// Validate nlaLen
		if nlaLen < 4 {
			return fmt.Errorf("invalid attribute length %d at offset %d", nlaLen, offset)
		}
		if int(offset+int(nlaLen)) > len(rawAttributes) {
			return fmt.Errorf("attribute length %d at offset %d exceeds data size", nlaLen, offset)
		}

		// Extract the attribute value
		valueStart := offset + 4
		valueEnd := offset + int(nlaLen)
		value := rawAttributes[valueStart:valueEnd]

		// Create the attribute
		attribute := Attribute{
			Type:  NlaFlags(nlaType),
			Value: value,
		}

		// Print indentation
		for i := 0; i < startingLevel; i++ {
			fmt.Print("> ")
		}

		nested := attribute.Type&NlaFNested != 0

		attribute.Print()

		// If nested, recursively process the nested attributes
		if nested {
			if err := printRawAttributes(attribute.Value, startingLevel+1); err != nil {
				return err
			}
		}

		// Advance the offset by the padded length of the attribute
		// Attributes are aligned to 4-byte boundaries
		alignedLen := (int(nlaLen) + 3) & ^3
		offset += alignedLen
	}
	return nil
}
