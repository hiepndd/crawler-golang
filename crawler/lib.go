package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"time"
)

type Selector struct {
	Titile        string
	PublishedDate string
	Author        string
	Content       string
}

type Data struct {
	Title         string
	PublishedDate time.Time
	Author        string
	Content       string
}

type ICrawler interface {
	Parse(res *http.Response) Data
}

type Crawler struct {
	selector Selector
	parser   Parser
}

func (self *Crawler) Parse(res *http.Response) Data {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}
	data := Data{}
	data.Title = self.parse.extractTitle(self.selector.Title, doc)
	data.Author = self.parse.extractAuthor(self.selector.Author, doc)
	data.Content = self.parse.extractContent(self.selector.Content, doc)
	data.PublishedDate = self.parse.extractPublishedDate(self.selector.PublishedDate, doc)
	return data

}
