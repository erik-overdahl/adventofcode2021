package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileToString(filename string) string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("Unable to read %s: %v", filename, err))
	}
	return strings.Trim(string(contents), "\n")
}

func ReadlinesInt(inputBlob string) []int {
	var lines []int
	reader := strings.NewReader(inputBlob)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		intLine, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("Failed to convert %s to int: %v", line, err))
		}
		lines = append(lines, intLine)
	}
	return lines
}

func ReadlinesStr(inputBlob string) []string {
	var lines []string
	reader := strings.NewReader(inputBlob)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
