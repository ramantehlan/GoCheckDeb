package main

import (
	"fmt"

	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	fmt.Println("DebGoGraph Starting...")
	// Level is used for sub dependencies
	//project := "github.com/ramantehlan/mateix"
	//project := "github.com/zyedidia/micro"
	project := "github.com/zaquestion/lab"

	fmt.Println("calculating...")

	gocheckdeb.Init()
	m, _ := gocheckdeb.GetDepTree(project, true)
	gocheckdeb.PrintDepMap(m, 0)
}
