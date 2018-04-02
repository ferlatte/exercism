package letter

import "sync"

// The FreqMap type gives you the mapping of runes to frequency
type FreqMap map[rune]int

// Frequency returns the mapping of runes to frequency.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency tries to calculate frequency using some goroutines.
func ConcurrentFrequency(texts []string) FreqMap {
	return concurrentFrequencyMapReduce(texts)
	// return concurrentFrequencyWithLocks(texts)
}

// goos: darwin
// goarch: amd64
// BenchmarkSequentialFrequency-8   	   50000	     33804 ns/op
// BenchmarkConcurrentFrequency-8   	   50000	     32557 ns/op
// PASS
func concurrentFrequencyMapReduce(texts []string) FreqMap {
	var chans []chan FreqMap
	var helper func(string, chan<- FreqMap)
	helper = func(s string, c chan<- FreqMap) {
		m := Frequency(s)
		c <- m
	}
	for _, t := range texts {
		c := make(chan FreqMap)
		chans = append(chans, c)
		go helper(t, c)
	}
	r := FreqMap{}
	for _, c := range chans {
		f := <-c
		for k, v := range f {
			r[k] += v
		}
	}
	return r
}

type safeFreqMap struct {
	v   FreqMap
	mux sync.Mutex
}

// goos: darwin
// goarch: amd64
// BenchmarkSequentialFrequency-8   	   50000	     33163 ns/op
// BenchmarkConcurrentFrequency-8   	   20000	     63481 ns/op
// PASS
func concurrentFrequencyWithLocks(texts []string) FreqMap {
	var f safeFreqMap
	f.v = FreqMap{}
	var chans []chan bool
	var helper func(string, chan<- bool)
	helper = func(s string, c chan<- bool) {
		for _, r := range s {
			f.mux.Lock()
			f.v[r]++
			f.mux.Unlock()
		}
		c <- true
	}
	for _, t := range texts {
		c := make(chan bool)
		chans = append(chans, c)
		go helper(t, c)
	}
	for _, c := range chans {
		<-c
	}
	return f.v
}
