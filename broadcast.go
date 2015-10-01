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
	pBType      = xmlpath.MustCompile("/html/head/meta[@name='og:type']/@content")
	pBLDesc     = xmlpath.MustCompile("//*/div[@class='content']/p/span[3]/text()")
	pBLDescNPO3 = xmlpath.MustCompile("//div[contains(@class,'meta-content')]/div[1]/p[1]/span[3]/text()")
	pBDate      = xmlpath.MustCompile("//span[@itemprop='startDate']/text()")
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
	longDesc, ok := pBLDesc.String(n)
	if !ok {
		// check for NPO3 video page contents
		longDesc, ok = pBLDescNPO3.String(n)
	}
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast long description")
	}

	// -- Date --
	datestr, ok := pBDate.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast date")
		return nil, err
	}

	date, err := time.Parse(broadcastDateLayout, datestr)
	if err != nil {
		return nil, err
	}

	// -- Type --
	typstr, ok := pBType.String(n)
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
	media, err := p.MediaURL(n)
	if err != nil {
		return nil, err
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
	MediaURL(*xmlpath.Node) (string, error)
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

func (p videoBroadcastParser) MediaURL(n *xmlpath.Node) (url string, err error) {
	url, ok := p.m.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast media URL (video)")
	}

	return
}

var broadcastParserV = videoBroadcastParser{
	l: xmlpath.MustCompile("/html/head/meta[@name='og:video:duration']/@content"),
	m: xmlpath.MustCompile("/html/head/meta[@name='og:video']/@content"),
}

type audioBroadcastParser struct {
	l *xmlpath.Path
	m *xmlpath.Path
}

func (p audioBroadcastParser) Length(n *xmlpath.Node) (len time.Duration, err error) {
	lenstr, ok := p.l.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast length (audio)")
		return
	}

	len, err = parseBroadcastLength(lenstr)
	return
}

func (p audioBroadcastParser) MediaURL(n *xmlpath.Node) (url string, err error) {
	url, ok := p.m.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast media URL (audio)")
	}

	return
}

var broadcastParserA = audioBroadcastParser{
	l: xmlpath.MustCompile("//span[@class='duration']/text()"),
	m: xmlpath.MustCompile("/html/head/meta[@name='og:audio']/@content"),
}

var broadcastLengthRegexp = regexp.MustCompile(`(\d:)?(\d+):(\d+)`)

func parseBroadcastLength(str string) (len time.Duration, err error) {
	parts := broadcastLengthRegexp.FindStringSubmatch(str)
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
