package main

import (
	"math/rand"
	"time"
	"fmt"
)

//生成随机六位数
func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func main(){
	fmt.Println("随机数：",GetRandomString())
	fmt.Println("当前时间：",time.Now())
	fmt.Println("当前时间戳：",time.Now().UnixNano())
	fmt.Println(time.Date(2018,8,24,9,35,11,22,time.Local))
}
