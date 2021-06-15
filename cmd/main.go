package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/meagar/lev-go"
)

//go:embed english3.txt.gz
var words []byte

func main() {
	if len(os.Args) != 2 {
		fmt.Println("lev: Print suggestions for the given word, along with Levenshtein distance score")
		return
	}

	candidateWord := os.Args[1]

	type Result struct {
		word  string
		score int
	}

	results := []Result{}

	eachWord(func(word string) {
		results = append(results, Result{
			word:  word,
			score: lev.Distance(candidateWord, word),
		})
	})

	sort.Slice(results, func(i, j int) bool {
		return results[i].score < results[j].score
	})

	fmt.Printf("Suggestions for %q\n", candidateWord)

	for n := 0; n < 10; n++ {
		fmt.Printf("%d. %s (%d)\n", n, results[n].word, results[n].score)
	}
}

func eachWord(fn func(word string)) {
	reader, err := gzip.NewReader(bytes.NewReader(words))
	if err != nil {
		log.Panicf("Error creating gzip reader: %s", err)
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fn(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}
}
