package main

// Define a struct for a pair of integers
type IntPair struct {
	First  int
	Second int
}

// Define a set using a map to store unique pairs
type IntPairSet struct {
	pairs map[IntPair]bool
}

// Function to create a new IntPairSet
func NewIntPairSet() *IntPairSet {
	return &IntPairSet{pairs: make(map[IntPair]bool)}
}

// Function to add a pair to the set
func (set *IntPairSet) AddPair(first, second int) {
	pair := IntPair{First: first, Second: second}
	set.pairs[pair] = true
}

// Function to check if a pair exists in the set
func (set *IntPairSet) Exists(first, second int) bool {
	pair := IntPair{First: first, Second: second}
	return set.pairs[pair]
}

// Function to list all pairs in the set
func (set *IntPairSet) ListPairs() []IntPair {
	keys := make([]IntPair, 0, len(set.pairs))
	for k := range set.pairs {
		keys = append(keys, k)
	}
	return keys
}

// Function to get length
func (set *IntPairSet) Length() int {
	return len(set.pairs)
}
