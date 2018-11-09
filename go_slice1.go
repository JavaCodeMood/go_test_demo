package main

import "fmt"

/*
切片长度不固定，可以追加元素，在追加时可能使切片的容量增大。
可以使用make函数来创建切片
var slice1[]type = make([]type, len)
或者
slice1 := make([]type,len)
也可以为切片指定容量，其中cap为可选参数  make([]type,len,cap)

初始化切片： slice1 := [] int {1,2,3}
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
*/

func main(){
	//定义一个切片,其中3是切片长度，5是切片容量
	var numbers = make([]int,3,5)
	printSlice(numbers)

	//一个切片在未初始化之前默认为 nil，长度为 0
	var slice1 []int
	printSlice(slice1)

	if (slice1 == nil){
		fmt.Println("切片slice1是空的")
	}
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)

}