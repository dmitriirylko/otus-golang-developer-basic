package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findFieldSize() int {
	const (
		defaultFieldSize = 8
		minFieldSize     = 1
		maxFieldSize     = 100
	)
	if len(os.Args) > 1 {
		cliFieldSize, err := strconv.Atoi(os.Args[1])
		if err == nil && cliFieldSize >= minFieldSize && cliFieldSize <= maxFieldSize {
			return cliFieldSize
		}
	}
	return defaultFieldSize
}

func makeField(size int) string {
	strLen := size*size + size
	var strBuilder strings.Builder
	strBuilder.Grow(strLen)
	for i := 0; i < size; i++ {
		var isBlack bool = i%2 == 1
		for j := 0; j < size; j++ {
			if isBlack {
				strBuilder.WriteByte('#')
			} else {
				strBuilder.WriteByte(' ')
			}
			isBlack = !isBlack
		}
		strBuilder.WriteByte('\n')
	}
	return strBuilder.String()
}

func main() {
	fmt.Println("Let's play chess")
	fieldSize := findFieldSize()
	fmt.Printf("Field size is %d\n", fieldSize)
	field := makeField(fieldSize)
	fmt.Print(field)
}
