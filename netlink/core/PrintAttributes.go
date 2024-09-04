package core

import "fmt"

func PrintAttributes(data []byte) {
	printAttributes(data, 0)
}

func printAttributes(data []byte, level int) {
	for attr := range ParseAttributes(data) {
		for i := 0; i < level; i++ {
			print("> ")
		}
		nested := attr.Type&NlaFNested != 0
		fmt.Printf("type=%d nested=%v len=%v %v\n", attr.Type&NlaTypeMask, nested, len(attr.Value), attr.Value)
		if nested {
			printAttributes(attr.Value, level+1)
		}
	}
}
