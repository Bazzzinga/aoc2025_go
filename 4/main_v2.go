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

type cell struct {
	empty      bool
	neighbours int
	marked     bool
}

var cells [][]*cell

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cells = make([][]*cell, 0)
	now := time.Now()
	for scanner.Scan() {
		line := scanner.Text()

		appendLine(strings.Split(line, ""))
	}

	res1 := 0
	res2 := 0
	removed := false

	for {
		marked := false
		for _, rows := range cells {
			for _, c := range rows {
				if !c.empty && c.neighbours < 4 {
					if !removed {
						res1++
					}
					res2++
					marked = true
					c.marked = true
				}
			}
		}

		if !marked {
			break
		}
		removed = true

		//printGrid()

		removeMarked()
	}
	now2 := time.Now()
	fmt.Println(now2.Sub(now))
	fmt.Println(res1)
	fmt.Println(res2)
}

func printGrid() {
	fmt.Println()
	for _, row := range cells {
		for _, c := range row {
			if c.marked {
				fmt.Print("x ")
			} else {
				if c.empty {
					fmt.Print(". ")
				} else {
					fmt.Print("@ ")
				}
			}
		}

		fmt.Println()
	}
}

func appendLine(line []string) {
	l := len(cells)

	nl := make([]*cell, len(line))

	for c, v := range line {
		empty := v == "."
		nc := &cell{
			empty: empty,
		}
		nl[c] = nc
		if empty {
			continue
		}
		if c > 0 {
			if !nl[c-1].empty {
				nc.neighbours++
				nl[c-1].neighbours++
			}
		}

		if l > 0 {
			if c > 0 {
				if !cells[l-1][c-1].empty {
					nc.neighbours++
					cells[l-1][c-1].neighbours++
				}
			}
			if !cells[l-1][c].empty {
				nc.neighbours++
				cells[l-1][c].neighbours++
			}
			if c < len(line)-1 {
				if !cells[l-1][c+1].empty {
					nc.neighbours++
					cells[l-1][c+1].neighbours++
				}
			}
		}
	}

	cells = append(cells, nl)
}

func removeMarked() {
	for i, row := range cells {
		for j, c := range row {
			if !c.marked {
				continue
			}

			c.empty = true
			c.marked = false
			c.neighbours = 0

			if i > 0 {
				//смотрим вверх
				if !cells[i-1][j].empty {
					cells[i-1][j].neighbours--
				}

				if j > 0 {
					//смотрим влево
					if !cells[i-1][j-1].empty {
						cells[i-1][j-1].neighbours--
					}
				}

				if j < len(row)-1 {
					//смотрим вправо
					if !cells[i-1][j+1].empty {
						cells[i-1][j+1].neighbours--
					}
				}
			}

			//cмотрим по бокам
			if j > 0 {
				//смотрим влево
				if !cells[i][j-1].empty {
					cells[i][j-1].neighbours--
				}
			}

			if j < len(row)-1 {
				//смотрим вправо
				if !cells[i][j+1].empty {
					cells[i][j+1].neighbours--
				}
			}

			if i < len(cells)-1 {
				//смотрим вниз
				if !cells[i+1][j].empty {
					cells[i+1][j].neighbours--
				}

				if j > 0 {
					//смотрим влево
					if !cells[i+1][j-1].empty {
						cells[i+1][j-1].neighbours--
					}
				}

				if j < len(row)-1 {
					//смотрим вправо
					if !cells[i+1][j+1].empty {
						cells[i+1][j+1].neighbours--
					}
				}
			}
		}
	}
}
