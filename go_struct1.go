package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  //匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main(){
	//初始化一个学生
	mark := Student{Human{"小明",25,170},"计算机专业"}

	fmt.Println(mark.name,"-",mark.age,"岁，",mark.weight,"cm，",mark.speciality)

	mark.name = "李白"
	mark.age = 50
	mark.speciality = "诗人"
	fmt.Println(mark.name,"-",mark.age,"岁，",mark.weight,"cm，",mark.speciality)

}
