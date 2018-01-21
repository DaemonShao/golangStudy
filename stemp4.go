package main

/*多值返回*/

import "fmt"

func swap(a int, b int) (int , int)  {

	return b,a
	
}

func main()  {
	a,b := swap(5,6)    //多值返回需要是 :=
	fmt.Println(a,b)
}
