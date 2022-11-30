package main

import "fmt"

type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}
func closeStream(stream Stream) {
	stream.close()
}

type File struct {
	text string
}

type Folder struct{}

func (f *File) read() string {
	return f.text
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}
func (f *Folder) close() {
	fmt.Println("Папка закрыта")
}

func StreamOperations() {
	myFile := &File{}
	myFolder := &Folder{}

	writeToStream(myFile, "hello")
	closeStream(myFile)
	myFolder.close()

}
