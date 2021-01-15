package main

import (
	"fmt"

)	

func main() {
	var operacao string
	var number1, number2 int

	fmt.Println("Número 1:\n")
	fmt.Scanln(&number1)
	fmt.Println("Número 2:")
	fmt.Scanln(&number2)

	fmt.Println("Selecione a operação (+,-,/,*):")
	fmt.Scanln(&operacao)
	output :=0
	switch operacao{
	case "+":
		output = number1 + number2
		break
	case "-":
		output = number1 - number2
		break
	case "*":
		output = number1 * number2
		break
	case "/":
		output = number1 / number2
		break

	default:
		fmt.Println("Operação inválida")
	}

	fmt.Printf("%d %s %d = %d\n", number1, operacao, number2, output)

}
