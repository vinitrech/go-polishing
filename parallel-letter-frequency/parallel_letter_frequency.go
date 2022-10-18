package letter

import "sync"

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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {

	sum := FreqMap{}
	channels := make(chan FreqMap, len(l))
	var wg sync.WaitGroup

	wg.Add(len(l))

	for _, x := range l {
		go func(x string) {
			channels <- Frequency(x)
			wg.Done()
		}(x)
	}

	go func() {
		wg.Wait()
		defer close(channels)
	}()

	for c := range channels {
		for index, value := range c {
			sum[index] += value
		}
	}

	return sum
}
