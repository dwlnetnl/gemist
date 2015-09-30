package gemist

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBroadcast_audio(t *testing.T) {
	r, err := os.Open("testdata/broadcast-audio.html.gz")
	require.NoError(t, err, "error opening broadcast page")

	cr, err := gzip.NewReader(r)
	require.NoError(t, err, "error opening broadcast page (gzip)")

	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	_ep := Broadcast{
		Title:           "Radio Bergeijk",
		Description:     "Met in de eerste nationale aflevering voor de regio Bergeijk ooit:\nDe rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk mag de muziek niet ontbreken.",
		LongDescription: " NieuwsWeerVerkeer: o.a. tarieven varkensbelasting\nMestberichten: De Giertijden\nStudiogast: Arno Vlemmings (Opvoedinstructies)\nmuziek: Onze vader zei (trad.) - Moek (van de CD: 'Brabant-bont-ste Vol. 2 Clipsound CCD 994)\nSportnieuws: Afgelastingen",
		URL:             "http://www.npo.nl/radio-bergeijk/03-04-2001/POMS_VPRO_396139",
		Date:            time.Date(2001, time.April, 3, 0, 44, 0, 0, l),
		Length:          885000000000,
		Type:            Audio,
		ImageURLs:       []string{"http://images.poms.omroep.nl/image/215303.png"},
		MediaURL:        "http://download.omroep.nl/vpro/29/08/57/39/POMS_VPRO_396139.mp3",
	}

	ep, err := ParseBroadcast(cr)
	assertParsedBroadcast(t, &_ep, ep, err)
}

func TestParseBroadcast_video(t *testing.T) {
	r, err := os.Open("testdata/broadcast-video.html.gz")
	require.NoError(t, err, "error opening broadcast page")

	cr, err := gzip.NewReader(r)
	require.NoError(t, err, "error opening broadcast page (gzip)")

	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	_ep := Broadcast{
		Description:     "Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.",
		LongDescription: " Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.",
		URL:             "http://www.npo.nl/zembla/18-03-2007/VARA_101141965",
		Date:            time.Date(2007, time.March, 18, 20, 35, 0, 0, l),
		Length:          3000000000000,
		Type:            Video,
		ImageURLs:       []string{"http://images.poms.omroep.nl/image/6848.png"},
		MediaURL:        "http://www.npo.nl/zembla/18-03-2007/VARA_101141965",
	}

	ep, err := ParseBroadcast(cr)
	assertParsedBroadcast(t, &_ep, ep, err)
}

func TestParseBroadcast_video_NPO3(t *testing.T) {
	r, err := os.Open("testdata/broadcast-video-npo3.html.gz")
	require.NoError(t, err, "error opening broadcast page")

	cr, err := gzip.NewReader(r)
	require.NoError(t, err, "error opening broadcast page (gzip)")

	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	_ep := Broadcast{
		Description:     "Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met technicus Ted van Lieshout.",
		LongDescription: " Serie verslagen van hun levenswerk: Radio Bergeijk. Een filmploeg over de radiovloer! Ze moesten even aan de gedachte wennen, de presentatoren van Radio Bergeijk. En het heeft even geduurd voordat Toon Spoorenberg en Peer van Eersel konden instemmen met die druktemakerij in hun zo vertrouwde en overzichtelijke radiobestaan. Maar het is er dan toch van gekomen. Pieter Verhoeff heeft een serie gemaakt van Radio Bergeijk, het ook landelijk bekende radiostation voor Bergeijk, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Ted van Lieshout. De pro Deo werkende radiomannen ontvangen in hun studio een keur aan plaatselijke gasten of doen rechtstreeks verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. De televisieserie wordt niet alleen een fel realistisch portret van de makers van het vermaarde radiostation, maar via hen ook van het dorp Bergeijk. Radio Bergeijk is en probeert een satire te zijn op een lokaal radiostation. In de verbeelding van de makers is het fictieve dorp Bergeijk een huiveringwekkend oord waar hedendaagse xenofobie, orthodox atheïsme, vrouwonvriendelijkheid en homo-afkeer om de voorrang strijden. Het radiostation probeert van deze benepen dorpse identiteit pseudo-professioneel verslag te doen. En hanteert als stijlmiddel de grove overdrijving. Teksten en spel: Pieter Bouwman en George van Houts. Gastacteurs o.a.: Alex Klaasen, Jeroen van Merwijk, Marjan Luif, Annet Malherbe, Anniek Pheifer, Margôt Ros en Lies Visschedijk.",
		URL:             "http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739",
		Date:            time.Date(2007, time.June, 25, 23, 10, 0, 0, l),
		Length:          2100000000000,
		Type:            Video,
		ImageURLs:       []string{},
		MediaURL:        "http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739",
	}

	ep, err := ParseBroadcast(cr)
	assertParsedBroadcast(t, &_ep, ep, err)
}

func assertParsedBroadcast(t *testing.T, _ep, ep *Broadcast, err error) {
	assert := assert.New(t)
	if assert.NoError(err) {
		assert.Equal(_ep.Description, ep.Description, "description not equal")
		assert.Equal(_ep.LongDescription, ep.LongDescription, "long description not equal")
		assert.Equal(_ep.URL, ep.URL, "broadcast URL not equal")
		assert.True(_ep.Date.Equal(ep.Date), "date not equal (%v != %v)", _ep.Date, ep.Date)
		assert.Equal(_ep.Length, ep.Length, "length not equal")
		assert.Equal(_ep.Type, ep.Type, "type not equal")
		assert.Equal(_ep.ImageURLs, ep.ImageURLs, "image URLs not equal")
		assert.Equal(_ep.MediaURL, ep.MediaURL, "media URL not equal")
	}
}

func ExampleGetBroadcast() {
	url := "http://www.npo.nl/radio-bergeijk/05-06-2004/POMS_VPRO_397233"

	if b, err := GetBroadcast(url); err == nil {
		fmt.Println(b.Title)
		fmt.Println(b.Date)
		fmt.Println(b.Type)
	}

	// Output:
	// Het radiostation voor Bergeijk - Radio Bergeijk
	// 2004-06-05 13:32:00 +0200 CEST
	// Audio
}

func ExampleParseBroadcast() {
	url := "http://www.npo.nl/radio-bergeijk/05-06-2004/POMS_VPRO_397233"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if b, err := ParseBroadcast(resp.Body); err == nil {
		fmt.Println(b.Title)
		fmt.Println(b.Date)
		fmt.Println(b.Type)
	}

	// Output:
	// Het radiostation voor Bergeijk - Radio Bergeijk
	// 2004-06-05 13:32:00 +0200 CEST
	// Audio
}
