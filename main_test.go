package main

// go test -run none -bench . -benchtime 3s -benchmem

import (
	"fmt"
	"testing"
)

var (
	records [][]string
	err     error
)

func BenchmarkReadParse(b *testing.B) {
	b.ResetTimer()

	records, err = readData(DataFile)
	if err != nil {
		fmt.Printf("Encountered error after %d records: %s", len(records), err)
	}

	//for i := 0; i < 10; i++ {
	parseData(records)
	//}
}
