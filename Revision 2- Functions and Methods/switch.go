//switch is an efficient way to transfer the execution to other parts of the code based on the value
//mainly there are two type of switches
//expression switch
//type switch

package main

import "fmt"

func main() {
	//experssion switch
	switch 3 {
	case 1:
		{
			fmt.Println("This is case 1")

		}
	case 2:
		{
			fmt.Println("This is case 2")
		}
	case 3:
		{
			fmt.Println("This is case 3")
		}
	}
	//switch case without optional statement and expression
	var a int = 1
	switch {
	case a == 1:

		fmt.Println("This is case 1")
	case a == 2:

		fmt.Println("This is case 2")
	case a == 3:

		fmt.Println("This is case 3")
	}
	//expression switch statement
	var value int = 3
	switch value {
	case 1, 4, 8:
		fmt.Println("This is case 1")
	case 2:

		fmt.Println("This is case 2")
	case 3, 6, 9:

		fmt.Println("This is case 3")

	}
	//type switch
	var b interface{}
	switch c := b.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", c)
	}
}
