package util

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func ReadInputFile(pathOfCaller string) (inputLines []string) {
	// try to get dir from which it's called via Caller object
	_, fn, _, ok := runtime.Caller(1)

	if !ok {
		panic("Could not get Caller of util.ReadInputFile")
	}

	// absolute path to the input file (shall be placed in the "day-xy"
	// directory alongside the source code)
	absPath := path.Join(path.Dir(fn), pathOfCaller)

	f, err := os.Open(absPath)
	if err != nil {
		panic("Could not open input file")
	}
	defer f.Close()

	fs := bufio.NewScanner(f)

	for fs.Scan() {
		line := fs.Text()
		inputLines = append(inputLines, line)
	}

	return inputLines
}

func PrepareExampleInput(input string) (inputLines []string) {
	fs := bufio.NewScanner(strings.NewReader(input))

	for fs.Scan() {
		line := fs.Text()
		line = strings.TrimSpace(line)
		inputLines = append(inputLines, line)
	}

	return inputLines
}

func ParseInt(input string) int {
	n, err := strconv.Atoi(input)

	if err != nil {
		panic("Trying to parse a string that is not a number")
	}

	return n
}
