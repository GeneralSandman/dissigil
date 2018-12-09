package main

import (
	"fmt"
	"io/ioutil"
    "bytes"
    "encoding/gob"
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

func persist () []byte{
    buffer := new(bytes.Buffer)
    encoder := gob.NewEncoder(buffer)

    a := 1
    b := 2
    c := 3
    d := 4
    encoder.Encode(a)
    encoder.Encode(b)
    encoder.Encode(c)
    encoder.Encode(d)

    data := buffer.Bytes()

    fmt.Printf("%s\n", data)
    return data
}

func readpersist (data [] byte) {
    buffer := bytes.NewBuffer(data)
    decoder := gob.NewDecoder(buffer)

    var a int
    var b int
    var c int
    var d int
    decoder.Decode(&a)
    decoder.Decode(&b)
    decoder.Decode(&c)
    decoder.Decode(&d)

    fmt.Printf("a=%d,b=%d,c=%d,d=%d\n", a, b, c, d)
}

func main() {
    data := persist()
    readpersist(data)
}
