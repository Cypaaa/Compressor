package main

import (
	"Compressor/printer"
	"Compressor/utils"
	"fmt"
	"os"
	"regexp"
	"time"
)

// get args from command line, if arg is -v or --version, print version
func main() {
	// get args from command line
	// if arg is -v or --version, print version
	workingDir, _ := os.Getwd()
	args := os.Args[1:]
	output := fmt.Sprintf("compressed.output.%v.txt", time.Now().UnixMilli())

	options := []string{"-v", "--version", "-h", "--help", "-f", "--files"}

	tempMap := make(map[string][]string)
	tempFlag := ""
	for i, arg := range args {
		// output name
		if ok, _ := regexp.MatchString(`^[\w\.]*$`, arg); i == 0 && ok {
			output = arg
			// options
		} else if utils.Contains(options, arg) {
			tempFlag = arg
			tempMap[arg] = []string{}
			// files
		} else if ok, _ := regexp.MatchString(`^[\w\.\*]*$`, arg); tempFlag != "" && ok {
			tempMap[tempFlag] = append(tempMap[tempFlag], arg)
			// invalid
		} else {
			printer.Usage()
			fmt.Printf("Invalid argument: %s\n", arg)
			return
		}
	}

	if len(tempMap) == 0 {
		printer.Usage()
		return
	}

	_, okversion := tempMap["--version"]
	_, okhelp := tempMap["--help"]
	_, okfiles := tempMap["--files"]
	if _, okv := tempMap["-v"]; okversion || okv {
		printer.Version()
	} else if _, okh := tempMap["-h"]; okhelp || okh {
		printer.Usage()
	} else if _, okf := tempMap["-f"]; okfiles || okf {
		if len(tempMap["-f"]) == 0 {
			printer.Usage()
			fmt.Println("No files")
			return
		} else {
			printer.Files(workingDir, output, tempMap)
		}
	} else {
		printer.Usage()
		fmt.Println("Invalid arguments")
		return
	}
}
