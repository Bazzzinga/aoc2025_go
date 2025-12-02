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

	var res1 uint64
	var res2 uint64

	for scanner.Scan() {
		line := scanner.Text()

		rngs := strings.Split(line, ",")
		for _, rng := range rngs {
			parts := strings.Split(rng, "-")

			from, _ := strconv.ParseUint(parts[0], 10, 64)
			to, _ := strconv.ParseUint(parts[1], 10, 64)

			for i := from; i <= to; i++ {
				if isInvalidInt1(i) {
					res1 += i
				}

				if isInvalidInt2(i) {
					res2 += i
				}
			}
		}
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func isInvalidInt1(n uint64) bool {
	s := strconv.FormatUint(n, 10)
	return isInvalidString1(s)
}

func isInvalidString1(n string) bool {
	if len(n)%2 != 0 {
		return false
	}

	return n[:len(n)/2] == n[len(n)/2:]
}

func isInvalidInt2(n uint64) bool {
	s := strconv.FormatUint(n, 10)
	return isInvalidString2(s)
}

func isInvalidString2(n string) bool {
	for i := 1; i <= len(n)/2; i++ {
		if len(n)%i == 0 {
			l := i
			parts := make([]string, len(n)/l)

			for j := 0; j < len(n)/l; j++ {
				parts[j] = n[j*l : (j+1)*l]
			}

			equals := true
			for j := 1; j < len(n)/l; j++ {
				if parts[j] != parts[0] {
					equals = false
				}
			}

			if equals {
				return true
			}
		}
	}

	return false
}
