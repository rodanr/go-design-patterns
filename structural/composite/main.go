package main

import "fmt"

type Component interface {
	search(string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword '%s' in file: %s\n", keyword, f.name)
}

type Folder struct {
	name       string
	components []Component
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword '%s' in folder: %s\n", keyword, f.name)
	for _, component := range f.components {
		component.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

func main() {
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}
	file3 := &File{name: "file3.txt"}

	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("keyword")
}
