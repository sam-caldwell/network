package core

import "fmt"

// PrintAttributes - Print the netlink attributes contained in the input byte slice
func PrintAttributes(rawAttributes []byte) {

	printAttributes(rawAttributes, 0)

}

// printAttributes - Given a byte slice and starting levels to print, write the attributes to stdout
func printAttributes(rawAttributes []byte, startingLevel int) {

	for attr := range ParseAttributes(rawAttributes) {

		for i := 0; i < startingLevel; i++ {
			print("> ")
		}

		nested := attr.Type&NlaFNested != 0

		fmt.Printf("type=%d nested=%v len=%v %v\n", attr.Type&NlaTypeMask, nested, len(attr.Value), attr.Value)

		if nested {
			printAttributes(attr.Value, startingLevel+1)
		}

	}

}
