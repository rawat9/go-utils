package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return num
}

func ReadInts(delim string) []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), delim)
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}
