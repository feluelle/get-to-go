package main

import (
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// EmojiTranslator translates words into emojis
// See: https://github.com/florinpop17/app-ideas/blob/master/Projects/2-Intermediate/Emoji-Translator-App.md
// References: https://unicode.org/emoji/charts/full-emoji-list.html
func EmojiTranslator(text string) (string, []string) {
	words := strings.Fields(text)
	emojisFile := DownloadFile("https://unicode.org/Public/emoji/13.0/emoji-test.txt")
	df := dataframe.ReadCSV(
		strings.NewReader(emojisFile),
		dataframe.WithDelimiter(';'),
		dataframe.Names("Code", "Name"),
		dataframe.HasHeader(false),
		dataframe.WithComments('#'),
	)
	// TODO: Cleanse dataframe to only contain relevant information
	// - Remove white spaces from "Code"
	// - Remove irrelevant information from "Name"
	var codes []string
	for _, word := range words {
		codes = append(codes, findCode(df, word))
	}
	var emojis []string
	for _, code := range codes {
		emojis = append(emojis, getEmojiFromCode(code))
	}
	emojisString := strings.Join(emojis, " ")
	return emojisString, codes
}

func findCode(df dataframe.DataFrame, word string) string {
	// TODO: Remove like func after it has landed in master branch / is released: https://github.com/go-gota/gota/issues/115
	like := func(str string) func(el series.Element) bool {
		return func(el series.Element) bool {
			if el.Type() == series.String {
				if val, ok := el.Val().(string); ok {
					return strings.Contains(val, str)
				}
			}
			return false
		}
	}
	dfFiltered := df.Filter(dataframe.F{
		Colname:    "Name",
		Comparator: series.CompFunc,
		Comparando: like(word),
	})
	numberOfResults := dfFiltered.Nrow()
	if numberOfResults > 0 {
		// Use one of the codes. There might be cases where many emojis work.
		return dfFiltered.Col("Code").Elem(rand.Intn(numberOfResults)).String()
	}
	// No emoji found for this word.
	return ""
}

func getEmojiFromCode(code string) string {
	codeParts := strings.Fields(code)
	var emoji []string
	for _, codePart := range codeParts {
		codePartInt, err := strconv.ParseInt(codePart, 16, 32)
		if err != nil {
			log.Fatal(err)
		} else {
			emoji = append(emoji, string(codePartInt))
		}
	}
	return strings.Join(emoji, " ")
}
