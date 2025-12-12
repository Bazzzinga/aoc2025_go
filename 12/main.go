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

type Map map[XY]bool

func NewMap() Map {
	return make(map[XY]bool)
}

type XY struct {
	X int
	Y int
}
type Shape struct {
	ID  int
	Map Map
}

type Field struct {
	MaxX     int
	MaxY     int
	Map      Map
	Presents map[int]int
	Total    int
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	shapeId := 0
	y := 0

	fields := make([]Field, 0)
	shapes := make(map[int]*Shape)
	var curShape *Shape
	for scanner.Scan() {
		line := scanner.Text()

		l := len(line)

		if l == 0 {
			shapes[curShape.ID] = curShape
			continue
		}

		if l == 2 {
			shapeId, _ = strconv.Atoi(strings.Trim(line, ":"))
			y = 0
			curShape = &Shape{
				ID:  shapeId,
				Map: NewMap(),
			}
		}

		if l == 3 {
			ps := strings.Split(line, "")
			for x := 0; x < l; x++ {
				if ps[x] == "#" {
					c := XY{x, y}
					curShape.Map[c] = true
				}
			}
			y++
		}

		if l > 3 {
			f := Field{
				Map:      NewMap(),
				Presents: make(map[int]int),
			}

			ps := strings.Split(line, ":")
			ms := strings.Split(ps[0], "x")
			mx, _ := strconv.Atoi(ms[0])
			my, _ := strconv.Atoi(ms[1])

			f.MaxX = mx
			f.MaxY = my

			ps2 := strings.Split(strings.TrimSpace(ps[1]), " ")
			for i := 0; i < len(ps2); i++ {
				v, _ := strconv.Atoi(ps2[i])
				f.Presents[i] = v
				f.Total += v
			}

			fields = append(fields, f)
		}
	}

	res1 := 0
	for _, f := range fields {
		res1 += f.calc(shapes)
	}

	fmt.Println(res1)
}

func (f Field) calc(shapes map[int]*Shape) int {
	size := f.MaxX * f.MaxY

	ps := 0
	for id, v := range f.Presents {
		ps += v * len(shapes[id].Map)
	}

	if ps > size {
		return 0
	}

	sq := f.MaxX / 3 * f.MaxY / 3
	if f.Total <= sq {
		return 1
	} else {
		fmt.Printf("problem: %dx%d\n", f.MaxX, f.MaxY)
		return 0
	}
}
