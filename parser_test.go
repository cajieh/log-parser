package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseCombinesDuplicateDomains(t *testing.T) {
	input := "example.com 10\nexample.com 5\nother.com 3\n"

	parsed, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parse() returned an unexpected error: %v", err)
	}

	if parsed.sum["example.com"] != 15 {
		t.Errorf("example.com visits = %d, want 15", parsed.sum["example.com"])
	}
	if parsed.sum["other.com"] != 3 {
		t.Errorf("other.com visits = %d, want 3", parsed.sum["other.com"])
	}
	if parsed.total != 18 {
		t.Errorf("total visits = %d, want 18", parsed.total)
	}
}

func TestParseRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "wrong number of fields", input: "broken line with too many fields\n"},
		{name: "non-numeric visits", input: "example.com many\n"},
		{name: "negative visits", input: "example.com -1\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := parse(strings.NewReader(test.input)); err == nil {
				t.Fatal("parse() returned nil error for invalid input")
			}
		})
	}
}

func TestPrintReportSortsDomains(t *testing.T) {
	parsed := parser{
		sum:   map[string]int{"z.com": 2, "a.com": 4},
		total: 6,
	}
	var output bytes.Buffer

	if err := printReport(&output, parsed); err != nil {
		t.Fatalf("printReport() returned an unexpected error: %v", err)
	}

	if a, z := strings.Index(output.String(), "a.com"), strings.Index(output.String(), "z.com"); a == -1 || z == -1 || a > z {
		t.Errorf("report domains are not sorted: %q", output.String())
	}
}
