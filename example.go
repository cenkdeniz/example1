package main

import "fmt"

func main() {

	var sayı int
	var sayı1 int

	fmt.Println("1 ile 10 arsında bir sayı söyle")

	if (0 <= sayı) && (sayı <= 10) {
		fmt.Scanln(&sayı)
		fmt.Println("bu sayıyı hangi sayıyla çarpmamı istersin")
		fmt.Scanln(&sayı1)
		var sonuc int
		sonuc = carp(sayı, sayı1)
		fmt.Printf("bu iki sayının çarpımının sonucu %d dir", sonuc)
	} else if ( sayı <= 10)&& (sayı == 0) {
		fmt.Println("lütfen 1 ile 10 arasında bir raka söyleyin")
	}
}

func carp(x int, y int) (int) {
	return x * y
}
