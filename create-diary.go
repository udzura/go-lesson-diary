package main

import (
//	"io"
	"os"
)

func main () {
	diary, err := os.Create("diary/sample.txt")
	if err != nil { panic(err) }
	defer func() {
		if err := diary.Close(); err != nil {
			panic(err)
		}
	}()

	write(diary, "Hello\n")
	write(diary, "World")
}

func write (f *os.File, str string) {
	_, e := f.Write([]byte(str))
	if e != nil { panic(e) }
}
