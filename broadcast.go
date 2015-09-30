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
	Title           string
	Description     string
	LongDescription string
	URL             string
	Date            time.Time
	Length          time.Duration
	Type            BroadcastType
	ImageURLs       []string
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
	pTitle  = xmlpath.MustCompile("/html/head/meta[@name='og:title']/@content")
	pDesc   = xmlpath.MustCompile("/html/head/meta[@name='og:description']/@content")
	pURL    = xmlpath.MustCompile("/html/head/meta[@name='og:url']/@content")
	pType   = xmlpath.MustCompile("/html/head/meta[@name='og:type']/@content")
	pImages = xmlpath.MustCompile("/html/head/meta[@name='og:image']/@content")
)

// ParseBroadcast parses content of a reader into a Broadcast.
func ParseBroadcast(r io.Reader) (*Broadcast, error) {
	n, err := xmlpath.ParseHTML(r)
	if err != nil {
		return nil, err
	}

	title, ok := pTitle.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast title")
	}

	desc, ok := pDesc.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast description")
	}

	longDesc, ok := parseLongDesc(n) // handle NPO3 video site correctly
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast long description")
	}

	url, ok := pURL.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast page URL")
	}

	typstr, ok := pType.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast type")
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

	var typ BroadcastType
	switch {
	case strings.HasPrefix(typstr, "music"):
		typ = Audio
	case strings.HasPrefix(typstr, "video"):
		typ = Video
	default:
		return nil, fmt.Errorf("gemist: unknown broadcast type %s", typstr)
	}

	var (
		dateFn  func(*xmlpath.Node) (time.Time, error)
		lenFn   func(*xmlpath.Node) (time.Duration, error)
		mediaFn func(*xmlpath.Node) (string, bool)
	)

	switch typ {
	case Video:
		dateFn = parseDateV
		lenFn = parseLengthV
		mediaFn = parseMediaV
	case Audio:
		dateFn = parseDateA
		lenFn = parseLengthA
		mediaFn = parseMediaA
	default:
		return nil, errors.New("gemist: unknown broadcast type")
	}

	date, err := dateFn(n)
	if err != nil {
		return nil, err
	}

	len, err := lenFn(n)
	if err != nil {
		return nil, err
	}

	media, ok := mediaFn(n)
	if !ok {
		return nil, errors.New("gemist: error parsing broadcast media URL")
	}

	b := Broadcast{
		Title:           title,
		Description:     desc,
		LongDescription: longDesc,
		URL:             url,
		Date:            date,
		Length:          len,
		Type:            typ,
		ImageURLs:       images,
		MediaURL:        media,
	}

	return &b, nil
}

var (
	pLDesc     = xmlpath.MustCompile("//*/div[@class='content']/p/span[3]/text()")
	pLDescNPO3 = xmlpath.MustCompile("//div[contains(@class,'meta-content')]/div[1]/p[1]/span[3]/text()")
)

func parseLongDesc(n *xmlpath.Node) (desc string, ok bool) {
	desc, ok = pLDesc.String(n)
	if !ok {
		// check for NPO3 video page contents
		desc, ok = pLDescNPO3.String(n)
	}

	return
}

var pDate = xmlpath.MustCompile("//span[@itemprop='startDate']/text()")

const dateLayoutV = "2006-01-02 15:04:05 -0700"

func parseDateV(n *xmlpath.Node) (date time.Time, err error) {
	datestr, ok := pDate.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast date (video)")
		return
	}

	date, err = time.Parse(dateLayoutV, datestr)
	return
}

var (
	pDateA     = xmlpath.MustCompile("//span[@itemprop='startDate']/text()")
	dateLoc, _ = time.LoadLocation("Europe/Amsterdam")
	dateRep    = strings.NewReplacer(
		"Ma", "Mon",
		"Di", "Tue",
		"Wo", "Wed",
		"Do", "Thu",
		"Vr", "Fri",
		"Za", "Sat",
		"Zo", "Sun",
		"jan", "Jan",
		"feb", "Feb",
		"mrt", "Mar",
		"apr", "Apr",
		"mei", "May",
		"jun", "Jun",
		"jul", "Jul",
		"aug", "Aug",
		"sep", "Sep",
		"okt", "Oct",
		"nov", "Nov",
		"dec", "Dec",
	)
)

const dateLayoutA = "2006-01-02 15:04:05 -0700"

func parseDateA(n *xmlpath.Node) (date time.Time, err error) {
	str, ok := pDateA.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast date (audio)")
		return
	}

	str = strings.TrimSpace(str)
	str = dateRep.Replace(str)
	date, err = time.ParseInLocation(dateLayoutA, str, dateLoc)
	return
}

var pLengthV = xmlpath.MustCompile("/html/head/meta[@name='og:video:duration']/@content")

func parseLengthV(n *xmlpath.Node) (len time.Duration, err error) {
	lenstr, ok := pLengthV.String(n)
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

var (
	pLengthA  = xmlpath.MustCompile("//span[@class='duration']/text()")
	reLengthA = regexp.MustCompile(`(\d:)?(\d+):(\d+)`)
)

func parseLengthA(n *xmlpath.Node) (len time.Duration, err error) {
	lenstr, ok := pLengthA.String(n)
	if !ok {
		err = errors.New("gemist: error parsing broadcast length (audio)")
		return
	}

	parts := reLengthA.FindStringSubmatch(lenstr)
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

var pMediaV = xmlpath.MustCompile("/html/head/meta[@name='og:video']/@content")

func parseMediaV(n *xmlpath.Node) (media string, ok bool) {
	return pMediaV.String(n)
}

var pMediaA = xmlpath.MustCompile("/html/head/meta[@name='og:audio']/@content")

func parseMediaA(n *xmlpath.Node) (media string, ok bool) {
	return pMediaA.String(n)
}
