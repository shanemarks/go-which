package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

type RunMode int
type Action = func(fileName string)

const (
	ModeUnknown RunMode = iota
	ModeNoOp
	ModeSingle
	ModeAll
)

func main() {

	allPointer := flag.Bool("a", false, "Returns all matches")
	flag.Parse()
	//any arguments that aren't flags come here
	arguments := flag.Args()
	target := ""
	if len(arguments) > 0 {
		target = arguments[0]
	}

	// Instead of having if/else of switches based on glads, we just calculate a mode and
	// invoke a function based on that mode.
	runMode := calculateRunMode(len(os.Args) >= 2, *allPointer)
	actions := buildActions()

	action, exists := actions[runMode]

	if exists {
		action(target)
	}

}

func calculateRunMode(hasArgs bool, allMode bool) RunMode {

	runMode := ModeUnknown

	if !hasArgs {
		runMode = ModeNoOp
	} else if allMode {
		runMode = ModeAll
	} else {
		runMode = ModeSingle
	}

	return runMode

}

func buildActions() map[RunMode]Action {

	actions := make(map[RunMode]Action)

	actions[ModeUnknown] = func(fileName string) {
		fmt.Println("If we hit an unknown state, something is wrong with the logic")
	}

	actions[ModeNoOp] = func(fileName string) {}

	actions[ModeSingle] = func(fileName string) {
		results := which(fileName, os.Getenv("PATH"), FileExists)
		if len(results) == 0 {
			fmt.Printf("%s not found", fileName)
		} else {
			fmt.Println(results[0])
		}
	}

	actions[ModeAll] = func(fileName string) {
		results := which(fileName, os.Getenv("PATH"), FileExists)
		if len(results) == 0 {
			fmt.Printf("%s not found", fileName)
		} else {
			for _, file := range results {
				fmt.Println(file)
			}
		}
	}
	return actions
}

// We pass in a delegate for the file existance check so its easy to mock out in our tests.
func which(bin string, pathsToParse string, fileExistsDelegate FileExistsDelegate) []string {

	paths := Map(strings.Split(pathsToParse, ":"), func(s string) string {
		return path.Join(s, bin)
	})

	matches := []string{}
	for _, singlePath := range paths {
		if fileExistsDelegate(singlePath) {
			matches = append(matches, singlePath)
		}
	}
	return matches
}
