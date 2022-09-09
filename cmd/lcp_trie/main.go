package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"

	"github.com/arpsch/truecaller/input/file"
	"github.com/arpsch/truecaller/trie"
)

func main() {

	inputFile := ""
	args := os.Args
	if len(args) < 2 {
		// default file
		inputFile = "./sample_prefixes.txt"
	} else {
		inputFile = args[1]
	}

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	quitCh := make(chan bool)

	go func(sigCh chan os.Signal, quitCh chan bool) {
		for {
			select {
			case <-sigCh:
				fmt.Println("Got interrupt signal")
				fmt.Println("Exiting signal handler..")
				os.Exit(1)
			case <-quitCh:
				fmt.Println("quit")
				os.Exit(1)
			}
		}

	}(sigCh, quitCh)

	prefixes := file.Read(f)
	trieRoot := buildTrie(prefixes)
	if trieRoot == nil {
		log.Fatal("failed to build trie")
	}

	for {

		fmt.Printf("Enter a word to search for longest matching prefix: ")
		var reader = bufio.NewReader(os.Stdin)
		searchWord, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		switch {
		case searchWord[:len(searchWord)-1] == "quit":
			quitCh <- true
		default:
			lcp := trieRoot.SearchLongestPrefix(searchWord)
			if lcp == "" {
				fmt.Println("Could not find a matching prefix")
			} else {
				fmt.Printf("Longest matched prefix: %s\n", lcp)
			}

		}
	}
}

func buildTrie(prefixes []string) *trie.Node {
	// sort prefixes by length
	sort.Slice(prefixes, func(i, j int) bool {
		return len(prefixes[i]) < len(prefixes[j])
	})

	root := trie.NewNode()
	for _, prefix := range prefixes {
		root.Build(prefix)
	}

	return root
}
