package main

import (
	"os"
)

func cat(filename string) (file string) {
	b, _ := os.ReadFile(filename)
	return (string(b))
}
