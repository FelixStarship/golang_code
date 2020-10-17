package main

import "fmt"

type Cat struct {
	Name string
	Age int
	Color string
}

func main() {

  var cat Cat
  cat.Name="小白"
  cat.Age=20
  cat.Color="灰色"

  fmt.Printf("cat=%v  type=%T cat=%p",cat,cat,&cat)


}
