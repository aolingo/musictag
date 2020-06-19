# Music Tag
[![Go Report Card](https://goreportcard.com/badge/github.com/aolingo/musictag)](https://goreportcard.com/report/github.com/aolingo/musictag)

A personal CLI tool to decode song names from old Apple devices. 

Read the blog post for more info: https://www.danielgong.com/project/mp3meta/

## Install
The project executable `musictag.exe` can be directly downloaded from the project [release](https://github.com/aolingo/musictag/releases/tag/v0.1-alpha) page.

## Building
Alternatively, to download and build the project yourself, run `go get github.com/aolingo/musictag`

The project can be then found in your `$GOPATH/src/<import-path> folder`. 

In Windows for example, it will most likely be in `C:\users\(you)\go\src\github.com\aolingo\musictag`

To build this project, run `go build` from the project root folder.

## Usage
After downloading or building the program, run the executable from the command line and input the path of the directory containing the mp3 files

`./musictag.exe input_path`

For example, parse and rename all the songs in a directory named `music`

```sh
./musictag.exe music
```

[![asciicast](https://asciinema.org/a/337143.svg)](https://asciinema.org/a/337143)
