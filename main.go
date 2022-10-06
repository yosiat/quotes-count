package main

import "fmt"

type test struct {
	eventIndex int    // current byte index into the event
	event      []byte // event being processed, treated as immutable
}

func (t test) increment() {
	t.eventIndex++
  fmt.Println(t.eventIndex)
}

func main() {
	t := test{}
	fmt.Printf("Before: %+v\n", t)
	t.increment()
	//t.eventIndex = 123
	fmt.Printf("After: %+v\n", t)
}
