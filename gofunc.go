package main

import (
	"fmt"
	"math"
)

/*
Go 语言函数定义格式如下：

func function_name( [parameter list] ) [return_types] {
   函数体
}
函数定义解析：

func：函数由 func 开始声明
function_name：函数名称，函数名和参数列表一起构成了函数签名。
parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。
                参数是可选的，也就是说函数也可以不包含参数。
return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
函数体：函数定义的代码集合。
*/

func main(){
	//定义局部变量
	var a int = 100
	var b int = 40
	var c int

	//调用函数返回最大值
	c = max(a,b)
	fmt.Println("最大值：%d\n", c)

	str1,str2 := swap("abc","bdf")
	fmt.Println(str1,str2)

	//Go 语言可以很灵活的创建函数，并作为值使用。
	//声明函数变量
	getSquareRoot := func(x float64) (float64,float64){
		return math.Sqrt(x),math.Abs(x)
	}
	fmt.Println(getSquareRoot(9))
}

//函数返回两个数的最大值
func max(num1,num2 int) int {
	//定义局部变量
	var result int

	if (num1 > num2){
		result = num1
	}else if(num1 < num2){
		result = num2
	}else{
		result = num2
	}
	return result
}

//函数返回多个值
func swap(x,y string)(string,string){
	return y,x
}