//Concurrency.go

package main

import (
	"fmt"
)
func f()  {
	fmt.Println("f function")
}

func main()  {
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
