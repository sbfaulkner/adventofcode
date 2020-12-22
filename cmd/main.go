package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/sbfaulkner/adventofcode/pkg/airline"
	"github.com/sbfaulkner/adventofcode/pkg/customs"
	"github.com/sbfaulkner/adventofcode/pkg/expense"
	"github.com/sbfaulkner/adventofcode/pkg/passport"
	"github.com/sbfaulkner/adventofcode/pkg/password"
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
	r, err := expense.ReadReport(input(1))
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
	p, _ := passport.ReadPassports(input(4))

	log.Println("4-1:", passport.ValidPassports(p, passport.RequireFields))

	log.Println("4-2:", passport.ValidPassports(p, passport.RequireValidFields))
}

func day5() {
	seats, err := airline.ReadSeats(input(5))
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
	groups, err := customs.ReadDeclarations(input(6))
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

func main() {
	day1()
	day2()
	day3()
	day4()
	day5()
	day6()
}
