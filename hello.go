package main

import (
	"fmt"
)

//全局变量
var arr4 [3][2]int
var arr5 [3][2]int = [...][2]int{{1,2},{3,4},{5,6}}

type Student struct {
	Age int
	Name string
}

type T struct {
	value int
}

func main()  {
	/**
	局部变量
	 */

	/**字符串**/
	//var a string = "hello world"
	//a = "hello world1"
	//b := "hello world"
	//var a,b ="hello world1","hello world"
	a,b := "hello world1", "hello world"
	fmt.Println(a,b)

	/**整形**/
	//var c int = 1
	//d := 2
	//var c,d = 1,2
	c,d := 1,2
	fmt.Println(c,d)

	//多种类型一起申明
	//var e,f = "hello world2", 2
	e,f := "hellow world2", 2
	fmt.Println(e,f)

	//匿名变量-暂时错误的
	//x,_ := "我是:","匿名变量"
	//fmt.Println(x)

	/**数组**/
	var arr1 [3]int = [3]int {1,2,3}
	var arr2 = [3]int {2,3,4}
	var arr3 = [...]int{4,5,6}
	strArr := [3]string{0:"hello world3", 1:"hello world4", 2:"hello world5"}
	strArr1 := [...]struct{
		name string
		age uint8
	}{
		{"name1",5},
		{"name2",3},
		{"name3",9},
	}
	//多维数组
	arr6 := [2][3]int{{1,2,3},{4,5,6}}

	fmt.Println(len(arr1),cap(arr2),arr3,strArr,strArr1, arr4, arr5, arr6)

	//数组遍历
	var arr7 [2][3]int = [...][3]int{{1,2,3},{7,8,9}}
	for k1,v1 := range arr7{
		for k2,v2 := range v1{
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
	}
	fmt.Println()

	/**切片**/
	slice := make([]int, 5) //初始化长度和容量都为5的切片
	slice1 := make([]int, 5, 10) //初始化长度为5容量为10的切片
	sliceShort := []int{1,2,3,4,5} //简短切片命名方式

	//使用数组初始化切片
	arrSlice := [5]int{1,2,3,4,5}
	sliceB := arrSlice[0:3] //左闭右开区间 sliceB最终为[1,2,3]

	//使用切片来初始化切片
	sliceC := []int{1,2,3,4,5}
	sliceD := sliceC[0:3] //左闭右开区间 sliceD最终为[1,2,3]

	fmt.Println(slice, slice1, sliceShort, sliceB, sliceD) //输出切片
	fmt.Println()

	sliceE := []int{1,2,3,4,5}
	newSlice := make([]int, len(sliceE))
	copy(newSlice, sliceE)
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", sliceE)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	newSlice[0] = 6
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", sliceE)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	//使用copy部分拷贝切片 [2,3,4]
	sliceF := []int{1,2,3,4,5}
	newSliceA := make([]int,3)
	copy(newSliceA,sliceF[1:4])
	fmt.Println(newSliceA)
	fmt.Println()

	/**Map**/
	var cMap map[string]int //定义一个map的变量 此时为nil
	fmt.Println(cMap)
	fmt.Println(cMap == nil)
	fmt.Println()
	//cMap["beijing"] = 1 //报错，因为 cMap 为 nil
	//fmt.Println(cMap)

	cMap1 := make(map[string]int)
	fmt.Println(cMap1)
	fmt.Println(cMap1 != nil) //map可以添加值，因为 cMap1 不为 nil
	cMap1["beijing"] = 1
	fmt.Println(cMap1)
	fmt.Println()

	//指定初始容量
	cMap2 := make(map[string]int,100)
	cMap2["shanghai"] = 2
	fmt.Println(cMap2)
	fmt.Println()

	//简短申明方式
	cMap3 := map[string]int{"sichuan":3}
	fmt.Println(cMap3)
	fmt.Println()

	//基本操作
	code := cMap1["beijing"] //获取key对应的value
	fmt.Println(code)

	code1 := cMap1["shanghai"] //获取key对应的value 无输出0
	fmt.Println(code1)

	code2,ok := cMap1["tianjing"] //判断key是否存在
	if ok {
		fmt.Println(code2)
	}else {
		fmt.Println("code2对应的map1里里面没有tianjing")
	}

	delete(cMap1,"beijing") //删除key
	fmt.Println(cMap1) //删除完为map[]
	fmt.Println()

	//map循环无序性校验
	cMap4 := map[string]int{"beijing":1,"shanghai":2,"shengzhen":3,"sichuan":4}
	for cMap4,code3 := range cMap4 {
		fmt.Println("%s:%d", cMap4,code3)
	}
	fmt.Println() //输出的map是无需的

	//线程不安全处理 暂时处理不来等待学会-. .-
	/*cMap5 := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		cMap5["beijing"] = 10
		wg.Done()
	}()

	go func() {
		cMap5["shichuan"] = 20
		wg.Done()
	}()

	wg.Wait()*/

	//map 嵌套
	provinces := make(map[string]map[string]int)
	provinces["北京"] = map[string]int{
		"东城区":1,
		"西城区":2,
		"朝阳区":3,
		"海淀区":4,
	}
	fmt.Println(provinces)
	fmt.Println()

	/**自定义类型**/
	type City string
	type Age int

	//批量声明
	type (
		B0 = int8
		B1 = int16
		B2 = int32
		B3 = int64
	)

	type (
		A0 int8
		A1 int16
		A2 int32
		A3 int64
	)

	city := City("咪咪蛋")
	fmt.Println(city)
	fmt.Println("我住在：", "成都的"+city)
	fmt.Println(len(city))
	fmt.Println()

	age := Age(12)
	if age >= 12 {
		fmt.Println("满了12岁就是一个青少年了哦！")
	}
	fmt.Println(age)
	fmt.Println("我今年是：",age)
	fmt.Println()

	//函数参数
	printAge(int(age))
	fmt.Println()

	/**结构体**/
	stu := Student{
		Age: 18,
		Name: "狗蛋",
	}
	fmt.Println(stu)
	fmt.Println(Student{20,"二丫"})//结构体赋值
	fmt.Println()

	/**函数**/
	fmt.Println(plus(7,2))
	fmt.Println(multi("狗蛋",2))
	fmt.Println(namedReturnValud())
	fmt.Println(sum(1,2,3,4,5,9,0))
	//匿名函数
	func(name string){
		fmt.Println(name)
	}("狗剩和狗蛋")
	//闭包函数
	addOne := addInt(1)
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())

	addTwo := addInt(2)
	fmt.Println(addTwo())
	fmt.Println(addTwo())
	fmt.Println(addTwo())

	//函数作为参数传递
	Logger(sayHello,"狗剩和狗蛋")

	// 传值和传引用
	str := "禾木课堂"
	fmt.Println("before calling sendValue, str : ", str)
	sendValue(str)
	fmt.Println("after calling sendValue, str : ", str)

	fmt.Println("before calling sendAddress, str : ", str)
	sendAddress(&str)
	fmt.Println("after calling sendAddress, str: ", str)
	fmt.Println()

	/**方法**/
	//值作为接收者（T） 不会修改结构体值，而指针 *T 可以修改。
	m := T{0}
	fmt.Println(m)

	//传值
	m.StayTheSame()
	fmt.Println(m)

	//传指针
	m.Update()
	fmt.Println(m)
	fmt.Println()

}

func printAge(age int)  {
	fmt.Println("Age is ", age)
}


//单返回值函数
func plus(a,b int) (res int){
	return a + b
}

//多返回值函数
func multi(a string, b int) (string,int)  {
	return a,b
}

//命名返回值 被命名的返回参数的值为该类型的默认零值
//该例子中 name 默认初始化为空字符串，height 默认初始化为 0
func namedReturnValud() (name string, height int) {
	name = "狗剩"
	height = 165
	return
}

//参数可变函数
func sum(nums ...int) int {
	res := 0
	for _,v := range nums{
		res += v
	}

	return res
}

//闭包函数
func addInt(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

func sayHello(name string)  {
	fmt.Println("hello", name)
}

//函数作为参数函数
func Logger(f func(string), name string)  {
	fmt.Println("开始调用sayHello方法")
	f(name)
	fmt.Println("结束调用sayHello方法")
}

//传值函数
func sendValue(name string)  {
	name = "狗蛋"
}

//传引用函数
func sendAddress(name *string)  {
	*name = "狗蛋&&狗剩"
}

//方法值接收
func (m T) StayTheSame() {
	m.value = 3
}

//指针修改
func (m *T) Update() {
	m.value = 3
}