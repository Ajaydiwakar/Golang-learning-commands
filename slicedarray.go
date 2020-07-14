package main
import "fmt"
func main () {

	EvenNum := [5]int{1,5,6,2,8}

	/*EvenNum[0] = 0
	EvenNum[1] = 4
	EvenNum[2] = 7
	EvenNum[3] = 45
	EvenNum[4] = 34
*/
 for _,value := range EvenNum {

 	fmt.Println(value)
 }

 numSlice :=[]int{4,7,5,2,8}
 sliced :=numSlice[3:5]
/*sliced :=numSlice[:5]   This line will print from 0 to 5th index  */
 fmt.Println(sliced)
	
}

