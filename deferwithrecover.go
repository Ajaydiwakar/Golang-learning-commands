package main
import "fmt"
func main () {

	fmt.Println(div(3,0))
	fmt.Println(div(3,5))
}

func div (num1 , num2 int) int {

	defer func() {
		fmt.Println(recover())

	}()
	solution := num1/num2

	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()
