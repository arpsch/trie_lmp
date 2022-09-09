package file

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := os.Open("./test_prefix.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	expContent := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		expContent = append(expContent, scanner.Text())
	}
	// seek to the start
	f.Seek(0, io.SeekStart)
	prefixes := Read(f)

	if len(prefixes) != len(expContent) {
		t.Fatalf("fail: expected len %d got len %d", len(expContent), len(prefixes))
	}
}
