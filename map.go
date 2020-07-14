package main
import "fmt"
func main () {

	StudentAge := make(map[string] int)

	StudentAge["Arya"] =23
	StudentAge["Ajay"] =21
	StudentAge["Amir"] =24
	StudentAge["Ankit"] =43
	StudentAge["Akriti"] =28
	StudentAge["Amit"] =22
	StudentAge["Aman"] =26

	fmt.Println(StudentAge)
	fmt.Println(StudentAge["Ankit"])
	fmt.Println(len(StudentAge))


}

