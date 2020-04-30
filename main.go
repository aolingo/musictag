package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

func main() {
	// First element in os.Args is always the program name,
	// 2nd argument = music directory path
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file directory")
		return
	}

	fmt.Println("Input file dir is " + os.Args[1])
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		p := filepath.Join(os.Args[1], f.Name())
		file, err := os.Open(p)
		if err != nil {
			fmt.Println("Can't open file:", f.Name())
			panic(err)
		}

		m, err := tag.ReadFrom(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("song name: ", m.Title())
	}

}
