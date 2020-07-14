package main
import "fmt"
func main (){

	var x int = 1500
	var y int = 1490

	var ret int 

	ret = max(x,y)
	fmt.Printf("max value is %d\n" ,ret)

}
func max(num1 , num2  int) int {
	var result int
	if ( num1 > num2 ){
		result = num1
	}else {
		result = num2
	}
	return result 
}
