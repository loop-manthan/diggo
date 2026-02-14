package main

import (
	"diggo/format"
	"diggo/scan"
	"fmt"
	"os"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	dirs, err := scan.Dir(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "diggo: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(format.DirList(dirs))
}
