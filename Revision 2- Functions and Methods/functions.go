//functions are the block of code that can be reused so that memory and time gets saved
//in GO functions are declared using func keyword followed by name of that function

//calling a function
package main

import "fmt"

func main() {
	fmt.Printf("THE DISTANCE IS %d meter\n", distance(10, 30))  //here we are calling a function
	fmt.Printf("The value of a now becomes %d\n", swap(10, 20)) //clling by value
	p := 10
	q := 20
	fmt.Printf("The value of q now becomes %d", swapping(&p, &q)) //clling by reference

}
func distance(speed, time int) int {
	dist := speed * time
	return dist
}
func swap(a, b int) int {
	a = a + b
	b = a - b
	a = a - b
	return a
}
func swapping(c, d *int) int {
	var o int
	o = *c
	*c = *d
	*d = o

	return o
}
