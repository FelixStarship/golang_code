package main

import "fmt"

type A struct {
	Name string
	Age int
}

func (a * A) SayOk()  {
	fmt.Println("A Sayok=%v",(*a).Name)
}

func (a * A) Hello()  {
	fmt.Println("A hello",a.Name,a.Age)
}

type B struct {
	A
	Name string
	Score int
}

type C struct {
   A
   B
}

type Goods struct {
	Name string
	Price float64
}

type  Brand struct {
	Name string
	Address string
}

//结构体多重嵌套
type TV struct {
	Goods    //匿名结构体
	Brand
}

type TV2 struct {
	*Goods    //匿名结构体
	*Brand
}

func main() {

	b:=B{}
	b.Name="访问匿名结构体"
	//通过匿名结构体名称访问
	b.A.Name="访问匿名结构体A"
	b.Age=30
	b.Score=90
	b.Hello()

	c:=C{}
	c.A.Name="aa"

	tv:=TV{Goods{
		Name: "金士顿",
		Price: 255.00,
	},
	Brand{
       Name: "ml",
       Address: "北京景昭业",
	}}
	fmt.Println("tv",tv)


	tv2:=TV2{
		&Goods{
			Name: "惠普",
			Price: 1200045.5,
		},
		&Brand{
			Name: "联想",
			Address: "深圳华强北",
		},
	}
	fmt.Printf("tv2=%v  tv2.Goods.Name=%p  tv2.Goods.Price=%p\n",tv2,&tv2.Goods.Name,&tv2.Goods.Price)
}
