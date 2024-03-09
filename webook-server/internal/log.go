package internal

import "fmt"

func Go(msg string, info interface{}) {
	fmt.Println("======================")
	fmt.Printf("%s %+v\n", msg, info)
	fmt.Println("======================")

}
