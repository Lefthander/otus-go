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
	return w[i].frequency < w[j].frequency
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
	re := regexp.MustCompile(`\w{2,}`)
	text := re.FindAllString(strings.ToLower(s), -1)
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
	sampletest2 := `Bacon ipsum dolor amet pork chop sunt swine, t-bone velit pastrami frankfurter adipisicing.
					Elit sunt sirloin sed eu. Eu qui swine cillum filet mignon ut prosciutto dolore.
					Deserunt reprehenderit pork, rump flank consequat magna consectetur. Mollit shoulder salami
					reprehenderit. Boudin flank turkey biltong beef porchetta. Brisket t-bone eiusmod beef ribs
					pig dolor ullamco spare ribs tail non strip steak shoulder sunt.
					Leberkas venison culpa prosciutto beef reprehenderit dolore strip steak. Eu pork chop irure
					drumstick in dolore. Commodo flank boudin, incididunt t-bone lorem sed strip steak eu.
					Tri-tip do et voluptate. In magna exercitation beef turkey quis, ball tip pancetta laborum
					nisi id. Meatloaf boudin porchetta sirloin brisket biltong aute adipisicing non turkey chuck tempor.
					Tongue dolore ipsum occaecat jerky fatback salami shankle. Et in meatloaf excepteur tempor pork.
					In jowl landjaeger short loin sed ut tail eiusmod elit strip steak ut ribeye duis ut. Short ribs
					beef ribs laboris sed tongue ea pig prosciutto lorem salami bacon est. Ut elit ipsum, dolor ea
					t-bone voluptate jowl chicken commodo jerky ex picanha esse. Dolore bacon andouille non irure
					tail pariatur.`
	sampletest3 := `Video provides a powerful way to help you prove your point. When you click Online Video, you can
				  paste in the embed code for the video you want to add. You can also type a keyword to search online
				  for the video that best fits your document. To make your document look professionally produced,
				  Word provides header, footer, cover page, and text box designs that complement each other. For 
				  example, you can add a matching cover page, header, and sidebar. Click Insert and then choose the
				  elements you want from the different galleries. Themes and styles also help keep your document
				  coordinated. When you click Design and choose a new Theme, the pictures, charts, and SmartArt
				  graphics change to match your new theme. When you apply styles, your headings change to match the
				  new theme. Save time in Word with new buttons that show up where you need them.`
	// Regexp to extract words from the text, all punctuations and spaces will be removed.
	//re := regexp.MustCompile("\\w+")
	//text := re.FindAllString(sampletest2, -1)

	fmt.Println(TextFrq(sampletest2))
	fmt.Println(TextFrq(testText))
	fmt.Println(TextFrq(sampletest3))
	//fmt.Println(textfrq(sampletest2))
}
