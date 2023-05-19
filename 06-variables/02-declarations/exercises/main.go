package main

import "fmt"

var packageLevelVariable int

func main() {

	var height int
	fmt.Println(height)

	var isOn bool
	fmt.Println(isOn)

	var brightness float64
	fmt.Println(brightness)

	// %T prints the type of the value
	// %q prints an empty string
	var name string
	fmt.Printf("(%T): %q\n", name, name)

	// 02-declarations/exercises/main.go:23:6: syntax error: unexpected literal 3, expected name
	// var 3speed int

	// 02-declarations/exercises/main.go:26:14: syntax error: unexpected !, expected name
	// var !speed int

	// 02-declarations/exercises/main.go:28:9: invalid character U+003F '?'
	// var spe?ed,

	// 02-declarations/exercises/main.go:32:6: syntax error: unexpected var, expected name
	// var var int

	// 02-declarations/exercises/main.go:34:6: syntax error: unexpected func, expected name
	// var func int

	// 02-declarations/exercises/main.go:36:6: syntax error: unexpected package, expected name
	// var package int

	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var r rune
	var b byte

	fmt.Println(i8, i16, i32, i64, f32, f64, c64, c128, r, b)

	var (
		active bool
		delta  int64
	)
	fmt.Println(active, delta)

	var firstName, lastName string
	fmt.Printf("%q %q\n", firstName, lastName)

	var isLiquid bool = false

	_ = isLiquid

	// 02-declarations/exercises/main.go:67:14: undefined: isLiquid2
	// 02-declarations/exercises/main.go:69:6: isLiquid2 declared and not used
	// fmt.Println(isLiquid2)

	// var isLiquid2 = false
}
