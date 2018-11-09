package main

import "fmt"

/*
type person struct {
	name string
	age int
}

var P person  // P现在就是person类型的变量了

P.name = "Astaxie"  // 赋值"Astaxie"给P的name属性.
P.age = 25  // 赋值"25"给变量P的age属性
fmt.Printf("The person's name is %s", P.name)  // 访问P的name属性.
除了上面这种P的声明使用之外，还有另外几种声明使用方式：

1.按照顺序提供初始化值

P := person{"Tom", 25}

2.通过field:value的方式初始化，这样可以任意顺序

P := person{age:24, name:"Tom"}

3.当然也可以通过new函数分配一个指针，此处P的类型为*person

P := new(person)
*/

//声明一个新的类型
type Person struct{
	name string
	age int
}

//比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
//struct也是传值的
func Older(p1,p2 Person) (Person,int){
	if p1.age>p2.age{
		return p1,p1.age-p2.age
	}
	return p2,p2.age - p1.age
}

func main(){
	var tom Person
	//赋值初始化
	tom.name,tom.age = "Tom",18
	//两个字段都写清楚的初始化
	bob := Person{age:25,name:"Bob"}
	//按照struct定义顺序初始化
	piter := Person{"Piter",30}

	tb_Older,tb_diff := Older(tom,bob)
	tp_Older,tp_diff := Older(tom,piter)
	bp_Older,bp_diff := Older(bob, piter)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",tom.name,bob.name,tb_Older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",tom.name,piter.name,tp_Older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",bob.name,piter.name,bp_Older.name, bp_diff)

}