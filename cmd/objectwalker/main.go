package main

import (
	"fmt"
	"log"
)

func main() {
	var count int64

	count, err := fswalk.WalkDir("directory")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(count)
}
