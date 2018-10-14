
package main

import (
	"fmt"
	"time"

)


func main() {

	fmt.Println("hello world")

	vector := [3]string{"a","b","c"}

	var chanel1 chan string = make(chan string)
	var chanel2 chan string = make(chan string)


	go func(resources [3]string){

		for _,value := range resources {
			fmt.Printf("[find] %s\n",value)
			chanel1 <- value
		}

	}(vector)

	go func() {

		select {
			case findOne := <-chanel1:
				fmt.Printf("[proce] get %s\n", findOne)
				chanel2 <- findOne
			default:
				fmt.Print("[proce] not get message\n")
		}

	}()

	go func() {
		for endOne := range chanel2 {
			fmt.Printf("[end] get from processer: %s\n", endOne)
		}

	}()

	<-time.After(time.Second * 3)
}
