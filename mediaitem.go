package gemist

import (
	"errors"

	"gopkg.in/xmlpath.v2"
)

// MediaItem represents a generic Uitzending Gemist media item.
type MediaItem struct {
	Title       string
	Description string
	ImageURLs   []string
	URL         string
}

var (
	pMITitle  = xmlpath.MustCompile("/html/head/meta[@name='og:title']/@content")
	pMIDesc   = xmlpath.MustCompile("/html/head/meta[@name='og:description']/@content")
	pMIURL    = xmlpath.MustCompile("/html/head/meta[@name='og:url']/@content")
	pMIImages = xmlpath.MustCompile("/html/head/meta[@name='og:image']/@content")
)

func parseMediaItem(n *xmlpath.Node) (mi MediaItem, err error) {
	title, ok := pMITitle.String(n)
	if !ok {
		err = errors.New("gemist: error parsing title of media item")
		return
	}

	desc, ok := pMIDesc.String(n)
	if !ok {
		err = errors.New("gemist: error parsing description of media item")
		return
	}

	url, ok := pMIURL.String(n)
	if !ok {
		err = errors.New("gemist: error parsing URL of media item")
		return
	}

	images := []string{} // we want to allocate the slice here!
	iter := pMIImages.Iter(n)
	for iter.Next() {
		img := iter.Node()
		url := img.String()
		if url != "" {
			images = append(images, url)
		}
	}

	mi.Title = title
	mi.Description = desc
	mi.ImageURLs = images
	mi.URL = url
	return
}
