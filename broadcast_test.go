package gemist

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBroadcast_audio(t *testing.T) {
	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	r := strings.NewReader(testDataBroadcastAudio)

	_b := Broadcast{
		MediaItem: MediaItem{
			Title:       "Radio bergeijk - Radio Bergeijk",
			Description: "Met in de eerste nationale aflevering voor de regio Bergeijk ooit:\nDe rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk mag de muziek niet ontbreken.",
			ImageURLs:   []string{"http://images.poms.omroep.nl/image/215303.png"},
			URL:         "http://www.npo.nl/radio-bergeijk/03-04-2001/POMS_VPRO_396139",
		},
		LongDescription: " NieuwsWeerVerkeer: o.a. tarieven varkensbelasting\nMestberichten: De Giertijden\nStudiogast: Arno Vlemmings (Opvoedinstructies)\nmuziek: Onze vader zei (trad.) - Moek (van de CD: 'Brabant-bont-ste Vol. 2 Clipsound CCD 994)\nSportnieuws: Afgelastingen",
		Date:            time.Date(2001, time.April, 3, 0, 44, 0, 0, l),
		Length:          885000000000,
		Type:            Audio,
		MediaURL:        "http://download.omroep.nl/vpro/29/08/57/39/POMS_VPRO_396139.mp3",
	}

	b, err := ParseBroadcast(r)
	assertParsedBroadcast(t, &_b, b, err)
}

func TestParseBroadcast_video(t *testing.T) {
	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	r := strings.NewReader(testDataBroadcastVideo)

	_b := Broadcast{
		MediaItem: MediaItem{
			Title:       "Het clusterbom gevoel - ZEMBLA",
			Description: "Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.",
			URL:         "http://www.npo.nl/zembla/18-03-2007/VARA_101141965",
			ImageURLs:   []string{"http://images.poms.omroep.nl/image/6848.png"},
		},
		LongDescription: " Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.",
		Date:            time.Date(2007, time.March, 18, 20, 35, 0, 0, l),
		Length:          3000000000000,
		Type:            Video,
		MediaURL:        "http://www.npo.nl/zembla/18-03-2007/VARA_101141965",
	}

	b, err := ParseBroadcast(r)
	assertParsedBroadcast(t, &_b, b, err)
}

func TestParseBroadcast_video_NPO3(t *testing.T) {
	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	r := strings.NewReader(testDataBroadcastVideoNPO3)

	_b := Broadcast{
		MediaItem: MediaItem{
			Title:       "Radio Bergeijk, toewijding in beeld",
			Description: "Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met technicus Ted van Lieshout.",
			ImageURLs:   []string{},
			URL:         "http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739",
		},
		LongDescription: " Serie verslagen van hun levenswerk: Radio Bergeijk. Een filmploeg over de radiovloer! Ze moesten even aan de gedachte wennen, de presentatoren van Radio Bergeijk. En het heeft even geduurd voordat Toon Spoorenberg en Peer van Eersel konden instemmen met die druktemakerij in hun zo vertrouwde en overzichtelijke radiobestaan. Maar het is er dan toch van gekomen. Pieter Verhoeff heeft een serie gemaakt van Radio Bergeijk, het ook landelijk bekende radiostation voor Bergeijk, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Ted van Lieshout. De pro Deo werkende radiomannen ontvangen in hun studio een keur aan plaatselijke gasten of doen rechtstreeks verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. De televisieserie wordt niet alleen een fel realistisch portret van de makers van het vermaarde radiostation, maar via hen ook van het dorp Bergeijk. Radio Bergeijk is en probeert een satire te zijn op een lokaal radiostation. In de verbeelding van de makers is het fictieve dorp Bergeijk een huiveringwekkend oord waar hedendaagse xenofobie, orthodox atheïsme, vrouwonvriendelijkheid en homo-afkeer om de voorrang strijden. Het radiostation probeert van deze benepen dorpse identiteit pseudo-professioneel verslag te doen. En hanteert als stijlmiddel de grove overdrijving. Teksten en spel: Pieter Bouwman en George van Houts. Gastacteurs o.a.: Alex Klaasen, Jeroen van Merwijk, Marjan Luif, Annet Malherbe, Anniek Pheifer, Margôt Ros en Lies Visschedijk.",
		Date:            time.Date(2007, time.June, 25, 23, 10, 0, 0, l),
		Length:          2100000000000,
		Type:            Video,
		MediaURL:        "http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739",
	}

	b, err := ParseBroadcast(r)
	assertParsedBroadcast(t, &_b, b, err)
}

func assertParsedBroadcast(t *testing.T, _b, b *Broadcast, err error) {
	assert := assert.New(t)
	if assert.NoError(err) {
		assert.Equal(_b.Title, b.Title, "title not equal")
		assert.Equal(_b.Description, b.Description, "description not equal")
		assert.Equal(_b.LongDescription, b.LongDescription, "long description not equal")
		assert.Equal(_b.URL, b.URL, "broadcast URL not equal")
		assert.True(_b.Date.Equal(b.Date), "date not equal (%v != %v)", _b.Date, b.Date)
		assert.Equal(_b.Length, b.Length, "length not equal")
		assert.Equal(_b.Type, b.Type, "type not equal")
		assert.Equal(_b.ImageURLs, b.ImageURLs, "image URLs not equal")
		assert.Equal(_b.MediaURL, b.MediaURL, "media URL not equal")
	}
}

var testDataBroadcastAudio = `<!DOCTYPE html>
<html>
<head>
<title>
Radio Bergeijk: Radio bergeijk kijk je op npo.nl
</title>
<script type="text/javascript">
//<![CDATA[
if (window != top) { top.location.href='http://www.npo.nl'; window.stop(); }
//]]>
</script>
<meta charset='UTF-8'>
<meta content='npo.nl' name='og:site_name'>
<meta content='//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb-fc021b141c9ccde825d868cd10a14ab6.png' name='image-placeholder'>
<meta content='IE=EmulateIE9' http-equiv='X-UA-Compatible'>
<meta content='width=device-width, initial-scale=1, maximum-scale=1' name='viewport'>
<meta content='http://www.npo.nl/suggesties' name='search-suggestions-url'>
<meta content='app-id=nl.uitzendinggemist' name='google-play-app'>
<meta content='//www-assets.npo.nl/assets/google_play_logo-c0f9cdc10f5633cc362e5b8113a3350e.png' name='google-play-logo'>
<meta content='16308a66-14c6-408b-bc68-2aebdb030636' name='msApplication-ID'>
<meta content='7200' name='utc-offset'>
<meta content='30' name='player-tick-length'>
<link href="//cookiesv2.publiekeomroep.nl/static/css/npo_cc_styles.css" media="screen" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/vendors-724ba1e7641e3d78aef2fcd9e1354f92.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/layout-d2942bb06b5bfc5e3869be130af53bc4.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/components-6c045005ff3e4f1125dd5e332d1ede15.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/application-ab6108433cf7b3dc4912c0aa2cd5c80f.css" media="all" rel="stylesheet" type="text/css" />
<!--[if lte IE 8]>
<link href="//www-assets.npo.nl/assets/ie-52a7ee466e06b2377f85de26a5c4efa4.css" media="all" rel="stylesheet" type="text/css" />
<![endif]-->
<!--[if lte IE 7]>
<link href="//www-assets.npo.nl/assets/ie7-87b61b941651ecf598c4db2055c80f41.css" media="all" rel="stylesheet" type="text/css" />
<![endif]-->
<script src="//www-assets.npo.nl/assets/application-64c4dbaacb4e6e1077795f2890897042.js" type="text/javascript"></script>
<script src="//cookiesv2.publiekeomroep.nl/data/script/cconsent.min.js" type="text/javascript"></script>
<link href="http://www.npo.nl/uitgelicht.rss" rel="alternate" title="Uitgelicht op npo.nl" type="application/rss+xml" />
<link href="http://www.npo.nl/artikelen.rss" rel="alternate" title="Artikelen op npo.nl" type="application/rss+xml" />
<meta content='{"c1":"2","c2":"17827132","ns_site":"po-totaal","potag1":"npoportal","potag3":"npo","potag4":"npo","potag5":"zenderportal","potag6":"video","potag7":"geen","potag8":"site","potag9":"site"}' name='scorecard-default-labels'>
<meta content='npo.softclick' name='scorecard-default-prefix'>
<meta content="Met in de eerste nationale aflevering voor de regio Bergeijk ooit: &#x000A;De rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk mag de muziek niet ontbreken. NieuwsWeerVerkeer: o.a. tarieven varkensbelasting &#x000A;Mestberichten: De Giertijden &#x000A;Studiogast: Arno Vlemmings (Opvoedinstructies) &#x000A;muziek: Onze vader zei (trad.) - Moek (van de CD: 'Brabant-bont-ste Vol. 2 Clipsound CCD 994) &#x000A;Sportnieuws: Afgelastingen" name='description'>
<script id='signup-popover' type='text/x-jquery-tmpl'><div class='signup-popover'>
<h3 class='omroep-bold'>Wil je dit programma bewaren als favoriet?</h3>
<p>Maak een eigen account aan en bewaar je favoriete programma’s. Een account zorgt ook voor tips op maat.</p>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2F03-04-2001%2FPOMS_VPRO_396139" class="button action-button">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken" class="button action-button">Account aanmaken</a>
</div>
</script>
<meta content='true' name='guest'>



<script type="text/javascript">
//<![CDATA[
npo_cookies.init();
//]]>
</script>
<script src="//www-assets.npo.nl/assets/radioplayer-756e90c15879396d9c200f005a9dc5e1.js" type="text/javascript"></script>
<meta content="authenticity_token" name="csrf-param" />
<meta content="PF9x3QovU2Mj1CeLgKDLRm07T8PsCYfIHL3UY+c18Vo=" name="csrf-token" />
<meta content="Radio bergeijk - Radio Bergeijk" name="og:title" />
<meta content="http://www.npo.nl/radio-bergeijk/03-04-2001/POMS_VPRO_396139" name="og:url" />
<meta content="http://images.poms.omroep.nl/image/215303.png" name="og:image" />
<meta content="Met in de eerste nationale aflevering voor de regio Bergeijk ooit:
De rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk mag de muziek niet ontbreken." name="og:description" />
<meta content="music.radio_station" name="og:type" />
<meta content="http://download.omroep.nl/vpro/29/08/57/39/POMS_VPRO_396139.mp3" name="og:audio" />
</head>
<body>
<!-- Begin comScore Inline Tag 1.1302.13 -->
<script type="text/plain" class="npo_cc_inline_analytics">
// <![CDATA[
function udm_(e){var t="comScore=",n=document,r=n.cookie,i="",s="indexOf",o="substring",u="length",a=2048,f,l="&ns_",c="&",h,p,d,v,m=window,g=m.encodeURIComponent||escape;if(r[s](t)+1)for(d=0,p=r.split(";"),v=p[u];d<v;d++)h=p[d][s](t),h+1&&(i=c+unescape(p[d][o](h+t[u])));e+=l+"_t="+ +(new Date)+l+"c="+(n.characterSet||n.defaultCharset||"")+"&c8="+g(n.title)+i+"&c7="+g(n.URL)+"&c9="+g(n.referrer),e[u]>a&&e[s](c)>0&&(f=e[o](0,a-8).lastIndexOf(c),e=(e[o](0,f)+l+"cut="+g(e[o](f+1)))[o](0,a)),n.images?(h=new Image,m.ns_p||(ns_p=h),h.src=e):n.write("<","p","><",'img src="',e,'" height="1" width="1" alt="*"',"><","/p",">")};
udm_('http'+(document.location.href.charAt(4)=='s'?'s://sb':'://b')+'.scorecardresearch.com/b?c1=2&c2=17827132&name=npo.programmas.radio-bergeijk.afleveringen.POMS_VPRO_396139&npo_ingelogd=no&npo_login_id=geen&ns_site=po-totaal&potag1=npoportal&potag2=programmas&potag3=npo&potag4=npo&potag5=zenderportal&potag6=video&potag7=geen&potag8=site&potag9=site');
// ]]>
</script>
<noscript><p><img src="http://b.scorecardresearch.com/p?c1=2&amp;c2=17827132&amp;name=npo.programmas.radio-bergeijk.afleveringen.POMS_VPRO_396139&amp;npo_ingelogd=no&amp;npo_login_id=geen&amp;ns_site=po-totaal&amp;potag1=npoportal&amp;potag2=programmas&amp;potag3=npo&amp;potag4=npo&amp;potag5=zenderportal&amp;potag6=video&amp;potag7=geen&amp;potag8=site&amp;potag9=site" height="1" width="1" alt="*"></p></noscript>
<!-- End comScore Inline Tag -->


<div class="hidden-item-props"><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/" itemprop="url"><span itemprop="title">NPO</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/radio-bergeijk/POMS_S_VPRO_396280" itemprop="url"><span itemprop="title">Radio Bergeijk</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="http://www.npo.nl/radio-bergeijk/03-04-2001/POMS_VPRO_396139" itemprop="url"><span itemprop="title">Radio bergeijk</span></a></div></div>
<div class='popup-overlay hide' id='popup-overlay'></div>
<div class='popup-alerts' id='popup-alerts'>
</div>
<div class='page-container programs'>
<div class='mobile-trigger visible-with-grid-collapse'></div>
<div class='page-header sub-navigation-collapse' id='page-header'>
<div class='container visible-for-desktop'>
<a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.logo&quot;}"><div class='logo'><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></div>
</a><div class='profile-box signed-out-box'>
<a href="http://www.npoplus.nl" data-scorecard="{&quot;name&quot;:&quot;home.npo-plus&quot;}">NPO Plus</a>
<span class='signed-out'>
<a href="https://mijn.npo.nl/inloggen">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken">Account aanmaken</a>
</span>
</div>
<div class='page-navigation'>
<ul>
<li><a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.start&quot;}" data-show-sub-navigation="home">Start</a></li>
<li><a href="http://www.npo.nl/live" data-show-sub-navigation="tv_channels">Tv</a></li>
<li><a href="http://www.npo.nl/radio" data-show-sub-navigation="radio_channels">Radio</a></li>
<li class="active"><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
</ul>
</div>

<form accept-charset="UTF-8" action="http://www.npo.nl/zoeken" id="search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='page-search-box'>
<div class='search-box'>
<span class="npo-glyph search"></span>
<input autocomplete="off" class="search-query-box" id="search-query" name="q" placeholder="zoeken binnen npo.nl&hellip;" tabindex="1" type="text" />
</div>
<div class='search-suggestions hide' id='search-suggestions'>
<h2>Snel naar een programma</h2>
<div class='suggestion-box'></div>
<a href="/zoeken" class="all-results" id="all-results" tabindex="-1">Bekijk alle zoekresultaten &raquo</a>
</div>
</div>
</form>


</div>
<div class='visible-for-desktop'>
<div class='page-sub-navigation' data-sub-navigation-category='tv_channels'>
<div class='container'>
<div class='content clearfix'>
<div class='slider slide-left'>◀</div>
<div class='sub-navigation-container'>
<ul>
<li><a href="http://www.npo.nl/live/npo-1" data-image="//www-assets.npo.nl/uploads/tv_channel/263/logo/regular_logo-npo1.png">NPO 1</a></li>
<li><a href="http://www.npo.nl/live/npo-2" data-image="//www-assets.npo.nl/uploads/tv_channel/264/logo/regular_logo-npo2.png">NPO 2</a></li>
<li><a href="http://www.npo.nl/npo3" data-image="//www-assets.npo.nl/uploads/tv_channel/265/logo/regular_npo3-logo.png" data-scorecard="{&quot;name&quot;:&quot;trending.logo&quot;,&quot;prefix&quot;:&quot;npo3.softclick&quot;}">NPO 3</a></li>
<li><a href="http://www.zapp.nl/nu-straks" data-image="//www-assets.npo.nl/assets/logos/logo-npozapp-regular-ea6be75d1d99c3f38e47faff9a3b78ed.png" data-scorecard="{&quot;name&quot;:&quot;clickout.zapp&quot;,&quot;prefix&quot;:&quot;npo.softclick&quot;}" target="_blank">Zapp <span class="npo-glyph arrow-jump-box"></span></a></li>
<li><a href="http://www.zappelin.nl/tv-kijken" data-image="//www-assets.npo.nl/assets/logos/logo-npozappelin-regular-7891aaa672e4eb3932469b15e0b7264c.png" data-scorecard="{&quot;name&quot;:&quot;clickout.zappelin&quot;,&quot;prefix&quot;:&quot;npo.softclick&quot;}" target="_blank">Zappelin <span class="npo-glyph arrow-jump-box"></span></a></li>
<li><a href="http://www.npo.nl/live/npo-nieuws" data-image="//www-assets.npo.nl/uploads/tv_channel/279/logo/regular_logonieuws.png">NPO Nieuws</a></li>
<li><a href="http://www.npo.nl/live/npo-cultura" data-image="//www-assets.npo.nl/uploads/tv_channel/280/logo/regular_logocultura.png">NPO Cultura</a></li>
<li><a href="http://www.npo.nl/live/npo-101" data-image="//www-assets.npo.nl/uploads/tv_channel/281/logo/regular_logo-101.png">NPO 101</a></li>
<li><a href="http://www.npo.nl/live/npo-politiek" data-image="//www-assets.npo.nl/uploads/tv_channel/282/logo/regular_logopolitiek.png">NPO Politiek</a></li>
<li><a href="http://www.npo.nl/live/npo-best" data-image="//www-assets.npo.nl/uploads/tv_channel/283/logo/regular_logobest.png">NPO Best</a></li>
<li><a href="http://www.npo.nl/live/npo-doc" data-image="//www-assets.npo.nl/uploads/tv_channel/284/logo/regular_logodoc.png">NPO Doc</a></li>
<li><a href="http://www.npo.nl/live/npo-zappxtra" data-image="//www-assets.npo.nl/uploads/tv_channel/288/logo/regular_logozappxtra.png">NPO Zapp Xtra</a></li>
<li><a href="http://www.npo.nl/live/npo-humor-tv" data-image="//www-assets.npo.nl/uploads/tv_channel/290/logo/regular_logohumor.png">NPO Humor TV</a></li>
</ul>
</div>
<div class='slider slide-right'>►</div>
</div>
</div>
</div>

</div>
<div class='main-navigation visible-for-small-screen'>
<div class='container'>
<a href="#" class="menu-button" id="menu-button"><span class="npo-glyph list contrast"></span>
Menu
</a><div class='border-wrapper left'>
<a href="http://www.npo.nl/" class="logo"><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></a>
</div>
<div class='border-wrapper right'>
<a href="http://www.npo.nl/profiel" class="profile"><span class="npo-glyph user gray"></span>
</a><a href="http://www.npo.nl/zoeken" class="search-button" id="page-search-button"><span class="npo-glyph search gray"></span>
</a></div>
</div>
</div>
<div class='sub-navigation visible-for-small-screen'>
<div class='container'>
<div class='page-navigation'>
<ul>
<li><a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.start&quot;}" data-show-sub-navigation="home">Start</a></li>
<li><a href="http://www.npo.nl/live" data-show-sub-navigation="tv_channels">Tv</a></li>
<li><a href="http://www.npo.nl/radio" data-show-sub-navigation="radio_channels">Radio</a></li>
<li class="active"><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
</ul>
</div>

</div>
</div>
<div class='sub-search visible-for-small-screen'>
<div class='container'>
<form accept-charset="UTF-8" action="http://www.npo.nl/zoeken" id="search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='page-search-box'>
<div class='search-box'>
<span class="npo-glyph search"></span>
<input autocomplete="off" class="search-query-box" id="search-query" name="q" placeholder="zoeken binnen npo.nl&hellip;" tabindex="1" type="text" />
</div>
<div class='search-suggestions hide' id='search-suggestions'>
<h2>Snel naar een programma</h2>
<div class='suggestion-box'></div>
<a href="/zoeken" class="all-results" id="all-results" tabindex="-1">Bekijk alle zoekresultaten &raquo</a>
</div>
</div>
</form>


</div>
</div>
</div>
<div class="schema-org-props hidden-item-props"><div itemscope="" itemtype="http://schema.org/TVSeries"><a href="/radio-bergeijk/POMS_S_VPRO_396280" itemprop="url"><span itemprop="name">Radio Bergeijk</span></a><span itemprop="description">Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een keur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.
 </span><div itemscope="" itemtype="http://schema.org/TVEpisode"><a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139" itemprop="url"><span itemprop="name">Radio bergeijk</span></a><span itemprop="description">Met in de eerste nationale aflevering voor de regio Bergeijk ooit:
De rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk mag de muziek niet ontbreken.</span><div itemscope="" itemtype="http://schema.org/BroadcastEvent"><span content="2001-04-03T00:44:00+02:00" itemprop="startDate">2001-04-03 00:44:00 +0200</span></div></div></div></div>
<div class='share-modal modal hide fade' data-id='POMS_VPRO_396139' id='share-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2 class='omroep-bold'>Radio Bergeijk</h2>
<h3 class='omroep-light'>Radio bergeijk</h3>
</div>
<div class='modal-body'><div class='row-fluid'>
<div class='span4 social-buttons'>
<h3>Deel via social media</h3>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/delen/facebook?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2F03-04-2001%2FPOMS_VPRO_396139" class="button facebook" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.facebook.radio-bergeijk.POMS_VPRO_396139&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph facebook"></span>Facebook</a>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/delen/google?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2F03-04-2001%2FPOMS_VPRO_396139" class="button google" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.google.radio-bergeijk.POMS_VPRO_396139&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph google"></span>Google</a>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/delen/twitter?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2F03-04-2001%2FPOMS_VPRO_396139" class="button twitter" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.twitter.radio-bergeijk.POMS_VPRO_396139&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph twitter"></span>Twitter</a>
<a href="whatsapp://send?text=Bekijk%20Radio%20Bergeijk:%20http://npo.nl/POMS_VPRO_396139" class="button whatsapp hide" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.whatsapp.radio-bergeijk.POMS_VPRO_396139&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph whatsapp"></span>Whatsapp</a>
</div>
<div class='span8'>
<div class='row-fluid switch'>
<div class='span6 switch-block switch-left'>
<a href="#" class="switch-link dark-well" data-block="link"><span class="npo-glyph globe"></span>
Link
</a></div>
<div class='span6 switch-block switch-right'>
<a href="#" class="switch-email dark-well inactive" data-block="email"><span class="npo-glyph email"></span>
E-mail vriend
</a></div>
</div>
<div class='link-block switch-list'>
<div class='row-fluid link-inputs'>
<div class='span12'>
<div class='time-inputs'>
<label>Starttijd</label>
<div class='share-action-wrapper'>
<a href="#" class="share-start-link">Start vanaf huidige positie</a>
</div>
<input id="share_start_position" name="share_start_position" placeholder="00:00:00" type="text" value="00:00:00" />
</div>
<div class='share-path'>
<label>Directe link</label>
<div class='share-action-wrapper' id='share-modal-url-copy-wrapper'>
<div class='share-action-link' id='share-modal-url-copy'>Kopieer link</div>
</div>
<input id="share_url" name="share_url" type="text" value="http://www.npo.nl/radio-bergeijk/03-04-2001/POMS_VPRO_396139" />
</div>
</div>
</div>
</div>
<div class='email-friend email-block switch-list'>
<h3>E-mail vriend</h3>
<form accept-charset="UTF-8" action="/shares" class="simple_form new_share" id="new_share_" method="post" novalidate="novalidate"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="PF9x3QovU2Mj1CeLgKDLRm07T8PsCYfIHL3UY+c18Vo=" /></div><div class='row-fluid'>
<div class='span5'>
<input id="id" name="id" type="hidden" value="POMS_VPRO_396139" />
<input id="type" name="type" type="hidden" value="program" />
<input id="mid" name="mid" type="hidden" value="POMS_VPRO_396139" />
<div class="input string required share_from"><label class="string required control-label" for="share_from">* Je naam</label><input class="string required" id="share_from" name="share[from]" size="50" type="text" /></div>
</div>
<div class='span7'>
<div class="input email required share_from_email"><label class="email required control-label" for="share_from_email">* Je e-mailadres</label><input class="string email required" id="share_from_email" name="share[from_email]" placeholder="naam@voorbeeld.nl" size="50" type="email" /></div>
</div>
</div>
<div class='row-fluid'>
<div class='span5'>
<div class="input string required share_to"><label class="string required control-label" for="share_to">* Naam vriend</label><input class="string required" id="share_to" name="share[to]" size="50" type="text" /></div>
</div>
<div class='span7'>
<div class="input email required share_to_email"><label class="email required control-label" for="share_to_email">* E-mailadres vriend</label><input class="string email required" id="share_to_email" name="share[to_email]" placeholder="naam@voorbeeld.nl" size="50" type="email" /></div>
</div>
</div>
<div class='row-fluid'>
<div class='span12'>
<div class="input text optional share_message"><label class="text optional control-label" for="share_message">Bericht (optioneel)</label><textarea class="text optional" cols="40" id="share_message" name="share[message]" rows="5">
</textarea></div>
</div>
</div>
<div class='row-fluid actions'>
<div class='span12'>
<input class="button right" name="commit" type="submit" value="Versturen" />
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139" class="button gray right" data-dismiss="modal">Annuleren</a>
</div>
</div>
</form>


</div>
</div>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='adaptive-modal modal hide fade' id='adaptive-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2>Instellingen</h2>
<h3>Videokwaliteit in het buitenland</h3>
</div>
<div class='modal-body'>
<p>De nieuwe player kan problemen veroorzaken met de connectie in het buitenland. Ervaart u problemen met npo.nl, zet onderstaande switch dan aan. De videospeler gaat bufferen en na enkele momenten kunt u de uitzending weer bekijken.</p>
<div class='switch-wrapper'>
Buitenland
<i class="npo-icon-radio-switch off adaptive-video-switch"></i>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='program-show-page'>
<div class="radio-showcase showcase showcase-fluid with-player program-player" id="header"><div class="showcase-background"><div class="header-gradient"></div><div class="showcase-wrapper"><div class='info no-overlay-for-small-screen'>
<div class='container'>
<div class='row-fluid title-and-broadcasters'>
<div class='broadcaster-logo'><a href="http://www.vpro.nl" data-scorecard="{&quot;name&quot;:&quot;vpro&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas.extern&quot;}" target="_blank" title="VPRO"><img alt="Logo van VPRO" class="no-shadow" src="//www-assets.npo.nl/uploads/broadcaster/40/logo/regular_regular_VPRO_alt2.png" /></a></div>
<h1 title="Radio Bergeijk"><a href="/radio-bergeijk/POMS_S_VPRO_396280">Radio Bergeijk</a></h1>
</div>
<div class='row-fluid videoplayer-and-metadata'>
<div class='span-main player-span'>
<div class='video-image-wrapper'>
<img alt="Afbeelding van Radio bergeijk" class="show-image" src="http://images.poms.omroep.nl/image/s564/c564x315/215303.png" />
</div>
<div class='row-fluid non-responsive' data-radio-player data-streams='http://download.omroep.nl/vpro/29/08/57/39/POMS_VPRO_396139.mp3' id='radio-player'>
<div class='span3'>
<div class='play-button-container'><a href="#" class="radio-player-icon play toggle  radio-player-play large" data-initial-state="play" data-playing-state="pause"></a></div>
</div>
<div class='span9'>
<div class='row-fluid non-responsive timebar'>
<div class='span4'>
<span class='runtime'>0:00</span>
/
<span class='duration'>14:45</span>
</div>
<div class='span8'>
<div class='time-bar-wrapper'>
<div class='time-bar'></div>
</div>
</div>
</div>
</div>
</div>
</div>
<div class='span-sidebar'>

<div class='metadata'>
<div class='content'>
<p class='overflow-description' title='Met in de eerste nationale aflevering voor de regio Bergeijk ooit: &#x000A;De rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk mag de muziek niet ontbreken.'><span class="hide-original">Met in de eerste nationale aflevering voor de regio Bergeijk ooit:
De rubriek NieuwsWeerVerkeer met zeer belangrijk nieuws over de varkensbelasting en blaasproblemen, het weer en meer. In mestberichten de giertijden want de klinker zit weer in de maand. In de studio te gast is Arno Vlemmings die ons bijpraat over de opvoedkunst. Peer van Eersel brengt het sportnieuws. En natuurlijk m...</span><span class="omission" data-scorecard="{&quot;name&quot;:&quot;programmas.radio-bergeijk.afleveringen.POMS_VPRO_396139.meer-beschrijving&quot;}"><a href="#" class="overflow-toggle">Meer</a></span><span class="omission-overflow hide"> NieuwsWeerVerkeer: o.a. tarieven varkensbelasting
Mestberichten: De Giertijden
Studiogast: Arno Vlemmings (Opvoedinstructies)
muziek: Onze vader zei (trad.) - Moek (van de CD: 'Brabant-bont-ste Vol. 2 Clipsound CCD 994)
Sportnieuws: Afgelastingen<a href="#" class="overflow-toggle less">Minder</a></span></p>

<div class='dark-well meta-data-box'>
<div class='data'>
<div class='broadcast-time'>
Di 3 apr 2001 00:44
</div>
<div class='genres'>

</div>
<div class='ratings'>
</div>
</div>
</div>
</div>
</div>

</div>
<div class='actions'>
<div class='action-buttons for-guest'>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2F03-04-2001%2FPOMS_VPRO_396139" class="button favorite-button guest-favorite-button"><span class="npo-glyph heart contrast"></span> <span class="glyph-label">Favoriet</span></a>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/delen" class="last button share-modal-button" data-id="POMS_VPRO_396139" rel="nofollow"><span class="npo-glyph share contrast"></span> <span class="glyph-label">Delen</span></a>
</div>

</div>
</div>
<div class='notifications'>

</div>

<div class='row-fluid program-fragments'>
<h2 class='carousel-header'>Fragmenten (1)</h2>
<div class='items-carousel tiles' data-dont-link-in-place>
<div class='container'>
<div class='items-navigation'>
<a href="#" class="carousel-left" id="carousel-left"><i class="npo-icon-navigate-left"></i>
</a><a href="#" class="carousel-right" id="carousel-right"><i class="npo-icon-navigate-right"></i>
</a></div>
<div class='items-carousel-scroller'>
<div class='items-carousel-wrapper'>
<div class="items-carousel-item " data-mid="POMS_VPRO_396140" data-video-mid="POMS_VPRO_396140">
<div class='list-item tile'>
<div class='contextual-main-title'>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/POMS_VPRO_396140"><h3>Radio bergeijk</h3>
</a></div>
<div class='image-container'>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/POMS_VPRO_396140"><img alt="Afbeelding van Radio Bergeijk Radiokwis vraag 1" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
</a></div>
<div class='description'>
<h4 title='Radio Bergeijk Radiokwis vraag 1'>
<a href="/radio-bergeijk/03-04-2001/POMS_VPRO_396139/POMS_VPRO_396140"><div class='default-title'>
Radio Bergeijk Radiokwis vraag 1
<span class='inactive'>(VPRO)</span>
</div>
<div class='contextual-title'>Radio Bergeijk Radiokwis vraag 1</div>
</a></h4>

<h5>
Di 3 apr 2001
</h5>
</div>
</div>


</div></div>
</div>
</div>
</div>

</div>
</div>
</div>
</div></div></div>
<div class='content' data-auto-load data-url='/radio-bergeijk/03-04-2001/POMS_VPRO_396139/show_content'></div>
</div>

<div class="page-footer" id="npo-footer"><div class='container'>
<div class='broadcasters-and-corporate-links'>
<div class='broadcasters'>
<strong>Omroepen</strong>
<ul>
<li><a href="http://www.avrotros.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.avrotros&quot;}" target="_blank">AVROTROS</a></li>
<li><a href="http://www.bnn.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.bnn&quot;}" target="_blank">BNN</a></li>
<li><a href="http://www.eo.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.eo&quot;}" target="_blank">EO</a></li>
<li><a href="http://www.omroephuman.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.human&quot;}" target="_blank">HUMAN</a></li>
<li><a href="http://www.kro-ncrv.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.kro-ncrv&quot;}" target="_blank">KRO-NCRV</a></li>
</ul>
<ul>
<li><a href="http://www.omroepmax.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.max&quot;}" target="_blank">MAX</a></li>
<li><a href="http://www.nos.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.nos&quot;}" target="_blank">NOS</a></li>
<li><a href="http://www.ntr.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.ntr&quot;}" target="_blank">NTR</a></li>
<li><a href="http://www.powned.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.powned&quot;}" target="_blank">PowNed</a></li>
<li><a href="http://www.vara.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.vara&quot;}" target="_blank">VARA</a></li>
</ul>
<ul>
<li><a href="http://www.vpro.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.vpro&quot;}" target="_blank">VPRO</a></li>
<li><a href="http://www.omroepwnl.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.wnl&quot;}" target="_blank">WNL</a></li>
<li><a href="http://www.zvk.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.zvk&quot;}" target="_blank">ZvK</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Meer NPO</strong>
<ul>
<li><a href="http://www.npoplus.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.npo-plus&quot;}" target="_blank">NPO Plus</a></li>
<li><a href="http://www.bvn.tv" data-scorecard="{&quot;name&quot;:&quot;footer.extern.bvn&quot;}" target="_blank">BVN</a></li>
<li><a href="http://www.npo.nl/specials/regionale-omroepen">Regionaal</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Organisatie</strong>
<ul>
<li><a href="http://www.npo.nl/overnpo">Over NPO</a></li>
<li><a href="http://www.npo.nl/overnpo/pers">Pers</a></li>
<li><a href="http://www.npo.nl/overnpo/vacatures-en-stages-npo">Vacatures</a></li>
<li><a href="http://www.nposales.com" data-scorecard="{&quot;name&quot;:&quot;footer.extern.npo-sales&quot;}" target="_blank">NPO Sales</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Volg de NPO</strong>
<ul>
<li><a href="http://omroep.dmd.omroep.nl/x/plugin/?pName=subscribe&amp;MIDRID=S7Y1swQAA98&amp;pLang=nl&amp;Z=543417650" data-scorecard="{&quot;name&quot;:&quot;footer.extern.aanmelden-nieuwsbrief&quot;}" target="_blank">Aanmelden nieuwsbrief</a></li>
<li><a href="https://instagram.com/npo.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.instagram&quot;}" target="_blank">Instagram</a></li>
<li><a href="https://www.facebook.com/NPO.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.facebook&quot;}" target="_blank">Facebook</a></li>
<li><a href="http://www.twitter.com/publiekeomroep" data-scorecard="{&quot;name&quot;:&quot;footer.extern.twitter&quot;}" target="_blank">Twitter</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Onze apps</strong>
<ul>
<li><a href="https://itunes.apple.com/nl/app/uitzending-gemist/id323998316?mt=8" data-scorecard="{&quot;name&quot;:&quot;footer.extern.ipad-iphone&quot;}" target="_blank">iPad &amp; iPhone</a></li>
<li><a href="https://play.google.com/store/apps/details?id=nl.uitzendinggemist" data-scorecard="{&quot;name&quot;:&quot;footer.extern.android&quot;}" target="_blank">Android</a></li>
<li><a href="https://help.npo.nl/faqs/smarttv-hbbtv">Smarttv en HbbTV</a></li>
</ul>
</div>
</div>
<div class='general-links'>
<ul class='disclaimers'>
<li><a href="http://www.npo.nl/contact">Contact</a></li>
<li><a href="http://help.npo.nl
" data-scorecard="{&quot;name&quot;:&quot;footer.extern.help&quot;}" target="_blank">Help</a></li>
<li><a href="http://www.npo.nl/uitgelicht.rss">RSS</a></li>
<li><a href="http://www.npo.nl/algemene-voorwaarden-privacy">Algemene Voorwaarden &amp; Privacy</a></li>
<li><a href="http://cookiesv2.publiekeomroep.nl/" data-scorecard="{&quot;name&quot;:&quot;footer.extern.cookiebeleid&quot;}" target="_blank">Cookiebeleid</a></li>
</ul>
<div class='logo small'><a href="http://www.npo.nl/"><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></a></div>
</div>
</div>
</div>
</div>
<script class="npo_cc_inline_analytics" language="JavaScript1.3" src="//b.scorecardresearch.com/c2/17827132/cs.js" type="text/plain"></script>

<script type="text/plain" class="npo_cc_inline_advertising">(function() {
  var _fbq = window._fbq || (window._fbq = []);
  if (!_fbq.loaded) {
    var fbds = document.createElement('script');
    fbds.async = true;
    fbds.src = '//connect.facebook.net/en_US/fbds.js';
    var s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(fbds, s);
    _fbq.loaded = true;
  }
  _fbq.push(['addPixelId', '1487544264866945']);
})();
window._fbq = window._fbq || [];
window._fbq.push(['track', 'PixelInitialized', {}]);
</script>
<noscript><img height="1" width="1" alt="" style="display:none" class="npo_cc_inline_advertising" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAA AICTAEAOw==" data-src="https://www.facebook.com/tr?id=1487544264866945&amp;ev=PixelInitialized" /></noscript>
</body>
</html>
`

var testDataBroadcastVideo = `<!DOCTYPE html>
<html>
<head>
<title>
ZEMBLA: Het clusterbom gevoel kijk je op npo.nl
</title>
<script type="text/javascript">
//<![CDATA[
if (window != top) { top.location.href='http://www.npo.nl'; window.stop(); }
//]]>
</script>
<meta charset='UTF-8'>
<meta content='npo.nl' name='og:site_name'>
<meta content='//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb-fc021b141c9ccde825d868cd10a14ab6.png' name='image-placeholder'>
<meta content='IE=EmulateIE9' http-equiv='X-UA-Compatible'>
<meta content='width=device-width, initial-scale=1, maximum-scale=1' name='viewport'>
<meta content='http://www.npo.nl/suggesties' name='search-suggestions-url'>
<meta content='app-id=nl.uitzendinggemist' name='google-play-app'>
<meta content='//www-assets.npo.nl/assets/google_play_logo-c0f9cdc10f5633cc362e5b8113a3350e.png' name='google-play-logo'>
<meta content='16308a66-14c6-408b-bc68-2aebdb030636' name='msApplication-ID'>
<meta content='7200' name='utc-offset'>
<meta content='30' name='player-tick-length'>
<link href="//cookiesv2.publiekeomroep.nl/static/css/npo_cc_styles.css" media="screen" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/vendors-724ba1e7641e3d78aef2fcd9e1354f92.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/layout-d2942bb06b5bfc5e3869be130af53bc4.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/components-6c045005ff3e4f1125dd5e332d1ede15.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/application-ab6108433cf7b3dc4912c0aa2cd5c80f.css" media="all" rel="stylesheet" type="text/css" />
<!--[if lte IE 8]>
<link href="//www-assets.npo.nl/assets/ie-52a7ee466e06b2377f85de26a5c4efa4.css" media="all" rel="stylesheet" type="text/css" />
<![endif]-->
<!--[if lte IE 7]>
<link href="//www-assets.npo.nl/assets/ie7-87b61b941651ecf598c4db2055c80f41.css" media="all" rel="stylesheet" type="text/css" />
<![endif]-->
<script src="//www-assets.npo.nl/assets/application-64c4dbaacb4e6e1077795f2890897042.js" type="text/javascript"></script>
<script src="//cookiesv2.publiekeomroep.nl/data/script/cconsent.min.js" type="text/javascript"></script>
<link href="http://www.npo.nl/uitgelicht.rss" rel="alternate" title="Uitgelicht op npo.nl" type="application/rss+xml" />
<link href="http://www.npo.nl/artikelen.rss" rel="alternate" title="Artikelen op npo.nl" type="application/rss+xml" />
<meta content='{"c1":"2","c2":"17827132","ns_site":"po-totaal","potag1":"npoportal","potag3":"npo","potag4":"npo","potag5":"zenderportal","potag6":"video","potag7":"geen","potag8":"site","potag9":"site"}' name='scorecard-default-labels'>
<meta content='npo.softclick' name='scorecard-default-prefix'>
<meta content='Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen. ' name='description'>
<script id='signup-popover' type='text/x-jquery-tmpl'><div class='signup-popover'>
<h3 class='omroep-bold'>Wil je dit programma bewaren als favoriet?</h3>
<p>Maak een eigen account aan en bewaar je favoriete programma’s. Een account zorgt ook voor tips op maat.</p>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fzembla%2F18-03-2007%2FVARA_101141965" class="button action-button">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken" class="button action-button">Account aanmaken</a>
</div>
</script>
<meta content='true' name='guest'>



<script type="text/javascript">
//<![CDATA[
npo_cookies.init();
//]]>
</script>
<script src="http://npoplayer.omroep.nl/csjs/npoplayer-min.js" type="text/javascript"></script>
<script src="//www-assets.npo.nl/assets/player-event-support-9bb142d8358270024970c91b7dd43104.js" type="text/javascript"></script>
<script defer="defer" src="http://topspin.npo.nl/divolte/divolte.js" type="text/javascript"></script>

<meta content="authenticity_token" name="csrf-param" />
<meta content="PF9x3QovU2Mj1CeLgKDLRm07T8PsCYfIHL3UY+c18Vo=" name="csrf-token" />
<meta content="Het clusterbom gevoel - ZEMBLA" name="og:title" />
<meta content="http://www.npo.nl/zembla/18-03-2007/VARA_101141965" name="og:url" />
<meta content="http://images.poms.omroep.nl/image/6848.png" name="og:image" />
<meta content="Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen." name="og:description" />
<meta content="video.episode" name="og:type" />
<meta content="http://www.npo.nl/zembla/18-03-2007/VARA_101141965" name="og:video" />
<meta content="3000" name="og:video:duration" />
<meta content="2007-03-18 20:35:00 +0100" name="og:video:release_date" />
</head>
<body>
<!-- Begin comScore Inline Tag 1.1302.13 -->
<script type="text/plain" class="npo_cc_inline_analytics">
// <![CDATA[
function udm_(e){var t="comScore=",n=document,r=n.cookie,i="",s="indexOf",o="substring",u="length",a=2048,f,l="&ns_",c="&",h,p,d,v,m=window,g=m.encodeURIComponent||escape;if(r[s](t)+1)for(d=0,p=r.split(";"),v=p[u];d<v;d++)h=p[d][s](t),h+1&&(i=c+unescape(p[d][o](h+t[u])));e+=l+"_t="+ +(new Date)+l+"c="+(n.characterSet||n.defaultCharset||"")+"&c8="+g(n.title)+i+"&c7="+g(n.URL)+"&c9="+g(n.referrer),e[u]>a&&e[s](c)>0&&(f=e[o](0,a-8).lastIndexOf(c),e=(e[o](0,f)+l+"cut="+g(e[o](f+1)))[o](0,a)),n.images?(h=new Image,m.ns_p||(ns_p=h),h.src=e):n.write("<","p","><",'img src="',e,'" height="1" width="1" alt="*"',"><","/p",">")};
udm_('http'+(document.location.href.charAt(4)=='s'?'s://sb':'://b')+'.scorecardresearch.com/b?c1=2&c2=17827132&name=npo.programmas.zembla.afleveringen.VARA_101141965&npo_ingelogd=no&npo_login_id=geen&ns_site=po-totaal&potag1=npoportal&potag2=programmas&potag3=npo&potag4=npo&potag5=zenderportal&potag6=video&potag7=geen&potag8=site&potag9=site');
// ]]>
</script>
<noscript><p><img src="http://b.scorecardresearch.com/p?c1=2&amp;c2=17827132&amp;name=npo.programmas.zembla.afleveringen.VARA_101141965&amp;npo_ingelogd=no&amp;npo_login_id=geen&amp;ns_site=po-totaal&amp;potag1=npoportal&amp;potag2=programmas&amp;potag3=npo&amp;potag4=npo&amp;potag5=zenderportal&amp;potag6=video&amp;potag7=geen&amp;potag8=site&amp;potag9=site" height="1" width="1" alt="*"></p></noscript>
<!-- End comScore Inline Tag -->


<div class="hidden-item-props"><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/" itemprop="url"><span itemprop="title">NPO</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/zembla/POMS_S_VARA_099718" itemprop="url"><span itemprop="title">ZEMBLA</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="http://www.npo.nl/zembla/18-03-2007/VARA_101141965" itemprop="url"><span itemprop="title">Het clusterbom gevoel</span></a></div></div>
<div class='popup-overlay hide' id='popup-overlay'></div>
<div class='popup-alerts' id='popup-alerts'>
</div>
<div class='page-container programs'>
<div class='mobile-trigger visible-with-grid-collapse'></div>
<div class='page-header sub-navigation-collapse' id='page-header'>
<div class='container visible-for-desktop'>
<a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.logo&quot;}"><div class='logo'><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></div>
</a><div class='profile-box signed-out-box'>
<a href="http://www.npoplus.nl" data-scorecard="{&quot;name&quot;:&quot;home.npo-plus&quot;}">NPO Plus</a>
<span class='signed-out'>
<a href="https://mijn.npo.nl/inloggen">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken">Account aanmaken</a>
</span>
</div>
<div class='page-navigation'>
<ul>
<li><a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.start&quot;}" data-show-sub-navigation="home">Start</a></li>
<li><a href="http://www.npo.nl/live" data-show-sub-navigation="tv_channels">Tv</a></li>
<li><a href="http://www.npo.nl/radio" data-show-sub-navigation="radio_channels">Radio</a></li>
<li class="active"><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
</ul>
</div>

<form accept-charset="UTF-8" action="http://www.npo.nl/zoeken" id="search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='page-search-box'>
<div class='search-box'>
<span class="npo-glyph search"></span>
<input autocomplete="off" class="search-query-box" id="search-query" name="q" placeholder="zoeken binnen npo.nl&hellip;" tabindex="1" type="text" />
</div>
<div class='search-suggestions hide' id='search-suggestions'>
<h2>Snel naar een programma</h2>
<div class='suggestion-box'></div>
<a href="/zoeken" class="all-results" id="all-results" tabindex="-1">Bekijk alle zoekresultaten &raquo</a>
</div>
</div>
</form>


</div>
<div class='visible-for-desktop'>
<div class='page-sub-navigation' data-sub-navigation-category='tv_channels'>
<div class='container'>
<div class='content clearfix'>
<div class='slider slide-left'>◀</div>
<div class='sub-navigation-container'>
<ul>
<li><a href="http://www.npo.nl/live/npo-1" data-image="//www-assets.npo.nl/uploads/tv_channel/263/logo/regular_logo-npo1.png">NPO 1</a></li>
<li><a href="http://www.npo.nl/live/npo-2" data-image="//www-assets.npo.nl/uploads/tv_channel/264/logo/regular_logo-npo2.png">NPO 2</a></li>
<li><a href="http://www.npo.nl/npo3" data-image="//www-assets.npo.nl/uploads/tv_channel/265/logo/regular_npo3-logo.png" data-scorecard="{&quot;name&quot;:&quot;trending.logo&quot;,&quot;prefix&quot;:&quot;npo3.softclick&quot;}">NPO 3</a></li>
<li><a href="http://www.zapp.nl/nu-straks" data-image="//www-assets.npo.nl/assets/logos/logo-npozapp-regular-ea6be75d1d99c3f38e47faff9a3b78ed.png" data-scorecard="{&quot;name&quot;:&quot;clickout.zapp&quot;,&quot;prefix&quot;:&quot;npo.softclick&quot;}" target="_blank">Zapp <span class="npo-glyph arrow-jump-box"></span></a></li>
<li><a href="http://www.zappelin.nl/tv-kijken" data-image="//www-assets.npo.nl/assets/logos/logo-npozappelin-regular-7891aaa672e4eb3932469b15e0b7264c.png" data-scorecard="{&quot;name&quot;:&quot;clickout.zappelin&quot;,&quot;prefix&quot;:&quot;npo.softclick&quot;}" target="_blank">Zappelin <span class="npo-glyph arrow-jump-box"></span></a></li>
<li><a href="http://www.npo.nl/live/npo-nieuws" data-image="//www-assets.npo.nl/uploads/tv_channel/279/logo/regular_logonieuws.png">NPO Nieuws</a></li>
<li><a href="http://www.npo.nl/live/npo-cultura" data-image="//www-assets.npo.nl/uploads/tv_channel/280/logo/regular_logocultura.png">NPO Cultura</a></li>
<li><a href="http://www.npo.nl/live/npo-101" data-image="//www-assets.npo.nl/uploads/tv_channel/281/logo/regular_logo-101.png">NPO 101</a></li>
<li><a href="http://www.npo.nl/live/npo-politiek" data-image="//www-assets.npo.nl/uploads/tv_channel/282/logo/regular_logopolitiek.png">NPO Politiek</a></li>
<li><a href="http://www.npo.nl/live/npo-best" data-image="//www-assets.npo.nl/uploads/tv_channel/283/logo/regular_logobest.png">NPO Best</a></li>
<li><a href="http://www.npo.nl/live/npo-doc" data-image="//www-assets.npo.nl/uploads/tv_channel/284/logo/regular_logodoc.png">NPO Doc</a></li>
<li><a href="http://www.npo.nl/live/npo-zappxtra" data-image="//www-assets.npo.nl/uploads/tv_channel/288/logo/regular_logozappxtra.png">NPO Zapp Xtra</a></li>
<li><a href="http://www.npo.nl/live/npo-humor-tv" data-image="//www-assets.npo.nl/uploads/tv_channel/290/logo/regular_logohumor.png">NPO Humor TV</a></li>
</ul>
</div>
<div class='slider slide-right'>►</div>
</div>
</div>
</div>

</div>
<div class='main-navigation visible-for-small-screen'>
<div class='container'>
<a href="#" class="menu-button" id="menu-button"><span class="npo-glyph list contrast"></span>
Menu
</a><div class='border-wrapper left'>
<a href="http://www.npo.nl/" class="logo"><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></a>
</div>
<div class='border-wrapper right'>
<a href="http://www.npo.nl/profiel" class="profile"><span class="npo-glyph user gray"></span>
</a><a href="http://www.npo.nl/zoeken" class="search-button" id="page-search-button"><span class="npo-glyph search gray"></span>
</a></div>
</div>
</div>
<div class='sub-navigation visible-for-small-screen'>
<div class='container'>
<div class='page-navigation'>
<ul>
<li><a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.start&quot;}" data-show-sub-navigation="home">Start</a></li>
<li><a href="http://www.npo.nl/live" data-show-sub-navigation="tv_channels">Tv</a></li>
<li><a href="http://www.npo.nl/radio" data-show-sub-navigation="radio_channels">Radio</a></li>
<li class="active"><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
</ul>
</div>

</div>
</div>
<div class='sub-search visible-for-small-screen'>
<div class='container'>
<form accept-charset="UTF-8" action="http://www.npo.nl/zoeken" id="search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='page-search-box'>
<div class='search-box'>
<span class="npo-glyph search"></span>
<input autocomplete="off" class="search-query-box" id="search-query" name="q" placeholder="zoeken binnen npo.nl&hellip;" tabindex="1" type="text" />
</div>
<div class='search-suggestions hide' id='search-suggestions'>
<h2>Snel naar een programma</h2>
<div class='suggestion-box'></div>
<a href="/zoeken" class="all-results" id="all-results" tabindex="-1">Bekijk alle zoekresultaten &raquo</a>
</div>
</div>
</form>


</div>
</div>
</div>
<div class="schema-org-props hidden-item-props"><div itemscope="" itemtype="http://schema.org/TVSeries"><a href="/zembla/POMS_S_VARA_099718" itemprop="url"><span itemprop="name">ZEMBLA</span></a><span itemprop="description">ZEMBLA is het actuele documentaireprogramma van VARA. De documentaire-reeks brengt opiniërende onderzoeksjournalistiek in actuele documentaires, met spraak- en nieuwsmakende onderwerpen.</span><div itemscope="" itemtype="http://schema.org/TVEpisode"><a href="/zembla/18-03-2007/VARA_101141965" itemprop="url"><span itemprop="name">Het clusterbom gevoel</span></a><span itemprop="description">Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.</span><div itemscope="" itemtype="http://schema.org/BroadcastEvent"><span content="2007-03-18T20:35:00+01:00" itemprop="startDate">2007-03-18 20:35:00 +0100</span></div></div></div></div>
<div class='share-modal modal hide fade' data-id='VARA_101141965' id='share-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2 class='omroep-bold'>ZEMBLA</h2>
<h3 class='omroep-light'>Zembla</h3>
</div>
<div class='modal-body'><div class='row-fluid'>
<div class='span4 social-buttons'>
<h3>Deel via social media</h3>
<a href="/zembla/18-03-2007/VARA_101141965/delen/facebook?url=http%3A%2F%2Fwww.npo.nl%2Fzembla%2F18-03-2007%2FVARA_101141965" class="button facebook" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.facebook.zembla.VARA_101141965&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph facebook"></span>Facebook</a>
<a href="/zembla/18-03-2007/VARA_101141965/delen/google?url=http%3A%2F%2Fwww.npo.nl%2Fzembla%2F18-03-2007%2FVARA_101141965" class="button google" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.google.zembla.VARA_101141965&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph google"></span>Google</a>
<a href="/zembla/18-03-2007/VARA_101141965/delen/twitter?url=http%3A%2F%2Fwww.npo.nl%2Fzembla%2F18-03-2007%2FVARA_101141965" class="button twitter" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.twitter.zembla.VARA_101141965&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph twitter"></span>Twitter</a>
<a href="whatsapp://send?text=Bekijk%20ZEMBLA:%20http://npo.nl/VARA_101141965" class="button whatsapp hide" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.whatsapp.zembla.VARA_101141965&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph whatsapp"></span>Whatsapp</a>
</div>
<div class='span8'>
<div class='row-fluid switch'>
<div class='span6 switch-block switch-left'>
<a href="#" class="switch-link dark-well" data-block="link"><span class="npo-glyph globe"></span>
Link
</a></div>
<div class='span6 switch-block switch-right'>
<a href="#" class="switch-email dark-well inactive" data-block="email"><span class="npo-glyph email"></span>
E-mail vriend
</a></div>
</div>
<div class='link-block switch-list'>
<div class='row-fluid link-inputs'>
<div class='span12'>
<div class='time-inputs'>
<label>Starttijd</label>
<div class='share-action-wrapper'>
<a href="#" class="share-start-link">Start vanaf huidige positie</a>
</div>
<input id="share_start_position" name="share_start_position" placeholder="00:00:00" type="text" value="00:00:00" />
</div>
<div class='share-path'>
<label>Directe link</label>
<div class='share-action-wrapper' id='share-modal-url-copy-wrapper'>
<div class='share-action-link' id='share-modal-url-copy'>Kopieer link</div>
</div>
<input id="share_url" name="share_url" type="text" value="http://www.npo.nl/zembla/18-03-2007/VARA_101141965" />
</div>
</div>
</div>
</div>
<div class='email-friend email-block switch-list'>
<h3>E-mail vriend</h3>
<form accept-charset="UTF-8" action="/shares" class="simple_form new_share" id="new_share_" method="post" novalidate="novalidate"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="PF9x3QovU2Mj1CeLgKDLRm07T8PsCYfIHL3UY+c18Vo=" /></div><div class='row-fluid'>
<div class='span5'>
<input id="id" name="id" type="hidden" value="VARA_101141965" />
<input id="type" name="type" type="hidden" value="program" />
<input id="mid" name="mid" type="hidden" value="VARA_101141965" />
<div class="input string required share_from"><label class="string required control-label" for="share_from">* Je naam</label><input class="string required" id="share_from" name="share[from]" size="50" type="text" /></div>
</div>
<div class='span7'>
<div class="input email required share_from_email"><label class="email required control-label" for="share_from_email">* Je e-mailadres</label><input class="string email required" id="share_from_email" name="share[from_email]" placeholder="naam@voorbeeld.nl" size="50" type="email" /></div>
</div>
</div>
<div class='row-fluid'>
<div class='span5'>
<div class="input string required share_to"><label class="string required control-label" for="share_to">* Naam vriend</label><input class="string required" id="share_to" name="share[to]" size="50" type="text" /></div>
</div>
<div class='span7'>
<div class="input email required share_to_email"><label class="email required control-label" for="share_to_email">* E-mailadres vriend</label><input class="string email required" id="share_to_email" name="share[to_email]" placeholder="naam@voorbeeld.nl" size="50" type="email" /></div>
</div>
</div>
<div class='row-fluid'>
<div class='span12'>
<div class="input text optional share_message"><label class="text optional control-label" for="share_message">Bericht (optioneel)</label><textarea class="text optional" cols="40" id="share_message" name="share[message]" rows="5">
</textarea></div>
</div>
</div>
<div class='row-fluid actions'>
<div class='span12'>
<input class="button right" name="commit" type="submit" value="Versturen" />
<a href="/zembla/18-03-2007/VARA_101141965" class="button gray right" data-dismiss="modal">Annuleren</a>
</div>
</div>
</form>


</div>
</div>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='adaptive-modal modal hide fade' id='adaptive-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2>Instellingen</h2>
<h3>Videokwaliteit in het buitenland</h3>
</div>
<div class='modal-body'>
<p>De nieuwe player kan problemen veroorzaken met de connectie in het buitenland. Ervaart u problemen met npo.nl, zet onderstaande switch dan aan. De videospeler gaat bufferen en na enkele momenten kunt u de uitzending weer bekijken.</p>
<div class='switch-wrapper'>
Buitenland
<i class="npo-icon-radio-switch off adaptive-video-switch"></i>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='program-show-page'>
<div class="showcase showcase-fluid with-player program-player" id="header"><div class="showcase-background" style="background-image: url(//www-assets.npo.nl/uploads/media_item/media_item/14/82/Zembla-_cover_blurred-1440149101.jpg )"><div class="header-gradient"><div class='info no-overlay-for-small-screen'>
<div class='container'>
<div class='row-fluid title-and-broadcasters'>
<div class='broadcaster-logo'><a href="http://www.vara.nl" data-scorecard="{&quot;name&quot;:&quot;vara&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas.extern&quot;}" target="_blank" title="VARA"><img alt="Logo van VARA" class="no-shadow" src="//www-assets.npo.nl/uploads/broadcaster/39/logo/regular_VARA.png" /></a></div>
<h1 title="ZEMBLA"><a href="/zembla/POMS_S_VARA_099718">ZEMBLA</a></h1>
</div>
<div class='row-fluid videoplayer-and-metadata'>
<div class='span-main player-span'>
<div class="video-player-container" data-prid="VARA_101141965" data-video-player="true" id="video-player-container" itemscope="" itemtype="http://schema.org/VideoObject"><meta content="Zembla" itemprop="name" /><meta content="Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen." itemprop="description" /><meta content="http://images.poms.omroep.nl/image/s174/c174x98/6848.png" itemprop="thumbnailUrl" /><meta content="PT50M0S" itemprop="duration" /><div class="video-player-holder" id="video-player"></div></div>
<div class="video-player-settings"><div class="duration-and-views"><span class="duration"><span class="npo-glyph camera"></span> 50:00</span></div><a href="#" class="toggle-player-size" data-scorecard="{&quot;name&quot;:&quot;beeld-vergroten&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas&quot;}" data-toggle-player-size="" title="Beeld vergroten / verkleinen"><span class="enlarge"><span class="npo-glyph expand"></span></span><span class="shrink"><span class="npo-glyph shrink"></span></span></a><a href="#" data-scorecard="{&quot;name&quot;:&quot;buitenland&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas&quot;}" data-toggle-player-settings="" title="Instellingen"><span class="npo-glyph cog"></span></a></div>
</div>
<div class='span-sidebar'>
<h2>Het clusterbom gevoel</h2>
<div class='metadata'>
<div class='content'>
<p class='overflow-description' title='Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.'><span class="hide-original">Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de ge...</span><span class="omission" data-scorecard="{&quot;name&quot;:&quot;programmas.zembla.afleveringen.VARA_101141965.meer-beschrijving&quot;}"><a href="#" class="overflow-toggle">Meer</a></span><span class="omission-overflow hide"> Nederlandse pensioenfondsen beleggen in wapenbedrijven die clusterbommen en landmijnen produceren. Ook investeren ze in bedrijven die het milieu ernstige schade toebrengen en in producten die worden gemaakt met behulp van kinderarbeid. Dat blijkt uit onderzoek van het televisieprogramma ZEMBLA en is te zien in de aflevering ‘Het clusterbom gevoel’. ZEMBLA onderzocht onder andere de geldstromen van de pensioenfondsen ABP, PGGM en het Spoorwegpensioenfonds die de controversiële investeringen bevestigen.<a href="#" class="overflow-toggle less">Minder</a></span></p>

<div class='dark-well meta-data-box'>
<a href="http://zembla.incontxt.nl/" class="external-link" data-scorecard="{&quot;name&quot;:&quot;programmas.VARA_101141965.extern.zembla&quot;}" target="_blank">Site
<span class="npo-glyph arrow-jump"></span>
</a><div class='data'>
<div class='broadcast-time'>
Di 24 apr 2007 09:40
</div>
<div class='genres'>
Documentaire
</div>
<div class='ratings'>
</div>
</div>
</div>
</div>
</div>

</div>
<div class='actions'>
<div class='action-buttons for-guest'>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fzembla%2F18-03-2007%2FVARA_101141965" class="button favorite-button guest-favorite-button"><span class="npo-glyph heart contrast"></span> <span class="glyph-label">Favoriet</span></a>
<a href="/zembla/18-03-2007/VARA_101141965/delen" class="last button share-modal-button" data-id="VARA_101141965" rel="nofollow"><span class="npo-glyph share contrast"></span> <span class="glyph-label">Delen</span></a>
</div>

</div>
</div>
<div class='notifications'>

</div>

<div class='row-fluid program-fragments'>
</div>
</div>
</div>
</div></div></div>
<div class='content' data-auto-load data-url='/zembla/18-03-2007/VARA_101141965/show_content'></div>
</div>

<div class="page-footer" id="npo-footer"><div class='container'>
<div class='broadcasters-and-corporate-links'>
<div class='broadcasters'>
<strong>Omroepen</strong>
<ul>
<li><a href="http://www.avrotros.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.avrotros&quot;}" target="_blank">AVROTROS</a></li>
<li><a href="http://www.bnn.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.bnn&quot;}" target="_blank">BNN</a></li>
<li><a href="http://www.eo.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.eo&quot;}" target="_blank">EO</a></li>
<li><a href="http://www.omroephuman.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.human&quot;}" target="_blank">HUMAN</a></li>
<li><a href="http://www.kro-ncrv.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.kro-ncrv&quot;}" target="_blank">KRO-NCRV</a></li>
</ul>
<ul>
<li><a href="http://www.omroepmax.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.max&quot;}" target="_blank">MAX</a></li>
<li><a href="http://www.nos.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.nos&quot;}" target="_blank">NOS</a></li>
<li><a href="http://www.ntr.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.ntr&quot;}" target="_blank">NTR</a></li>
<li><a href="http://www.powned.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.powned&quot;}" target="_blank">PowNed</a></li>
<li><a href="http://www.vara.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.vara&quot;}" target="_blank">VARA</a></li>
</ul>
<ul>
<li><a href="http://www.vpro.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.vpro&quot;}" target="_blank">VPRO</a></li>
<li><a href="http://www.omroepwnl.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.wnl&quot;}" target="_blank">WNL</a></li>
<li><a href="http://www.zvk.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.zvk&quot;}" target="_blank">ZvK</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Meer NPO</strong>
<ul>
<li><a href="http://www.npoplus.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.npo-plus&quot;}" target="_blank">NPO Plus</a></li>
<li><a href="http://www.bvn.tv" data-scorecard="{&quot;name&quot;:&quot;footer.extern.bvn&quot;}" target="_blank">BVN</a></li>
<li><a href="http://www.npo.nl/specials/regionale-omroepen">Regionaal</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Organisatie</strong>
<ul>
<li><a href="http://www.npo.nl/overnpo">Over NPO</a></li>
<li><a href="http://www.npo.nl/overnpo/pers">Pers</a></li>
<li><a href="http://www.npo.nl/overnpo/vacatures-en-stages-npo">Vacatures</a></li>
<li><a href="http://www.nposales.com" data-scorecard="{&quot;name&quot;:&quot;footer.extern.npo-sales&quot;}" target="_blank">NPO Sales</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Volg de NPO</strong>
<ul>
<li><a href="http://omroep.dmd.omroep.nl/x/plugin/?pName=subscribe&amp;MIDRID=S7Y1swQAA98&amp;pLang=nl&amp;Z=543417650" data-scorecard="{&quot;name&quot;:&quot;footer.extern.aanmelden-nieuwsbrief&quot;}" target="_blank">Aanmelden nieuwsbrief</a></li>
<li><a href="https://instagram.com/npo.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.instagram&quot;}" target="_blank">Instagram</a></li>
<li><a href="https://www.facebook.com/NPO.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.facebook&quot;}" target="_blank">Facebook</a></li>
<li><a href="http://www.twitter.com/publiekeomroep" data-scorecard="{&quot;name&quot;:&quot;footer.extern.twitter&quot;}" target="_blank">Twitter</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Onze apps</strong>
<ul>
<li><a href="https://itunes.apple.com/nl/app/uitzending-gemist/id323998316?mt=8" data-scorecard="{&quot;name&quot;:&quot;footer.extern.ipad-iphone&quot;}" target="_blank">iPad &amp; iPhone</a></li>
<li><a href="https://play.google.com/store/apps/details?id=nl.uitzendinggemist" data-scorecard="{&quot;name&quot;:&quot;footer.extern.android&quot;}" target="_blank">Android</a></li>
<li><a href="https://help.npo.nl/faqs/smarttv-hbbtv">Smarttv en HbbTV</a></li>
</ul>
</div>
</div>
<div class='general-links'>
<ul class='disclaimers'>
<li><a href="http://www.npo.nl/contact">Contact</a></li>
<li><a href="http://help.npo.nl
" data-scorecard="{&quot;name&quot;:&quot;footer.extern.help&quot;}" target="_blank">Help</a></li>
<li><a href="http://www.npo.nl/uitgelicht.rss">RSS</a></li>
<li><a href="http://www.npo.nl/algemene-voorwaarden-privacy">Algemene Voorwaarden &amp; Privacy</a></li>
<li><a href="http://cookiesv2.publiekeomroep.nl/" data-scorecard="{&quot;name&quot;:&quot;footer.extern.cookiebeleid&quot;}" target="_blank">Cookiebeleid</a></li>
</ul>
<div class='logo small'><a href="http://www.npo.nl/"><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></a></div>
</div>
</div>
</div>
</div>
<script class="npo_cc_inline_analytics" language="JavaScript1.3" src="//b.scorecardresearch.com/c2/17827132/cs.js" type="text/plain"></script>

<script type="text/plain" class="npo_cc_inline_advertising">(function() {
  var _fbq = window._fbq || (window._fbq = []);
  if (!_fbq.loaded) {
    var fbds = document.createElement('script');
    fbds.async = true;
    fbds.src = '//connect.facebook.net/en_US/fbds.js';
    var s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(fbds, s);
    _fbq.loaded = true;
  }
  _fbq.push(['addPixelId', '1487544264866945']);
})();
window._fbq = window._fbq || [];
window._fbq.push(['track', 'PixelInitialized', {}]);
</script>
<noscript><img height="1" width="1" alt="" style="display:none" class="npo_cc_inline_advertising" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAA AICTAEAOw==" data-src="https://www.facebook.com/tr?id=1487544264866945&amp;ev=PixelInitialized" /></noscript>
</body>
</html>
`

var testDataBroadcastVideoNPO3 = `<!DOCTYPE html>
<html>
<head>
<title>
Radio Bergeijk, toewijding in beeld kijk je op npo.nl
</title>
<script type="text/javascript">
//<![CDATA[
if (window != top) { top.location.href='http://www.npo.nl'; window.stop(); }
//]]>
</script>
<meta charset='UTF-8'>
<meta content='npo.nl' name='og:site_name'>
<meta content='//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb-fc021b141c9ccde825d868cd10a14ab6.png' name='image-placeholder'>
<meta content='IE=EmulateIE9' http-equiv='X-UA-Compatible'>
<meta content='width=device-width, initial-scale=1, maximum-scale=1' name='viewport'>
<meta content='http://www.npo.nl/suggesties' name='search-suggestions-url'>
<meta content='app-id=nl.uitzendinggemist' name='google-play-app'>
<meta content='//www-assets.npo.nl/assets/google_play_logo-c0f9cdc10f5633cc362e5b8113a3350e.png' name='google-play-logo'>
<meta content='16308a66-14c6-408b-bc68-2aebdb030636' name='msApplication-ID'>
<meta content='7200' name='utc-offset'>
<meta content='30' name='player-tick-length'>
<link href="//cookiesv2.publiekeomroep.nl/static/css/npo_cc_styles.css" media="screen" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/vendors-724ba1e7641e3d78aef2fcd9e1354f92.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/layout-d2942bb06b5bfc5e3869be130af53bc4.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/components-6c045005ff3e4f1125dd5e332d1ede15.css" media="all" rel="stylesheet" type="text/css" />
<link href="//www-assets.npo.nl/assets/application-ab6108433cf7b3dc4912c0aa2cd5c80f.css" media="all" rel="stylesheet" type="text/css" />
<!--[if lte IE 8]>
<link href="//www-assets.npo.nl/assets/ie-52a7ee466e06b2377f85de26a5c4efa4.css" media="all" rel="stylesheet" type="text/css" />
<![endif]-->
<!--[if lte IE 7]>
<link href="//www-assets.npo.nl/assets/ie7-87b61b941651ecf598c4db2055c80f41.css" media="all" rel="stylesheet" type="text/css" />
<![endif]-->
<script src="//www-assets.npo.nl/assets/application-64c4dbaacb4e6e1077795f2890897042.js" type="text/javascript"></script>
<script src="//cookiesv2.publiekeomroep.nl/data/script/cconsent.min.js" type="text/javascript"></script>
<link href="http://www.npo.nl/uitgelicht.rss" rel="alternate" title="Uitgelicht op npo.nl" type="application/rss+xml" />
<link href="http://www.npo.nl/artikelen.rss" rel="alternate" title="Artikelen op npo.nl" type="application/rss+xml" />
<meta content='{"c1":"2","c2":"17827132","ns_site":"po-totaal","potag1":"npoportal","potag3":"npo","potag4":"npo","potag5":"zenderportal","potag6":"video","potag7":"npo3","potag8":"site","potag9":"site","potag2":"npo3"}' name='scorecard-default-labels'>
<meta content='npo3.softclick' name='scorecard-default-prefix'>
<meta content='Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met technicus Ted van Lieshout. Serie verslagen van hun levenswerk: Radio Bergeijk. Een filmploeg over de radiovloer! Ze moesten even aan de gedachte wennen, de presentatoren van Radio Bergeijk. En het heeft even geduurd voordat Toon Spoorenberg en Peer van Eersel konden instemmen met die druktemakerij in hun zo vertrouwde en overzichtelijke radiobestaan. Maar het is er dan toch van gekomen. Pieter Verhoeff heeft een serie gemaakt van Radio Bergeijk, het ook landelijk bekende radiostation voor Bergeijk, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Ted van Lieshout. De pro Deo werkende radiomannen ontvangen in hun studio een keur aan plaatselijke gasten of doen rechtstreeks verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. De televisieserie wordt niet alleen een fel realistisch portret van de makers van het vermaarde radiostation, maar via hen ook van het dorp Bergeijk. Radio Bergeijk is en probeert een satire te zijn op een lokaal radiostation. In de verbeelding van de makers is het fictieve dorp Bergeijk een huiveringwekkend oord waar hedendaagse xenofobie, orthodox atheïsme, vrouwonvriendelijkheid en homo-afkeer om de voorrang strijden. Het radiostation probeert van deze benepen dorpse identiteit pseudo-professioneel verslag te doen. En hanteert als stijlmiddel de grove overdrijving. Teksten en spel: Pieter Bouwman en George van Houts. Gastacteurs o.a.: Alex Klaasen, Jeroen van Merwijk, Marjan Luif, Annet Malherbe, Anniek Pheifer, Margôt Ros en Lies Visschedijk.' name='description'>
<script id='signup-popover' type='text/x-jquery-tmpl'><div class='signup-popover'>
<h3 class='omroep-bold'>Wil je dit programma bewaren als favoriet?</h3>
<p>Maak een eigen account aan en bewaar je favoriete programma’s. Een account zorgt ook voor tips op maat.</p>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk-toewijding-in-beeld%2F25-06-2007%2FVPRO_1122739" class="button action-button">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken" class="button action-button">Account aanmaken</a>
</div>
</script>
<meta content='true' name='guest'>



<script type="text/javascript">
//<![CDATA[
npo_cookies.init();
//]]>
</script>
<script src="http://npoplayer.omroep.nl/csjs/npoplayer-min.js" type="text/javascript"></script>
<script src="//www-assets.npo.nl/assets/player-event-support-9bb142d8358270024970c91b7dd43104.js" type="text/javascript"></script>
<script defer="defer" src="http://topspin.npo.nl/divolte/divolte.js" type="text/javascript"></script>

<meta content="authenticity_token" name="csrf-param" />
<meta content="PF9x3QovU2Mj1CeLgKDLRm07T8PsCYfIHL3UY+c18Vo=" name="csrf-token" />
<meta content="Radio Bergeijk, toewijding in beeld" name="og:title" />
<meta content="http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739" name="og:url" />
<meta content="Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met technicus Ted van Lieshout." name="og:description" />
<meta content="video.episode" name="og:type" />
<meta content="http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739" name="og:video" />
<meta content="2100" name="og:video:duration" />
<meta content="2007-06-25 23:10:00 +0200" name="og:video:release_date" />
</head>
<body>
<!-- Begin comScore Inline Tag 1.1302.13 -->
<script type="text/plain" class="npo_cc_inline_analytics">
// <![CDATA[
function udm_(e){var t="comScore=",n=document,r=n.cookie,i="",s="indexOf",o="substring",u="length",a=2048,f,l="&ns_",c="&",h,p,d,v,m=window,g=m.encodeURIComponent||escape;if(r[s](t)+1)for(d=0,p=r.split(";"),v=p[u];d<v;d++)h=p[d][s](t),h+1&&(i=c+unescape(p[d][o](h+t[u])));e+=l+"_t="+ +(new Date)+l+"c="+(n.characterSet||n.defaultCharset||"")+"&c8="+g(n.title)+i+"&c7="+g(n.URL)+"&c9="+g(n.referrer),e[u]>a&&e[s](c)>0&&(f=e[o](0,a-8).lastIndexOf(c),e=(e[o](0,f)+l+"cut="+g(e[o](f+1)))[o](0,a)),n.images?(h=new Image,m.ns_p||(ns_p=h),h.src=e):n.write("<","p","><",'img src="',e,'" height="1" width="1" alt="*"',"><","/p",">")};
udm_('http'+(document.location.href.charAt(4)=='s'?'s://sb':'://b')+'.scorecardresearch.com/b?c1=2&c2=17827132&name=npo3.programmas.radio-bergeijk-toewijding-in-beeld.afleveringen.VPRO_1122739&npo_ingelogd=no&npo_login_id=geen&ns_site=po-totaal&potag1=npoportal&potag2=npo3&potag3=npo&potag4=npo&potag5=zenderportal&potag6=video&potag7=npo3&potag8=site&potag9=site');
// ]]>
</script>
<noscript><p><img src="http://b.scorecardresearch.com/p?c1=2&amp;c2=17827132&amp;name=npo3.programmas.radio-bergeijk-toewijding-in-beeld.afleveringen.VPRO_1122739&amp;npo_ingelogd=no&amp;npo_login_id=geen&amp;ns_site=po-totaal&amp;potag1=npoportal&amp;potag2=npo3&amp;potag3=npo&amp;potag4=npo&amp;potag5=zenderportal&amp;potag6=video&amp;potag7=npo3&amp;potag8=site&amp;potag9=site" height="1" width="1" alt="*"></p></noscript>
<!-- End comScore Inline Tag -->


<div class="hidden-item-props"><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/" itemprop="url"><span itemprop="title">NPO</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994" itemprop="url"><span itemprop="title">Radio Bergeijk, toewijding in beeld</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739" itemprop="url"><span itemprop="title">Radio Bergeijk, toewijding in beeld</span></a></div></div>
<div class='popup-overlay hide' id='popup-overlay'></div>
<div class='popup-alerts' id='popup-alerts'>
</div>
<div class='npo3 page-container programs'>
<div class='mobile-trigger visible-with-grid-collapse'></div>
<div class='page-header sub-navigation-collapse' id='page-header'>
<div class='container visible-for-desktop'>
<a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.logo&quot;}"><div class='logo'><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></div>
</a><div class='profile-box signed-out-box'>
<a href="http://www.npoplus.nl" data-scorecard="{&quot;name&quot;:&quot;home.npo-plus&quot;}">NPO Plus</a>
<span class='signed-out'>
<a href="https://mijn.npo.nl/inloggen">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken">Account aanmaken</a>
</span>
</div>
<div class='page-navigation'>
<ul>
<li><a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.start&quot;}" data-show-sub-navigation="home">Start</a></li>
<li><a href="http://www.npo.nl/live" data-show-sub-navigation="tv_channels">Tv</a></li>
<li><a href="http://www.npo.nl/radio" data-show-sub-navigation="radio_channels">Radio</a></li>
<li class="active"><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
</ul>
</div>

<form accept-charset="UTF-8" action="http://www.npo.nl/zoeken" id="search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='page-search-box'>
<div class='search-box'>
<span class="npo-glyph search"></span>
<input autocomplete="off" class="search-query-box" id="search-query" name="q" placeholder="zoeken binnen npo.nl&hellip;" tabindex="1" type="text" />
</div>
<div class='search-suggestions hide' id='search-suggestions'>
<h2>Snel naar een programma</h2>
<div class='suggestion-box'></div>
<a href="/zoeken" class="all-results" id="all-results" tabindex="-1">Bekijk alle zoekresultaten &raquo</a>
</div>
</div>
</form>


</div>
<div class='visible-for-desktop'>
<div class='page-sub-navigation' data-sub-navigation-category='tv_channels' data-sub-navigation-default='true'>
<div class='container'>
<div class='content clearfix'>
<div class='slider slide-left'>◀</div>
<div class='sub-navigation-container'>
<ul>
<li><a href="http://www.npo.nl/live/npo-1" data-image="//www-assets.npo.nl/uploads/tv_channel/263/logo/regular_logo-npo1.png">NPO 1</a></li>
<li><a href="http://www.npo.nl/live/npo-2" data-image="//www-assets.npo.nl/uploads/tv_channel/264/logo/regular_logo-npo2.png">NPO 2</a></li>
<li><a href="http://www.npo.nl/npo3" data-image="//www-assets.npo.nl/uploads/tv_channel/265/logo/regular_npo3-logo.png" data-scorecard="{&quot;name&quot;:&quot;trending.logo&quot;,&quot;prefix&quot;:&quot;npo3.softclick&quot;}">NPO 3</a></li>
<li><a href="http://www.zapp.nl/nu-straks" data-image="//www-assets.npo.nl/assets/logos/logo-npozapp-regular-ea6be75d1d99c3f38e47faff9a3b78ed.png" data-scorecard="{&quot;name&quot;:&quot;clickout.zapp&quot;,&quot;prefix&quot;:&quot;npo.softclick&quot;}" target="_blank">Zapp <span class="npo-glyph arrow-jump-box"></span></a></li>
<li><a href="http://www.zappelin.nl/tv-kijken" data-image="//www-assets.npo.nl/assets/logos/logo-npozappelin-regular-7891aaa672e4eb3932469b15e0b7264c.png" data-scorecard="{&quot;name&quot;:&quot;clickout.zappelin&quot;,&quot;prefix&quot;:&quot;npo.softclick&quot;}" target="_blank">Zappelin <span class="npo-glyph arrow-jump-box"></span></a></li>
<li><a href="http://www.npo.nl/live/npo-nieuws" data-image="//www-assets.npo.nl/uploads/tv_channel/279/logo/regular_logonieuws.png">NPO Nieuws</a></li>
<li><a href="http://www.npo.nl/live/npo-cultura" data-image="//www-assets.npo.nl/uploads/tv_channel/280/logo/regular_logocultura.png">NPO Cultura</a></li>
<li><a href="http://www.npo.nl/live/npo-101" data-image="//www-assets.npo.nl/uploads/tv_channel/281/logo/regular_logo-101.png">NPO 101</a></li>
<li><a href="http://www.npo.nl/live/npo-politiek" data-image="//www-assets.npo.nl/uploads/tv_channel/282/logo/regular_logopolitiek.png">NPO Politiek</a></li>
<li><a href="http://www.npo.nl/live/npo-best" data-image="//www-assets.npo.nl/uploads/tv_channel/283/logo/regular_logobest.png">NPO Best</a></li>
<li><a href="http://www.npo.nl/live/npo-doc" data-image="//www-assets.npo.nl/uploads/tv_channel/284/logo/regular_logodoc.png">NPO Doc</a></li>
<li><a href="http://www.npo.nl/live/npo-zappxtra" data-image="//www-assets.npo.nl/uploads/tv_channel/288/logo/regular_logozappxtra.png">NPO Zapp Xtra</a></li>
<li><a href="http://www.npo.nl/live/npo-humor-tv" data-image="//www-assets.npo.nl/uploads/tv_channel/290/logo/regular_logohumor.png">NPO Humor TV</a></li>
</ul>
</div>
<div class='slider slide-right'>►</div>
</div>
</div>
</div>

</div>
<div class='main-navigation visible-for-small-screen'>
<div class='container'>
<a href="#" class="menu-button" id="menu-button"><span class="npo-glyph list contrast"></span>
Menu
</a><div class='border-wrapper left'>
<a href="http://www.npo.nl/" class="logo"><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></a>
</div>
<div class='border-wrapper right'>
<a href="http://www.npo.nl/profiel" class="profile"><span class="npo-glyph user gray"></span>
</a><a href="http://www.npo.nl/zoeken" class="search-button" id="page-search-button"><span class="npo-glyph search gray"></span>
</a></div>
</div>
</div>
<div class='sub-navigation visible-for-small-screen'>
<div class='container'>
<div class='page-navigation'>
<ul>
<li><a href="http://www.npo.nl/" data-scorecard="{&quot;name&quot;:&quot;home.start&quot;}" data-show-sub-navigation="home">Start</a></li>
<li><a href="http://www.npo.nl/live" data-show-sub-navigation="tv_channels">Tv</a></li>
<li><a href="http://www.npo.nl/radio" data-show-sub-navigation="radio_channels">Radio</a></li>
<li class="active"><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
</ul>
</div>

</div>
</div>
<div class='sub-search visible-for-small-screen'>
<div class='container'>
<form accept-charset="UTF-8" action="http://www.npo.nl/zoeken" id="search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='page-search-box'>
<div class='search-box'>
<span class="npo-glyph search"></span>
<input autocomplete="off" class="search-query-box" id="search-query" name="q" placeholder="zoeken binnen npo.nl&hellip;" tabindex="1" type="text" />
</div>
<div class='search-suggestions hide' id='search-suggestions'>
<h2>Snel naar een programma</h2>
<div class='suggestion-box'></div>
<a href="/zoeken" class="all-results" id="all-results" tabindex="-1">Bekijk alle zoekresultaten &raquo</a>
</div>
</div>
</form>


</div>
</div>
</div>
<div class="schema-org-props hidden-item-props"><div itemscope="" itemtype="http://schema.org/TVSeries"><a href="/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994" itemprop="url"><span itemprop="name">Radio Bergeijk, toewijding in beeld</span></a><span itemprop="description">Serie verslagen van hun levenswerk: Radio Bergeijk. Een filmploeg over de radiovloer! Ze moesten even aan de gedachte wennen, de presentatoren van Radio Bergeijk. En het heeft even geduurd voordat Toon Spoorenberg en Peer van Eersel konden instemmen met die druktemakerij in hun zo vertrouwde en overzichtelijke radiobestaan. Maar het is er dan toch van gekomen. Pieter Verhoeff heeft een serie gemaakt van Radio Bergeijk, het ook landelijk bekende radiostation voor Bergeijk, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Ted van Lieshout. De pro Deo werkende radiomannen ontvangen in hun studio een keur aan plaatselijke gasten of doen rechtstreeks verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. De televisieserie wordt niet alleen een fel realistisch portret van de makers van het vermaarde radiostation, maar via hen ook van het dorp Bergeijk. Radio Bergeijk is en probeert een satire te zijn op een lokaal radiostation. In de verbeelding van de makers is het fictieve dorp Bergeijk een huiveringwekkend oord waar hedendaagse xenofobie, orthodox atheïsme, vrouwonvriendelijkheid en homo-afkeer om de voorrang strijden. Het radiostation probeert van deze benepen dorpse identiteit pseudo-professioneel verslag te doen. En hanteert als stijlmiddel de grove overdrijving. Teksten en spel: Pieter Bouwman en George van Houts. Gastacteurs o.a.: Alex Klaasen, Jeroen van Merwijk, Marjan Luif, Annet Malherbe, Anniek Pheifer, Margôt Ros en Lies Visschedijk. Samenstelling en regie: Pieter Verhoeff.</span><div itemscope="" itemtype="http://schema.org/TVEpisode"><a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739" itemprop="url"><span itemprop="name">Radio Bergeijk, toewijding in beeld</span></a><span itemprop="description">Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met technicus Ted van Lieshout.</span><div itemscope="" itemtype="http://schema.org/BroadcastEvent"><span content="2007-06-25T23:10:00+02:00" itemprop="startDate">2007-06-25 23:10:00 +0200</span></div></div></div></div>
<div class='share-modal modal hide fade' data-id='VPRO_1122739' id='share-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2 class='omroep-bold'>Radio Bergeijk, toewijding in beeld</h2>
</div>
<div class='modal-body'><div class='row-fluid'>
<div class='span4 social-buttons'>
<h3>Deel via social media</h3>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739/delen/facebook?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk-toewijding-in-beeld%2F25-06-2007%2FVPRO_1122739" class="button facebook" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.facebook.radio-bergeijk-toewijding-in-beeld.VPRO_1122739&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph facebook"></span>Facebook</a>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739/delen/google?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk-toewijding-in-beeld%2F25-06-2007%2FVPRO_1122739" class="button google" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.google.radio-bergeijk-toewijding-in-beeld.VPRO_1122739&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph google"></span>Google</a>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739/delen/twitter?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk-toewijding-in-beeld%2F25-06-2007%2FVPRO_1122739" class="button twitter" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.twitter.radio-bergeijk-toewijding-in-beeld.VPRO_1122739&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph twitter"></span>Twitter</a>
<a href="whatsapp://send?text=Bekijk%20Radio%20Bergeijk,%20toewijding%20in%20beeld:%20http://npo.nl/VPRO_1122739" class="button whatsapp hide" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.whatsapp.radio-bergeijk-toewijding-in-beeld.VPRO_1122739&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph whatsapp"></span>Whatsapp</a>
</div>
<div class='span8'>
<div class='row-fluid switch'>
<div class='span6 switch-block switch-left'>
<a href="#" class="switch-link dark-well" data-block="link"><span class="npo-glyph globe"></span>
Link
</a></div>
<div class='span6 switch-block switch-right'>
<a href="#" class="switch-email dark-well inactive" data-block="email"><span class="npo-glyph email"></span>
E-mail vriend
</a></div>
</div>
<div class='link-block switch-list'>
<div class='row-fluid link-inputs'>
<div class='span12'>
<div class='time-inputs'>
<label>Starttijd</label>
<div class='share-action-wrapper'>
<a href="#" class="share-start-link">Start vanaf huidige positie</a>
</div>
<input id="share_start_position" name="share_start_position" placeholder="00:00:00" type="text" value="00:00:00" />
</div>
<div class='share-path'>
<label>Directe link</label>
<div class='share-action-wrapper' id='share-modal-url-copy-wrapper'>
<div class='share-action-link' id='share-modal-url-copy'>Kopieer link</div>
</div>
<input id="share_url" name="share_url" type="text" value="http://www.npo.nl/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739" />
</div>
</div>
</div>
</div>
<div class='email-friend email-block switch-list'>
<h3>E-mail vriend</h3>
<form accept-charset="UTF-8" action="/shares" class="simple_form new_share" id="new_share_" method="post" novalidate="novalidate"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="PF9x3QovU2Mj1CeLgKDLRm07T8PsCYfIHL3UY+c18Vo=" /></div><div class='row-fluid'>
<div class='span5'>
<input id="id" name="id" type="hidden" value="VPRO_1122739" />
<input id="type" name="type" type="hidden" value="program" />
<input id="mid" name="mid" type="hidden" value="VPRO_1122739" />
<div class="input string required share_from"><label class="string required control-label" for="share_from">* Je naam</label><input class="string required" id="share_from" name="share[from]" size="50" type="text" /></div>
</div>
<div class='span7'>
<div class="input email required share_from_email"><label class="email required control-label" for="share_from_email">* Je e-mailadres</label><input class="string email required" id="share_from_email" name="share[from_email]" placeholder="naam@voorbeeld.nl" size="50" type="email" /></div>
</div>
</div>
<div class='row-fluid'>
<div class='span5'>
<div class="input string required share_to"><label class="string required control-label" for="share_to">* Naam vriend</label><input class="string required" id="share_to" name="share[to]" size="50" type="text" /></div>
</div>
<div class='span7'>
<div class="input email required share_to_email"><label class="email required control-label" for="share_to_email">* E-mailadres vriend</label><input class="string email required" id="share_to_email" name="share[to_email]" placeholder="naam@voorbeeld.nl" size="50" type="email" /></div>
</div>
</div>
<div class='row-fluid'>
<div class='span12'>
<div class="input text optional share_message"><label class="text optional control-label" for="share_message">Bericht (optioneel)</label><textarea class="text optional" cols="40" id="share_message" name="share[message]" rows="5">
</textarea></div>
</div>
</div>
<div class='row-fluid actions'>
<div class='span12'>
<input class="button right" name="commit" type="submit" value="Versturen" />
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739" class="button gray right" data-dismiss="modal">Annuleren</a>
</div>
</div>
</form>


</div>
</div>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='adaptive-modal modal hide fade' id='adaptive-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2>Instellingen</h2>
<h3>Videokwaliteit in het buitenland</h3>
</div>
<div class='modal-body'>
<p>De nieuwe player kan problemen veroorzaken met de connectie in het buitenland. Ervaart u problemen met npo.nl, zet onderstaande switch dan aan. De videospeler gaat bufferen en na enkele momenten kunt u de uitzending weer bekijken.</p>
<div class='switch-wrapper'>
Buitenland
<i class="npo-icon-radio-switch off adaptive-video-switch"></i>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='npo3-page program-player program-show-page with-player'>
<div class="showcase showcase-fluid visible-for-desktop"><div class="showcase-background"><div class="showcase-wrapper"></div></div></div><div class='npo3-menu' id='npo3-menu'>
<div class="page-content"><div class="container"><div class='main-menu'>
<ul>
<li class='channel-logo'><img alt="Smaller_npo3-logo" src="//www-assets.npo.nl/uploads/tv_channel/265/logo/smaller_npo3-logo.png" /></li>
<li class='no-submenu'><a href="/npo3">Trending</a></li>
<li class='toggle' data-scorecard='{"name":"menu.programma-s"}' data-target='#npo3-menu-1'>Programma's</li>
<li class='no-submenu'><a href="http://www.npo.nl/npo3/nederlands-film-festival" data-scorecard="{&quot;name&quot;:&quot;menu.nff&quot;}">NFF</a></li>
<li class='toggle' data-scorecard='{"name":"menu.meer"}' data-target='#npo3-menu-2'>Meer</li>
<li class='social'><a href="http://www.facebook.com/npo3" data-scorecard="{&quot;name&quot;:&quot;menu.extern.facebook&quot;}" target="_blank"><span class="npo-glyph facebook2"></span></a></li>
<li class='social'><a href="http://twitter.com/npo3" data-scorecard="{&quot;name&quot;:&quot;menu.extern.twitter&quot;}" target="_blank"><span class="npo-glyph twitter"></span></a></li>
<li class='social'><a href="http://instagram.com/npo_3" data-scorecard="{&quot;name&quot;:&quot;menu.extern.instagram&quot;}" target="_blank"><span class="npo-glyph instagram"></span></a></li>
</ul>
</div>
<div class='sub-menu'>
<div class='collapse-menu hide' id='npo3-menu-1'>
<div class='visible-no-grid-collapse'>
<div class='filters'>
<ul>
<li class='active' data-scorecard='{"name":"menu.programma-s.genre.alles"}'>Alles</li>
<li data-scorecard='{"name":"menu.programma-s.genre.film"}'>Film</li>
<li data-scorecard='{"name":"menu.programma-s.genre.humor"}'>Humor</li>
<li data-scorecard='{"name":"menu.programma-s.genre.muziek"}'>Muziek</li>
<li data-scorecard='{"name":"menu.programma-s.genre.reizen"}'>Reizen</li>
<li data-scorecard='{"name":"menu.programma-s.genre.serie"}'>Serie</li>
<li data-scorecard='{"name":"menu.programma-s.genre.webonly"}'>Webonly</li>
</ul>
</div>
<div class='menu-item-list'>
<ul class="menu-group "><li class="no-logo" data-tags="[&quot;Reizen&quot;]"><a href="/3-op-reis/POMS_S_BNN_097362">3 op Reis</a></li>
<li class="no-logo" data-tags="[&quot;Reizen&quot;]"><a href="/3-op-reis-backpack/POMS_S_BNN_868638">3 op Reis Backpack</a></li>
<li class="no-logo" data-tags="[]"><a href="/3doc/POMS_S_EO_098006">3Doc</a></li>
<li class="no-logo" data-tags="[]"><a href="http://www.npo.nl/npo3/3lab">3Lab</a></li>
<li class="no-logo" data-tags="[]"><a href="/3onderzoekt/POMS_S_EO_420681">3Onderzoekt</a></li>
<li class="no-logo" data-tags="[&quot;Serie&quot;]"><a href="http://www.npo.nl/4jim/POMS_S_KN_1032073">4JIM</a></li>
<li class="no-logo" data-tags="[]"><a href="/aanmodderfakker/POMS_S_AT_2125302">Aanmodderfakker</a></li>
<li class="no-logo" data-tags="[&quot;Film&quot;]"><a href="http://www.npo.nl/aanmodderfakker/28-09-2015/AT_2044870">Aanmodderfakker</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/ali-b-bekend-t/POMS_S_AT_1488792">Ali B Bekend(t)</a></li>
<li class="no-logo" data-tags="[]"><a href="/all-stars-2-old-stars/POMS_S_VARA_1684367">All Stars 2 Old Stars</a></li>
<li class="no-logo" data-tags="[]"><a href="/alles-is-familie/POMS_S_VARA_2124339">Alles is familie</a></li>
<li class="no-logo" data-tags="[&quot;Webonly&quot;]"><a href="http://www.npo.nl/beam-naar-de-balkan/POMS_S_EO_1957580">Beam naar de Balkan</a></li>
<li class="no-logo" data-tags="[]"><a href="/beste-nederlandse-film-simon/POMS_S_VPRO_113385">Beste Nederlandse Film: Simon</a></li>
<li class="no-logo" data-tags="[]"><a href="/het-beste-van-101/POMS_S_BNN_123416">Het Beste van 101</a></li>
<li class="no-logo" data-tags="[&quot;Muziek&quot;]"><a href="/de-beste-zangers-van-nederland/POMS_S_TROS_123317">De Beste Zangers van Nederland</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/bij-dino-in-de-straat/POMS_S_NTR_858513">Bij Dino in de Straat</a></li>
<li class="no-logo" data-tags="[&quot;Film&quot;]"><a href="/bride-flight/POMS_S_NCRV_1299373">Bride Flight</a></li>
<li class="no-logo" data-tags="[]"><a href="/bully/POMS_S_BNN_141165">Bully</a></li>
<li class="no-logo" data-tags="[]"><a href="/bureau-sport/POMS_S_VARA_124275">Bureau Sport</a></li>
<li class="no-logo" data-tags="[]"><a href="/club-27/POMS_S_AT_1673854">Club van 27</a></li>
<li class="no-logo" data-tags="[]"><a href="/college-tour/POMS_S_NTR_124382">College Tour</a></li>
<li class="no-logo" data-tags="[]"><a href="/dag-tegen-pesten/POMS_S_BNN_141164">Dag Tegen Pesten</a></li>
<li class="no-logo" data-tags="[]"><a href="/di-rect-concert/POMS_S_AT_1905717">DI-RECT Concert</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/draadstaal-2015/POMS_S_AT_844275">Draadstaal 2015</a></li>
<li class="no-logo" data-tags="[]"><a href="/eo-jongerendag/POMS_S_EO_097554">EO-Jongerendag</a></li>
<li class="no-logo" data-tags="[&quot;Film&quot;]"><a href="/escort/POMS_S_VARA_382597">Escort</a></li>
<li class="no-logo" data-tags="[]"><a href="/filmlab-voetzoeker/POMS_S_VPRO_2026622">Filmlab: Voetzoeker</a></li>
<li class="no-logo" data-tags="[&quot;Webonly&quot;,&quot;Serie&quot;]"><a href="http://www.npo.nl/heemennekes-en-hellehonden/POMS_S_VPRO_1405182">Heemennekes en Hellehonden</a></li>
<li class="no-logo" data-tags="[]"><a href="/de-heineken-ontvoering/POMS_S_VARA_460439">De Heineken ontvoering</a></li></ul>
<ul class="menu-group "><li class="no-logo" data-tags="[]"><a href="/hemel-op-aarde/POMS_S_AVRO_2064471">Hemel op aarde</a></li>
<li class="no-logo" data-tags="[]"><a href="/hotel-di-rect/POMS_S_TROS_1702109">Hotel DI-RECT</a></li>
<li class="no-logo" data-tags="[]"><a href="/hufterproef/POMS_S_EO_860345">Hufterproef</a></li>
<li class="no-logo" data-tags="[]"><a href="/humortv-op-3/POMS_S_VARA_099941">HumorTV op 3</a></li>
<li class="no-logo" data-tags="[]"><a href="/into-the-great-wide-open-2015/POMS_S_VPRO_1956627">Into The Great Wide Open (2015)</a></li>
<li class="no-logo" data-tags="[]"><a href="/jasper-van-kuijk/POMS_S_VARA_1986396">Jasper van Kuijk</a></li>
<li class="no-logo" data-tags="[]"><a href="/je-zal-het-maar-zijn/POMS_S_BNN_123419">Je Zal Het Maar Zijn</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="http://www.npo.nl/jeuk/POMS_S_VARA_461066">Jeuk</a></li>
<li class="no-logo" data-tags="[]"><a href="/keuringsdienst-van-waarde/POMS_S_KRO_123908">Keuringsdienst van Waarde</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/larie/POMS_S_KRO_808174">Larie</a></li>
<li class="no-logo" data-tags="[]"><a href="/lauren/POMS_S_AT_697765">Lauren!</a></li>
<li class="no-logo" data-tags="[]"><a href="/lek/POMS_S_VARA_382592">Lek</a></li>
<li class="no-logo" data-tags="[]"><a href="/loft/POMS_S_BNN_123423">Loft</a></li>
<li class="no-logo" data-tags="[]"><a href="/mindf-ck/POMS_S_AT_728505">Mindf*ck</a></li>
<li class="no-logo" data-tags="[&quot;Serie&quot;]"><a href="/motief-voor-moord/POMS_S_EO_1249971">Motief voor Moord</a></li>
<li class="no-logo" data-tags="[&quot;Muziek&quot;]"><a href="http://www.npo.nl/bnn-presents-mysteryland/30-08-2014/BNN_101363640">Mysteryland</a></li>
<li class="no-logo" data-tags="[]"><a href="/de-nationale-iq-test/POMS_S_BNN_097257">De Nationale IQ test</a></li>
<li class="no-logo" data-tags="[]"><a href="http://www.npo.nl/nederlands-film-festival-op-3/POMS_S_VPRO_2099142">Nederlands Film Festival op 3</a></li>
<li class="no-logo" data-tags="[]"><a href="/nos-op-3/POMS_S_NOS_059929">NOS op 3</a></li>
<li class="no-logo" data-tags="[&quot;Film&quot;]"><a href="/one-night-stand/POMS_S_VARA_099984">One night stand</a></li>
<li class="no-logo" data-tags="[]"><a href="/penoza/POMS_S_KRO_123909">Penoza</a></li>
<li class="no-logo" data-tags="[]"><a href="/plan-c/POMS_S_AVRO_460376">Plan C</a></li>
<li class="no-logo" data-tags="[]"><a href="/rambam/POMS_S_VARA_124202">Rambam</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/ranking-the-stars/POMS_S_BNN_097277">Ranking the Stars</a></li>
<li class="no-logo" data-tags="[]"><a href="/de-rekenkamer/POMS_S_KRO_123979">De Rekenkamer</a></li>
<li class="no-logo" data-tags="[]"><a href="/rolling-stones-sweet-summer-sun/POMS_S_NTR_449961">Rolling Stones: Sweet Summer Sun</a></li>
<li class="no-logo" data-tags="[]"><a href="/roue-verveer-ff-wat-anders/POMS_S_VARA_1914354">Roué Verveer - FF wat anders</a></li>
<li class="no-logo" data-tags="[]"><a href="/rundfunk/POMS_S_KN_1666829">Rundfunk</a></li>
<li class="no-logo" data-tags="[]"><a href="/sara-kroos/POMS_S_VARA_124118">Sara Kroos</a></li></ul>
<ul class="menu-group "><li class="no-logo" data-tags="[&quot;Muziek&quot;]"><a href="/solar/POMS_S_BNN_097354">Solar</a></li>
<li class="no-logo" data-tags="[]"><a href="http://www.npo.nl/sophie-in-de-kreukels/01-09-2015/BNN_101375894">Sophie in de kreukels</a></li>
<li class="no-logo" data-tags="[]"><a href="/spuiten-en-slikken/POMS_S_BNN_097254">Spuiten en Slikken</a></li>
<li class="no-logo" data-tags="[]"><a href="/spuiten-en-slikken-op-reis/POMS_S_BNN_097352">Spuiten en Slikken op Reis</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/streetlab/POMS_S_KRO_797293">Streetlab</a></li>
<li class="no-logo" data-tags="[]"><a href="/studio-powned/POMS_S_POWN_775803">Studio PowNed</a></li>
<li class="no-logo" data-tags="[&quot;Serie&quot;]"><a href="http://www.npo.nl/super-stream-me/POMS_S_VPRO_1790127">Super Stream Me</a></li>
<li class="no-logo" data-tags="[]"><a href="/tbs/POMS_S_BNN_1613878">TBS</a></li>
<li class="no-logo" data-tags="[]"><a href="/teledoc-voetbalmiljonair-uit-oost/POMS_S_BNN_1914803">Teledoc: Voetbalmiljonair Uit Oost</a></li>
<li class="no-logo" data-tags="[]"><a href="/the-freestyle-games/POMS_S_BNN_1633995">The Freestyle Games</a></li>
<li class="no-logo" data-tags="[]"><a href="/the-undateables/POMS_S_BNN_399127">The Undateables</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/theatershow-jandino/POMS_S_NTR_125623">Theatershow Jandino</a></li>
<li class="no-logo" data-tags="[&quot;Webonly&quot;]"><a href="http://www.npo.nl/de-thuiszorgman/POMS_S_VPRO_1593549">De Thuiszorgman</a></li>
<li class="no-logo" data-tags="[]"><a href="/troy/POMS_S_BNN_1404983">Troy</a></li>
<li class="no-logo" data-tags="[]"><a href="/uit-de-kast/POMS_S_KRO_123939">Uit de Kast</a></li>
<li class="no-logo" data-tags="[&quot;Humor&quot;]"><a href="/uitzonderlijk-vervoer/POMS_S_VARA_788350">Uitzonderlijk Vervoer</a></li>
<li class="no-logo" data-tags="[]"><a href="/van-de-straat/POMS_S_NCRV_121157">Van de straat</a></li>
<li class="no-logo" data-tags="[]"><a href="/van-god-los-de-film/POMS_S_BNN_388012">Van God Los de Film</a></li>
<li class="no-logo" data-tags="[]"><a href="/van-luizenleven-naar-luiers/POMS_S_EO_1592532">Van luizenleven naar luiers</a></li>
<li class="no-logo" data-tags="[&quot;Serie&quot;]"><a href="/het-verhaal-van/POMS_S_TROS_137027">Het verhaal van...</a></li>
<li class="no-logo" data-tags="[&quot;Film&quot;]"><a href="/verliefd-op-ibiza/POMS_S_AT_1340616">Verliefd op Ibiza</a></li>
<li class="no-logo" data-tags="[]"><a href="/villa-single-mama/POMS_S_EO_1673428">Villa Single Mama</a></li>
<li class="no-logo" data-tags="[&quot;Webonly&quot;]"><a href="http://www.npo.nl/vriende/POMS_S_VPRO_2086252">Vriende</a></li>
<li class="no-logo" data-tags="[]"><a href="/de-week-van-filemon/POMS_S_BNN_123420">De Week van Filemon</a></li>
<li class="no-logo" data-tags="[]"><a href="/de-wereld-draait-door/POMS_S_VARA_059933">De Wereld Draait Door</a></li>
<li class="no-logo" data-tags="[]"><a href="/wolf/POMS_S_NTR_2089397">Wolf</a></li>
<li class="no-logo" data-tags="[]"><a href="/de-zomer-voorbij/POMS_S_TROS_098831">De Zomer Voorbij</a></li>
<li class="no-logo" data-tags="[]"><a href="http://www.npo.nl/zondag-met-lubach/POMS_S_VPRO_686783">Zondag met Lubach</a></li></ul>
</div>
</div>
<div class='mobile-dropdowns'>
<div class="npo3-dropdown-container visible-with-grid-collapse tag-filter"><div class="select-text">&nbsp;</div><select class="npo3-dropdown" name="tag-filter"><option value="">Genre</option>
<option value="Film">Film</option>
<option value="Humor">Humor</option>
<option value="Muziek">Muziek</option>
<option value="Reizen">Reizen</option>
<option value="Serie">Serie</option>
<option value="Webonly">Webonly</option></select></div>
<div class="npo3-dropdown-container visible-with-grid-collapse menu-item-dropdown"><div class="select-text">&nbsp;</div><select class="npo3-dropdown" name="menu-item"><option value="" selected="selected" disabled="disabled">Kies programma</option>
<option value="http://www.npo.nl/3-op-reis/POMS_S_BNN_097362" data-tags="[&quot;Reizen&quot;]">3 op Reis</option>
<option value="http://www.npo.nl/3-op-reis-backpack/POMS_S_BNN_868638" data-tags="[&quot;Reizen&quot;]">3 op Reis Backpack</option>
<option value="http://www.npo.nl/3doc/POMS_S_EO_098006" data-tags="[]">3Doc</option>
<option value="http://www.npo.nl/npo3/3lab" data-tags="[]">3Lab</option>
<option value="http://www.npo.nl/3onderzoekt/POMS_S_EO_420681" data-tags="[]">3Onderzoekt</option>
<option value="http://www.npo.nl/4jim/POMS_S_KN_1032073" data-tags="[&quot;Serie&quot;]">4JIM</option>
<option value="http://www.npo.nl/aanmodderfakker/POMS_S_AT_2125302" data-tags="[]">Aanmodderfakker</option>
<option value="http://www.npo.nl/aanmodderfakker/28-09-2015/AT_2044870" data-tags="[&quot;Film&quot;]">Aanmodderfakker</option>
<option value="http://www.npo.nl/ali-b-bekend-t/POMS_S_AT_1488792" data-tags="[&quot;Humor&quot;]">Ali B Bekend(t)</option>
<option value="http://www.npo.nl/all-stars-2-old-stars/POMS_S_VARA_1684367" data-tags="[]">All Stars 2 Old Stars</option>
<option value="http://www.npo.nl/alles-is-familie/POMS_S_VARA_2124339" data-tags="[]">Alles is familie</option>
<option value="http://www.npo.nl/beam-naar-de-balkan/POMS_S_EO_1957580" data-tags="[&quot;Webonly&quot;]">Beam naar de Balkan</option>
<option value="http://www.npo.nl/beste-nederlandse-film-simon/POMS_S_VPRO_113385" data-tags="[]">Beste Nederlandse Film: Simon</option>
<option value="http://www.npo.nl/het-beste-van-101/POMS_S_BNN_123416" data-tags="[]">Het Beste van 101</option>
<option value="http://www.npo.nl/de-beste-zangers-van-nederland/POMS_S_TROS_123317" data-tags="[&quot;Muziek&quot;]">De Beste Zangers van Nederland</option>
<option value="http://www.npo.nl/bij-dino-in-de-straat/POMS_S_NTR_858513" data-tags="[&quot;Humor&quot;]">Bij Dino in de Straat</option>
<option value="http://www.npo.nl/bride-flight/POMS_S_NCRV_1299373" data-tags="[&quot;Film&quot;]">Bride Flight</option>
<option value="http://www.npo.nl/bully/POMS_S_BNN_141165" data-tags="[]">Bully</option>
<option value="http://www.npo.nl/bureau-sport/POMS_S_VARA_124275" data-tags="[]">Bureau Sport</option>
<option value="http://www.npo.nl/club-27/POMS_S_AT_1673854" data-tags="[]">Club van 27</option>
<option value="http://www.npo.nl/college-tour/POMS_S_NTR_124382" data-tags="[]">College Tour</option>
<option value="http://www.npo.nl/dag-tegen-pesten/POMS_S_BNN_141164" data-tags="[]">Dag Tegen Pesten</option>
<option value="http://www.npo.nl/di-rect-concert/POMS_S_AT_1905717" data-tags="[]">DI-RECT Concert</option>
<option value="http://www.npo.nl/draadstaal-2015/POMS_S_AT_844275" data-tags="[&quot;Humor&quot;]">Draadstaal 2015</option>
<option value="http://www.npo.nl/eo-jongerendag/POMS_S_EO_097554" data-tags="[]">EO-Jongerendag</option>
<option value="http://www.npo.nl/escort/POMS_S_VARA_382597" data-tags="[&quot;Film&quot;]">Escort</option>
<option value="http://www.npo.nl/filmlab-voetzoeker/POMS_S_VPRO_2026622" data-tags="[]">Filmlab: Voetzoeker</option>
<option value="http://www.npo.nl/heemennekes-en-hellehonden/POMS_S_VPRO_1405182" data-tags="[&quot;Webonly&quot;,&quot;Serie&quot;]">Heemennekes en Hellehonden</option>
<option value="http://www.npo.nl/de-heineken-ontvoering/POMS_S_VARA_460439" data-tags="[]">De Heineken ontvoering</option>
<option value="http://www.npo.nl/hemel-op-aarde/POMS_S_AVRO_2064471" data-tags="[]">Hemel op aarde</option>
<option value="http://www.npo.nl/hotel-di-rect/POMS_S_TROS_1702109" data-tags="[]">Hotel DI-RECT</option>
<option value="http://www.npo.nl/hufterproef/POMS_S_EO_860345" data-tags="[]">Hufterproef</option>
<option value="http://www.npo.nl/humortv-op-3/POMS_S_VARA_099941" data-tags="[]">HumorTV op 3</option>
<option value="http://www.npo.nl/into-the-great-wide-open-2015/POMS_S_VPRO_1956627" data-tags="[]">Into The Great Wide Open (2015)</option>
<option value="http://www.npo.nl/jasper-van-kuijk/POMS_S_VARA_1986396" data-tags="[]">Jasper van Kuijk</option>
<option value="http://www.npo.nl/je-zal-het-maar-zijn/POMS_S_BNN_123419" data-tags="[]">Je Zal Het Maar Zijn</option>
<option value="http://www.npo.nl/jeuk/POMS_S_VARA_461066" data-tags="[&quot;Humor&quot;]">Jeuk</option>
<option value="http://www.npo.nl/keuringsdienst-van-waarde/POMS_S_KRO_123908" data-tags="[]">Keuringsdienst van Waarde</option>
<option value="http://www.npo.nl/larie/POMS_S_KRO_808174" data-tags="[&quot;Humor&quot;]">Larie</option>
<option value="http://www.npo.nl/lauren/POMS_S_AT_697765" data-tags="[]">Lauren!</option>
<option value="http://www.npo.nl/lek/POMS_S_VARA_382592" data-tags="[]">Lek</option>
<option value="http://www.npo.nl/loft/POMS_S_BNN_123423" data-tags="[]">Loft</option>
<option value="http://www.npo.nl/mindf-ck/POMS_S_AT_728505" data-tags="[]">Mindf*ck</option>
<option value="http://www.npo.nl/motief-voor-moord/POMS_S_EO_1249971" data-tags="[&quot;Serie&quot;]">Motief voor Moord</option>
<option value="http://www.npo.nl/bnn-presents-mysteryland/30-08-2014/BNN_101363640" data-tags="[&quot;Muziek&quot;]">Mysteryland</option>
<option value="http://www.npo.nl/de-nationale-iq-test/POMS_S_BNN_097257" data-tags="[]">De Nationale IQ test</option>
<option value="http://www.npo.nl/nederlands-film-festival-op-3/POMS_S_VPRO_2099142" data-tags="[]">Nederlands Film Festival op 3</option>
<option value="http://www.npo.nl/nos-op-3/POMS_S_NOS_059929" data-tags="[]">NOS op 3</option>
<option value="http://www.npo.nl/one-night-stand/POMS_S_VARA_099984" data-tags="[&quot;Film&quot;]">One night stand</option>
<option value="http://www.npo.nl/penoza/POMS_S_KRO_123909" data-tags="[]">Penoza</option>
<option value="http://www.npo.nl/plan-c/POMS_S_AVRO_460376" data-tags="[]">Plan C</option>
<option value="http://www.npo.nl/rambam/POMS_S_VARA_124202" data-tags="[]">Rambam</option>
<option value="http://www.npo.nl/ranking-the-stars/POMS_S_BNN_097277" data-tags="[&quot;Humor&quot;]">Ranking the Stars</option>
<option value="http://www.npo.nl/de-rekenkamer/POMS_S_KRO_123979" data-tags="[]">De Rekenkamer</option>
<option value="http://www.npo.nl/rolling-stones-sweet-summer-sun/POMS_S_NTR_449961" data-tags="[]">Rolling Stones: Sweet Summer Sun</option>
<option value="http://www.npo.nl/roue-verveer-ff-wat-anders/POMS_S_VARA_1914354" data-tags="[]">Roué Verveer - FF wat anders</option>
<option value="http://www.npo.nl/rundfunk/POMS_S_KN_1666829" data-tags="[]">Rundfunk</option>
<option value="http://www.npo.nl/sara-kroos/POMS_S_VARA_124118" data-tags="[]">Sara Kroos</option>
<option value="http://www.npo.nl/solar/POMS_S_BNN_097354" data-tags="[&quot;Muziek&quot;]">Solar</option>
<option value="http://www.npo.nl/sophie-in-de-kreukels/01-09-2015/BNN_101375894" data-tags="[]">Sophie in de kreukels</option>
<option value="http://www.npo.nl/spuiten-en-slikken/POMS_S_BNN_097254" data-tags="[]">Spuiten en Slikken</option>
<option value="http://www.npo.nl/spuiten-en-slikken-op-reis/POMS_S_BNN_097352" data-tags="[]">Spuiten en Slikken op Reis</option>
<option value="http://www.npo.nl/streetlab/POMS_S_KRO_797293" data-tags="[&quot;Humor&quot;]">Streetlab</option>
<option value="http://www.npo.nl/studio-powned/POMS_S_POWN_775803" data-tags="[]">Studio PowNed</option>
<option value="http://www.npo.nl/super-stream-me/POMS_S_VPRO_1790127" data-tags="[&quot;Serie&quot;]">Super Stream Me</option>
<option value="http://www.npo.nl/tbs/POMS_S_BNN_1613878" data-tags="[]">TBS</option>
<option value="http://www.npo.nl/teledoc-voetbalmiljonair-uit-oost/POMS_S_BNN_1914803" data-tags="[]">Teledoc: Voetbalmiljonair Uit Oost</option>
<option value="http://www.npo.nl/the-freestyle-games/POMS_S_BNN_1633995" data-tags="[]">The Freestyle Games</option>
<option value="http://www.npo.nl/the-undateables/POMS_S_BNN_399127" data-tags="[]">The Undateables</option>
<option value="http://www.npo.nl/theatershow-jandino/POMS_S_NTR_125623" data-tags="[&quot;Humor&quot;]">Theatershow Jandino</option>
<option value="http://www.npo.nl/de-thuiszorgman/POMS_S_VPRO_1593549" data-tags="[&quot;Webonly&quot;]">De Thuiszorgman</option>
<option value="http://www.npo.nl/troy/POMS_S_BNN_1404983" data-tags="[]">Troy</option>
<option value="http://www.npo.nl/uit-de-kast/POMS_S_KRO_123939" data-tags="[]">Uit de Kast</option>
<option value="http://www.npo.nl/uitzonderlijk-vervoer/POMS_S_VARA_788350" data-tags="[&quot;Humor&quot;]">Uitzonderlijk Vervoer</option>
<option value="http://www.npo.nl/van-de-straat/POMS_S_NCRV_121157" data-tags="[]">Van de straat</option>
<option value="http://www.npo.nl/van-god-los-de-film/POMS_S_BNN_388012" data-tags="[]">Van God Los de Film</option>
<option value="http://www.npo.nl/van-luizenleven-naar-luiers/POMS_S_EO_1592532" data-tags="[]">Van luizenleven naar luiers</option>
<option value="http://www.npo.nl/het-verhaal-van/POMS_S_TROS_137027" data-tags="[&quot;Serie&quot;]">Het verhaal van...</option>
<option value="http://www.npo.nl/verliefd-op-ibiza/POMS_S_AT_1340616" data-tags="[&quot;Film&quot;]">Verliefd op Ibiza</option>
<option value="http://www.npo.nl/villa-single-mama/POMS_S_EO_1673428" data-tags="[]">Villa Single Mama</option>
<option value="http://www.npo.nl/vriende/POMS_S_VPRO_2086252" data-tags="[&quot;Webonly&quot;]">Vriende</option>
<option value="http://www.npo.nl/de-week-van-filemon/POMS_S_BNN_123420" data-tags="[]">De Week van Filemon</option>
<option value="http://www.npo.nl/de-wereld-draait-door/POMS_S_VARA_059933" data-tags="[]">De Wereld Draait Door</option>
<option value="http://www.npo.nl/wolf/POMS_S_NTR_2089397" data-tags="[]">Wolf</option>
<option value="http://www.npo.nl/de-zomer-voorbij/POMS_S_TROS_098831" data-tags="[]">De Zomer Voorbij</option>
<option value="http://www.npo.nl/zondag-met-lubach/POMS_S_VPRO_686783" data-tags="[]">Zondag met Lubach</option></select></div>
</div>
</div>
<div class='collapse-menu hide' id='npo3-menu-4'>
<div class='visible-no-grid-collapse'>
<div class='menu-item-list'>
<ul class="menu-group "><li class="no-logo" data-tags="[]"><a href="http://www.npo.nl/npo3/nederlands-film-festival">NFF</a></li></ul>
</div>
</div>
<div class='mobile-dropdowns'>

<div class="npo3-dropdown-container visible-with-grid-collapse menu-item-dropdown"><div class="select-text">&nbsp;</div><select class="npo3-dropdown" name="menu-item"><option value="" selected="selected" disabled="disabled">Kies programma</option>
<option value="http://www.npo.nl/npo3/nederlands-film-festival" data-tags="[]">NFF</option></select></div>
</div>
</div>
<div class='collapse-menu hide' id='npo3-menu-2'>
<div class='visible-no-grid-collapse'>
<div class='menu-item-list'>

<ul class="menu-group "><li class="title">Websites</li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.spuiten-slikken&quot;}" data-tags="[]"><a href="http://spuitenenslikken.bnn.nl/" target="_blank">Spuiten &amp; Slikken</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.101barz&quot;}" data-tags="[]"><a href="http://101barz.bnn.nl/" target="_blank">101Barz</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.de-beste-singer-songwriter&quot;}" data-tags="[]"><a href="http://www.debestesingersongwriter.nl/" target="_blank">De Beste Singer-Songwriter</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.3lab&quot;}" data-tags="[]"><a href="http://www.npo3lab.nl/" target="_blank">3LAB</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.nos-op-3&quot;}" data-tags="[]"><a href="http://nos.nl/op3/" target="_blank">NOS op 3</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.3-op-reis&quot;}" data-tags="[]"><a href="http://3opreis.bnn.nl/home" target="_blank">3 op Reis</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.powned&quot;}" data-tags="[]"><a href="http://www.powned.tv/" target="_blank">PowNed</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.beam&quot;}" data-tags="[]"><a href="http://www.eo.nl/beam/" target="_blank">BEAM</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.proefkonijnen&quot;}" data-tags="[]"><a href="http://proefkonijnen.bnn.nl/" target="_blank">Proefkonijnen</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.keuringsdienst-van-waarde&quot;}" data-tags="[]"><a href="http://keuringsdienstvanwaarde.kro.nl/" target="_blank">Keuringsdienst van Waarde</a></li></ul>
<ul class="menu-group "><li class="title">Themakanalen</li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.humortv&quot;}" data-tags="[]"><a href="http://humortv.vara.nl/" target="_blank">HumorTV</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.101tv&quot;}" data-tags="[]"><a href="http://www.101.tv/" target="_blank">101TV</a></li></ul>
<ul class="menu-group "><li class="title">Contact</li>
<li class="no-logo" data-tags="[]"><a href="http://www.npo.nl/npo3/contact--2">Contact met NPO 3</a></li></ul>
<ul class="menu-group "><li class="title">Radiozenders</li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.npo-funx&quot;}" data-tags="[]"><a href="http://www.funx.nl/" target="_blank">NPO FunX</a></li>
<li class="no-logo" data-scorecard="{&quot;name&quot;:&quot;menu.extern.npo-3fm&quot;}" data-tags="[]"><a href="http://www.3fm.nl/home" target="_blank">NPO 3FM</a></li></ul>
</div>
</div>
<div class='mobile-dropdowns'>

<div class="npo3-dropdown-container visible-with-grid-collapse menu-item-dropdown"><div class="select-text">&nbsp;</div><select class="npo3-dropdown" name="menu-item"><option value="" selected="selected" disabled="disabled">Kies programma</option>
<option value="http://spuitenenslikken.bnn.nl/" data-tags="[]">Spuiten &amp; Slikken</option>
<option value="http://101barz.bnn.nl/" data-tags="[]">101Barz</option>
<option value="http://www.npo.nl/npo3/contact--2" data-tags="[]">Contact met NPO 3</option>
<option value="http://www.debestesingersongwriter.nl/" data-tags="[]">De Beste Singer-Songwriter</option>
<option value="http://humortv.vara.nl/" data-tags="[]">HumorTV</option>
<option value="http://www.npo3lab.nl/" data-tags="[]">3LAB</option>
<option value="http://www.101.tv/" data-tags="[]">101TV</option>
<option value="http://nos.nl/op3/" data-tags="[]">NOS op 3</option>
<option value="http://3opreis.bnn.nl/home" data-tags="[]">3 op Reis</option>
<option value="http://www.powned.tv/" data-tags="[]">PowNed</option>
<option value="http://www.funx.nl/" data-tags="[]">NPO FunX</option>
<option value="http://www.eo.nl/beam/" data-tags="[]">BEAM</option>
<option value="http://proefkonijnen.bnn.nl/" data-tags="[]">Proefkonijnen</option>
<option value="http://keuringsdienstvanwaarde.kro.nl/" data-tags="[]">Keuringsdienst van Waarde</option>
<option value="http://www.3fm.nl/home" data-tags="[]">NPO 3FM</option></select></div>
</div>
</div>
</div>
</div></div></div>
<div class="main-content page-content"><div class="container"><div class="row-fluid no-gutter"><div class="span12"><div class="row-fluid no-gutter"><div class='row-fluid span12'>
<h1><h1 title="Radio Bergeijk, toewijding in beeld"><a href="/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994">Radio Bergeijk, toewijding in beeld</a></h1></h1>
</div>
<div class="span9 npo3-main"><div class="row-fluid no-gutter highlight-group focus-left"><div class="span12 focus"><div class="player-container highlight-container"><div class="highlight"><img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="background-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_carousel_medium-b248249edfa754063c3703b6b3f021f0.png" />
<div class="video-player-container" data-auto-play="true" data-manual="true" data-prid="VPRO_1122739" data-skin="zapp" data-video-player="true" id="video-player-container" itemscope="" itemtype="http://schema.org/VideoObject"><meta content="Radio Bergeijk, toewijding in beeld" itemprop="name" /><meta content="Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met technicus Ted van Lieshout." itemprop="description" /><meta itemprop="thumbnailUrl" /><meta content="PT35M0S" itemprop="duration" /><div class="video-player-holder" id="video-player"></div></div>
<div class="video-player-settings"><a href="#" class="toggle-player-size" data-scorecard="{&quot;name&quot;:&quot;beeld-vergroten&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas&quot;}" data-toggle-player-size="" title="Beeld vergroten / verkleinen"><span class="enlarge"><span class="npo-glyph expand"></span></span><span class="shrink"><span class="npo-glyph shrink"></span></span></a><a href="#" data-scorecard="{&quot;name&quot;:&quot;buitenland&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas&quot;}" data-toggle-player-settings="" title="Instellingen"><span class="npo-glyph cog"></span></a></div>
<div class='play-button'></div>
<div class="meta-container"><div class="meta first"><div class="md-label"><span class="npo-glyph triangle-right"></span></div><div class="md-value">35:00</div></div></div>
<div class='highlight-text'>
<div class='tags'>
<div class='tag g-w'>Aflevering</div>
</div>
</div>
</div></div></div></div><div class='program-metadata-container'>
<div class='notifications'>

</div>

</div>
<div class='item-list show-container' data-source='/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994/search'>
<div class="npo3-dropdown-container visible-with-grid-collapse switch-dropdown" id="switch-dropdown"><div class="select-text">&nbsp;</div><select class="npo3-dropdown" name="media_type"><option value="" selected="selected" disabled="disabled">Laat zien</option>
<option value="video">Alles</option>
<option value="broadcast">Afleveringen</option></select></div>
<form accept-charset="UTF-8" action="search" id="program-search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='npo3-filter-block clearfix'>
<div class="video-block switch-block filter "><a href="#" class="omroep-light " data-media-type="video" data-scorecard="{&quot;name&quot;:&quot;programmas.content.alles&quot;,&quot;soft&quot;:true}">Alles <span class="num-found"></span></a></div>
<div class="broadcast-block switch-block filter "><a href="#" class="omroep-light inactive" data-media-type="broadcast" data-scorecard="{&quot;name&quot;:&quot;programmas.content.afleveringen&quot;,&quot;soft&quot;:true}">Afleveringen <span class="num-found"></span></a></div>


</div>
</form>

<div class='video-block video-container-block' data-media-type='video' id='videos-block'>
<div class='content'><div class="search-results" data-num-found="5" data-rows="8" data-start="0"><div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 25 jun 2007 23:10</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/18-06-2007/VPRO_1122738"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/18-06-2007/VPRO_1122738">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 18 jun 2007 23:05</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/11-06-2007/VPRO_1122737"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/11-06-2007/VPRO_1122737">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 11 jun 2007 22:55</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/04-06-2007/VPRO_1122736"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/04-06-2007/VPRO_1122736">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 4 jun 2007 23:15</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/21-05-2007/VPRO_1121819"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/21-05-2007/VPRO_1121819">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 21 mei 2007 22:55</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>

</div></div>
<div class='more-button-container'><div class="more-button"><a href="/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994?media_type=video" class="with-icon left"><span class="npo-glyph plus"></span><span>Meer </span></a></div></div>
</div>
<div class='broadcast-block video-container-block' data-media-type='broadcast' id='broadcasts-block'>
<div class='content'><div class="search-results" data-num-found="5" data-rows="8" data-start="0"><div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 25 jun 2007 23:10</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/18-06-2007/VPRO_1122738"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/18-06-2007/VPRO_1122738">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 18 jun 2007 23:05</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/11-06-2007/VPRO_1122737"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/11-06-2007/VPRO_1122737">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 11 jun 2007 22:55</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/04-06-2007/VPRO_1122736"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/04-06-2007/VPRO_1122736">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 4 jun 2007 23:15</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>
<div class='row-fluid item'>
<div class='span4 image'>
<a href="/radio-bergeijk-toewijding-in-beeld/21-05-2007/VPRO_1121819"><div class="meta-container"><div class="meta first"><div class="md-label no-text"><span class="npo-glyph triangle-right"></span></div></div></div>
<img alt="Afbeelding van Radio Bergeijk, toewijding in beeld" class="program-image" src="//www-assets.npo.nl/assets/placeholders/nederland_npo_thumb_large-d01f19cd515e73fc516be7303224d3a6.png" />
</a></div>
<div class='span8 item-description'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/21-05-2007/VPRO_1121819">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='tag w-g'>Aflevering</div>
<h4><span class="inactive">Datum uitzending:</span> Ma 21 mei 2007 22:55</h4>
<p>Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</p>
</div>
</div>

</div></div>
<div class='more-button-container'><div class="more-button"><a href="/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994?media_type=broadcast" class="with-icon left"><span class="npo-glyph plus"></span><span>Meer afleveringen</span></a></div></div>
</div>
</div>

</div><div class="span3 npo3-sidebar"><div class="sidebar-item broadcaster-box double-height" data-mid="VPRO_1122739" data-parameterized-title="radio-bergeijk-toewijding-in-beeld" data-title="Radio Bergeijk, toewijding in beeld"><div class="sidebar-item-content"><div class='broadcaster-box-content'>
<div class='broadcaster-logo'><a href="http://www.vpro.nl" data-scorecard="{&quot;name&quot;:&quot;vpro&quot;,&quot;prefix&quot;:&quot;npo.softclick.programmas.extern&quot;}" target="_blank" title="VPRO"><img alt="Logo van VPRO" class="no-shadow" src="//www-assets.npo.nl/uploads/broadcaster/40/logo/regular_regular_VPRO_alt2.png" /></a></div>
<h2><a href="/radio-bergeijk-toewijding-in-beeld/POMS_S_VPRO_083994">Radio Bergeijk, toewijding in beeld</a></h2>
<div class='ratings'>
<span class="npo-glyph rating-12"></span>
<span class="npo-glyph rating-discrimination" title="Discriminatie"></span><span class="npo-glyph rating-strong-language" title="Grof taalgebruik"></span>
</div>
<div class='sort-date'>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739"><div class='broadcast-prefix'>Laatste uitzending:</div>
Ma 25 jun 2007 23:10 - 23:45
</a></div>
<div class='meta-content visible-for-desktop'>
<h3><a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739">Radio Bergeijk, toewijding in beeld</a></h3>
<div class='meta-scroller iscroll'>
<p class='overflow-description'><span class="hide-original">Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</span><span class="omission" data-scorecard="{&quot;name&quot;:&quot;programmas.radio-bergeijk-toewijding-in-beeld.afleveringen.VPRO_1122739.meer-beschrijving&quot;}"><a href="#" class="overflow-toggle">Meer</a></span><span class="omission-overflow hide"> Serie verslagen van hun levenswerk: Radio Bergeijk. Een filmploeg over de radiovloer! Ze moesten even aan de gedachte wennen, de presentatoren van Radio Bergeijk. En het heeft even geduurd voordat Toon Spoorenberg en Peer van Eersel konden instemmen met die druktemakerij in hun zo vertrouwde en overzichtelijke radiobestaan. Maar het is er dan toch van gekomen. Pieter Verhoeff heeft een serie gemaakt van Radio Bergeijk, het ook landelijk bekende radiostation voor Bergeijk, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Ted van Lieshout. De pro Deo werkende radiomannen ontvangen in hun studio een keur aan plaatselijke gasten of doen rechtstreeks verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. De televisieserie wordt niet alleen een fel realistisch portret van de makers van het vermaarde radiostation, maar via hen ook van het dorp Bergeijk. Radio Bergeijk is en probeert een satire te zijn op een lokaal radiostation. In de verbeelding van de makers is het fictieve dorp Bergeijk een huiveringwekkend oord waar hedendaagse xenofobie, orthodox atheïsme, vrouwonvriendelijkheid en homo-afkeer om de voorrang strijden. Het radiostation probeert van deze benepen dorpse identiteit pseudo-professioneel verslag te doen. En hanteert als stijlmiddel de grove overdrijving. Teksten en spel: Pieter Bouwman en George van Houts. Gastacteurs o.a.: Alex Klaasen, Jeroen van Merwijk, Marjan Luif, Annet Malherbe, Anniek Pheifer, Margôt Ros en Lies Visschedijk.<a href="#" class="overflow-toggle less">Minder</a></span></p>
</div>
<div class='mobile-description'>
<p class='overflow-description'><span class="hide-original">Serie verslagen over Radio Bergeijk, een satire op een lokaal radiostation, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis ...</span><span class="omission" data-scorecard="{&quot;name&quot;:&quot;programmas.radio-bergeijk-toewijding-in-beeld.afleveringen.VPRO_1122739.meer-beschrijving&quot;}"><a href="#" class="overflow-toggle">Meer</a></span><span class="omission-overflow hide"> Serie verslagen van hun levenswerk: Radio Bergeijk. Een filmploeg over de radiovloer! Ze moesten even aan de gedachte wennen, de presentatoren van Radio Bergeijk. En het heeft even geduurd voordat Toon Spoorenberg en Peer van Eersel konden instemmen met die druktemakerij in hun zo vertrouwde en overzichtelijke radiobestaan. Maar het is er dan toch van gekomen. Pieter Verhoeff heeft een serie gemaakt van Radio Bergeijk, het ook landelijk bekende radiostation voor Bergeijk, dat wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Ted van Lieshout. De pro Deo werkende radiomannen ontvangen in hun studio een keur aan plaatselijke gasten of doen rechtstreeks verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. De televisieserie wordt niet alleen een fel realistisch portret van de makers van het vermaarde radiostation, maar via hen ook van het dorp Bergeijk. Radio Bergeijk is en probeert een satire te zijn op een lokaal radiostation. In de verbeelding van de makers is het fictieve dorp Bergeijk een huiveringwekkend oord waar hedendaagse xenofobie, orthodox atheïsme, vrouwonvriendelijkheid en homo-afkeer om de voorrang strijden. Het radiostation probeert van deze benepen dorpse identiteit pseudo-professioneel verslag te doen. En hanteert als stijlmiddel de grove overdrijving. Teksten en spel: Pieter Bouwman en George van Houts. Gastacteurs o.a.: Alex Klaasen, Jeroen van Merwijk, Marjan Luif, Annet Malherbe, Anniek Pheifer, Margôt Ros en Lies Visschedijk.<a href="#" class="overflow-toggle less">Minder</a></span></p>
</div>
</div>
<div class='clickouts-and-actions'>
<div class='clickouts meta-content visible-for-desktop'>
<a href="http://www.radiobergeijk.tv" class="external-link pull-right" data-scorecard="{&quot;name&quot;:&quot;programmas.extern.radio-bergeijk-toewijding-in-beeld.VPRO_1122739&quot;}" target="_blank"><span class="npo-glyph arrow-jump"></span>
Site
</a></div>
<a href="#" class="meta-trigger visible-for-small-screen"><span class='show-on-closed'>+ Meer</span>
<span class='show-on-open hide'>- Minder</span>
</a><div class='actions'>
<div class='action-buttons for-guest'>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk-toewijding-in-beeld%2FPOMS_S_VPRO_083994" class="button favorite-button guest-favorite-button"><span class="npo-glyph heart contrast"></span> <span class="glyph-label">Favoriet</span></a>
<a href="/radio-bergeijk-toewijding-in-beeld/25-06-2007/VPRO_1122739/delen" class="last button share-modal-button" data-id="VPRO_1122739" rel="nofollow"><span class="npo-glyph share contrast"></span> <span class="glyph-label">Delen</span></a>
</div>

</div>
</div>
</div>
</div></div>
<div class="sidebar-item black-sidebar-item countdown"><div class="sidebar-item-content"><div class='countdown-header'>NPO 3 kijk je <span class="green">live</span> over</div>
<div class='counter' data-count-down='2015-09-30T19:55:00+02:00'></div>
<div class='tv-guide'>
<div class='tv-guide-entry'>
<a href="/nederlands-film-festival-op-3/30-09-2015/VPWON_1245766" data-scorecard="{&quot;name&quot;:&quot;zijbalk.zometeen.nederlands-film-festival-op-3.afleveringen.VPWON_1245766&quot;}"><div class='tv-guide-time'>19:55</div>
<div class='tv-guide-program'>
<div class='primary'>Nederlands Film Festival op 3</div>
<div class='secondary'></div>
</div>
</a></div>
</div>
</div></div><div class="sidebar-item white-sidebar-item double-height tonight"><div class="sidebar-item-content"><h2>Vanavond op 3</h2>
<div class='tv-guide'>
<div class='tv-guide-entry'>
<a href="/nederlands-film-festival-op-3/30-09-2015/VPWON_1245766" data-scorecard="{&quot;name&quot;:&quot;zijbalk.vanavond.nederlands-film-festival-op-3.afleveringen.VPWON_1245766&quot;}"><div class='tv-guide-time'>19:55</div>
<div class='tv-guide-program'>
<div class='secondary'>Nederlands Film Festival op 3</div>
<div class='secondary'></div>
</div>
</a></div>
<div class='tv-guide-entry'>
<a href="/nos-op-3-flits/30-09-2015/POW_01076024" data-scorecard="{&quot;name&quot;:&quot;zijbalk.vanavond.nos-op-3-flits.afleveringen.POW_01076024&quot;}"><div class='tv-guide-time'>20:25</div>
<div class='tv-guide-program'>
<div class='secondary'>NOS op 3 flits</div>
<div class='secondary'></div>
</div>
</a></div>
<div class='tv-guide-entry'>
<a href="/de-zomer-voorbij/30-09-2015/AT_2041452" data-scorecard="{&quot;name&quot;:&quot;zijbalk.vanavond.de-zomer-voorbij.afleveringen.AT_2041452&quot;}"><div class='tv-guide-time'>20:30</div>
<div class='tv-guide-program'>
<div class='secondary'>De Zomer Voorbij</div>
<div class='secondary'></div>
</div>
</a></div>
<div class='tv-guide-entry'>
<a href="/nos-op-3-flits/30-09-2015/POW_01076318" data-scorecard="{&quot;name&quot;:&quot;zijbalk.vanavond.nos-op-3-flits.afleveringen.POW_01076318&quot;}"><div class='tv-guide-time'>21:10</div>
<div class='tv-guide-program'>
<div class='secondary'>NOS op 3 flits</div>
<div class='secondary'></div>
</div>
</a></div>
<div class='tv-guide-entry'>
<a href="/the-undateables/30-09-2015/BNN_101375929" data-scorecard="{&quot;name&quot;:&quot;zijbalk.vanavond.the-undateables.afleveringen.BNN_101375929&quot;}"><div class='tv-guide-time'>21:15</div>
<div class='tv-guide-program'>
<div class='secondary'>The Undateables</div>
<div class='secondary'></div>
</div>
</a></div>
<div class='tv-guide-entry'>
<a href="/nos-op-3-flits/30-09-2015/POW_01076134" data-scorecard="{&quot;name&quot;:&quot;zijbalk.vanavond.nos-op-3-flits.afleveringen.POW_01076134&quot;}"><div class='tv-guide-time'>21:58</div>
<div class='tv-guide-program'>
<div class='secondary'>NOS op 3 flits</div>
<div class='secondary'></div>
</div>
</a></div>
</div>
<div class='guide-link'>
<a href="/gids" data-scorecard="{&quot;name&quot;:&quot;zijbalk.volledige-gids&quot;,&quot;potag2&quot;:&quot;npo3&quot;}">Toon volledige gids <span class="npo-glyph arrow-jump right"></span></a>
</div>
</div></div><div class="sidebar-item white-sidebar-item double-height top-5"><div class="sidebar-item-content"><h2>Top 5 kort</h2>
<a href="/rambam/28-01-2015/VARA_101373256/POMS_VARA_828806" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.top5kort.rambam.VARA_101373256.fragment.POMS_VARA_828806&quot;}"><div class='image-container'>
<div class='image-wrapper'><img alt="595387" src="http://images.poms.omroep.nl/image/s145/c145x120/595387.png" /></div>
</div>
<div class='content'>
<div class='title'>De laptopmonteur!</div>
<div class='views'>
565 X bekeken
</div>
</div>
</a><a href="/streetlab/28-09-2015/KN_1673142/POMS_KN_2014960" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.top5kort.streetlab.KN_1673142.fragment.POMS_KN_2014960&quot;}"><div class='image-container'>
<div class='image-wrapper'><img alt="658178" src="http://images.poms.omroep.nl/image/s145/c145x120/658178.png" /></div>
</div>
<div class='content'>
<div class='title'>Vriendin Stijn doet je make-uppie vandaag</div>
<div class='views'>
465 X bekeken
</div>
</div>
</a><a href="/sophie-in-de-kreukels/22-09-2015/BNN_101375888/POMS_BNN_1869795" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.top5kort.sophie-in-de-kreukels.BNN_101375888.fragment.POMS_BNN_1869795&quot;}"><div class='image-container'>
<div class='image-wrapper'><img alt="655913" src="http://images.poms.omroep.nl/image/s145/c145x120/655913.png" /></div>
</div>
<div class='content'>
<div class='title'>Sophie op bezoek bij haar moeder</div>
<div class='views'>
453 X bekeken
</div>
</div>
</a><a href="/jeuk/27-09-2015/VARA_101375984/POMS_VARA_2035590" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.top5kort.jeuk.VARA_101375984.fragment.POMS_VARA_2035590&quot;}"><div class='image-container'>
<div class='image-wrapper'><img alt="657826" src="http://images.poms.omroep.nl/image/s145/c145x120/657826.png" /></div>
</div>
<div class='content'>
<div class='title'>'Mancave'</div>
<div class='views'>
351 X bekeken
</div>
</div>
</a><a href="/tv-lab-streetlab/21-08-2014/KN_1662356/POMS_KRO_1898796" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.top5kort.tv-lab-streetlab.KN_1662356.fragment.POMS_KRO_1898796&quot;}"><div class='image-container'>
<div class='image-wrapper'><img alt="652457" src="http://images.poms.omroep.nl/image/s145/c145x120/652457.png" /></div>
</div>
<div class='content'>
<div class='title'>Waar vallen vrouwen op? (Grapjes)</div>
<div class='views'>
331 X bekeken
</div>
</div>
</a></div></div><div class="sidebar-item now-playing" data-now-playing-radiobox-id="3"><div class="sidebar-item-content"><a href="http://www.3fm.nl/live" data-scorecard="{&quot;name&quot;:&quot;zijbalk.clickout.npo-3fm&quot;}" target="_blank"><img alt="3fm" class="now-playing-logo" src="//www-assets.npo.nl/assets/npo3/3fm-a23cd9654c20386be4a0b6961e5da75b.png" />
<div class='now-playing-header'>
<span class='header green'></span>
op npo 3fm
</div>
<div class='now-playing-content'>
<div class='artist green'></div>
<div class='track'></div>
</div>
</a></div></div><div class="sidebar-item white-sidebar-item double-height"><div class="sidebar-item-content"><h2>Meest bekeken</h2>
<a href="/penoza-iv/27-09-2015/KN_1673186" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.meest-gezien.penoza-iv.afleveringen.KN_1673186&quot;}"><div class='image-container'></div>
<div class='content'>
<div class='title'>Penoza IV</div>
<div class='views'>21.791 X bekeken - 27 september</div>
</div>
</a><a href="/streetlab/28-09-2015/KN_1673142" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.meest-gezien.streetlab.afleveringen.KN_1673142&quot;}"><div class='image-container'><img alt="658058" src="http://images.poms.omroep.nl/image/s145/c145x120/658058.png" /></div>
<div class='content'>
<div class='title'>Streetlab</div>
<div class='views'>14.857 X bekeken - 28 september</div>
</div>
</a><a href="/sophie-in-de-kreukels/29-09-2015/BNN_101375886" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.meest-gezien.sophie-in-de-kreukels.afleveringen.BNN_101375886&quot;}"><div class='image-container'><img alt="658510" src="http://images.poms.omroep.nl/image/s145/c145x120/658510.png" /></div>
<div class='content'>
<div class='title'>Sophie In De Kreukels</div>
<div class='views'>7.299 X bekeken - 29 september</div>
</div>
</a><a href="/zondag-met-lubach/27-09-2015/VPWON_1243303" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.meest-gezien.zondag-met-lubach.afleveringen.VPWON_1243303&quot;}"><div class='image-container'><img alt="657487" src="http://images.poms.omroep.nl/image/s145/c145x120/657487.png" /></div>
<div class='content'>
<div class='title'>Zondag met Lubach</div>
<div class='views'>6.692 X bekeken - 27 september</div>
</div>
</a><a href="/aanmodderfakker/28-09-2015/AT_2044870" class="most-views-tile" data-scorecard="{&quot;name&quot;:&quot;zijbalk.meest-gezien.aanmodderfakker.afleveringen.AT_2044870&quot;}"><div class='image-container'><img alt="577015" src="http://images.poms.omroep.nl/image/s145/c145x120/577015.png" /></div>
<div class='content'>
<div class='title'>Aanmodderfakker</div>
<div class='views'>5.738 X bekeken - 28 september</div>
</div>
</a></div></div>
</div></div></div></div></div></div></div>

<div class="page-footer" id="npo-footer"><div class='container'>
<div class='broadcasters-and-corporate-links'>
<div class='broadcasters'>
<strong>Omroepen</strong>
<ul>
<li><a href="http://www.avrotros.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.avrotros&quot;}" target="_blank">AVROTROS</a></li>
<li><a href="http://www.bnn.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.bnn&quot;}" target="_blank">BNN</a></li>
<li><a href="http://www.eo.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.eo&quot;}" target="_blank">EO</a></li>
<li><a href="http://www.omroephuman.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.human&quot;}" target="_blank">HUMAN</a></li>
<li><a href="http://www.kro-ncrv.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.kro-ncrv&quot;}" target="_blank">KRO-NCRV</a></li>
</ul>
<ul>
<li><a href="http://www.omroepmax.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.max&quot;}" target="_blank">MAX</a></li>
<li><a href="http://www.nos.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.nos&quot;}" target="_blank">NOS</a></li>
<li><a href="http://www.ntr.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.ntr&quot;}" target="_blank">NTR</a></li>
<li><a href="http://www.powned.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.powned&quot;}" target="_blank">PowNed</a></li>
<li><a href="http://www.vara.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.vara&quot;}" target="_blank">VARA</a></li>
</ul>
<ul>
<li><a href="http://www.vpro.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.vpro&quot;}" target="_blank">VPRO</a></li>
<li><a href="http://www.omroepwnl.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.wnl&quot;}" target="_blank">WNL</a></li>
<li><a href="http://www.zvk.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.zvk&quot;}" target="_blank">ZvK</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Meer NPO</strong>
<ul>
<li><a href="http://www.npoplus.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.npo-plus&quot;}" target="_blank">NPO Plus</a></li>
<li><a href="http://www.bvn.tv" data-scorecard="{&quot;name&quot;:&quot;footer.extern.bvn&quot;}" target="_blank">BVN</a></li>
<li><a href="http://www.npo.nl/specials/regionale-omroepen">Regionaal</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Organisatie</strong>
<ul>
<li><a href="http://www.npo.nl/overnpo">Over NPO</a></li>
<li><a href="http://www.npo.nl/overnpo/pers">Pers</a></li>
<li><a href="http://www.npo.nl/overnpo/vacatures-en-stages-npo">Vacatures</a></li>
<li><a href="http://www.nposales.com" data-scorecard="{&quot;name&quot;:&quot;footer.extern.npo-sales&quot;}" target="_blank">NPO Sales</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Volg de NPO</strong>
<ul>
<li><a href="http://omroep.dmd.omroep.nl/x/plugin/?pName=subscribe&amp;MIDRID=S7Y1swQAA98&amp;pLang=nl&amp;Z=543417650" data-scorecard="{&quot;name&quot;:&quot;footer.extern.aanmelden-nieuwsbrief&quot;}" target="_blank">Aanmelden nieuwsbrief</a></li>
<li><a href="https://instagram.com/npo.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.instagram&quot;}" target="_blank">Instagram</a></li>
<li><a href="https://www.facebook.com/NPO.nl" data-scorecard="{&quot;name&quot;:&quot;footer.extern.facebook&quot;}" target="_blank">Facebook</a></li>
<li><a href="http://www.twitter.com/publiekeomroep" data-scorecard="{&quot;name&quot;:&quot;footer.extern.twitter&quot;}" target="_blank">Twitter</a></li>
</ul>
</div>
<div class='corporate-links'>
<strong>Onze apps</strong>
<ul>
<li><a href="https://itunes.apple.com/nl/app/uitzending-gemist/id323998316?mt=8" data-scorecard="{&quot;name&quot;:&quot;footer.extern.ipad-iphone&quot;}" target="_blank">iPad &amp; iPhone</a></li>
<li><a href="https://play.google.com/store/apps/details?id=nl.uitzendinggemist" data-scorecard="{&quot;name&quot;:&quot;footer.extern.android&quot;}" target="_blank">Android</a></li>
<li><a href="https://help.npo.nl/faqs/smarttv-hbbtv">Smarttv en HbbTV</a></li>
</ul>
</div>
</div>
<div class='general-links'>
<ul class='disclaimers'>
<li><a href="http://www.npo.nl/contact">Contact</a></li>
<li><a href="http://help.npo.nl
" data-scorecard="{&quot;name&quot;:&quot;footer.extern.help&quot;}" target="_blank">Help</a></li>
<li><a href="http://www.npo.nl/uitgelicht.rss">RSS</a></li>
<li><a href="http://www.npo.nl/algemene-voorwaarden-privacy">Algemene Voorwaarden &amp; Privacy</a></li>
<li><a href="http://cookiesv2.publiekeomroep.nl/" data-scorecard="{&quot;name&quot;:&quot;footer.extern.cookiebeleid&quot;}" target="_blank">Cookiebeleid</a></li>
</ul>
<div class='logo small'><a href="http://www.npo.nl/"><img alt="NPO logo" src="//www-assets.npo.nl/assets/npo_logo-c0f9cdc10f5633cc362e5b8113a3350e.png" /></a></div>
</div>
</div>
</div>
</div>
<script class="npo_cc_inline_analytics" language="JavaScript1.3" src="//b.scorecardresearch.com/c2/17827132/cs.js" type="text/plain"></script>

<script type="text/plain" class="npo_cc_inline_advertising">(function() {
  var _fbq = window._fbq || (window._fbq = []);
  if (!_fbq.loaded) {
    var fbds = document.createElement('script');
    fbds.async = true;
    fbds.src = '//connect.facebook.net/en_US/fbds.js';
    var s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(fbds, s);
    _fbq.loaded = true;
  }
  _fbq.push(['addPixelId', '1487544264866945']);
})();
window._fbq = window._fbq || [];
window._fbq.push(['track', 'PixelInitialized', {}]);
</script>
<noscript><img height="1" width="1" alt="" style="display:none" class="npo_cc_inline_advertising" src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAA AICTAEAOw==" data-src="https://www.facebook.com/tr?id=1487544264866945&amp;ev=PixelInitialized" /></noscript>
</body>
</html>
`
