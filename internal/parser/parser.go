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

	linkSet := make(map[string]struct{})

	var traverse func(*html.Node)

	traverse = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {

			for _, attr := range n.Attr {

				if attr.Key == "href" {

					href := strings.TrimSpace(attr.Val)

					if href == "" {
						continue
					}

					if strings.HasPrefix(href, "#") {
						continue
					}

					if strings.HasPrefix(href, "javascript:") {
						continue
					}

					if strings.HasPrefix(href, "mailto:") {
						continue
					}

					if strings.HasPrefix(href, "tel:") {
						continue
					}

					link := resolveURL(
						baseURL,
						href,
					)

					if link != "" {
						linkSet[link] = struct{}{}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	links := make([]string, 0, len(linkSet))

	for link := range linkSet {
		links = append(links, link)
	}

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

	resolved := baseURL.ResolveReference(refURL)

	resolved.Fragment = ""

	return resolved.String()
}
