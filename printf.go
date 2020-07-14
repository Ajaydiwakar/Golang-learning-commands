package main
import "fmt"
func main () {
	
	var name string ="Array paul"
	const pi float64 = 3.1412345
	win :=true

	fmt.Println(len(name))
	fmt.Println(name + " is a chill dude")


	fmt.Printf("%f \n" ,pi)
	fmt.Printf("%.3f \n" ,pi)
	fmt.Printf("%T \n",name)
	fmt.Printf("%T \n" ,pi)
	fmt.Printf("%t \n", win)
	fmt.Printf("%b \n" ,25)
	fmt.Printf("%c \n" ,65)
	fmt.Printf("%x \n",15)
	fmt.Printf("%e \n" ,pi)
}
