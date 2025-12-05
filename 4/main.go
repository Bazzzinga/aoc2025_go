package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const inputFileName = "input"

var grid [][]string
var dirs [][]int

func mainV1() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid = make([][]string, 0)

	dirs = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	now := time.Now()
	for scanner.Scan() {
		line := scanner.Text()

		gridLine := strings.Split(line, "")
		grid = append(grid, gridLine)
	}

	res1 := 0
	res2 := 0
	removed := false
	for {
		remove := make([][]int, 0)

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if check(i, j) {
					if !removed {
						res1++
					}
					res2++
					remove = append(remove, []int{i, j})
				}
			}
		}

		if len(remove) == 0 {
			break
		}

		removed = true

		for i := 0; i < len(remove); i++ {
			grid[remove[i][0]][remove[i][1]] = "."
		}
	}
	now2 := time.Now()

	fmt.Println(now2.Sub(now))

	fmt.Println(res1)
	fmt.Println(res2)
}

func check(i, j int) bool {
	if grid[i][j] != "@" {
		return false
	}

	cnt := 0

	maxi := len(grid) - 1
	maxj := len(grid[i]) - 1

	for _, d := range dirs {
		ni := i + d[0]
		nj := j + d[1]

		if ni >= 0 && ni <= maxi && nj >= 0 && nj <= maxj {
			if grid[ni][nj] == "@" {
				cnt++
			}
		}
	}

	return cnt < 4
}
