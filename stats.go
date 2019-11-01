package main

// Statistics represents cahce statistics
type Statistics struct {
	Writes uint64 // Number of writes
	Reads  uint64 // Numebr of reads
	Hits   uint64 // Number of Hits
	Misses uint64 // Number of Misses
}

// MissRatio returns the miss ration of this set of stats, eg,
// the number of misses divided by the total number of hits and misses
func (s *Statistics) MissRatio() float64 {
	return float64(s.Misses) / float64(s.Hits+s.Misses)
}
