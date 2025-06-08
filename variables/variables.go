package main

import "fmt"

func main() {

	// signed integers
	var a int = 32
	var b int8 = 16
	var c int16 = 554
	var d int32 = 3535345
	var e int64 = 54353534534

	// unsigned integers
	var f uint = 888
	var g uint16 = 45454
	var h uint32 = 454549534
	var i uint64 = 455353453435
	var j uint8 = 45

	//floating point variables
	var k float32 = 3.45444
	var l float64 = 455.655455

	//strings
	var m string = "I'm a simple string"

	// booleans
	var n bool = true
	var o bool = false

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
}
