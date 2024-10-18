package simple

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) (*File, func()) { // The key to initialize a cleanup function is to RETURNS A FUNCTION
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}

func (f *File) Close() {
	fmt.Println("File " + f.Name + " has been closed")
}
