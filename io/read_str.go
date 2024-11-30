package io

import (
	"bufio"
	"os"
	"strings"
)

func ReadStr() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ReadStrs(delim string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.Split(scanner.Text(), delim)
}
