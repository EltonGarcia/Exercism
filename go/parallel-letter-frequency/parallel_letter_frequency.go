package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// SafeFreqMap safe records the frequency of each rune in a given text.
type SafeFreqMap struct {
	freqMap FreqMap
	mux     sync.Mutex
}

// ConcurrentFrequency counts the frequency of each rune concurrently in a given text
// and returns this data as a FreqMap.
func ConcurrentFrequency(arr []string) FreqMap {
	safeFreqMap := SafeFreqMap{freqMap: make(map[rune]int)}
	for _, s := range arr {
		SafeFrequency(s, &safeFreqMap)
	}
	safeFreqMap.mux.Lock()
	defer safeFreqMap.mux.Unlock()
	return safeFreqMap.freqMap
}

//SafeFrequency safe counts the frequency of runes in a given string
func SafeFrequency(s string, m *SafeFreqMap) {
	for _, r := range s {
		go m.Inc(r)
	}
}

//Inc increments the counter for the given key.
func (m *SafeFreqMap) Inc(key rune) {
	m.mux.Lock()
	m.freqMap[key]++
	m.mux.Unlock()
}
