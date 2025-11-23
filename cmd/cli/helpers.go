package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(prompt string) (string, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	l, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(l), nil
}

func strConv(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %q to int", s)
	}
	return n, nil
}
