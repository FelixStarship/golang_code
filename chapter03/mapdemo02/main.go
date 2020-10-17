package main

import "fmt"

func main() {


	m:=make(map[string]string)
	m["city01"]="北京"
	m["city02"]="上海"
	m["city03"]="深圳"
	m["city04"]="广州"
	delete(m,"city01")
	//一次删除所有map
	//m=make(map[string]string)

	if val,ok:=m["city02"];ok {
		fmt.Printf("m[city02]=%v\n",val)
	}else {
		fmt.Println("m[city01]",val)
	}
	
	fmt.Println(m)


	studMap:=make(map[string]map[string]string)
	studMap["stud01"]= make(map[string]string)

	studMap["stud01"]["name"]="彭于晏"
	studMap["stud01"]["sex"]="男"
	studMap["stud01"]["address"]="南山文体中心"

	studMap["stud02"]=make(map[string]string)
	studMap["stud02"]["name"]="政府"
	studMap["stud02"]["sex"]="女"
	studMap["stud02"]["address"]="保安体育中心"

	for k1,v1:= range studMap{
		fmt.Println("k1=",k1)
		for k2,v2:= range v1{
			fmt.Printf("\t k2=%v v2=%v",k2,v2)
		}
		fmt.Println()
	}
}
