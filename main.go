package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	log.Print("DebGoGraph Starting...")

	// No sub-command
	if len(os.Args) < 2 {
		fmt.Println("Sub command required")
		os.Exit(1)
	}

	// Level is used for sub dependencies
	//project := "github.com/ramantehlan/mateix"
	//project := "github.com/zyedidia/micro/cmd/micro"

	//./GoCheckDeb git... -tree -
	project := "github.com/zaquestion/lab"

	fmt.Println("calculating...")

	// List | Graph | Tree
	m2, err := gocheckdeb.GetDep(project, "tree")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---Border---")
	// DepMaps | DebFilder - display deb unpacked | displayAll - Display all unpacked or just head | inc start
	gocheckdeb.PrintDep(m2, true, true, 0)
	log.Print("DebGoGraph Ending...")
}
