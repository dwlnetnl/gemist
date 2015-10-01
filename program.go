package gemist

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gopkg.in/xmlpath.v2"
)

// Program represents a program on Uitzending Gemist.
type Program struct {
	MediaItem

	bs []*BroadcastProxy
}

// GetProgram gets the page content from url, parses it and returns a Program.
func GetProgram(url string) (*Program, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return ParseProgram(r.Body)
}

// ParseProgram parses content of a reader into a Program.
func ParseProgram(r io.Reader) (*Program, error) {
	n, err := xmlpath.ParseHTML(r)
	if err != nil {
		return nil, err
	}

	mi, err := parseMediaItem(n)
	if err != nil {
		return nil, err
	}

	// Get broadcast list node.
	bs, err := parseProgramBroadcasts(n)
	if err != nil {
		return nil, err
	}

	p := Program{
		MediaItem: mi,
		bs:        bs,
	}

	return &p, nil
}

var pPBData = xmlpath.MustCompile("//div[@id='broadcasts-block']/div[1]/div[1]")

func parseProgramBroadcasts(n *xmlpath.Node) ([]*BroadcastProxy, error) {
	iter := pPBData.Iter(n)
	iter.Next()
	l := iter.Node()

	bs, err := newBroadcastProxySlice(l)
	if err != nil {
		return nil, err
	}

	err = parseProgramList(l, &bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

type BroadcastProxy struct {
	MediaItem
	SubTitle string
	Date     time.Time
	Length   time.Duration
}

var pPListNum = xmlpath.MustCompile("@data-num-found")

func newBroadcastProxySlice(n *xmlpath.Node) ([]*BroadcastProxy, error) {
	snum, ok := pPListNum.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing program list length")
	}

	num, err := strconv.Atoi(snum)
	if err != nil {
		return nil, err
	}

	bps := make([]*BroadcastProxy, 0, num)
	return bps, nil
}

var (
	pPListItems = xmlpath.MustCompile("div")
)

func parseProgramList(n *xmlpath.Node, s *[]*BroadcastProxy) error {
	iter := pPListItems.Iter(n)

	var (
		bps []*BroadcastProxy
		err error
	)

	for iter.Next() {
		bp, lerr := parseProgramListItem(iter.Node())
		if lerr != nil {
			break
		}

		bps = append(bps, bp)
	}

	if err != nil {
		return err
	}

	for _, bp := range bps {
		*s = append(*s, bp)
	}

	return nil
}

var (
	pPListItemTitle     = xmlpath.MustCompile("div[2]/a/h4/text()")
	pPListItemDesc      = xmlpath.MustCompile("div[2]/a/p/text()")
	pPListItemImage     = xmlpath.MustCompile("div[1]/div/a/img/@src")
	pPListItemURL       = xmlpath.MustCompile("div[1]/div/a/@href")
	pPListItemInfo      = xmlpath.MustCompile("div[2]/a/h5/text()")
	pPListItemLen       = xmlpath.MustCompile("div[1]/div/a/div/text()")
	pListItemDateLoc, _ = time.LoadLocation("Europe/Amsterdam")
	pListItemDateRep    = strings.NewReplacer(
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

const urlBase = "http://www.npo.nl"
const pListItemDateLayout = "Mon _2 Jan 2006 15:04"

func parseProgramListItem(n *xmlpath.Node) (*BroadcastProxy, error) {
	title, ok := pPListItemTitle.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing program list title")
	}

	// No need to exist, no description is valid.
	desc, _ := pPListItemDesc.String(n)

	img, ok := pPListItemImage.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing program list image")
	}

	path, ok := pPListItemURL.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing program list url")
	}

	infostr, ok := pPListItemInfo.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing program list info")
	}

	info := strings.Split(infostr, " · ")

	datestr := pListItemDateRep.Replace(info[1])
	date, err := time.ParseInLocation(pListItemDateLayout, datestr, pListItemDateLoc)
	if err != nil {
		return nil, err
	}

	lenstr, ok := pPListItemLen.String(n)
	if !ok {
		return nil, errors.New("gemist: error parsing program list length")
	}

	len, err := parseBroadcastLength(lenstr)
	if err != nil {
		return nil, err
	}

	bp := BroadcastProxy{
		MediaItem: MediaItem{
			Title:       strings.TrimSpace(title),
			Description: desc,
			ImageURLs:   []string{img},
			URL:         urlBase + path,
		},
		SubTitle: info[0],
		Date:     date,
		Length:   len,
	}

	return &bp, nil
}
