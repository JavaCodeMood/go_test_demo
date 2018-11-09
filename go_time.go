package main

import (
	"time"
	"fmt"
)

func main(){
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)

	endDate := "2018-08-20"


	if endDate <= now {
		fmt.Println("截至时间不能小于当前时间")
	}else{
		fmt.Println("验证通过")
	}
}
