package trie

import (
	"reflect"
	"sort"
	"testing"
)

func TestBuild(t *testing.T) {
	tt := []struct {
		name   string
		input  []string
		want   []string
		length int
	}{
		{
			name:   "valid: prefixes array, multiple elements",
			input:  []string{"mousepad", "mouse", "pad"},
			want:   []string{"mousepad", "mouse", "pad"},
			length: len([]string{"mousepad", "mouse", "pad"}),
		},
		{
			name:   "valid: prefixes array with one element",
			input:  []string{"mousepad"},
			want:   []string{"mousepad"},
			length: len([]string{"mousepad"}),
		},

		{
			name:   "invalid: prefix",
			input:  nil,
			want:   nil,
			length: 0,
		},
	}

	for _, tc := range tt {
		root := NewNode()
		t.Run(tc.name, func(t *testing.T) {
			for _, prefix := range tc.input {
				root.Build(prefix)
			}

			// verify
			got := []string{}

			for i := 0; i < len(tc.input); i++ {
				cur := root
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

			if tc.input != nil && !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("Failed testcase #1 want: %s got %s", tc.want, got)
			}
			if tc.length != len(got) {
				t.Fatalf("Failed testcase #1 want: %d got %d", tc.length, len(got))
			}

		})
	}
}

func TestSearchLongestPrefix(t *testing.T) {

	tt := []struct {
		name       string
		input      []string
		searchWord string
		want       string
	}{
		{
			name:       "valid: matching last prefix",
			input:      []string{"mob", "mo", "money", "monitr", "mouse", "mousep"},
			searchWord: "mousepad",
			want:       "mousep",
		},
		{
			name:       "valid: matching first prefix",
			input:      []string{"hav", "havina", "hava"},
			searchWord: "have",
			want:       "hav",
		},
		{
			name:       "valid: trace missing prefix",
			input:      []string{"hav", "havina", "hava"},
			searchWord: "havi",
			want:       "hav",
		},
		{
			name:       "valid: match prefix and word",
			input:      []string{"test", "testing", "true", "truecaller"},
			searchWord: "testinger",
			want:       "testing",
		},
	}

	for _, tc := range tt {
		//prepare for test
		sort.Slice(tc.input, func(i, j int) bool {
			return len(tc.input[i]) < len(tc.input[j])
		})

		root := NewNode()
		for _, prefix := range tc.input {
			root.Build(prefix)
		}

		// test
		t.Run(tc.name, func(t *testing.T) {
			if got := root.SearchLongestPrefix(tc.searchWord); tc.want != got {
				t.Fatalf("Failed: want: %s got %s", tc.want, got)
			}
		})
	}
}
