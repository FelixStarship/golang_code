package main

import (
	"fmt"
	_ "go_code/chapter02/init/utils"   //https://www.cnblogs.com/TimLiuDream/p/9929934.html
)

var age=test()
var age1=test

func test()int  {
	fmt.Println("test()")
	return 90
}

func init()  {
	fmt.Println("init()")
}

func main() {
	fmt.Printf("age1=%T\n",age1)
	fmt.Println("main()....age=",age)
	//fmt.Printf("Age=%v,Name=%v",utils.Age,utils.Name)
}
