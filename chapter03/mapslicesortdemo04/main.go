package main

import (
	"fmt"
	"sort"
)

func main() {

   m:=make(map[int]int)
   m[1]=10
   m[2]=1
   m[0]=56
   m[3]=3

   fmt.Println(m)

    var key [] int
	for k,_:=range m{
		key=append(key,k)
	}

    sort.Ints(key)
	fmt.Println(key)
	for k,_:=range key{
		fmt.Printf("m[%v]=%v\n",k,m[k])
	}

}
