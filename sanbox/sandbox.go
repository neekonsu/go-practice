package main

import "fmt"

func addItems(x float64, y float64) float64 {
	return x + y
}

func add(x, y float32) float32 {
	return x + y
}

func main() {
	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
	var bool bool = true
	fmt.Println(bool)
	var n1 float64 = 6.6
	var n2 float64 = 9.9
	fmt.Println(addItems(n1, n2))
	var n3, n4 float64 = 1.1, 2.2
	fmt.Println(n3, n4)
	n5, n6 := 5.6, 9.5
	fmt.Println(addItems(n5, n6))
	fmt.Println(add(float32(n5), float32(n6)))
}
