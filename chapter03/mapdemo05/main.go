package main

import "fmt"

type stu struct {
	Name    string
	Age     int
	Address string
}

func modify(m map[int]int) {
	m[2] = 100
}

func main() {
	m := make(map[int]int)
	m[0] = 10
	m[1] = 20
	m[2] = 30
	m[3] = 40
	modify(m)
	fmt.Println(m)

	students := make(map[int]stu)
	students[0] = stu{"李", 20, "银田新村二排七号301"}
	students[1] = stu{"邱", 60, "美宜佳超市二排1号"}
	fmt.Println(students)
}
