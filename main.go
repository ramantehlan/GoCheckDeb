package main

import (
	"fmt"
	"log"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	log.Print(aurora.Bold(aurora.Green("[DebGoGraph Starting]")))

	project := "github.com/zaquestion/lab"
	returnType := "graph"
	debFilter := true
	displayAll := true

	// No sub-command
	if len(os.Args) < 1 {
		fmt.Println("Sub command required")
		os.Exit(1)
	}

	fmt.Printf("\nGolang package: %s\n", aurora.Bold(project))
	fmt.Printf("Output type: %s\n", aurora.Bold(returnType))
	fmt.Printf("Deb filter: %v\n", aurora.Bold(debFilter))
	fmt.Printf("Display all deb (false for only main package): %v\n\n", aurora.Bold(displayAll))

	fmt.Printf("Fetching dependencies of %s | It may take a while.\n\n", aurora.Bold(aurora.BrightBlue(project)))

	// List | Graph | Tree
	m2, err := gocheckdeb.GetDep(project, returnType)
	if err != nil {
		fmt.Println(err)
	}
	// DepMaps | DebFilder - display deb unpacked | displayAll - Display all unpacked or just head | inc start
	gocheckdeb.PrintDep(m2, debFilter, displayAll, 0)

	fmt.Println("")
	log.Print(aurora.Bold(aurora.Green("[DebGoGraph Ending]")))
}
