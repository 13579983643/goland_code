package main
import(
	"fmt"
)
/*
func main(){
	i:=2
	switch i {
		case 0 :
			fmt.Printf("0")
		case 1:
6			fmt.Printf("1")
		case 2:
			fallthrough  //fallthrough会强制执行后面的case代码
		case 3:
			fmt.Printf("3")
		default:
			fmt.Printf("Default")
	}
}
*/


func main(){
	Num := 6
	var grade string = "B"
	var marks int  = 90
	switch{
		case 0 <= Num && Num <=3:
			fmt.Println("0-3")
		case 4<= Num && Num <= 6:
			fmt.Println("4-6")
		case 7<= Num && Num <= 9:
			fmt.Println("7-9")
	}
	 switch marks {
	 	case 90: grade ="A"
	 	case 80: grade ="B"
	 	case 50,60,70: grade = "C"
	 	default: grade = "D"
	 }
	switch {
		case grade == "A":
			fmt.Println("Excellent!")
		case grade == "B",grade == "C":
			fmt.Println("Well done" )
		case grade == "D" :
			fmt.Println("You passed" )
		case grade == "F":
			fmt.Println("Better try again" )
		default:
			fmt.Println("Invalid grade" );
	}
	fmt.Println("Your grade is :", grade );
}



















