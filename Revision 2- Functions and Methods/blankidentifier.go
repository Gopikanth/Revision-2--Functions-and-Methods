package main

import "fmt"

func main() {
	z := [5]int{1, 23, 4, 6}
	for _, x := range z { //here blank identifier is used so only the values get printed
		fmt.Println(x)
	}
}
