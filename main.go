package main

import (
	"flag"
	"fmt"

	"github.com/MustafaNafizDurukan/HashWalker/pkg/constants"
	"github.com/MustafaNafizDurukan/HashWalker/pkg/csv"
	"github.com/MustafaNafizDurukan/HashWalker/pkg/hasher"
)

func main() {
	dirFlag := flag.String("dir", "", "Directory to walk and hash")
	baselineFlag := flag.String("baseline", "", "Baseline CSV file for comparison")
	csvFlag := flag.String("csv", "", "Second CSV file for comparison")
	flag.Parse()

	ok := walk(*dirFlag)
	if !ok {
		return
	}

	ok = compare(*baselineFlag, *csvFlag)
	if !ok {
		return
	}

	fmt.Println("Usage:")
	flag.PrintDefaults()
}

func walk(dir string) bool {
	if dir == "" {
		return true
	}

	fileHashes, err := hasher.Walk(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	err = csv.Save(constants.FileHashesName, fileHashes)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return false
}

func compare(baselinePath, csvPath string) bool {
	if baselinePath == "" || csvPath == "" {
		return true
	}

	baseline, err := csv.Read(baselinePath)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	fileHashes, err := csv.Read(csvPath)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	comparedHashes := csv.Compare(baseline, fileHashes)

	err = csv.Save(constants.ComparedHashesName, comparedHashes)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return false
}
