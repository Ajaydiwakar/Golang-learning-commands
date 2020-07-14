package main
import "fmt"
func main () {
	
	x :=10

	changevalue(&x)
	fmt.Println(x)
}

func changevalue(x *int) {
	*x=7;
}
