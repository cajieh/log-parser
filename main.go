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

func main() {
	fmt.Println(greeting.Message())
	

	var (
		sum map[string]int
		domains []string
		total int
		lines int
	)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "invalid input: %q(line #%d)\n", in.Text(), lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Fprintf(os.Stderr, "invalid number of visits: %q(line #%d)\n", fields[1], lines)
			return
		}

		if sum == nil {
			sum = make(map[string]int)
		}
		
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits	
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "Domain", "visit counts:")

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "Total", total)
}
