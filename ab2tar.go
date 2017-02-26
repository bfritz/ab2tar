package main

import (
    "compress/zlib"
    "github.com/alexflint/go-arg"
    "io"
    "os"
)

// populated with linker flag `-X main.version=${VERSION}`
var version string

type args struct {
    File  string `arg:"positional,required,help:name of Android backup file"`
}

func (args) Description() string {
    return "Convert unencrypted Android backup files into tar archives on stdout."
}

func (args) Version() string {
    return "ab2tar " + version
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var args args
    arg.MustParse(&args)

    file, err := os.Open(args.File)
    check(err)

    _, err = file.Seek(24, 0)
    check(err)

    r, err := zlib.NewReader(file)
    check(err)

    io.Copy(os.Stdout, r)
    r.Close()
}
