package main

import (
	"fmt"
	"io/ioutil"
)


func readFile1 (file string) {

	content, error := ioutil.ReadFile(file)

	if error != nil {
		fmt.Println("open file [%s] error", file)
		fmt.Println(error)
	}

	fmt.Print(content)

	var str string
	str = string(content)

	fmt.Print(str)
}


func writeFile1 (file string, content string) {
	
	var str string
	str = "write to file:"
	str += content

	tmp := []byte(str)

	error := ioutil.WriteFile(file, tmp, 0644)
	// delete file content and write to file 
	if error != nil {
		panic(error)
	}

}


func main() {
	fmt.Println("hello world!")

	var file string
	file = "test.log"
	readFile1(file)

	writeFile1(file,"zhenhuli")
}
