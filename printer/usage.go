package printer

import "fmt"

func Usage() {
	fmt.Println("Usage: Compressor [output] [options]")
	fmt.Println("Options:")
	fmt.Println("  -v, --version\t\tPrint version and exit")
	fmt.Println("  -h, --help\t\tPrint this help and exit")
	fmt.Println("  -f, --files\t\tFiles to compile")
}
