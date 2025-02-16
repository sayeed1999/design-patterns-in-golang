package main

func main() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{name: "folder1"}
	folder1.add(file1)
	folder1.add(file2)
	folder1.add(file3)

	folder2 := &Folder{name: "folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
