package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/sbfaulkner/adventofcode/pkg/report"
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

func main() {
	r, err := report.NewReport(input(1))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("1-1: ", r.ProductOfCombinationWithSum(2020, 2))

	log.Println("1-2: ", r.ProductOfCombinationWithSum(2020, 3))
}
