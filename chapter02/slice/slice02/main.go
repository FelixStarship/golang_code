package main

import "fmt"



func fbn(n int) [] uint64 {
	fbnSlice:=make([]uint64,n)
	fbnSlice[0]=1
	fbnSlice[1]=1
	for i:=2;i<n;i++ {
		fbnSlice[i]=fbnSlice[i-1]+fbnSlice[i-2]
	}
	return fbnSlice
}


func main() {

	//切片
	var slice[] int=make([]int,5,10)
	slice[2]=90
	slice[4]=100
    fmt.Printf("slice的长度%v  slice的容量%v  slice=%T  slice=%v  slice=%p   slice=%p\n",
    	len(slice),cap(slice),slice,slice,&slice,&slice[0])

	var slice1 [] string=[]string{"tom","jack","test"}
	fmt.Printf("slice01=%v\n",slice1)

	for i:=0;i<len(slice1);i++ {
		fmt.Println(slice1[i])
	}

	for i,v:= range slice1{
		fmt.Println(i,v)
	}

	//slice动态扩容append
	slice1=append(slice1,"am","pm")
	fmt.Printf("slice1的容量%v\n",cap(slice1))


	str:="hello@hxz"
	strSlice:=str[6:]
	fmt.Println(strSlice)

	arr:=[]rune(str)
    arr[6]='北'
    str=string(arr)
    fmt.Println(str)
	fmt.Println(strSlice)

   fbnSlice:=fbn(20)
   fmt.Println(fbnSlice)
}


