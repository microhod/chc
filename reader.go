package main

import (
	"bufio"
	"io"
	"regexp"
	"sort"
)

type CharCount struct {
	Char  rune
	Count int
}

type Reader struct {
	include *regexp.Regexp
	counts  map[rune]int
}

func NewReader(include *regexp.Regexp) *Reader {
	return &Reader{
		include: include,
		counts:  make(map[rune]int),
	}
}

func (c *Reader) Read(r io.Reader) error {
	reader := bufio.NewReader(r)
	for {
		ch, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if !c.include.MatchString(string(ch)) {
			continue
		}
		c.counts[ch]++
	}
	return nil
}

func (c *Reader) Counts() []CharCount {
	sorted := make([]CharCount, 0, len(c.counts))
	for ch, count := range c.counts {
		sorted = append(sorted, CharCount{Char: ch, Count: count})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count == sorted[j].Count {
			return sorted[i].Char > sorted[j].Char
		}
		return sorted[i].Count > sorted[j].Count
	})
	return sorted
}
