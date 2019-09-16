package main

import (
	"fmt"
	"testing"
)

func TestBasicTest(t *testing.T) {
	testText := "aaaa bbb aaaa bbb cccc ddddd aaaa xxxx zzzzz fffff rrrrr eeee aaaa bbb rrrrr"
	expectedResult := map[string]int{
		"aaaa":  4,
		"bbb":   3,
		"rrrrr": 2,
		"ddddd": 1,
		"zzzzz": 1,
		"fffff": 1,
		"xxxx":  1,
		"eeee":  1,
		"cccc":  1,
	}
	result := TextFrq(testText)
	fmt.Println(result)
	fmt.Println(expectedResult)
	if len(result) != len(expectedResult) {
		t.Errorf("Mismatch len of result %d end expected result %d", len(result), len(expectedResult))
	}
	for k := range result {
		if result[k].frequency != expectedResult[result[k].word] {
			t.Error("Elements not matched...", result[k], expectedResult[result[k].word])
		}
	}
}
func TestAdvancedTest(t *testing.T) {

	testText := `Bacon ipsum dolor amet pork chop sunt swine, t-bone velit pastrami frankfurter adipisicing.
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
	expectedResult := map[string]int{
		"ut":     5,
		"beef":   5,
		"dolore": 5,
		"sed":    4,
		"eu":     4,
		"ribs":   4,
		"bone":   4,
		"steak":  4,
		"pork":   4,
		"strip":  4,
	}
	result := TextFrq(testText)
	fmt.Println(result)
	fmt.Println(expectedResult)
	if len(result) != len(expectedResult) {
		t.Errorf("Mismatch len of result %d end expected result %d", len(result), len(expectedResult))
	}
	for k := range result {
		if result[k].frequency != expectedResult[result[k].word] {
			t.Error("Elements not matched...", result[k], expectedResult[result[k].word])
		}
	}
}
func TestRegularRandomText(t *testing.T) {
	testText := `Video provides a powerful way to help you prove your point. When you click Online Video, you can
	paste in the embed code for the video you want to add. You can also type a keyword to search online
	for the video that best fits your document. To make your document look professionally produced,
	Word provides header, footer, cover page, and text box designs that complement each other. For 
	example, you can add a matching cover page, header, and sidebar. Click Insert and then choose the
	elements you want from the different galleries. Themes and styles also help keep your document
	coordinated. When you click Design and choose a new Theme, the pictures, charts, and SmartArt
	graphics change to match your new theme. When you apply styles, your headings change to match the
	new theme. Save time in Word with new buttons that show up where you need them.`
	expectedResult := map[string]int{
		"you":      10,
		"the":      7,
		"to":       6,
		"your":     6,
		"and":      6,
		"video":    4,
		"new":      4,
		"when":     3,
		"document": 3,
		"theme":    3,
	}
	result := TextFrq(testText)
	fmt.Println(result)
	fmt.Println(expectedResult)
	if len(result) != len(expectedResult) {
		t.Errorf("Mismatch len of result %d end expected result %d", len(result), len(expectedResult))
	}
	for k := range result {
		if result[k].frequency != expectedResult[result[k].word] {
			t.Error("Elements not matched...", result[k], expectedResult[result[k].word])
		}
	}
}
