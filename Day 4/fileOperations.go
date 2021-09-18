package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err1 := ioutil.ReadFile("ourFile.txt")
	checkError(err1)
	fmt.Println("Content of the file is : ", string(content))

	f, er := os.Open("ourFile.txt")
	defer f.Close()
	checkError(er)
	bucket := make([]byte, 10)
	bytesRead, e := f.Read(bucket)
	checkError(e)
	fmt.Println("Content of the file(limited) : ", string(bucket[:bytesRead]))
	fmt.Println("Bytes read : ", bytesRead)

	//writeFile

	data := []byte("hello to go")
	err := ioutil.WriteFile("ourFile1.txt", data, 1)
	checkError(err)
	fmt.Println("Done Writing")

	f, er1 := os.Create("ourFile2.txt")
	defer f.Close()
	checkError(er1)
	bytesWritter, e := f.WriteString("Hello")
	checkError(e)
	fmt.Println("Bytes writtens are : ", bytesWritter)
	fmt.Println("Done Writing")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
