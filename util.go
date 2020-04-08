package main

import (
    "io/ioutil"
    "log"
    "os"
    "os/user"
    "path/filepath"
    "strings"
)

var (
    colorGreen  = "\033[32m"
    colorReset  = "\033[0m"
    colorYellow = "\033[33m"
    colorPurple = "\033[35m"
)

// GetCurrentDirectory get current directory
func GetCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    return strings.Replace(dir, "\\", "/", -1)
}

// ListDir list directories
func ListDir(dirPth string) ([]string, error) {
    var files []string

    dir, err := ioutil.ReadDir(dirPth)
    if err != nil {
        return nil, err
    }

    for _, fi := range dir {
        if !strings.HasPrefix(fi.Name(), ".") && fi.IsDir() {
            files = append(files, fi.Name())
        }
    }

    return files, nil
}

// Exist check it is exist
func Exist(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

// HomeDir user's home directory
func HomeDir() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return usr.HomeDir
}
