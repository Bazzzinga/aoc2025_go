package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileName = "input"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res1, res2 uint

	for scanner.Scan() {
		line := scanner.Text()

		res1 += getMaxJoltage1(line)
		res2 += getMaxJoltage2(line)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func getMaxJoltage1(bank string) uint {
	bs := strings.Split(bank, "")

	maxl := 0
	idx := 0
	maxr := 0
	for i := 0; i < len(bs)-1; i++ {
		jl, _ := strconv.Atoi(bs[i])

		if maxl < jl {
			maxl = jl
			idx = i
		}
	}

	for i := idx + 1; i < len(bs); i++ {
		jl, _ := strconv.Atoi(bs[i])

		if maxr < jl {
			maxr = jl
		}
	}

	return uint(maxl*10 + maxr)
}

func getMaxJoltage2(bank string) uint {
	n := 12
	bs := strings.Split(bank, "")

	var res uint

	for i := n; i > 0; i-- {
		m, idx := getMaxWithRightOffset(bs, i)
		res = res*10 + m
		bs = bs[idx+1:]
	}

	return res
}

func getMaxWithRightOffset(bs []string, offset int) (uint, int) {
	maxr := 0
	idx := 0
	for i := 0; i < len(bs)-offset+1; i++ {
		jl, _ := strconv.Atoi(bs[i])

		if maxr < jl {
			maxr = jl
			idx = i
		}
	}

	return uint(maxr), idx
}
