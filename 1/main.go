package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFileName = "input"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := 50
	zeros := 0

	dial2 := 50
	zeros2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		// L -> -
		// R -> +

		//part 1 start
		dirC := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		num1 := num
		dir := 1
		if dirC == 'L' {
			num1 *= -1
			dir = -1
		}

		dial += num1
		dial = dial % 100
		if dial == 0 {
			zeros++
		}
		//part 1 end

		//part 2 start
		for i := 0; i < num; i++ {
			dial2 += dir
			dial2 = dial2 % 100
			if dial2 == 0 {
				zeros2++
			}
		}
		//part 2 end
	}

	fmt.Println(zeros)
	fmt.Println(zeros2)
}
