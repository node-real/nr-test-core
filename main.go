package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

// GetFoo comments I can find easely
func GetFoo() {
	// Comment I would like to access
	test := 1
	fmt.Println(test)
}

func main() {
	fset := token.NewFileSet() // positions are relative to fset
	d, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, f := range d {
		fmt.Println("package", k)
		for n, f := range f.Files {
			fmt.Printf("File name: %q\n", n)
			for i, c := range f.Comments {
				fmt.Printf("Comment Group %d\n", i)
				for i2, c1 := range c.List {
					fmt.Printf("Comment %d: Position: %d, Text: %q\n", i2, c1.Slash, c1.Text)
				}
			}
		}

	}
}
