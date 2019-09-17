// Sergey Olisov (c) 2019
// Lesson 3

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Word is a struct which represents pair of word and it's frequency of appears.
type Word struct {
	word      string // specific word
	frequency int    // number of appearence of word in the text
}

// Words is a slice of Word struct used for sorting
type Words []Word // array of words sorted by number of apperance

// Three functions below are used for fullfill the sort.Sort interface contract.
// Len return lenght of array words
func (w Words) Len() int {
	return len(w)
}

// Less compare two elements of array words by frequency.
func (w Words) Less(i, j int) bool {
	// Additional check is added to avoid unpredictable result of Sort function when the frequency of word are the same.
	return w[i].frequency < w[j].frequency || (w[i].frequency == w[j].frequency && w[i].word > w[j].word)
}

// Swap two elements of array words
func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

// TextFrq does frequency analysis of text in terms of calculation of apperance freuncy of each word.
// TextFrq output the top 10 words sorted by apperance. In case of imput text is to short the error or
// warning has returned. I.e. Error - when the input text is empty. Waring in case of lenght of imput
// text less then 10 words.
func TextFrq(s string) Words {
	re := regexp.MustCompile(`\w{2,}`)               // Considering the word is greater then one character. Can be tuned.
	text := re.FindAllString(strings.ToLower(s), -1) // All input passed through ToLower in order to filter out such cases as Word != word
	mapWords := map[string]int{}
	for _, w := range text {
		mapWords[w]++
	}
	sliceOfWords := make(Words, len(mapWords))
	i := 0
	for w, f := range mapWords {
		sliceOfWords[i] = Word{w, f}
		i++
	}
	sort.Sort(sort.Reverse(sliceOfWords))
	if len(sliceOfWords) <= 10 {
		return sliceOfWords
	}
	return sliceOfWords[:10]
}
func main() {
	testText := "aaaa bbb aaaa bbb cccc ddddd aaaa xxxx zzzzz fffff rrrrr eeee aaaa bbb rrrrr"
	fmt.Println(TextFrq(testText))
}
