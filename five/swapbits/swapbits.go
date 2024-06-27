package main

import "fmt"

func SwapBits(octet byte) byte {
	return (octet << 4) | (octet >> 4)
}

func main() {
	octet := byte(0b01000001) // Example byte
	swapped := SwapBits(octet)
	fmt.Printf("Original byte: %08b\n", octet)
	fmt.Printf("Swapped byte: %08b\n", swapped)
}
