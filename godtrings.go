package main

import (
	"strings"
	"fmt"
)

func main(){
	var users string ="18295514401,18295514402,18295514403"

	us := strings.Split(users,",")
	u := ""

	for _,u = range us {
		fmt.Println(u)
	}
}
