package html2steam

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func Replace(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}
	html := doc.Selection
	body := doc.Filter("body")
	if body.Length() > 0 {
		html = body.First().Children()
	}
	html.Find("pre").Each(func(_ int, s *goquery.Selection) {
		s.ReplaceWithSelection(s.Children())
	})
	html.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		text := s.Text()
		tag := fmt.Sprintf("[url=%s]%s[/url]", url, text)
		s.ReplaceWithHtml(tag)
	})
	html.Find("blockquote, q").Each(func(_ int, s *goquery.Selection) {
		city, _ := s.Attr("city")
		text := s.Text()
		tag := fmt.Sprintf("[quote=%s]%s[/quote]", city, text)
		s.ReplaceWithHtml(tag)
	})
	html.Find("br").Each(func(_ int, s *goquery.Selection) {
		s.ReplaceWithHtml("\n")
	})
	normal := map[string]string{
		"h1":         "h1",
		"b":          "b",
		"u":          "u",
		"i":          "i",
		"s, spoiler": "spoiler",
		"ul":         "list",
		"ol":         "olist",
		"li":         "*",
		"code":       "code",
	}
	for k, v := range normal {
		html.Find(k).Each(func(_ int, s *goquery.Selection) {
			if v != "*" {
				s.ReplaceWithHtml(fmt.Sprintf("[%s]%s[/%s]", v, s.Text(), v))
			} else {
				s.ReplaceWithHtml(fmt.Sprintf("[*]%s", s.Text()))
			}
		})
	}
	return html.Text(), nil
}
