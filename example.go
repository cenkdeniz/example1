package main

import

"fmt"

func topla( x int, y int)(int) {

	return x + y

}

func main(){
	var number int
	var number2 int
	var total int
	fmt.Println("bir sayı söle ")
	fmt.Scanln(&number)
	fmt.Println("ikinci bir sayı daha")
	fmt.Scanln(&number2)
	total = topla(number, number2)
	fmt.Println(total)


}