package main

import "math"

var Pi float64
var Pii float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
	Pii = Pi
}
