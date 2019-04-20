// Package graph creates a ItemGraph data structure for the Item type
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/willf/pad"
)

// GoDebBinaryStruct is the structere of json
type GoDebBinaryStruct struct {
	Binary         string `json:"binary"`
	XSGoImportPath string `json:"metadata_value"`
	Source         string `json:"source"`
}

// DepMap is the map of dependencies
type DepMap struct {
	deps map[string]DepMap
}

// LevelMap is a single level dependencies map
type LevelMap map[string]bool

// Global stdMap
var stdMap LevelMap
var finalMap DepMap

// GoDebBinariesURL is the url of binary list of go lang
const GoDebBinariesURL = "https://api.ftp-master.debian.org/binary/by_metadata/Go-Import-Path"

// Err is to log the error
func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetGoPath is to get $GOPATH environment variable
func GetGoPath() (string, error) {
	if os.Getenv("GOPATH") == "" {
		return "", errors.New("$GOPATH not set")
	}
	return os.Getenv("GOPATH"), nil
}

// GetProjectPath is to get full project path
func GetProjectPath(project string) (string, error) {
	path, e := GetGoPath()
	if e != nil {
		return "", e
	}
	return path + "/src/" + project, nil
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
func GetURLStatus(project string) (bool, error) {
	res, err := http.Get("http://" + project)
	if err != nil {
		return false, errors.New("Can't get " + "http://" + project)
	} else if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return true, nil
	}

	return false, errors.New("Can't get " + "http://" + project)
}

// HandleProject is to get project
func HandleProject(project string) error {
	projectPath, err := GetProjectPath(project)
	if err != nil {
		return err
	}
	// Project is already downloaded
	if FileExist(projectPath) {
		return nil
	}
	// Project don't exist, get it
	if status, err := GetURLStatus(project); status {
		if err != nil {
			return err
		}
		cmd := exec.Command("go", "get", project)
		_, err := cmd.CombinedOutput()
		if err != nil {
			return errors.New("Error in 'go get " + project + "'")
		}
	}

	return nil
}

// GetImports is to get first level dependencies of a project
func GetImports(project, importType string) ([]string, error) {
	cmd := exec.Command("go", "list", "-f", "'{{ join .Imports `\n` }}'", project)
	switch importType {
	case "deps":
		cmd = exec.Command("go", "list", "-f", "'{{ join .Deps `\n` }}'", project)
	case "std":
		cmd = exec.Command("go", "list", "std")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New("Error in getting 'go list" + importType + "'")
	}

	// Prepare the slice for Output
	libs := strings.Replace(string(out), "'", "", -1)
	slice := strings.Split(libs, "\n")
	return slice, nil
}

// SliceToMap is to convert a slice into a map
func SliceToMap(slice []string) LevelMap {
	m := make(LevelMap)
	for i := 0; i < len(slice); i++ {
		m[slice[i]] = true
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

// PrintDepMap is to print the DepMap
func PrintDepMap(m DepMap, i int) {
	for key, value := range m.deps {
		fmt.Println(pad.Left("- "+key, len(key)+(i+1)*2, " "))
		PrintDepMap(value, i+1)
	}
	i++
}

// SliceToDepMap is to convert a slice into a DepMap
func SliceToDepMap(slice []string) DepMap {
	var m DepMap
	m.deps = make(map[string]DepMap)
	for i := 0; i < len(slice); i++ {
		var dummy DepMap
		m.deps[slice[i]] = dummy
	}
	// Delete the empty elements
	delete(m.deps, "")
	return m
}

// GetDepTree is to get the recursive tree of dependencies
func GetDepTree(project string) (DepMap, error) {
	// Handle path, if it don't exist, get it.
	HandleProject(project)
	// Convert slice to map, since it's fast in searching.
	importSlice, err := GetImports(project, "imports")
	if err != nil {
		var m DepMap
		return m, err
	}
	importMap := SliceToMap(importSlice)
	// Remove standard libs from users libs
	importMap = RemoveMap(importMap, stdMap)
	// Convert importMap to slice again
	importSlice = MapToSlice(importMap)
	// Convert slice to DepMap now
	importDepMap := SliceToDepMap(importSlice)

	for key := range importDepMap.deps {
		importDepMap.deps[key], _ = GetDepTree(key)
	}

	return importDepMap, nil
}

// GetGoDebBinaries is to get the complete list of all the binaries packaged in debian
func GetGoDebBinaries() (map[string]string, error) {
	golangBinaries := make(map[string]string)
	resp, err := http.Get(GoDebBinariesURL)
	var pkgs []GoDebBinaryStruct

	if err != nil {
		return nil, fmt.Errorf("getting %q: %v", GoDebBinariesURL, err)
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		return nil, fmt.Errorf("unexpected HTTP status code: got %d, want %d", got, want)
	}

	if err := json.NewDecoder(resp.Body).Decode(&pkgs); err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		if !strings.HasSuffix(pkg.Binary, "-dev") {
			continue // skip -dbgsym packages etc.
		}
		for _, importPath := range strings.Split(pkg.XSGoImportPath, ",") {
			// XS-Go-Import-Path can be comma-separated and contain spaces.
			golangBinaries[strings.TrimSpace(importPath)] = pkg.Binary
		}
	}

	return golangBinaries, nil
}

func main() {
	fmt.Println("DebGoGraph Starting...")
	// Level is used for sub dependencies
	project := "github.com/ramantehlan/mateix"
	stdSlice, err := GetImports(project, "std")
	if err != nil {
		os.Exit(-1)
	}
	// Get standard libraries in map
	stdMap = SliceToMap(stdSlice)
	finalMap.deps = make(map[string]DepMap)

	fmt.Println("calculating...")

	m, _ := GetDepTree(project)
	PrintDepMap(m, 0)
}
