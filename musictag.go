package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/dhowden/tag"
)

var logFile *os.File
var files []string
var reservedChars = [9]string{"\"", "<", ">", ":", "/", "|", "\\", "?", "*"}

// renameSongs will replace the name of all mp3 files with
// the song title in its ID3 metadata tag, if it has one
func renameSongs() {
	fmt.Printf("%d songs detected, beginning renaming\n", len(files))

	renameCount := 0
	// create and start a new progress bar
	bar := pb.StartNew(len(files))

	for _, f := range files {
		bar.Increment()
		time.Sleep(time.Millisecond * 10)

		// extra safeguard to skip a mp3 file if it's file name is already decoded
		// eg. it doesn't have the form ABCD, where ABCD are 4 uppercase letters
		oldname := strings.Replace(filepath.Base(f), ".mp3", "", 1)
		if len(oldname) != 4 || strings.ToUpper(oldname) != oldname {
			continue
		}

		file, err := os.Open(f)
		if err != nil {
			fmt.Println("Can't open file:", f)
			panic(err)
		}

		// get the ID3 tag from the opened mp3 file
		m, err := tag.ReadFrom(file)
		if err != nil {
			log.Fatal(err)
		}

		// rename file if a title can be found from the file's ID3 tag
		if m.Title() != "" {

			// Handle case where title has illegal chars eg. Kaleo "Save Yourself".mp3
			// will cause program to crash when renaming file to the song title
			title := cleanTitle(m.Title())

			// retrive song title from tag and do some cleaning before renaming
			newbase := strings.TrimSpace(title)
			if !strings.Contains(newbase, ".mp3") {
				newbase += ".mp3"
			}
			newpath := filepath.Join(filepath.Dir(f), newbase)

			// log the name conversion for reference
			logStr := file.Name() + " -> " + newbase
			fmt.Fprintln(logFile, logStr)

			// close and rename the file
			file.Close()

			err := os.Rename(file.Name(), newpath)
			if err != nil {
				log.Fatal(err)
			}
			renameCount++
		}
	}

	bar.Finish()

	if renameCount == len(files) {
		fmt.Printf("Succesfully renamed all detected songs, results are logged in results.txt\n")
	} else {
		skipped := len(files) - renameCount
		fmt.Printf("Renamed %d songs, skipped %d songs due to incomplete ID3 tag or already named\n", renameCount, skipped)
		fmt.Println("File renaming results are logged in results.txt")
	}
}

// cleanTitle checks and removes if inputted title string contains reserved characters
// for file and folder names in Windows such as colons and double quotes
func cleanTitle(title string) string {
	var newTitle string
	for _, symbol := range reservedChars {
		if strings.ContainsAny(title, symbol) {
			newTitle = strings.ReplaceAll(title, symbol, "")
		}
	}
	return newTitle
}

// visit is the helper walkFn called by filepath.Walk,
// it tracks all the .mp3 files found in the input
// directory and its subdirectories
func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && filepath.Ext(path) == ".mp3" {
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

	// Can optimize the program's time complexity from O(2n) to O(n) in the future
	// by implementing renameSongs() in the visit walkFn helper (scan and rename
	// at the same time instead of rename after scanning all the files)
	renameSongs()

	// Close log file
	err := logFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
