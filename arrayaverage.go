package main
import "fmt"
func main (){
	var a=[]int{100,1200,5,2,67}
	var avg float64
	avg = getavg(a,5)
	fmt.Printf("Average is %f\n ",avg)
}
func getavg(arr []int , size int) float64 {
	var i,sum int
	var avg float64
	for i=0;i<size;i++{
		sum += arr[i]

	}
	avg=float64(sum/size)
	return avg
}