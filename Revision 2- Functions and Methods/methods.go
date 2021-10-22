//the method contains a receiver argument
package main

import "fmt"

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {

	fmt.Println(a.name)      //returning name
	fmt.Println(a.branch)    //returning branch
	fmt.Println(a.particles) //returning particles
	fmt.Println(a.salary)    //returning salary
}

func main() {

	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}

	res.show()
}
