package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	numWords uint64
	dict     string
	delim    string
	help     bool
)

func init() {
	flag.Uint64Var(&numWords, "n", 5, "number of words in password")
	flag.StringVar(&delim, "delim", "-", "word delimiter")
	flag.StringVar(&dict, "dict", "/usr/share/dict/words", "dictionary file")
	flag.BoolVar(&help, "help", false, "print help text")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	words, err := GetWords(dict, numWords)
	if err != nil {
		log.Fatal(err)
	}

	pw := strings.Join(words, delim)

	fmt.Println(pw)
}
