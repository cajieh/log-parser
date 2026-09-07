package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func parse(r io.Reader) (parser, error) {
	parsed := parser{sum: make(map[string]result)}
	in := bufio.NewScanner(r)

	for in.Scan() {
		parsed.lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			return parser{}, fmt.Errorf("invalid input: %q (line #%d)", in.Text(), parsed.lines)
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			return parser{}, fmt.Errorf("invalid number of visits: %q (line #%d)", fields[1], parsed.lines)
		}

		if _, ok := parsed.sum[domain]; !ok {
			parsed.domains = append(parsed.domains, domain)
		}
		parsed.total += visits
		parsed.sum[domain] = result{domain: domain, visits: visits + parsed.sum[domain].visits}
	}

	if err := in.Err(); err != nil {
		return parser{}, fmt.Errorf("reading standard input: %w", err)
	}

	return parsed, nil
}
