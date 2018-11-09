package main

import "fmt"

//所有的内置类型和自定义类型都是可以作为匿名字段的

type Skills []string   //定义一个切片类型

type Human struct {
	name string
	age int
	height int
}

type Studnet struct {
	Human  //匿名字段
	Skills //匿名字段，自定义的类型
	int   //内置类型作为匿名字段
	speciality string
}

func main(){
	//初始化一个学生
	mood := Studnet{Human:Human{"李白",35, 170},speciality:"诗人"}

	fmt.Println(mood.name,"-",mood.age,"-",mood.height,"cm","职业：",mood.speciality)

	//修改skills
	mood.Skills = []string{"唐朝"}
	fmt.Println(mood.Skills)

	//添加
	mood.Skills = append(mood.Skills,"王维","白居易")
	fmt.Println(mood.Skills)

	mood.int = 100
	fmt.Println(mood.int)
}
