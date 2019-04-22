package main

import (
	"fmt"
	"log"

	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	log.Print("DebGoGraph Starting...")
	// Level is used for sub dependencies
	//project := "github.com/ramantehlan/mateix"
	//project := "github.com/zyedidia/micro/cmd/micro"
	project := "github.com/zaquestion/lab"

	fmt.Println("calculating...")

	// List | Graph | Tree
	m2, err := gocheckdeb.GetDep(project, "graph")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---Border---")
	gocheckdeb.PrintDep(m2, true, 0)

	log.Print("DebGoGraph Ending...")
}
