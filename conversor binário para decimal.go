package main

import (
	"fmt"
	"strconv"
)

//  programa 1

func main(){
	for{
	fmt.Println("Conversor binario para decimal\n")

	var binarios string
	fmt.Println("Insira um numero binario de até 8 dígito: ")
	fmt.Scanln(&binarios)
	fmt.Println(len(binarios))
	if len(binarios) > 8{
		fmt.Println("maximo 8")
		break
	} else{
		output, err := strconv.ParseInt(binarios,2,  64)
		if err != nil{
			fmt.Println(err)
			fmt.Println("Insira apenas 0 e 1\n")
			return
		}
		fmt.Printf("Output %d\n", output)

	}
	}
	//output, err := strconv.ParseInt(binarios,2,  64)
	//if err != nil{
		//fmt.Println(err)
		//fmt.Println("Insira apenas 0 e 1\n")
		//return
	//}
	//if len(binarios) > 8{
	//	fmt.Println("Máximo de 8")

	//}

	//fmt.Printf("Output %d\n", output)
}
