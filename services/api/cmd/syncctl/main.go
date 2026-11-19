package main

import (
	"fmt"
	"os"
)

const version = "0.1.0-f0"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("syncctl %s\n", version)
		return
	}

	fmt.Printf("syncctl %s — CLI operasional Config-as-Code SINERGI (akan diimplementasikan pada Fase F5)\n", version)
}
