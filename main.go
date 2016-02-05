package main

import (
    "os"
    "log"
    "fmt"
    "strings"
    "io/ioutil"
    "path/filepath"
)

func main() {

    args := os.Args
    if len(args) != 3 {
        panic("Usage program src dest")
    }

    // read the file content
    file, err := os.Open(args[1])
    if err != nil {
        panic(err)
    }

    defer file.Close()

    data = readFile(file)
    log.Printf("Source data read from file %s: %s\n ", file.Name(), data)

    // write to folder recursively
    filepath.Walk(args[2], walkpath)
    fmt.Printf("Total file operated %d\n", count)
}

var data []byte
var count int = 0

func readFile(file *os.File) []byte {

    data, err := ioutil.ReadAll(file)
    if err != nil {
        panic(err)
    }
    return data
}

func walkpath(path string, f os.FileInfo, err error) error {
    // if dir walk through
    if f.IsDir() {
        files, err := ioutil.ReadDir(path)
        if err != nil {
            panic(err)
        }
        for _, file := range files {
            if file.IsDir() {
                s := []string{path, string(filepath.Separator), file.Name()}
                walkpath(strings.Join(s, ""), file, err)
            }
        }
    } else { // if file open to write
        // skip .DS_Store file
        fmt.Printf("file name %s\n", f.Name())
        if strings.Compare(f.Name(), ".DS_Store") == 0 {
            return nil
        }
        fmt.Printf("%s with %d bytes\n", path, f.Size())
        count++
    }
    return nil
}

