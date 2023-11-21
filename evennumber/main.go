package main

import (
	"fmt"
	"math"
)
func main(){
	numberRange :=[]int{0,1,2,3,4,5,6,7,8,9,10}
	for _,num:=range numberRange{
		m:= math.Mod(float64(num),2)
		if m == 0 {
			fmt.Printf("%v is an even number \n",num)
		}
		if m != 0 {
			fmt.Printf("%v is an odd number \n",num)
		}
	}
}