package main
import "fmt"
func main (){
	var a=[5][2]int{}
	var i int 
	var j int
	for i =1 ; i<=5 ;i++{
		for j=1;i<=2 ;j++{
			fmt.Printf(j + " ")
		}
		fmt.Println(" ")
	}
}