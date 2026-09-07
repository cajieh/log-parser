package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type result struct {
	// domain is kept with the count so each map entry contains a complete result.
	domain string
	// visits is the combined number of visits for this domain.
	visits int
}

type parser struct {
	// sum stores the accumulated result for each domain.
	sum map[string]result
	// domains keeps track of domains in the order they first appear.
	domains []string
	// total is the number of visits across all domains.
	total int
	// lines records how many input lines have been processed, including the current line.
	lines int
}

// parse reads log entries from r and returns their accumulated visit counts.
func parse(r io.Reader) (parser, error) {
	// A map lets us quickly find and update an existing domain.
	parsed := parser{sum: make(map[string]result)}
	in := bufio.NewScanner(r)

	for in.Scan() {
		// Count lines early so validation errors can identify their input line.
		parsed.lines++

		// Fields splits on whitespace, so extra spaces do not affect parsing.
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			return parser{}, fmt.Errorf("invalid input: %q (line #%d)", in.Text(), parsed.lines)
		}

		domain := fields[0]
		// A visit count must be a whole number and cannot be negative.
		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			return parser{}, fmt.Errorf("invalid number of visits: %q (line #%d)", fields[1], parsed.lines)
		}

		// Keep a domain once, when it first appears; main.go later sorts this list for display.
		if _, ok := parsed.sum[domain]; !ok {
			parsed.domains = append(parsed.domains, domain)
		}
		// Add this line's visits to both the domain total and the overall total.
		parsed.total += visits
		parsed.sum[domain] = result{domain: domain, visits: visits + parsed.sum[domain].visits}
	}

	// Scanner errors are different from invalid log data, so return them separately.
	if err := in.Err(); err != nil {
		return parser{}, fmt.Errorf("reading standard input: %w", err)
	}

	return parsed, nil
}
