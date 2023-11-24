package main

import (
	"fmt"
	"testing"
)

var rxLines = []string{
	`{"Time":"2023-11-23T21:20:02.620626-08:00","Action":"output","Package":"github.com/kevinburke/whynocache/getenv","Output":"ok  \tgithub.com/kevinburke/whynocache/getenv\t(cached)\n"}`,
	`{"Time":"2023-11-23T21:32:12.262785-08:00","Action":"output","Package":"github.com/kevinburke/whynocache/getenv","Output":"ok  \tgithub.com/kevinburke/whynocache/getenv\t0.096s\n"}`,
}

func TestRxMatch(t *testing.T) {
	for _, line := range rxLines {
		matches := outputRx.FindStringSubmatch(line)
		if matches == nil {
			t.Fatalf("no matches found for line %q", line)
		}
		fmt.Printf("matches: %#v\n", matches)
	}
	t.Fail()
}
