package main

import (
	"fmt"
    "bufio"
	"os"
	"strconv"
	"strings"
	"sort"
	"go-log-parser/greeting"
)

type parser struct {
		sum map[string]int
	    domains []string
		total int
		lines int
}

func main() {
	fmt.Println(greeting.Message())
	
	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s < input.txt\n", os.Args[0])
		return
	}

	p := parser{sum: make(map[string]int)}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "invalid input: %q(line #%d)\n", in.Text(), p.lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Fprintf(os.Stderr, "invalid number of visits: %q(line #%d)\n", fields[1], p.lines)
			return
		}

		
		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
		p.total += visits	
		p.sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "Domain", "visit counts:")

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		visits := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "Total", p.total)


if err := in.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "reading standard input: %s\n", err)
}
}
