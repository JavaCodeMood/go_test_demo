package main

import "fmt"

/*
append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
*/

func main(){
	//创建切片
	var numbers []int
	printSlice(numbers)

	//向切片追加元素
	numbers = append(numbers,1,2,3,4)
	printSlice(numbers)

	//创建切片numbers1是之前切片的两倍容量
	numbers1 := make([]int,len(numbers),cap(numbers)*2)
	printSlice(numbers1)

	//拷贝numbers的内容到numbers1
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
