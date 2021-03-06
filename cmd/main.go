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
	bps, err := advent.ReadBoardingPasses(input(5))
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(bps, func(i, j int) bool { return bps[i].ID < bps[j].ID })

	log.Println("5-1:", bps[len(bps)-1].ID)

	for i, bp := range bps[1:] {
		if bp.ID == (i + bps[0].ID + 2) {
			log.Println("5-2:", i+bps[0].ID+1)
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

func day11() {
	seating, err := advent.ReadSeating(input(11))
	if err != nil {
		log.Fatal(err)
	}

	var occupied int
	var evolved bool

	for s := seating; ; {
		s, occupied, evolved = s.Evolve(1, 4)
		if !evolved {
			break
		}
	}

	log.Println("11-1:", occupied)

	for s := seating; ; {
		s, occupied, evolved = s.Evolve(len(*seating), 5)
		if !evolved {
			break
		}
	}

	log.Println("11-2:", occupied)
}

func day12() {
	f := advent.NewFerry()

	err := f.Navigate(input(12), advent.SimpleNavigator())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("12-1:", f.ManhattanDistance())

	f = advent.NewFerry()

	err = f.Navigate(input(12), advent.WaypointNavigator(10, 1))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("12-2:", f.ManhattanDistance())
}

func day13() {
	n := advent.ScheduleNotes{}

	if err := n.Read(input(13)); err != nil {
		log.Fatal(err)
	}

	id, wait := n.FindBus()

	log.Println("13-1:", id*wait)

	log.Println("13-2:", n.FindSyncTime())
}

func day14() {
	dp := advent.DockingProgram{}

	if err := dp.Initialize(input(14), &advent.DecoderV1{}); err != nil {
		log.Fatal(err)
	}

	log.Println("14-1:", dp.CheckSum())

	dp = advent.DockingProgram{}

	if err := dp.Initialize(input(14), &advent.DecoderV2{}); err != nil {
		log.Fatal(err)
	}

	log.Println("14-2:", dp.CheckSum())
}

func day15() {
	game := advent.NewMemoryGame(0, 12, 6, 13, 20, 1, 17)

	log.Println("15-1:", game.Play(uint(2020)))
	log.Println("15-2:", game.Play(uint(30000000)))
}

func day16() {
	notes, err := advent.ReadTicketNotes(input(16))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("16-1:", notes.TicketScanningErrorRate())

	log.Println("16-2:", notes.GetDepartureFieldCheckSum())
}

func day17() {
	var count int

	d, err := advent.ReadPocketDimension(input(17))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 6; i++ {
		count = d.Cycle(3)
	}

	log.Println("17-1:", count)

	d, err = advent.ReadPocketDimension(input(17))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 6; i++ {
		count = d.Cycle(4)
	}

	log.Println("17-2:", count)
}

func day18() {
	h, err := advent.ReadMathHomework(input(18))
	if err != nil {
		log.Fatal(err)
	}

	s := 0

	for _, p := range h {
		i, err := p.Evaluate(advent.LeftToRightPrecedence)
		if err != nil {
			log.Fatal(err)
		}
		s += i
	}

	log.Println("18-1:", s)

	s = 0

	for _, p := range h {
		i, err := p.Evaluate(advent.AdvancedMathPrecedence)
		if err != nil {
			log.Fatal(err)
		}
		s += i
	}

	log.Println("18-2:", s)
}

func day19() {
	r, m := advent.ReadMessages(input(19))
	log.Println("19-1:", r.CountValid(m))
}

func main() {
	// day1()
	// day2()
	// day3()
	// day4()
	// day5()
	// day6()
	// day7()
	// day8()
	// day9()
	// day10()
	// day11()
	// day12()
	// day13()
	// day14()
	// day15()
	// day16()
	// day17()
	// day18()
	day19()
}
