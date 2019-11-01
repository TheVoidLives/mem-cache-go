package main

import (
	"fmt"
	"testing"
)

// InitTest32 tests Cache Initialization (32KB)
func TestInit32(t *testing.T) {
	options := &Options{
		Size:        32768, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write back
		Debug:       true,  // Debug Tracing
	}
	cache := &Cache{Options: options}
	err := cache.Init()
	fmt.Printf("\nError: %v\n", err)
}

// InitTest64 tests Cache Initialization (64KB)
func TestInit64(t *testing.T) {
	options := &Options{
		Size:        65536, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write back
		Debug:       true,  // Debug Tracing
	}

	cache := &Cache{Options: options}
	err := cache.Init()
	fmt.Printf("\nError: %v\n", err)
}

// ParseTest32 tests the cache address parsing with a 32KB Cache
func TestParse32(t *testing.T) {
	options := &Options{
		Size:        32768, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write back
		Debug:       false, // Debug Tracing
	}

	cache := &Cache{Options: options}
	err := cache.Init()

	tag, set, offset, err := cache.Parse("0x62b8")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n\n           Expected - Result\n")
	fmt.Printf("%7s:   %-10s - %b\n", "Tag", "1", tag)
	fmt.Printf("%7s:   %-10s - %b\n", "Set", "10001010", set)
	fmt.Printf("%7s:   %-10s - %b\n", "Offset", "111000", offset)
}

// ParseTest32 tests the cache address parsing with a 32KB Cache
func TestParseBreak32(t *testing.T) {
	options := &Options{
		Size:        32768, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write back
		Debug:       false, // Debug Tracing
	}

	cache := &Cache{Options: options}
	err := cache.Init()

	tag, set, offset, err := cache.Parse("0x62b8")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n\n           Expected - Result\n")
	fmt.Printf("%7s:   %-10s - %b\n", "Tag", "1", tag)
	fmt.Printf("%7s:   %-10s - %b\n", "Set", "10001010", set)
	fmt.Printf("%7s:   %-10s - %b\n", "Offset", "111000", offset)
}

// ParseTest64 tests the cache address parsing with a 32KB Cache
func TestParse64(t *testing.T) {
	options := &Options{
		Size:        65536, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write back
		Debug:       false, // Debug Tracing
	}

	cache := &Cache{Options: options}
	err := cache.Init()

	tag, set, offset, err := cache.Parse("0x62b8")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n\n           Expected - Result\n")
	fmt.Printf("%7s:   %-10s - %b\n", "Tag", "0", tag)
	fmt.Printf("%7s:   %-10s - %b\n", "Set", "10001010", set)
	fmt.Printf("%7s:   %-10s - %b\n", "Offset", "111000", offset)
}

func TestWriteBack32(t *testing.T) {
	options := &Options{
		Size:        32768, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write back
		Debug:       false, // Debug Tracing
	}

	cache := &Cache{Options: options}
	_ = cache.Init()

	err := cache.Execute(WRITE, "0x62b8")
	if err != nil {
		t.Errorf("%v", err)
	}

	// validation
	if cache.Stats.Misses != 1 {
		t.Errorf("Misses incorrect: %v", cache.Stats.Misses)
	} else if cache.Stats.Hits > 0 {
		t.Errorf("Hits incorrect: %v", cache.Stats.Hits)
	} else if cache.Stats.Reads > 0 {
		t.Errorf("Reads incorrect: %v", cache.Stats.Hits)
	} else if cache.Stats.Writes != 0 {
		t.Errorf("Writes incorrect: %v", cache.Stats.Hits)
	}
}

func TestWriteThrough32(t *testing.T) {
	options := &Options{
		Size:        32768, // 32 KB
		Assoc:       2,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   1,     // Write Through
		Debug:       false, // Debug Tracing
	}

	cache := &Cache{Options: options}
	_ = cache.Init()
	cache.Options.Debug = true

	err := cache.Execute(WRITE, "0x62b8")
	if err != nil {
		t.Errorf("%v", err)
	}

	// validation
	if cache.Stats.Misses != 1 {
		t.Errorf("Misses incorrect: %v", cache.Stats.Misses)
	} else if cache.Stats.Hits > 0 {
		t.Errorf("Hits incorrect: %v", cache.Stats.Hits)
	} else if cache.Stats.Reads > 0 {
		t.Errorf("Reads incorrect: %v", cache.Stats.Reads)
	} else if cache.Stats.Writes != 1 {
		t.Errorf("Writes incorrect: %v", cache.Stats.Writes)
	}

	err = cache.Execute(WRITE, "0x62b8")
	if err != nil {
		t.Errorf("%v", err)
	}

	// validation
	if cache.Stats.Misses != 1 {
		t.Errorf("Misses incorrect: %v", cache.Stats.Misses)
	} else if cache.Stats.Hits != 1 {
		t.Errorf("Hits incorrect: %v", cache.Stats.Hits)
	} else if cache.Stats.Reads > 0 {
		t.Errorf("Reads incorrect: %v", cache.Stats.Reads)
	} else if cache.Stats.Writes != 2 {
		t.Errorf("Writes incorrect: %v", cache.Stats.Writes)
	}

	err = cache.Execute(READ, "0x62b8")
	if err != nil {
		t.Errorf("%v", err)
	}

	// validation
	if cache.Stats.Misses != 1 {
		t.Errorf("Misses incorrect: %v", cache.Stats.Misses)
	} else if cache.Stats.Hits != 2 {
		t.Errorf("Hits incorrect: %v", cache.Stats.Hits)
	} else if cache.Stats.Reads != 0 {
		t.Errorf("Reads incorrect: %v", cache.Stats.Reads)
	} else if cache.Stats.Writes != 2 {
		t.Errorf("Writes incorrect: %v", cache.Stats.Writes)
	}

	err = cache.Execute(READ, "0x63b8")
	if err != nil {
		t.Errorf("%v", err)
	}

	// validation
	if cache.Stats.Misses != 2 {
		t.Errorf("Misses incorrect: %v", cache.Stats.Misses)
	} else if cache.Stats.Hits != 2 {
		t.Errorf("Hits incorrect: %v", cache.Stats.Hits)
	} else if cache.Stats.Reads != 1 {
		t.Errorf("Reads incorrect: %v", cache.Stats.Reads)
	} else if cache.Stats.Writes != 2 {
		t.Errorf("Writes incorrect: %v", cache.Stats.Writes)
	}
}

func TestLRU(t *testing.T) {
	options := &Options{
		Size:        32768, // 32 KB
		Assoc:       8,     // 2 way
		BlockSize:   64,    // B
		Replacement: 0,     // LRU
		WriteBack:   0,     // Write Through
		Debug:       false, // Debug Tracing
	}

	// Init
	cache := &Cache{Options: options}
	_ = cache.Init()
	cache.Options.Debug = true

	// Execute writes
	vars := make([]string, 0)
	vars = append(vars, "0x1fff6c5b7b80")
	vars = append(vars, "0x2fff6c5b7b80")
	vars = append(vars, "0x3fff6c5b7b80")
	vars = append(vars, "0x4fff6c5b7b80")
	vars = append(vars, "0x5fff6c5b7b80")
	vars = append(vars, "0x6fff6c5b7b80")
	vars = append(vars, "0x7fff6c5b7b80")
	vars = append(vars, "0x9fff6c5b7b80")
	vars = append(vars, "0xafff6c5b7b80")
	vars = append(vars, "0xafff6c5b7b80")
	vars = append(vars, "0x3fff6c5b7b80")

	for i := range vars {
		_ = cache.Execute(WRITE, vars[i])
	}
}
