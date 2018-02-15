package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"errors"
	"io"
	"math/big"
	"os"
	"sort"
)

// countLines counts the number of lines in a file.
func countLines(r io.Reader) (uint64, error) {
	buf := make([]byte, 64*1024)
	count := uint64(0)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += uint64(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// readLines reads lines from a file. It takes a list of line numbers, sorts
// them so that they can be read sequentially, and reads the contents of those
// lines. It returns the list of words on the provided line numbers, in the
// same order they were received.
func readLines(r io.Reader, lineNums []uint64) (lines []string, err error) {
	sortedLineNums := make([]uint64, len(lineNums))
	copy(sortedLineNums, lineNums)
	sort.Slice(sortedLineNums, func(i, j int) bool {
		return sortedLineNums[i] < sortedLineNums[j]
	})

	i := 0
	lineNum := sortedLineNums[i]
	scanner := bufio.NewScanner(r)
	linesByNum := make(map[uint64]string, len(sortedLineNums))

	scanner.Split(bufio.ScanLines)

	for n := uint64(0); scanner.Scan(); n++ {
		if n == lineNum {
			line, err := scanner.Text(), scanner.Err()
			if err != nil {
				return nil, err
			}
			linesByNum[n] = line

			i++
			if i == len(sortedLineNums) {
				break
			}
			lineNum = sortedLineNums[i]
		}
	}

	lines = make([]string, len(lineNums))
	for i = 0; i < len(lineNums); i++ {
		lines[i] = linesByNum[lineNums[i]]
	}

	return lines, nil
}

// GetWords gets numWords words from a dictionary file.
func GetWords(dict string, numWords uint64) (words []string, err error) {
	file, err := os.Open(dict)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	numLines, err := countLines(file)
	if err != nil {
		return nil, err
	}

	max := big.NewInt(int64(numLines))
	lineNums := make([]uint64, numWords)

	for i := uint64(0); i < numWords; i++ {
		b, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}

		if !b.IsUint64() {
			return nil, errors.New("Could not get random dictionary words")
		}

		lineNum := b.Uint64()
		lineNums[i] = lineNum
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	words, err = readLines(file, lineNums)
	if err != nil {
		return nil, err
	}

	return words, nil
}
