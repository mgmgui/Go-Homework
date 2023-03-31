package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal
//
// EXPECTED OUTPUT
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// NOTES
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system

func main() {
	// 1, 2, 3, 4, 5, 6, 7, 8, 9
	fmt.Println(0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	// 10, 11, 12, 13, 14, 15
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	// 17
	fmt.Println(0x11)
	// 25
	fmt.Println(0x19)
	// 50
	fmt.Println(0x32)
	// 100
	fmt.Println(0x64)
}
