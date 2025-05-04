package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(reader *bufio.Reader) string {
	matrix := read_test(reader)
	rows := make([]uint32, len(matrix))
	columns := make([]uint32, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			columns[j] += matrix[i][j]
			rows[i] += matrix[i][j]
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i] > rows[j] })
	sort.Slice(columns, func(i, j int) bool { return rows[i] > rows[j] })
	for i := 0; i < len(rows); i++ {
		if rows[i] != columns[i] {
			return "no"
		}
	}
	return "yes"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)
	defer writer.Flush()

	temp, _ := reader.ReadString('\n')
	n_of_tests, _ := strconv.Atoi(strings.TrimSpace(temp))
	for t := 0; t < n_of_tests; t++ {
		fmt.Fprintln(writer, solution(reader))
	}

}

func read_test(reader *bufio.Reader) [][]uint32 {
	temp, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(temp))
	// fmt.Println(n)
	matrix := make([][]uint32, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]uint32, n)
		// fmt.Println(matrix)
		str, _ := reader.ReadString('\n')
		// fmt.Println(str)
		row := strings.Fields(strings.TrimSpace(str))
		// fmt.Println(row)
		for r := 0; r < int(n); r++ {
			temp, _ := strconv.Atoi(row[r])
			matrix[i][r] = uint32(temp)
		}
	}
	return matrix
}
