package main

import (
	"fmt"
	"strconv"
)

func main() {

	//map slice
	monsters := make([]map[int]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[int]string)
		monsters[0][1] = "牛魔王"
		monsters[0][2] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[int]string)
		monsters[1][3] = "玉兔精"
		monsters[1][4] = "600"
	}

	//index out of range [2] with length 2
	//if monsters[2]==nil {
	//	monsters[1]=make(map[string]string)
	//	monsters[1]["name"]="玉兔精"
	//	monsters[1]["age"]="600"
	//}

	//动态扩容
	monsters = append(monsters, map[int]string{5: "火云谢生~", 6: "20"})

	fmt.Println(monsters)

	m := make(map[int]string)
	for i := 0; i < 30; i++ {
		m[i] = strconv.Itoa(i)
	}

	for k, _ := range m {
		fmt.Println(k)
	}

	//切片没有初始化
	var a []int
	a[0] = 100
	fmt.Println(a)
}
