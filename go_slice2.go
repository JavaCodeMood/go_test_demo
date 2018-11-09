package main

import "fmt"

//可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]

func main(){
	//创建切片,并初始化
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}
	printSlice(numbers)

	//打印原始切片
	fmt.Println("numbers=",numbers)

	//打印子切片从索引1（包含）到索引4（不包含）
	fmt.Println("numbers[1:4] ==",numbers[1:4])

	//默认下限为0
	fmt.Println("numbers[:7]==",numbers[:7])

	//默认上限为len(s)
	fmt.Println("numbers[4:]==",numbers[4:])

	//使用make函数创建一个切片，并指定容量
	numbers1 := make([]int, 0 , 5)
	printSlice(numbers1)

	//打印子切片从索引0（包含）到索引2(不包含）
	number2 := numbers[:2]
	printSlice(number2)

	//打印子切片从索引2（包含）到索引5（不包含）
	number3 := numbers[2:5]
	printSlice(number3)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
