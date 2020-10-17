package main

import (
	"fmt"
)

func test()  {

  defer func(){
		if err:=recover();err!=nil{
			fmt.Println("panic")
		}
	}()

	num1:=10
	num2:=0
	res:=num1/num2
	fmt.Println("res=",res)

}

func main() {
	test()
}

