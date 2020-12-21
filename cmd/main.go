package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/sbfaulkner/adventofcode/pkg/passport"
	"github.com/sbfaulkner/adventofcode/pkg/password"
	"github.com/sbfaulkner/adventofcode/pkg/report"
	"github.com/sbfaulkner/adventofcode/pkg/tree"
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
	r, err := report.ReadReport(input(1))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("1-1:", r.ProductOfCombinationWithSum(2020, 2))

	log.Println("1-2:", r.ProductOfCombinationWithSum(2020, 3))
}

func day2() {
	db, err := password.ReadDatabase(input(2))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("2-1:", db.CountValid(password.SledPolicy))

	log.Println("2-2:", db.CountValid(password.TobogganPolicy))
}

func day3() {
	m, err := tree.ReadMap(input(3))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("3-1:", m.Count(tree.Slope{DX: 3, DY: 1}))

	slopes := []tree.Slope{
		{DX: 1, DY: 1},
		{DX: 3, DY: 1},
		{DX: 5, DY: 1},
		{DX: 7, DY: 1},
		{DX: 1, DY: 2},
	}

	log.Println("3-2:", m.ProductOfCounts(slopes))
}

func day4() {
	p1, _ := passport.ReadPassports(input(4), passport.RequireFields)

	log.Println("4-1:", len(p1))

	p2, _ := passport.ReadPassports(input(4), passport.RequireValidFields)

	log.Println("4-2:", len(p2))
}

func main() {
	day1()
	day2()
	day3()
	day4()
}
