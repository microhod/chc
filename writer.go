package main

import (
	"fmt"
	"io"
	"text/tabwriter"
)

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(counts []CharCount) error {
	writer := tabwriter.NewWriter(w.w, 1, 4, 1, ' ', 0)

	for _, count := range counts {
		char := w.escape(count.Char)
		_, err := writer.Write([]byte(fmt.Sprintf("%s\t%d\n", char, count.Count)))
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func (w *Writer) escape(ch rune) string {
	char := string(ch)
	escape := map[string]string{
		"\t": "\\t",
		"\v": "\\v",
		"\n": "\\n",
		"\f": "\\f",
	}
	if escaped, required := escape[char]; required {
		char = escaped
	}
	return char
}
