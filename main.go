package main

import (
	"fmt"
	"go-log-parser/greeting"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(greeting.Message())

	p, err := parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if err := printReport(os.Stdout, p); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// printReport writes the parsed totals in a stable alphabetical order.
func printReport(w io.Writer, p parser) error {
	domains := make([]string, 0, len(p.sum))
	for domain := range p.sum {
		domains = append(domains, domain)
	}
	sort.Strings(domains)

	if _, err := fmt.Fprintf(w, "%-30s %10s\n", "Domain", "Visit counts:"); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "%-30s %10s\n", strings.Repeat("-", 30), strings.Repeat("-", 10)); err != nil {
		return err
	}
	for _, domain := range domains {
		if _, err := fmt.Fprintf(w, "%-30s %10d\n", domain, p.sum[domain]); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintf(w, "\n%-30s %10d\n", "Total", p.total)
	return err
}
