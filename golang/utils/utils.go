package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadlinesInt(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read %s: %v", filename, err)
	}
	defer f.Close()
	var lines []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		intLine, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert %s to int: %v", line, err)
		}
		lines = append(lines, intLine)
	}
	return lines, nil
}

func ReadlinesStr(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read %s: %v", filename, err)
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}
