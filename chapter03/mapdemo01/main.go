package main

import "fmt"

func main() {

	var  m map[string]string
	m=make(map[string]string,2)
	m["n1"]="李"
	m["n2"]="王"
	m["n4"]="张"
	m["n4"]="某"
	m["n5"]="张"
	m["n6"]="张"

	fmt.Printf("m=%v  m=%v\n",m,len(m))

	var m1 map[string] string= map[string]string{"1":"1"}

	fmt.Printf("m1=%v m1=%v m1=%T\n",m1,len(m1),m1)


	studMap:=make(map[string]map[string]string)
	studMap["stud01"]= make(map[string]string)

	studMap["stud01"]["name"]="彭于晏"
	studMap["stud01"]["sex"]="男"
	studMap["stud01"]["address"]="南山文体中心"

	studMap["stud02"]=make(map[string]string)
	studMap["stud02"]["name"]="政府"
	studMap["stud02"]["sex"]="女"
	studMap["stud02"]["address"]="保安体育中心"

	fmt.Println(studMap)

}
