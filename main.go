package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb"
)

func main() {
	log.Print(aurora.Bold(aurora.Green("[DebGoGraph Starting]")))

	project := "github.com/zaquestion/lab"
	returnType := "graph"
	debFilter := true
	displayAll := true

	projectFlag := flag.String("project", "github.com/zaquestion/lab", "(String) Return type can be Graph, Tree and List")
	returnFlag := flag.String("return", "graph", "(String) Return type can be Graph, Tree and List")

	flag.Parse()
	project = *projectFlag
	returnType = *returnFlag

	fmt.Printf("\nGolang package: %s\n", aurora.Bold(project))
	fmt.Printf("Output type: %s\n", aurora.Bold(returnType))
	fmt.Printf("Deb filter: %v\n", aurora.Bold(debFilter))
	fmt.Printf("Display all deb (false for only main package): %v\n\n", aurora.Bold(displayAll))

	fmt.Printf("Fetching dependencies of %s | It may take a while.\n\n", aurora.Bold(aurora.BrightBlue(project)))

	// List | Graph | Tree
	m, err := gocheckdeb.GetDep(project, "imports", returnType)
	if err != nil {
		fmt.Println(err)
	}
	m2, err := gocheckdeb.GetDep(project, "test", returnType)
	if err != nil {
		fmt.Println(err)
	}
	// DepMaps | DebFilder - display deb unpacked | displayAll - Display all unpacked or just head | inc start
	fmt.Println(aurora.Bold(aurora.BrightBlue("--Project Dependencies--")))
	gocheckdeb.PrintDep(m, debFilter, displayAll, 0)
	fmt.Println(aurora.Bold(aurora.BrightBlue("--Test Dependencies--")))
	gocheckdeb.PrintDep(m2, debFilter, displayAll, 0)

	fmt.Println("")
	log.Print(aurora.Bold(aurora.Green("[DebGoGraph Ending]")))
}
