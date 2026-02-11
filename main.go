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

	r, _ := rand.Int(rand.Reader, big.NewInt(10))

	// slice[1] = int(r.Int64())

	for i := 0; size <= len(slice)-1; i++ {
		slice[i] = int(r.Int64())
	}

	fmt.Println(slice)
}