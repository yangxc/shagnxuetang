package main

import "fmt"

func main() {
	var name, age string
	fmt.Println("请输入您的姓名和年龄：")
	fmt.Scanln(&age, &name)
	fmt.Println(age, name)
}
