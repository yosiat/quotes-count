package main

import "io"

type heapQuotesCounter struct {
	eventIndex int    // current byte index into the event
	event      []byte // event being processed, treated as immutable
}

func (qc *heapQuotesCounter) countQuotesWithStepping(event []byte) int {
	qc.eventIndex = 0
	qc.event = event

	count := 0

	for {
		ch := qc.ch()

		if ch == '"' {
			count++
		}

		if qc.step() != nil {
			return count
		}
	}
}

func (qc *heapQuotesCounter) countQuotesWithNext(event []byte) int {
	qc.eventIndex = 0
	qc.event = event

	count := 0

	for qc.next() {
		ch := qc.ch()

		if ch == '"' {
			count++
		}

		qc.step()
	}

	return count
}

func (qc *heapQuotesCounter) countQuotesWithFor(event []byte) int {
	qc.eventIndex = 0
	qc.event = event

	count := 0
	i := 0

	for i < len(qc.event) {
		ch := qc.event[i]

		if ch == '"' {
			count++
		}

		i++
	}

	return count
}

func (qc *heapQuotesCounter) countQuotesWithForIncrementIndex(event []byte) int {
	qc.eventIndex = 0
	qc.event = event

	count := 0

	for qc.eventIndex < len(qc.event) {
		ch := qc.event[qc.eventIndex]

		if ch == '"' {
			count++
		}

		qc.eventIndex++
	}

	return count
}

// ch fetches the next byte from the event. It doesn't check array bounds,
// so it's the caller's responsibility to ensure we haven't run off the end of the event.
func (qc *heapQuotesCounter) ch() byte {
	return qc.event[qc.eventIndex]
}

func (qc *heapQuotesCounter) next() bool {
	return qc.eventIndex < len(qc.event)
}

func (qc *heapQuotesCounter) inc() {
	qc.eventIndex++
}

// step advances the event pointer and returns an error if you've run off the end of the event
func (qc *heapQuotesCounter) step() error {
	qc.eventIndex++

	if qc.eventIndex >= len(qc.event) {
		// In real qumina we return good error, since we are reading till end of input
		// we can return EOF, which is more efficient than doing `fmt.Errorf`
		return io.EOF
	}

	return nil
}
