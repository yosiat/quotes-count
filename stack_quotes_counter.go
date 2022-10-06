package main

type stackQuotesCounter struct {
	eventIndex int    // current byte index into the event
	event      []byte // event being processed, treated as immutable
}

func (qc stackQuotesCounter) countQuotesWithFor(event []byte) int {
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

func (qc stackQuotesCounter) countQuotesWithForIncrementIndex(event []byte) int {
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
