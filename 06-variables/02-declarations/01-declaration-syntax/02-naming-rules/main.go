// package main

// // VARIABLE NAMING RULES

// func main() {
// 	// CORRECT DECLARATIONS
// 	var speed int
// 	var SpeeD int

// 	// underscore is allowed but not recommended
// 	var _speed int

// 	// Unicode letters are allowed
// 	var 速度 int

//		// keep the compiler happy
//		_, _, _, _ = speed, SpeeD, _speed, 速度
//	}
package main

func main() {
	var speed int
	var SpeedD int

	var _speed int

	var 速度 int

	_, _, _, _ = speed, SpeedD, _speed, 速度
}
