package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type result struct {
	visits int
}

type parser struct {
	// sum stores the total visits for each domain.
	sum map[string]int
	// total stores visits across all domains.
	total int
}

// parse reads one "domain visits" entry per line and combines duplicate domains.
func parse(r io.Reader) (parser, error) {
	parsed := parser{sum: make(map[string]int)}
	in := bufio.NewScanner(r)
	lineNumber := 0

	for in.Scan() {
		lineNumber++

		// Fields accepts any amount of whitespace between the domain and count.
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			return parser{}, fmt.Errorf("invalid input: %q (line #%d)", in.Text(), lineNumber)
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		// Visit counts must be valid non-negative integers.
		if err != nil || visits < 0 {
			return parser{}, fmt.Errorf("invalid number of visits: %q (line #%d)", fields[1], lineNumber)
		}

		parsed.total += visits
		parsed.sum[domain] += visits
	}

	// Scanner errors indicate a problem reading the input, rather than invalid log data.
	if err := in.Err(); err != nil {
		return parser{}, fmt.Errorf("reading standard input: %w", err)
	}

	return parsed, nil
}
