package main

import (
	"fmt"
)

var c, d int = 1, 2
var e, f = 123, "hello"
//!!注意：下行这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
func main() {

	var v21 int32  //被定义初始化为0
	var v22 int = 2
	var v23 = 3 //被自动识别为int类型
	v24 := 4    //简易声明&定义的方式
	v21 = int32(v23) //强制转换

	g, h := 123, "hello"
	fmt.Println("v21 is", v21) //v21被赋新值
	fmt.Println("v22 is", v22)
	fmt.Println("v23 is", v23)
	fmt.Println("v24 is", v24)
	fmt.Println(c, d, e, f, g, h)
}





