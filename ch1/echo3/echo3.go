// Echo3 exibe seus argumentos de linha de comando
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%d elapsed\n", time.Since(start).Nanoseconds())
}
