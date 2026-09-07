package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	"go-log-parser/greeting"
)

func main() {
	fmt.Println(greeting.Message())
	
	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s < input.txt\n", os.Args[0])
		return
	}

	p, err := parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("%-30s %10s\n", "Domain", "Visit counts:")
	fmt.Printf("%-30s %10s\n", strings.Repeat("-", 30), strings.Repeat("-", 10))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "Total", p.total)
}
