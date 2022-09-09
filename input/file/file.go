package file

import (
	"bufio"
	"io"
)

// Read reads all from the given reader line-by-line
// returns all content
func Read(r io.Reader) []string {
	prefChan := make(chan string)

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			prefChan <- scanner.Text()
		}
		close(prefChan)
	}()

	contentBuf := []string{}
	for v := range prefChan {
		contentBuf = append(contentBuf, v)
	}

	return contentBuf
}
