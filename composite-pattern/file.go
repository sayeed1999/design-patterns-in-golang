package main

import "fmt"

// Leaf: implements Component via composition
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in file '%s'\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
