package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files on golang")

	//create file
	content := "lets girnd up in go Section"

	file, err := os.Create("./mylogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	ErrorHandler(err) // i use instead of this

	length, err := io.WriteString(file, content)
	ErrorHandler(err)

	fmt.Println("length is:", length)
	defer file.Close()

	readFile("./mylogofile.txt")

}

//read files

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	ErrorHandler(err)

	fmt.Println("Text data inside the file is: ", string(databyte))
}

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

/**
1.  os = interact with the operating system.
2.  io = read or write data
3.  when you read the file it not in the string format, Data always read into bytes format.
	eg- when you read the data on internat
4.  ioutil → older package of input/output helper functions; prefer modern

*/
