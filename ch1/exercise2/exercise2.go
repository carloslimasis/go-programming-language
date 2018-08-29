// Exercise2 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Program: %s", os.Args[0])

	for i, arg := range os.Args[1:] {
		fmt.Printf("\nParam[%d]: %s", i, arg)
	}
}
