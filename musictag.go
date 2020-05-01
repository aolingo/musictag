package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

var logFile *os.File
var files []string

// renameSongs will replace the name of all mp3 files with
// the song title in its ID3 metadata tag, if it has one
func renameSongs() {
	fmt.Printf("%d songs detected, beginning renaming\n", len(files))

	renameCount := 0

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

		// rename file if a title can be found from the file's ID3 tag
		if m.Title() != "" {
			// log the name conversion for reference
			logStr := file.Name() + " -> " + m.Title()
			fmt.Fprintln(logFile, logStr)
			// close and rename the file
			file.Close()

			// dir := filepath.Dir(f)
			base := filepath.Base(f)
			fmt.Println(base)

			// fmt.Printf("mv %q %q\n", file.Name(), newpath)
			// err := os.Rename(file.Name(), m.Title())
			// if err != nil {
			// 	log.Fatal(err)
			// }
			renameCount++
		}
	}

	fmt.Printf("Successfully renamed %d files", renameCount)

}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		files = append(files, path)
	}

	return nil
}

func main() {
	// First element in os.Args is always the program name,
	// 2nd argument = name of the iPod music directory
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file directory")
		return
	}

	fmt.Println("Input dir is " + os.Args[1])

	// Initialize a log file for tracking all renamed songs and their names
	var logErr error
	logFile, logErr = os.Create("results.txt")
	if logErr != nil {
		fmt.Println(logErr)
		logFile.Close()
		return
	}
	fmt.Fprintln(logFile, "old name -> new name")

	// recursively scan for all songs contained in all the F subdirectories
	walkErr := filepath.Walk(os.Args[1], visit)
	if walkErr != nil {
		log.Println(walkErr)
	}

	renameSongs()

	// Close log file
	err := logFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
