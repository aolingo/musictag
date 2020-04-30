package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

func main() {
	// First element in os.Args is always the program name,
	// 2nd argument = name of the iPod music directory
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file directory")
		return
	}

	fmt.Println("Input dir is " + os.Args[1])

	var files []string

	// recursively scan for all songs contained in all the F subdirectories
	err := filepath.Walk(os.Args[1],
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%d songs detected, beginning renaming\n", len(files))

	// Initialize a log file for tracking all renamed songs and their names
	logFile, err := os.Create("results.txt")
	if err != nil {
		fmt.Println(err)
		logFile.Close()
		return
	}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			fmt.Println("Can't open file:", f)
			panic(err)
		}

		m, err := tag.ReadFrom(file)
		if err != nil {
			log.Fatal(err)
		}

		logStr := file.Name() + " -> " + m.Title()
		fmt.Fprintln(logFile, logStr)
		fmt.Println("song name: ", m.Title())
	}

	// Close log file
	err = logFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
