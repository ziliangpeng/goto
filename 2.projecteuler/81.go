package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	const n_col, n_row = 80, 80

	matrix := make([][]int, n_row)
	for i := range matrix {
		matrix[i] = make([]int, n_col)
	}

	content, _ := os.ReadFile("81.in")
	rows := strings.Split(string(content), "\n")
	for i, row := range rows {
		cols := strings.Split(row, ",")
		for j, col := range cols {
			matrix[i][j], _ = strconv.Atoi(col)
		}
	}

	sum := make([][]int, n_row)
	for i := range sum {
		sum[i] = make([]int, n_col)
	}

	for i := 0; i < n_row; i++ {
		for j := 0; j < n_col; j++ {
			min_pre := math.MaxInt
			if i > 0 {
				min_pre = min(min_pre, sum[i-1][j])
			}
			if j > 0 {
				min_pre = min(min_pre, sum[i][j-1])
			}
			if min_pre == math.MaxInt {
				min_pre = 0
			}
			sum[i][j] = min_pre + matrix[i][j]
		}
	}

	fmt.Println(sum[n_row-1][n_col-1])
}
