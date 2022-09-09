package main

import (
	"reflect"
	"testing"
)

func TestBuildTrie(t *testing.T) {
	tt := []struct {
		name  string
		input []string
	}{
		{
			name:  "valid: with multiple prefixes",
			input: []string{"mob", "mo", "money", "monitr", "mouse", "mousep"},
		},
		{
			name:  "valid: with one prefix",
			input: []string{"hav"},
		},
		{
			name:  "valid: with no prefix",
			input: []string{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			trieRoot := buildTrie(tc.input)
			// verify
			got := []string{}

			for i := 0; i < len(tc.input); i++ {
				cur := trieRoot
				pfx := ""
				for _, c := range tc.input[i] {
					if cur.Children[string(c)] != nil {
						pfx += string(c)
					}
					ch := string(c)
					cur = cur.Children[ch]
				}
				got = append(got, pfx)
			}

			if tc.input != nil && !reflect.DeepEqual(tc.input, got) {
				t.Fatalf("Failed testcase #1 want: %v got %v", tc.input, got)
			}
			if len(tc.input) != len(got) {
				t.Fatalf("Failed testcase #1 want: %d got %d", len(tc.input), len(got))
			}

		})
	}
}
