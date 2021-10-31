package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gleanerio/objectwalker/internal/checks"
	"github.com/gleanerio/objectwalker/internal/fswalk"
)

func main() {

	// argsnp := os.Args[1:]
	patharg := os.Args[1]

	log.Printf("Scanning directory %s", patharg)
	r, err := fswalk.WalkDirNames(patharg)
	if err != nil {
		log.Println(err)
	}

	for _, v := range r {
		fmt.Println(v)
	}

	checks.DoChecks(r)

}
