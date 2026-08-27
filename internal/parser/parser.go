package parser

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(htmlBody []byte, baseURL string) ([]string, error) {

	doc, err := html.Parse(
		strings.NewReader(string(htmlBody)),
	)

	if err != nil {
		return nil, err
	}

	var links []string

	var traverse func(*html.Node)

	traverse = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {

			for _, attr := range n.Attr {

				if attr.Key == "href" {

					link := resolveURL(
						baseURL,
						attr.Val,
					)

					if link != "" {
						links = append(
							links,
							link,
						)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	return links, nil
}

func resolveURL(
	base string,
	href string,
) string {

	baseURL, err := url.Parse(base)

	if err != nil {
		return ""
	}

	refURL, err := url.Parse(href)

	if err != nil {
		return ""
	}

	return baseURL.ResolveReference(
		refURL,
	).String()
}
