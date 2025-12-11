package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFileName = "input"

var data map[string][]string
var cache map[string]int

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data = make(map[string][]string)
	cache = make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()

		ps1 := strings.Split(line, ":")
		ps2 := strings.Split(strings.TrimSpace(ps1[1]), " ")
		data[ps1[0]] = ps2
	}

	fmt.Printf("%v\n", dfs("you"))
	fmt.Printf("%v\n", dfs2("svr", false, false))
}

func dfs(from string) int {
	if from == "out" {
		return 1
	}

	res := 0

	for _, c := range data[from] {
		res += dfs(c)
	}

	return res
}

func dfs2(from string, fft, dac bool) int {
	key := fmt.Sprintf("%s:%v:%v", from, fft, dac)
	r, ok := cache[key]
	if ok {
		return r
	}

	if from == "out" {
		if fft && dac {
			return 1
		}
		return 0
	}

	if from == "fft" {
		fft = true
	}

	if from == "dac" {
		dac = true
	}

	res := 0

	for _, c := range data[from] {
		res += dfs2(c, fft, dac)
	}

	cache[key] = res

	return res
}
