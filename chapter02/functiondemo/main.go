package main

import "fmt"

func main() {
	fmt.Println(sum(1,2,3,45,56))
}


//slice切片  【可变参数】只能放到形参的最后一个参数
func sum(n1 int,args ...int)int  {
	sum:=n1
	for i:=0;i<len(args);i++ {
        sum+=args[i]
	}
	return sum
}