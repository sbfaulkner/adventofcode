package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/sbfaulkner/adventofcode/pkg/advent"
)

func input(n int) *os.File {
	_, sourcePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to get source path")
	}

	sourceDir := filepath.Dir(sourcePath)
	file, err := os.Open(path.Join(sourceDir, fmt.Sprintf("../testdata/%02d", n)))
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func day1() {
	r, err := advent.ReadReport(input(1))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("1-1:", r.ProductOfCombinationWithSum(2020, 2))

	log.Println("1-2:", r.ProductOfCombinationWithSum(2020, 3))
}

func day2() {
	db, err := advent.ReadDatabase(input(2))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("2-1:", db.CountValid(advent.SledPolicy))

	log.Println("2-2:", db.CountValid(advent.TobogganPolicy))
}

func day3() {
	m, err := advent.ReadMap(input(3))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("3-1:", m.Count(advent.Slope{DX: 3, DY: 1}))

	slopes := []advent.Slope{
		{DX: 1, DY: 1},
		{DX: 3, DY: 1},
		{DX: 5, DY: 1},
		{DX: 7, DY: 1},
		{DX: 1, DY: 2},
	}

	log.Println("3-2:", m.ProductOfCounts(slopes))
}

func day4() {
	p, _ := advent.ReadPassports(input(4))

	log.Println("4-1:", advent.ValidPassports(p, advent.RequireFields))

	log.Println("4-2:", advent.ValidPassports(p, advent.RequireValidFields))
}

func day5() {
	seats, err := advent.ReadSeats(input(5))
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(seats, func(i, j int) bool { return seats[i].ID < seats[j].ID })

	log.Println("5-1:", seats[len(seats)-1].ID)

	for i, s := range seats[1:] {
		if s.ID == (i + seats[0].ID + 2) {
			log.Println("5-2:", i+seats[0].ID+1)
			break
		}
	}
}

func day6() {
	groups, err := advent.ReadDeclarations(input(6))
	if err != nil {
		log.Fatal(err)
	}

	var any, all int

	for _, g := range groups {
		any += g.Any()
		all += g.All()
	}

	log.Println("6-1:", any)
	log.Println("6-2:", all)
}

func day7() {
	rules, err := advent.ReadRules(input(7))
	if err != nil {
		log.Fatal(err)
	}

	containing := rules.FindContaining("shiny gold")
	log.Println("7-1:", len(containing))

	contents := rules.ContentsOf("shiny gold")
	log.Println("7-2:", len(contents))
}

func day8() {
	p, err := advent.LoadProgram(input(8))
	if err != nil {
		log.Fatal(err)
	}

	cpu := &advent.CPU{}

	cpu.Execute(*p, cpu.DetectLoop())
	log.Println("8-1:", cpu.ACC)

	for ii, i := range *p {
		switch i.Op {
		case "jmp":
			(*p)[ii].Op = "nop"
		case "nop":
			(*p)[ii].Op = "jmp"
		default:
			continue
		}

		cpu := &advent.CPU{}

		if cpu.Execute(*p, cpu.DetectLoop()) {
			log.Println("8-2:", cpu.ACC)
			break
		}

		(*p)[ii].Op = i.Op
	}
}

func day9() {
	x, err := advent.LoadXMAS(25, input(9))
	if err != nil {
		log.Fatal(err)
	}

	n := x.FindInvalid()

	log.Println("9-1:", n)

	min, max := x.FindRangeMinMax(n)
	log.Println("9-2:", min+max)
}

func day10() {
	a, err := advent.ReadAdapters(input(10))
	if err != nil {
		log.Fatal(err)
	}

	c := a.CountAdapters()
	log.Println("10-1:", c[1]*c[3])

	adapter := advent.Adapter{}

	n := adapter.CountChains(a)
	log.Println("10-2:", n)
}

func main() {
	day1()
	day2()
	day3()
	day4()
	day5()
	day6()
	day7()
	day8()
	day9()
	day10()
}
