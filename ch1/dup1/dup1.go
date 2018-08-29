// Dup1 exibe o texto de toda linha que aparece mais de
// uma vez na entrada-padrão, precedida por sua contagem.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Ctrl+D para sair
	// Criando um map de chaves string e valores int
	// Neste codigo utilizo o make para criar um map vazio
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTA: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
