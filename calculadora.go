package platzigocalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	a := pasaAEntero(entradaLimpia[0])
	b := pasaAEntero(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println("Suma")
		return a + b
	case "-":
		fmt.Println("Resta")
		return a - b
	case "*":
		fmt.Println("Multiplica")
		return a * b
	case "/":
		fmt.Println("Divide")
		return a / b
	default:
		fmt.Println("Operacion no permitida")
		return 0
	}
}

func LeerTeclado() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func pasaAEntero(texto string) int {
	operador, _ := strconv.Atoi(texto)
	return operador
}

// Como es un modulo no usa main

// func main() {
// 	operacion := leerTeclado()
// 	operador := leerTeclado()
// 	c := calc{}
// 	fmt.Println("Resultado: ", c.operate(operacion, operador))
// }
