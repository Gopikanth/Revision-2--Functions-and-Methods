// defer methods are methods that delay the execution untill nearby functions executed
package main

import "fmt"

func main() {
	defer fmt.Println(1)
	fmt.Println(2) //it is executed first because it is not a defer method
	defer fmt.Println(3)

}
