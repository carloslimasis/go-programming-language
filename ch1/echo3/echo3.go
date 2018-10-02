// Echo3 show command line arguments
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
