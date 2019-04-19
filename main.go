// Package graph creates a ItemGraph data structure for the Item type
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/willf/pad"
)

// DepMaps is the recursive map for multi layer dependencies
type DepMaps struct {
	graph map[string]DepMaps
}

// LevelMap is a single level dependencies map
type LevelMap map[string]bool

// Global stdMap
var stdMap LevelMap
var finalMap DepMaps
var levelSlice [][]string

// Err is to log the error
func Err(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Fatal(msg)
		os.Exit(1)
	}
}

// GetGoPath is to get $GOPATH environment variable
func GetGoPath() string {
	if os.Getenv("GOPATH") == "" {
		Err(nil, "GOPATH Not set")
	}
	return os.Getenv("GOPATH")
}

// GetProjectPath is to get full project path
func GetProjectPath(project string) string {
	return GetGoPath() + "/src/" + project
}

// FileExist is check if file exist
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// GetURLStatus is to get the status of a package
func GetURLStatus(project string) bool {
	res, err := http.Get("http://" + project)
	fmt.Println(project)
	Err(err, "GetURLStatus Error")
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return true
	}
	return false
}

// HandleProject is to get project
func HandleProject(project string) {
	if !FileExist(GetProjectPath(project)) {
		if GetURLStatus(project) {
			log.Printf("Starting: go get %s", project)
			cmd := exec.Command("go", "get", project)
			_, err := cmd.CombinedOutput()
			Err(err, "Go get error")
			log.Printf("Done: go get %s", project)
		}
	}
}

// GetStd is to get a slice of standard libs
func GetStd() []string {
	cmd := exec.Command("go", "list", "std")
	out, err := cmd.CombinedOutput()
	Err(err, "Go List STD error")
	libs := strings.Replace(string(out), "'", "", -1)
	slice := strings.Split(libs, "\n")
	return slice
}

// GetImports is to get first level dependencies of a project
func GetImports(project string) []string {
	cmd := exec.Command("go", "list", "-f", "'{{ join .Imports `\n` }}'", project)
	out, err := cmd.CombinedOutput()
	Err(err, "go list error")
	libs := strings.Replace(string(out), "'", "", -1)
	slice := strings.Split(libs, "\n")
	return slice
}

// SliceToDepMap is to convert a slice into a DepMap
func SliceToDepMap(slice []string) DepMaps {
	var m DepMaps
	m.graph = make(map[string]DepMaps)
	for i := 0; i < len(slice); i++ {
		var dummy DepMaps
		m.graph[slice[i]] = dummy
	}
	// Delete the empty elements
	delete(m.graph, "")
	return m
}

// SliceToLevelMap is to convert a slice into a map
func SliceToLevelMap(slice []string) LevelMap {
	m := make(LevelMap)
	for i := 0; i < len(slice); i++ {
		m[slice[i]] = false
	}
	// Delete the empty elements
	delete(m, "")
	return m
}

// RemoveMap is to remove key of mainMap which are present in needleMap
func RemoveMap(mainMap, needleMap LevelMap) LevelMap {
	for key := range mainMap {
		if _, ok := needleMap[key]; ok {
			delete(mainMap, key)
		}
	}
	return mainMap
}

// LevelMapToSlice is to convert a LevelMap into slice
func LevelMapToSlice(m LevelMap) []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// GetDep is
func GetDep(level []string) {
	project := level[len(level)-1]
	// Handle path, if it don't exist, get it.
	HandleProject(project)
	// Convert slice to map, since it's fast in searching.
	importMap := SliceToLevelMap(GetImports(project))
	// Remove standard libs from users libs
	importMap = RemoveMap(importMap, stdMap)
	// Convert importMap to slice again
	importSlice := LevelMapToSlice(importMap)
	// Convert slice to DepMaps now
	importDepMaps := SliceToDepMap(importSlice)

	for key := range importDepMaps.graph {
		level = append(level, key)
		levelSlice = append(levelSlice, level)
		GetDep(level)
	}
}

// InsertDep is to insert dependencies in a recursive map
func InsertDep(slice []string) DepMaps {
	var m DepMaps
	m.graph = make(map[string]DepMaps)
	if len(slice) > 0 {
		m.graph[slice[0]] = InsertDep(slice[1:])
	}
	return m
}

// PrintDepMaps is to print the DepMaps
func PrintDepMaps(m DepMaps, i int) {
	for key, value := range m.graph {
		fmt.Println(pad.Left("- "+key, len(key) + (i+1)*2, " "))
		PrintDepMaps(value, i+1)
	}
	i++
}

// MergeDepMaps is to merge two DepMaps 
func MergeDepMaps(m1, m2 DepMaps) DepMaps {
	return m1 
}


// Maybe we should go get the main project first
// like mannually do the go get github.com/zyedidia/micro
// and the use the inner folder where the list will work

func main() {
	fmt.Println("DebGoGraph Starting...")
	// Final map of all the dependencies
	finalMap.graph = make(map[string]DepMaps)
	// Level is used for sub dependencies
	level := []string{"github.com/ramantehlan/mateix"}
	// Get standard libraries in map
	stdMap = SliceToLevelMap(GetStd())


	fmt.Println("calculating...")
	GetDep(level)
	fmt.Println("[Output]")
	finalMap = InsertDep(levelSlice[4])
	fmt.Println(levelSlice)
	PrintDepMaps(finalMap,0)
}
