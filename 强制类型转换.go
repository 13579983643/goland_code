package main
import "fmt"

func main(){
	var sum int = 17
	var count int = 5
	var mean float32      //没有double类型，所以float64就是double,  float 4字节 有效位6-7 ；double 8字节 有效位15-16
	mean = float32(sum)/float32(count)
	fmt.Print("mean的值为：%f\n",mean)
}
