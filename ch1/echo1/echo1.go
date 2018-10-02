// Echo1 show command line arguments
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	fmt.Printf("%d elapsed\n", time.Since(start).Nanoseconds())
}
