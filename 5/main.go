package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFileName = "input"

var fresh [][]uint64
var ids []uint64

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fresh = make([][]uint64, 0)
	ids = make([]uint64, 0)

	doneWithFresh := false
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			doneWithFresh = true
			continue
		}

		if !doneWithFresh {
			parts := strings.Split(line, "-")
			p1, _ := strconv.ParseUint(parts[0], 10, 64)
			p2, _ := strconv.ParseUint(parts[1], 10, 64)

			fresh = append(fresh, []uint64{p1, p2})
		} else {
			id, _ := strconv.ParseUint(line, 10, 64)

			ids = append(ids, id)
		}
	}

	fmt.Println(countFresh1())
	fmt.Println(countFresh2())
}

func countFresh1() int {
	res := 0

cycle:
	for _, id := range ids {
		for _, f := range fresh {
			if id >= f[0] && id <= f[1] {
				res++
				continue cycle
			}
		}
	}

	return res
}

func countFresh2() uint64 {
	var res uint64

	sort.SliceStable(fresh, func(i, j int) bool {
		return fresh[i][0] < fresh[j][0]
	})

	var lastMax uint64

	for _, f := range fresh {
		if lastMax < f[0] {
			res += f[1] - f[0] + 1
		} else {
			newMin := lastMax + 1
			if newMin <= f[1] {
				res += f[1] - newMin + 1
			}
		}

		if lastMax < f[1] {
			lastMax = f[1]
		}
	}

	return res
}
