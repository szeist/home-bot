package facebook

import (
	"errors"

	"github.com/antchfx/htmlquery"
)

func GetVideoURL(pageUrl string) (string, error) {
	doc, err := htmlquery.LoadURL(pageUrl)
	if err != nil {
		return "", err
	}

	node := htmlquery.FindOne(doc, "//meta[@property='og:video']/@content")
	if node == nil {
		return "", errors.New("Video url not found")
	}

	return htmlquery.InnerText(node), nil
}
