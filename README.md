# Music Tag
[![Go Report Card](https://goreportcard.com/badge/github.com/aolingo/musictag)](https://goreportcard.com/report/github.com/aolingo/musictag)

A personal CLI tool to decode song names from old Apple devices. 

Read the blog post for more info: https://www.danielgong.com/project/mp3meta/

## Building
To download the project, run `go get github.com/aolingo/musictag`

The project can be then found in your `$GOPATH/src/<import-path> folder`. 

In Windows for example, it will most likely be in `C:\users\(you)\go\src\github.com\OpenDiablo2\OpenDiablo2`

To build this project, run `go build` from the project root folder.

## Usage
After building or downloading the executable, run the program and input the path of the directory containing the mp3 files

`./musictag.exe input_path`

For example, parse and rename all the songs in a directory named `music`

```sh
./musictag.exe music
```

[![asciicast](https://asciinema.org/a/337143.svg)](https://asciinema.org/a/337143)
