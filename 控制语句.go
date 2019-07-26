package main
import "fmt"

/*
func main(){
	//基本的例子
	if 7%2 == 0{
	    fmt.Print("7 is even")
	}else{
		fmt.Println("7 is odd")
	}
	//只有if条件的情况
	if 8%4 == 0{
		fmt.Println("8 is divisible by 4")
	}
	//if 条件可以包含一个初始化表达式，这个表达式是这个判断结构的局部变量
	if num :=9; num <0 {
		fmt.Println(num, "is negative")
	}else if num <10{
		fmt.Println(num, "has 1 digit")
	} else{
		fmt.Println(num, "has multiple digits ")
	}
}
*/
/*
func main(){
	sum :=0
	for i :=0;i<10;i++ {
		sum += i
	}
	fmt.Println(sum)
}
*/
func main(){
	arr :=[...]int{6,7,8}
	for i , v := range arr {
		fmt.Println(i, v)
	}
}






