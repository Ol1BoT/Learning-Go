package main

import "fmt"

// A type should have one primary responsibility
// Random Tidbits:
// 	Separation of Concerns
//		e.g. Save, load etc should be separate
//	Avoid God Object

var entryCount = 0

type journal struct {
	entries []string
}

func (j *journal) addEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func main() {

	j := journal{}
	j.addEntry("Today I went to work")
	j.addEntry("I then went to the gym")
	fmt.Println(j.entries)
}
