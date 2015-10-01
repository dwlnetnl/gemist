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
	pTitle  = xmlpath.MustCompile("/html/head/meta[@name='og:title']/@content")
	pDesc   = xmlpath.MustCompile("/html/head/meta[@name='og:description']/@content")
	pURL    = xmlpath.MustCompile("/html/head/meta[@name='og:url']/@content")
	pImages = xmlpath.MustCompile("/html/head/meta[@name='og:image']/@content")
)

func parseMediaItem(n *xmlpath.Node) (mi MediaItem, err error) {
	title, ok := pTitle.String(n)
	if !ok {
		err = errors.New("gemist: error parsing title of media item")
		return
	}

	desc, ok := pDesc.String(n)
	if !ok {
		err = errors.New("gemist: error parsing description of media item")
		return
	}

	url, ok := pURL.String(n)
	if !ok {
		err = errors.New("gemist: error parsing URL of media item")
		return
	}

	images := []string{} // we want to allocate the slice here!
	iter := pImages.Iter(n)
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
