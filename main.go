package main

import (
	"fmt"

	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	fmt.Println("DebGoGraph Starting...")
	// Level is used for sub dependencies
	//project := "github.com/ramantehlan/mateix"
	project := "github.com/zyedidia/micro/cmd/micro"
	//project := "github.com/zaquestion/lab"

	fmt.Println("calculating...")

	// List | Graph | Tree
	m2, _ := gocheckdeb.GetDep(project, "graph")
	fmt.Println("---Border---")
	gocheckdeb.PrintDep(m2, true, 0)
}
