package utils

import "fmt"

var Age int
var Name string

func init()  {
	fmt.Println("utils包的 init()....")
	Age=20
	Name="tom~"
}