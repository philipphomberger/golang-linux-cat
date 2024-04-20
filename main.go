package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("%s\n", cat(flag.Arg(i)))
	}
}
