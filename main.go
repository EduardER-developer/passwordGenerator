package main

import(
	"fmt"
	"os"
	"math/big"
	"crypto/rand"
)

func main()  {
	// password := 0
	size := 0
	// var lenghtPassword []int

	fmt.Println("Какая длина пароля? Введите:")
	fmt.Scanln(&size)

	slice := make([]int, size);
	if len(slice) == 0 {
		fmt.Println("Вы ввели длину 0. Введите более 1")
		os.Exit(0)
	}
	// slice[1] = int(r.Int64())

	// if (size == len(slice)){
	// 	fmt.Printf("успешно")
	// }

	for i := 0; len(slice) > i; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(10))
		slice[i] = int(r.Int64())
	}

	for _, v := range slice {
		fmt.Print(v) 
	}
}