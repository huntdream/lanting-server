package util

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// ExtractText extract plain text from html markup
func ExtractText(raw string) string {
	var text string

	tokenizer := html.NewTokenizer(strings.NewReader(raw))

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()

			if err == io.EOF {

				//end of the file, break out of the loop
				break
			}
		}

		if tokenType == html.TextToken {
			node := tokenizer.Token().Data

			text += node
		}

	}

	return strings.ReplaceAll(text, "\n", " ")
}
