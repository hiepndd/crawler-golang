package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

type ExtractString = func(selector string, doc *goquery.Document) string

type ExtractTime = func(selector string, doc *goquery.Document) time.Time

type Parser struct {
	extractPublishDate ExtractTime
	extractContent     ExtractString
	extractAuthor      ExtractString
	extractTitle       ExtractString
}

func extractNotYetImplement(selector string, doc *goquery.Document) time.Time {
	panic("Should implement at yours")
}

func extract(selector string, doc *goquery.Document) string {
	if selector == "" {
		return ""
	}
	return doc.Find(selector).Text()
}

func extractSimple(selector string, doc *goquery.Document) string {
	value := extract(selector, doc)
	return strings.TrimSpace(value)
}

func createDefaultParser() Parser {
	parser := Parser{
		extractPublishDate: extractNotYetImplement,
		extractContent:     extractSimple,
		extractAuthor:      extractSimple,
		extractTitle:       extractSimple,
	}
	return parser
}
