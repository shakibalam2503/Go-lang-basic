package main
import
	(
		"fmt"
		"learning/myutils"
		"learning/adder"
	)
	
func main(){
	fmt.Println("hello world")
	fmt.Println("----------------")
	myutils.PrintMessage("hello this is from print message")
	adder.Add(2,4)
}