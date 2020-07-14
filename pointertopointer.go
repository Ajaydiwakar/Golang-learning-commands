package main
import "fmt"
func main (){
	var a int
	var b *int
	var pptr **int
	a= 100
	b = &a
	pptr =&b


	fmt.Printf("value of a is %d\n", a)

	fmt.Printf("value of a is %d\n", *b)

	fmt.Printf("value of a is %d\n", **pptr)

	fmt.Printf("value of a is %d\n", b)
	fmt.Printf("value of a is %d\n", pptr)

}