package main

import "fmt"

// Composite: implements Component via composition
type Folder struct {
	components []Component // can be Folder or File
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("\nSearching recursively for keyword '%s' in folder '%s'\n..\n", keyword, f.name)

	for _, component := range f.components {
		// can be file or folder
		component.search(keyword)
	}
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}
