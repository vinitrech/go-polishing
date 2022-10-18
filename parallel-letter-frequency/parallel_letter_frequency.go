package letter

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
	defer close(channels)

	for _, x := range l {
		go func(x string) {
			items := FreqMap{}
			for _, r := range x {
				items[r]++
			}
			channels <- items
		}(x)
	}

	for x := 0; x < len(l); x++ {
		runeMap := <-channels

		for index, value := range runeMap {
			sum[index] += value
		}
	}

	return sum
}
