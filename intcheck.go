package main

import (
	"fmt"
	
)

func main() {

     var n int = 42
     var m int8 = 54
     
	fmt.Printf("%v , %T\n",n, n)
	
	
	fmt.Printf("%v , %T\n",m,m)

	fmt.Println(int(m)+n)
}


