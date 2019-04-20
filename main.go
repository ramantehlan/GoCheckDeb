package main

import (
	"fmt"

	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	fmt.Println("DebGoGraph Starting...")
	// Level is used for sub dependencies
	project := "github.com/ramantehlan/mateix"

	fmt.Println("calculating...")

	m, _ := gocheckdeb.GetDepTree(project)
	gocheckdeb.PrintDepMap(m, 0)
}
