package main

import(
	"fmt"
	"os"
	"math/big"
	"crypto/rand"
)

func main()  {
	size := 0
	
	fmt.Println("Какая длина пароля? Введите:")
	fmt.Scanln(&size)

	slice := make([]int, size);
	if len(slice) == 0 {
		fmt.Println("Вы ввели длину 0. Введите более 1")
		os.Exit(0)
	}

	for i := 0; len(slice) > i; i++ {
		x, _ := rand.Int(rand.Reader, big.NewInt(5347852498))	

		if x.Int64() % 2 == 0 {
			r, _ := rand.Int(rand.Reader, big.NewInt(10))

			slice[i] = int(r.Int64())
		} else {
			t, _ := rand.Int(rand.Reader, big.NewInt(94))		
			slice[i] = int(t.Int64())+33
		}
	}
	
	for _, v := range slice {
		if (v < 10) {
			fmt.Print(v) 
		} else {
			ch := rune(v)
			fmt.Printf("%c", ch)
		}
	}	
}