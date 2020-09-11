package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"sync"
	"unicode"
)

func parseArgs() (int, string) {
	threads := flag.Int("n", 4, "a string")
	filepath := flag.String("p", "data.txt", "a string")

	flag.Parse()

	return *threads, *filepath
}

func readTextFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func countLetter(letter rune, filepath string) (int, error) {
	data, err := readTextFile(filepath)
	if err != nil {
		return 0, err
	}

	var count int
	for _, c := range data {
		if unicode.ToUpper(c) == unicode.ToUpper(letter) {
			count++
		}
	}

	return count, nil
}

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Counter struct {
	Letter rune
	Total  int
	Err    error
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	threads, filepath := parseArgs()

	result := []*Counter{}
	mutex := &sync.Mutex{}

	in := make(chan rune, threads)
	errc := make(chan error, threads)

	wait := &sync.WaitGroup{}
	wait.Add(threads)

	quit := make(chan bool)
	var err error
	go func() {
		select {
		case <-quit:
			return
		case e := <-errc:
			err = e
			close(quit)
		}
	}()

	for i := 0; i < threads; i++ {
		go func() {
			defer wait.Done()

			for {
				select {
				case <-quit:
					return
				case letter, ok := <-in:
					if !ok {
						return
					}

					total, err := countLetter(letter, filepath)
					if err != nil {
						errc <- err
						return
					}

					mutex.Lock()
					result = append(result, &Counter{
						Letter: letter,
						Total:  total,
					})
					mutex.Unlock()
				}
			}

		}()
	}

	go func() {
		for _, l := range Alphabet {
			select {
			case <-quit:
				return
			case in <- l:
			}
		}
		close(in)
	}()

	wait.Wait()

	if err != nil {
		fmt.Println("Got error ", err)
		return
	}

	close(quit)
	showResult(result)
}

func showResult(result []*Counter) {

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Letter < result[j].Letter
	})

	for _, l := range result {
		fmt.Printf("%c :  %d \n", l.Letter, l.Total)
	}
}
