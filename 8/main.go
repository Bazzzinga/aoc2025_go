package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputFileName = "input"

type box struct {
	x int
	y int
	z int
	c int
}

type edge struct {
	a *box
	b *box
	d float64
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	boxes := make([]*box, 0)

	for scanner.Scan() {
		line := scanner.Text()

		cs := strings.Split(line, ",")

		x, _ := strconv.Atoi(cs[0])
		y, _ := strconv.Atoi(cs[1])
		z, _ := strconv.Atoi(cs[2])

		boxes = append(boxes, &box{x, y, z, 0})
	}

	fmt.Println(calc1(boxes, 1000))
}

func calc1(boxes []*box, n int) (int, int) {
	cId := 0
	circuits := make(map[int][]*box)

	edges := make([]*edge, 0, len(boxes)*(len(boxes)-1))

	for i, a := range boxes {
		for j, b := range boxes {
			if i != j {
				edges = append(edges, &edge{
					a: a,
					b: b,
					d: d(a, b),
				})
			}
		}
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].d < edges[j].d
	})

	ans2 := 0
	cnt := 0
	for k := 0; k < len(edges); k += 2 {
		e := edges[k]

		//for part 1
		/*if cnt == n {
			break
		}*/

		cnt++

		if e.a.c > 0 && e.b.c == e.a.c {
			continue
		}

		ans2 = e.a.x * e.b.x

		if e.a.c > 0 && e.b.c > 0 {
			o := e.b.c
			for _, x := range circuits[o] {
				x.c = e.a.c
			}
			circuits[e.a.c] = append(circuits[e.a.c], circuits[o]...)
			delete(circuits, o)
		} else if e.a.c > 0 {
			e.b.c = e.a.c
			circuits[e.b.c] = append(circuits[e.b.c], e.b)
		} else if e.b.c > 0 {
			e.a.c = e.b.c
			circuits[e.a.c] = append(circuits[e.a.c], e.a)
		} else {
			cId++
			circuits[cId] = make([]*box, 0)

			e.a.c = cId
			e.b.c = cId
			circuits[cId] = append(circuits[cId], e.a, e.b)
		}
	}

	connections := make([]int, len(circuits))
	i := 0
	for _, c := range circuits {
		connections[i] = len(c)
		i++
	}

	l := len(connections)
	sort.Ints(connections)

	m1 := connections[l-1]
	m2 := 1
	m3 := 1

	if l-2 >= 0 {
		m2 = connections[l-2]
	}

	if l-3 >= 0 {
		m3 = connections[l-3]
	}

	return m1 * m2 * m3, ans2
}

func d(a, b *box) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z

	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
