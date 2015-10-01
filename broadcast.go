// Package gemist provides functionality for working with Uitzending Gemist content.
package gemist

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/xmlpath.v2"
)

// Broadcast represents a single broadcast of a program on Uitzending Gemist.
type Broadcast struct {
	MediaItem
	LongDescription string
	Date            time.Time
	Length          time.Duration
	Type            BroadcastType
	MediaURL        string
}

// BroadcastType indicates the type of media (audio or video).
type BroadcastType int

// Broadcast media types.
const (
	Video BroadcastType = iota
	Audio
)

// GetBroadcast gets the page content from url, parses it and returns a Broadcast.
func GetBroadcast(url string) (*Broadcast, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return ParseBroadcast(r.Body)
}

var (
	pType      = xmlpath.MustCompile("/html/head/meta[@name='og:type']/@content")
	pLDesc     = xmlpath.MustCompile("//*/div[@class='content']/p/span[3]/text()")
	pLDescNPO3 = xmlpath.MustCompile("//div[contains(@class,'meta-content')]/div[1]/p[1]/span[3]/text()")
	pDate      = xmlpath.MustCompile("//span[@itemprop='startDate']/text()")
)

const broadcastDateLayout = "2006-01-02 15:04:05 -0700"

// ParseBroadcast parses content of a reader into a Broadcast.
func ParseBroadcast(r io.Reader) (*Broadcast, error) {
	n, err := xmlpath.ParseHTML(r)
	if err != nil {
		return nil, err
	}

	// -- MediaItem --
	mi, err := parseMediaItem(n)
	if err != nil {
		return nil, err
	}

	// -- LongDescription --
	longDesc, ok := pLDesc.String(n)
	if !ok {
		// check for NPO3 video page contents
		longDesc, ok = pLDescNPO3.String(n)
	}
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast long description")
	}

	// -- Date --
	datestr, ok := pDate.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast date")
		return nil, err
	}

	date, err := time.Parse(broadcastDateLayout, datestr)
	if err != nil {
		return nil, err
	}

	// -- Type --
	typstr, ok := pType.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast type")
	}

	var (
		typ BroadcastType
		p   broadcastParser
	)

	switch {
	case strings.HasPrefix(typstr, "music"):
		typ = Audio
		p = broadcastParserA
	case strings.HasPrefix(typstr, "video"):
		typ = Video
		p = broadcastParserV
	default:
		return nil, fmt.Errorf("gemist: unknown broadcast type %s", typstr)
	}

	// -- Length --
	len, err := p.Length(n)
	if err != nil {
		return nil, err
	}

	// -- Media --
	media, ok := p.MediaURL(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast media URL")
	}

	b := Broadcast{
		MediaItem:       mi,
		LongDescription: longDesc,
		Date:            date,
		Length:          len,
		Type:            typ,
		MediaURL:        media,
	}

	return &b, nil
}

type broadcastParser interface {
	Length(*xmlpath.Node) (time.Duration, error)
	MediaURL(*xmlpath.Node) (string, bool)
}

type videoBroadcastParser struct {
	l *xmlpath.Path
	m *xmlpath.Path
}

func (p videoBroadcastParser) Length(n *xmlpath.Node) (len time.Duration, err error) {
	lenstr, ok := p.l.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast length (video)")
		return
	}

	lenint, err := strconv.Atoi(lenstr)
	if err != nil {
		return
	}

	len = time.Duration(lenint) * time.Second // length in source is in seconds.
	return
}

func (p videoBroadcastParser) MediaURL(n *xmlpath.Node) (string, bool) {
	return p.m.String(n)
}

var broadcastParserV = videoBroadcastParser{
	l: xmlpath.MustCompile("/html/head/meta[@name='og:video:duration']/@content"),
	m: xmlpath.MustCompile("/html/head/meta[@name='og:video']/@content"),
}

type audioBroadcastParser struct {
	lp  *xmlpath.Path
	lre *regexp.Regexp
	m   *xmlpath.Path
}

func (p audioBroadcastParser) Length(n *xmlpath.Node) (len time.Duration, err error) {
	lenstr, ok := p.lp.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast length (audio)")
		return
	}

	parts := p.lre.FindStringSubmatch(lenstr)
	// parts[0] -> matched string
	// parts[1] -> hours
	// parts[2] -> minutes
	// parts[3] -> seconds

	if parts[1] != "" {
		if hour, err := strconv.Atoi(parts[1]); err != nil {
			len += time.Duration(hour) * 60 * 60 * time.Second
		}
	}

	min, err := strconv.Atoi(parts[2])
	if err != nil {
		return
	}

	len += time.Duration(min) * 60 * time.Second

	sec, err := strconv.Atoi(parts[3])
	if err != nil {
		return
	}

	len += time.Duration(sec) * time.Second
	return
}

func (p audioBroadcastParser) MediaURL(n *xmlpath.Node) (string, bool) {
	return p.m.String(n)
}

var broadcastParserA = audioBroadcastParser{
	lp:  xmlpath.MustCompile("//span[@class='duration']/text()"),
	lre: regexp.MustCompile(`(\d:)?(\d+):(\d+)`),
	m:   xmlpath.MustCompile("/html/head/meta[@name='og:audio']/@content"),
}
