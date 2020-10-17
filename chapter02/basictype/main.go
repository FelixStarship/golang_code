package main

import (
	"fmt"
	"strconv"
)

func main() {

	var name1 =10
	var name2=12.0
	var name3=true
	var str string

	str=fmt.Sprintf("%d",name1)
	fmt.Printf("str=%q str=%T\n",str,str)

	str=fmt.Sprintf("%f",name2)
	fmt.Printf("str=%q type=%T\n",str,str)

	str=fmt.Sprintf("%t",name3)
	fmt.Printf("str=%q type=%T\n",str,str)

	var num1=99
	var num2=99.89
	var num3=true
	str=strconv.FormatInt(int64(num1),10)
	fmt.Printf("str=%q type=%T\n",str,str)

	str=strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("str=%q  type=%T\n",str,str)

	str=strconv.FormatBool(num3)
	fmt.Printf("str=%q  type=%T\n",str,str)


	var  num5=56
	str=strconv.Itoa(num5)
	fmt.Printf("str=%q  type=%T\n",str,str)
}

