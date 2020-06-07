# Music Tag

A personal CLI tool to decode song names from old Apple devices. 

Read the blog post for more info: https://www.danielgong.com/project/mp3meta/

## Install

```sh
go get github.com/aolingo/musictag
```

## Usage
After compiling into executable, run the program and input the path of directory containing the mp3 files

`./musictag.exe input_path`

For example, parse and rename all the songs in a directory named music

```sh
./musictag.exe music
```

[![asciicast](https://asciinema.org/a/337143.svg)](https://asciinema.org/a/337143)
