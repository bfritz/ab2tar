package main

import (
    "compress/zlib"
    "io"
    "os"
    "flag"
)

var infile = flag.String("f", "", "infile")

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
    flag.Parse()

    file, err := os.Open(*infile)
    check(err)

    _, err = file.Seek(24, 0)
    check(err)

    r, err := zlib.NewReader(file)
    check(err)

    io.Copy(os.Stdout, r)
    r.Close()
}
