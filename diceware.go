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
	verbose  bool
	help     bool
)

func init() {
	flag.Uint64Var(&numWords, "n", 5, "number of words in password")
	flag.StringVar(&delim, "delim", "-", "word delimiter")
	flag.StringVar(&dict, "dict", "/usr/share/dict/words", "dictionary file")
	flag.BoolVar(&verbose, "verbose", false, "display extra information")
	flag.BoolVar(&help, "help", false, "print help text")
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	words, ent, err := getWords(dict, numWords)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		fmt.Printf("Entropy: %f bits\n", ent)
	}

	pw := strings.Join(words, delim)
	fmt.Println(pw)
}
