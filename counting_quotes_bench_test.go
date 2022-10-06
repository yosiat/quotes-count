package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	topCount int
	event    []byte
)

func init() {
	var err error
	event, err = os.ReadFile("./testdata/status.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to read file: %s", err))
	}
}

func Benchmark_QuotesCount_Heap_Stepping(b *testing.B) {

	var localCount int

	counter := &heapQuotesCounter{}
	count := counter.countQuotesWithStepping(event)
	assertCount(b, count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := counter.countQuotesWithStepping(event)
		localCount = count
	}

	topCount = localCount
}

func Benchmark_QuotesCount_Heap_For_LocalI(b *testing.B) {
	var localCount int

	counter := &heapQuotesCounter{}
	count := counter.countQuotesWithFor(event)
	assertCount(b, count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := counter.countQuotesWithFor(event)
		localCount = count
	}

	topCount = localCount
}

func Benchmark_QuotesCount_Heap_For_EventIndex(b *testing.B) {
	var localCount int

	counter := &heapQuotesCounter{}
	count := counter.countQuotesWithForIncrementIndex(event)
	assertCount(b, count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := counter.countQuotesWithForIncrementIndex(event)
		localCount = count
	}

	topCount = localCount
}

func Benchmark_QuotesCount_Heap_Next(b *testing.B) {
	var localCount int

	counter := &heapQuotesCounter{}
	count := counter.countQuotesWithFor(event)
	assertCount(b, count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := counter.countQuotesWithNext(event)
		localCount = count
	}

	topCount = localCount
}

func Benchmark_QuotesCount_Stack_For_LocalI(b *testing.B) {
	var localCount int

	counter := stackQuotesCounter{}
	count := counter.countQuotesWithFor(event)
	assertCount(b, count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := counter.countQuotesWithFor(event)
		localCount = count
	}

	topCount = localCount
}

func Benchmark_QuotesCount_Stack_For_EventIndex(b *testing.B) {
	var localCount int

	counter := stackQuotesCounter{}
	count := counter.countQuotesWithForIncrementIndex(event)
	assertCount(b, count)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		count := counter.countQuotesWithForIncrementIndex(event)
		localCount = count
	}

	topCount = localCount
}

func assertCount(b *testing.B, count int) {
  b.Helper()

	if count != 642 {
		b.Fatalf("Expected quotes count to be %d but it is %d", 642, count)
	}

}
