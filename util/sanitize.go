package util

import "github.com/microcosm-cc/bluemonday"

// Sanitize html
func Sanitize(html string) string {
	p := bluemonday.UGCPolicy()

	return p.Sanitize(html)
}
