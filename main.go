package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	h bool
	i string
	d string
)

// init
func init() {
	flag.BoolVar(&h, "h", false, `Desc: help document
`)
	flag.StringVar(&i, "i", "", `Desc: add or alter a description 
Usage: ls-pro -i directory description
`)
	flag.StringVar(&d, "d", "", `Desc: delete a description
Usage: ls-pro -d directory
`)
	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `ls-pro version: 0.0.1
Usage: ls-pro [-h] [-i add or alter item] [-d delete item]

Options:
`)
	flag.PrintDefaults()
}

// main
func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	params := os.Args

	db, err := CreateConnection(HomeDir() + "/ls_pro.db")
	if err != nil {
		log.Fatal("path: ", err)
		return
	}
	defer db.Close()

	currDir := GetCurrentDirectory()

	_, _ = fmt.Fprintf(os.Stderr, `
%sCurrent directory: %s %s %s

`, string(colorPurple), string(colorYellow), currDir, string(colorReset))

	srv := &service{
		repo:    &LsRepository{db: db},
		currDir: currDir,
		id:      Md5Encode(currDir),
	}

	// check table

	if i != "" {
		srv.Insert(params)
		return
	}

	if d != "" {
		srv.Del(params)
		return
	}

	srv.GetAll()
}
