package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file youtube.com"

	file, err := os.Create("./myTestfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	defer file.Close() // reccomended so that this code will be executed at the end

	readFile("./myTestfile.txt")
}

func readFile(filename string) {
	byte_data, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Test data inside the file is \n", string(byte_data))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
