package main

import (
	"math"
	"fmt"
)

/*
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。
func (variable_name variable_data_type) function_name() [return_type]{
   // 函数体
}
*/

//定义结构体
type Circle struct {
	radius float64
}

//该方法属于Circle类型对象中的方法、
func (c Circle) getArea() float64 {
	//c.radius即为Circle类型对象中的属性
	return math.Pi * c.radius * c.radius
}

//定义结构体
type User struct {
	name string
	age int
	phone string
}

func(u User) getUserInfo() string {
	username := u.name
	age := u.age
	phone := u.phone
	userInfo := "姓名：" + username + " 年龄：" + string(age) + " 手机号码：" + phone
	return userInfo
}

func main(){
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆面积：",c1.getArea())
	fmt.Println("圆周率：",math.Pi)

	var u User
	u.name = "张三"
	u.age = 21
	u.phone = "18295666666"
	fmt.Println(u.getUserInfo())
}
