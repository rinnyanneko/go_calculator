package main
import "fmt"

func main(){
	fmt.Println("Author: rinnyanneko")
	fmt.Println("This program is for practicing Golang")
	fmt.Println("There might be many bugs in this program, Please don't use it for production")
	fmt.Println("Thank you")
	fmt.Println("=====================================")
	var run bool = true
	var choice int
	var num1, num2 float64
	var num1Ptr, num2Ptr *float64
	for run {
		fmt.Println("Please choose the calculate operation")
		fmt.Println("1. Plus, 2. Minus, 3. Multiply, 4. Divide, 5. Exit")
		fmt.Scanln(&choice)
		if choice == 5 {
			run = false
			break
		}else{
			fmt.Println("Please input the first number")
			fmt.Scanln(&num1)
			fmt.Println("Please input the second number")
			fmt.Scanln(&num2)
			numPtr1=&num1
			numPtr2=&num2
			calculate(choice, num1Ptr, num2Ptr)
		}
}
func calculate(choice int, a *float64, b *float64) {
	switch choice {
		case 1:
			*a+=*b
		case 2:
			*a-=*b
		case 3:
			*a*=*b
		case 4:
			*a/=*b
		default:
			fmt.Println("Please choose the correct operation")
	}
}