package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"
)

func main() {
    counts := make(map[string]map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines("stdin", os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup: %v\n", err)
                continue
            }
            countLines(arg, f, counts)
            f.Close()
        }
    }

    for text, results := range counts {
        var files []string
        occurrences := 0
        for file, o := range results {
            files = append(files, file)
            occurrences += o
        }
        if occurrences > 1 {
            fmt.Printf("%d\t%s\t%s\n", occurrences, text, strings.Join(files, ","))
        }
    }
}

func countLines(filename string, f *os.File, counts map[string]map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        text := input.Text()
        if counts[text] == nil {
            counts[text] = make(map[string]int)
        }
        counts[text][filename]++
    }
    // Note: ignoring potential errors from input.Err()
}
