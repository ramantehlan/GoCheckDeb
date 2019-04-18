// Package graph creates a ItemGraph data structure for the Item type
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// DepMaps is the recursive map for multi layer dependencies
type DepMaps struct {
	graph map[string]DepMaps
}

// LevelMap is a single level dependencies map
type LevelMap map[string]bool

// Err is to log the error
func Err(msg error) {
	if msg != nil {
		log.Fatal(msg)
		os.Exit(1)
	}
}

// GetGoPath is to get $GOPATH environment variable
func GetGoPath() string {
	if os.Getenv("GOPATH") == "" {
		Err(errors.New("GOPATH Not set"))
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
	Err(err)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return true
	}
	return false
}

// HandleProject is to get project
func HandleProject(project string) {
	if !FileExist(GetProjectPath(project)) {
		if GetURLStatus(project) {
			cmd := exec.Command("go", "get", project)
			out, err := cmd.CombinedOutput()
			Err(err)
			log.Printf("%s", out)
		}
	}
}

// GetStd is to get a slice of standard libs
func GetStd() []string {
	cmd := exec.Command("go", "list", "std")
	out, err := cmd.CombinedOutput()
	Err(err)
	libs := strings.Replace(string(out), "'", "", -1)
	slice := strings.Split(libs, "\n")
	return slice
}

// GetImports is to get first level dependencies of a project
func GetImports(project string) []string {
	cmd := exec.Command("go", "list", "-f", "'{{ join .Deps `\n` }}'", project)
	out, err := cmd.CombinedOutput()
	Err(err)
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

// SliceToMap is to convert a slice into a map
func SliceToMap(slice []string) LevelMap {
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

// MapToSlice is to convert a LevelMap into slice
func MapToSlice(m LevelMap) []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	log.Print("DebGoGraph Starting...")
	// Get the project
	var project = "github.com/ramantehlan/mateix"

	// Handle path, if it don't exist, get it.
	HandleProject(project)
	// We convert slice to map, since it's fast in searching.
	importMap := SliceToMap(GetImports(project))
	//stdMap := SliceToMap(GetStd())

	fmt.Println(importMap)

	// loop
	// Get deps of this project
	// check if they are present in the debian
	// if not present add them to list // seperate list if present
	// repeat the loop unless goal reached
}
