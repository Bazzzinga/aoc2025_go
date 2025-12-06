package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputFileName = "input"

var data1 [][]uint64
var lines []string
var actionLine string
var actions []string

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data1 = make([][]uint64, 0)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		if slices.Contains([]string{"+", "*"}, parts[0]) {
			actionLine = line
			actions = parts
			break
		}

		lineData := make([]uint64, len(parts))
		for i, p := range parts {
			lineData[i], _ = strconv.ParseUint(p, 10, 64)
		}
		lines = append(lines, line)
		data1 = append(data1, lineData)
	}

	fmt.Println(calc1())
	fmt.Println(calc2())
}

func calc2() uint64 {
	var res uint64
	idx := 0
	a := actionLine[0]

	maxLineLen := 0
	for _, line := range lines {
		if len(line) > maxLineLen {
			maxLineLen = len(line)
		}
	}
	//fmt.Println("maxLineLen=", maxLineLen)
	for j, line := range lines {
		for i := 0; i < maxLineLen-len(line)+3; i++ {
			lines[j] = lines[j] + " "
		}
	}

	for i := 0; i < maxLineLen-len(actionLine)+3; i++ {
		actionLine = actionLine + " "
	}

	for i := 1; i < len(actionLine); i++ {
		if actionLine[i] == '+' || actionLine[i] == '*' || i == len(actionLine)-1 {
			var bres uint64
			if a == '*' {
				bres = 1
			}

			for j := i - 2; j >= idx; j-- {
				var num uint64
				for l := 0; l < len(lines); l++ {
					ch := lines[l][j]
					if ch == ' ' {
						continue
					}

					num *= 10
					num += uint64(ch) - '0'
				}

				if a == '*' {
					bres *= num
				}

				if a == '+' {
					bres += num
				}
			}
			res += bres

			idx = i
			a = actionLine[i]
		}
	}

	return res
}

func calc1() uint64 {
	var res uint64

	for i, action := range actions {
		var lres uint64
		if action == "*" {
			lres = 1
		}
		for _, nums := range data1 {
			if action == "+" {
				lres += nums[i]
			}

			if action == "*" {
				lres *= nums[i]
			}
		}

		res += lres
	}

	return res
}
