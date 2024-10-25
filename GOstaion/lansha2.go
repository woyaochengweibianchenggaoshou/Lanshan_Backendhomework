package main

import "math"

import "fmt"

func get_S(r float64) float64 {
	return r * r * math.Pi
}

func main() {
	var r float64
	fmt.Printf("please input the radius:")
	fmt.Scan(&r)
	s := get_S(r)
	fmt.Printf("the S is: %f\n", s)
}
