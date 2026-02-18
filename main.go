package main

import (
	"diggo/format"
	"diggo/scan"
	"flag"
	"fmt"
	"os"
)

func main() {
	depth := flag.Int("depth", -1, "max depth (0=root only, 1=root+children, ...); default full tree")
	flag.Parse()

	root := "."
	if flag.NArg() > 0 {
		root = flag.Arg(0)
	}

	dirs, err := scan.Dir(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "diggo: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(format.Tree(dirs, root, *depth))
}
