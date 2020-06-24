package main

import (
	"fmt"
	"math"
)

//定义接口
//type geometry interface {
//	area() float32
//	perim() float32
//}

type rect struct {
	len, wid float32
}

type cycle struct {
	radius float32
}

/**接口种内嵌接口**/
//接口种内嵌接口
type tmp interface {
	area() float32
}

type geometry interface {
	tmp
	perim() float32
}

func main()  {
	rec := rect{
		len: 2,
		wid: 4,
	}

	show("rect", rec)

	cir := cycle{
		radius: 6,
	}

	show("cycle", cir)

	show("test", "test param")
}

//求面积方法
func (r rect) area() float32  {
	return r.len * r.wid
}

//求直径方法
func (r rect) perim() float32  {
	return 2*(r.len+r.wid)
}

//求园行面积
func (c cycle) area() float32 {
	return math.Pi*c.radius*c.radius
}

//求原型直径
func (c cycle) perim() float32  {
	return 2*math.Pi*c.radius
}

func show(name string, param interface{})  {
	switch param.(type) {
	case geometry:
		//类型断言不同的数据
		fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, param.(geometry).perim())
	default:
		fmt.Println("类型错误")
	}
}