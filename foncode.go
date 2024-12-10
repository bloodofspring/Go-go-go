package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var err error

const toZ int32 = 96

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func hash(str string) int {
	totalHash := 0
	for n, s := range str {
		totalHash += int(s-toZ) * int(math.Pow(37, float64(n)))
	}

	return totalHash
}

func HashTaskMain() {
	var str string

	_, _ = fmt.Scanf("%d")
	handleError(err)
	_, err = fmt.Scanf("%s", &str)
	handleError(err)

	fmt.Println(hash(str))
}

func AllDevours() {
	var n int
	_, err = fmt.Scanf("%d", &n)
	handleError(err)

	for i := 1; i <= n/2+1; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("%d", n)
}

func Matrix() {
	var n int
	_, err = fmt.Scanf("%d", &n)
	handleError(err)

	graph := make(map[int][]string)
	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= n; i++ {
		var data string
		data, err = reader.ReadString('\n')
		handleError(err)

		graph[i] = strings.Split(strings.Trim(data, "\n"), " ")[1:]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Print("0 ")
			} else if slices.Contains(graph[i], strconv.Itoa(j)) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Print("\n")
	}
}

func isSimple(n int) bool {
	for i := 2; i <= n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindSimpleNumbers() {
	var n int
	_, err = fmt.Scanf("%d", &n)
	handleError(err)

	for i := 1; i <= n; i++ {
		var start, end int
		_, err = fmt.Scanf("%d %d", &start, &end)
		handleError(err)

		simple := 0
		for num := start; num <= end; num++ {
			if isSimple(num) {
				simple++
			}
		}
		fmt.Println(simple)
	}
}

const delim int = 12345

func CountMod() {
	var a, n float64
	_, err = fmt.Scanf("%f %f", &a, &n)
	handleError(err)

	fmt.Println(int(math.Pow(a, n)) % delim)
}

type dot struct {
	x, y float64
}

type triangle struct {
	dotA, dotB, dotC dot
}

type equation struct {
	k, b float64
}

func (e equation) check(d dot) bool {
	fmt.Printf("%.3f <= %.3f * %.3f + %.3f %t\n", d.y, e.k, d.x, e.b, d.y <= e.k*d.x+e.b)
	return d.y <= e.k*d.x+e.b
}

// Определяет уравнение прямой по 2 точкам
func (d1 dot) lineEquation(d2 dot) equation {
	m := (d1.y - d2.y) / (d1.x - d2.x)
	b := d1.y - m*d1.x

	return equation{k: m, b: b}
}

func PointInATriangle() {
	var t triangle
	var dotD dot
	_, err = fmt.Scanf("%f %f %f %f %f %f %f %f", &t.dotA.x, &t.dotA.y, &t.dotB.x, &t.dotB.y, &t.dotC.x, &t.dotC.y, &dotD.x, &dotD.y)
	handleError(err)
	fmt.Println(t.dotA.lineEquation(t.dotB).check(dotD), t.dotB.lineEquation(t.dotC).check(dotD), t.dotC.lineEquation(t.dotA).check(dotD))
}
