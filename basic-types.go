package main

import "fmt"

func basicTypes() {
	var b bool = false
	var s string = `hello`
	var i int = 1
	// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
	var f32 float32 = 3.14
	var f64 float64 = float64(f32) // No implicit conversion between typed values
	var r rune = '汉'               // Rob Pike & Ken Thompson designed UTF-8
	fmt.Println(b, s, i, f32, f64, r)
}

func shortVars() {
	b := false
	s := `hello`
	i := 1
	f64 := 3.14
	r := '汉'
	fmt.Println(b, s, i, f64, r)
}

// TODO constants
