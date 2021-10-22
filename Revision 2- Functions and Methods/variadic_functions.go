package main

import "fmt"

func main() {
	sum(1, 2, 3, 4)
	sum(5, 6, 7, 8)
}
func sum(a ...int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println(total)
}
