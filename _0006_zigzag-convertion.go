// https://leetcode.com/problems/zigzag-conversion/

package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]string, min(numRows, len(s)))
	currentRow := 0
	movingDown := false

	for _, i := range s {
		rows[currentRow] += string(i)

		if currentRow == 0 || currentRow == numRows-1 {
			movingDown = !movingDown
		}

		if movingDown {
			currentRow++
		} else {
			currentRow--
		}

		fmt.Println(currentRow)
	}

	return strings.Join(rows, "")
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
