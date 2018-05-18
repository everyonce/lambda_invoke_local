package main

import (
	"fmt"
	"log"

	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {
	response, err := golambdainvoke.Run(8001, "{'payload'}")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")

	fmt.Println(string(response))
}
