package main

import (
	"fmt"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// EmojiTranslator ...
// References:
// https://github.com/florinpop17/app-ideas/blob/master/Projects/2-Intermediate/Emoji-Translator-App.md
// https://unicode.org/emoji/charts/full-emoji-list.html
// https://unicode.org/Public/emoji/13.0/emoji-test.txt
func EmojiTranslator(text string) string {
	emojis := DownloadFile("https://unicode.org/Public/emoji/13.0/emoji-test.txt")
	df := dataframe.ReadCSV(
		strings.NewReader(emojis),
		dataframe.WithDelimiter(';'),
		dataframe.Names("Code", "Name"),
		dataframe.HasHeader(false),
		dataframe.WithComments('#'),
	)
	// TODO: Cleanse dataframe to only contain relevant information
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
	df = df.Filter(dataframe.F{
		Colname:    "Name",
		Comparator: series.CompFunc,
		Comparando: like(text),
	})
	// TODO: Return only string of matched codes separated by spaces
	fmt.Println(df)
	return ""
}
