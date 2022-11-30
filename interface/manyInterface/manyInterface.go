package manyInterface

import (
	"fmt"
)

type Reader interface {
	read()
}
type Writer interface {
	write(string)
}
type ReaderWriter interface {
	Reader
	Writer
}

func writeToStream(writer ReaderWriter, text string) {
	writer.write(text)
}
func readFromStream(reader ReaderWriter) {
	reader.read()
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}
func ManyInterface() {
	myFile := &File{}
	writeToStream(myFile, "hello manyInterface")
	readFromStream(myFile)
	writeToStream(myFile, "write manyInterface")
	readFromStream(myFile)
}
