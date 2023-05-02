package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Files(workingDir, output string, tempMap map[string][]string) {
	content := []byte{}
	fmt.Printf("Working directory: %s\n", workingDir)
	dir, err := os.ReadDir(workingDir)
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())
		return
	}
	for _, filename := range tempMap["-f"] {
		if ok, _ := regexp.MatchString(`\*`, filename); ok {
			match := "^" + strings.ReplaceAll(filename, "*", `[\w\.]*`) + "$"
			for _, f := range dir {
				if ok, _ := regexp.MatchString(match, f.Name()); ok {
					fileContent, err := os.ReadFile(filepath.Join(workingDir, f.Name()))
					if err != nil {
						fmt.Printf("Error %s\n", err.Error())
						return
					}
					fmt.Printf("%s: content added\n", f.Name())
					content = append(content, append(fileContent, []byte("\n")...)...)
				}
			}
		} else {
			fileContent, err := os.ReadFile(filepath.Join(workingDir, filename))
			if err != nil {
				fmt.Printf("Error %s\n", err.Error())
				return
			}
			fmt.Printf("%s: content added\n", filename)
			content = append(content, append(fileContent, []byte("\n")...)...)
		}
	}
	os.WriteFile(filepath.Join(workingDir, output), content, 0644)
	fmt.Println("Saved to", filepath.Join(workingDir, output))
}
