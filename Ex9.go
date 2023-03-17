package main

import "fmt"

func main() {

	var (
		preço float64
	)
	fmt.Print("Qual o preço do seu produto:")
	fmt.Scan(&preço)

	desconto := (preço*10)/100 - preço

	fmt.Println("O preço do seu produto sera", desconto)
}
