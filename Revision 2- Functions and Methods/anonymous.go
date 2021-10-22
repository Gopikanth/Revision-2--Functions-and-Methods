//special type of function which is used to make inline function
package main

import "fmt"

func main() {
	e := func(c, d string) string {
		return c + d + "hello"
	}
	S(e)
	func(a string) {
		fmt.Println(a)

	}("granting") //passing values to anonymus functions
	b := func() {
		fmt.Println("Hai Maddy")
	}
	b() //assigning variables to functions

}
func S(i func(c, d string) string) {
	fmt.Println(i("i", "hi")) // passing arguments to anonymous functions

}
