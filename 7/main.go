package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	res1 := 0
	splitters := make([]map[int]struct{}, 0)
	var start int
	beams := make(map[int]struct{})
	for scanner.Scan() {
		line := scanner.Text()

		spl := make(map[int]struct{})

		chs := strings.Split(line, "")

		for i := 0; i < len(chs); i++ {
			if chs[i] == "S" {
				beams[i] = struct{}{}
				start = i
				break
			}

			if chs[i] == "^" {
				spl[i] = struct{}{}
				_, ok := beams[i]
				if ok {
					res1++
					delete(beams, i)
					beams[i-1] = struct{}{}
					beams[i+1] = struct{}{}
				}
			}
		}

		if len(spl) > 0 {
			splitters = append(splitters, spl)
		}
	}

	fmt.Println(res1)

	data := make(map[string]int)

	res2 := timelines(start, 0, splitters, &data)
	fmt.Println(res2)
}

func timelines(pos, line int, splitters []map[int]struct{}, data *map[string]int) int {
	//если все строки со сплиттерами кончились - возвращаем 1
	if line == len(splitters) {
		return 1
	}
	//если в кеше есть инфа для этой позиции - возвращаем ее
	key := fmt.Sprintf("%d-%d", pos, line)
	r, ok := (*data)[key]
	if ok {
		return r
	}

	//если сплиттер - добавляем результат для правой и левой позиции и кладем его в кеш
	_, sp := splitters[line][pos]
	if sp {
		c := timelines(pos-1, line+1, splitters, data) + timelines(pos+1, line+1, splitters, data)
		(*data)[key] = c
		return c
	} else {
		//иначе возвращаем результат для следующей строки
		return timelines(pos, line+1, splitters, data)
	}
}
