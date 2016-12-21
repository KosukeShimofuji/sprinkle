package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"
const verb_file = "verb.list"
const noun_file = "noun.list"

func gen_transforms() []string {
	var transforms []string
	// Generation base on verb
	fp, err := os.Open(verb_file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		transforms = append(transforms, scanner.Text()+otherWord)
	}

	fp.Close()

	// Generation base on noun
	fp, err = os.Open(noun_file)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(fp)
	for scanner.Scan() {
		transforms = append(transforms, otherWord+scanner.Text())
	}
	return transforms
}

func main() {
	transforms := gen_transforms()
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
