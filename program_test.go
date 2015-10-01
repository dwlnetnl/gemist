package gemist

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseProgram(t *testing.T) {
	assert := assert.New(t)

	l, err := time.LoadLocation("Europe/Amsterdam")
	require.NoError(t, err, "error loading time location")

	r := strings.NewReader(testDataProgramBroadcast)

	_p := Program{
		MediaItem: MediaItem{
			Title:       "Radio Bergeijk",
			Description: "Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een keur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.\n ",
			ImageURLs: []string{
				"http://images.poms.omroep.nl/image/215303.png",
				"http://images.poms.omroep.nl/image/193478.png",
				"http://images.poms.omroep.nl/image/215304.png",
			},
			URL: "http://www.npo.nl/radio-bergeijk/POMS_S_VPRO_396280",
		},
		bs: []*BroadcastProxy{
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk : De allerlaatste !",
					Description: "Aan het begin van de uitzending wordt Tedje van Lieshout dood aantroffen in zijn werkhok. Een groots gevoel van machteloze moedeloosheid overvalt presen...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk-de-allerlaatste/06-10-2007/POMS_VPRO_396279",
				},
				SubTitle: "Zaterdagavond 6 okt om Half 7 op Radio 1",
				Date:     time.Date(2007, time.October, 6, 18, 32, 0, 0, l),
				Length:   1502000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk: De één na laatste!",
					Description: "Hoe kan het nu opeens zo snel gaan! Met de acties tégen Radio Bergeijk.\nToon leent 17 euro van Peer. Krijgt hij volgende week weer terug.\nDe vader van...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk-de-een-na-laatste/29-09-2007/POMS_VPRO_396278",
				},
				SubTitle: "Zaterdagavond 29 sept om Half 7 op Radio 1",
				Date:     time.Date(2007, time.September, 29, 18, 32, 0, 0, l),
				Length:   1600000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk",
					Description: "Er hebben zich inmiddels veel meer dan 200 inwoners aangesloten bij het gemeentelijk actiecomité tégen Radio Bergeijk. Toon heeft daarover een 18 pagina...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk/22-09-2007/POMS_VPRO_396277",
				},
				SubTitle: "Elke zaterdagavond om Half 7 op Radio 1",
				Date:     time.Date(2007, time.September, 22, 18, 32, 0, 0, l),
				Length:   1522000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk",
					Description: "Toon en Peer zijn vorige week uit GrandCafé Beeks gezet en zijn nu teruggekeerd\nin hun oude loods op het industrieterrein. En alsof er niets gebeurd is...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk/15-09-2007/POMS_VPRO_396276",
				},
				SubTitle: "Elke zaterdagavond om Half 7 op Radio 1",
				Date:     time.Date(2007, time.September, 15, 18, 32, 0, 0, l),
				Length:   1351000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk",
					Description: "Deze laatste zomerspeciaal van 2007 wordt ge-evalueerd door cafébaas Jan Beeks.\nIn de rubriek vacaturebank zoekt uitvaartverzorging van der Stappen een...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk/08-09-2007/POMS_VPRO_396275",
				},
				SubTitle: "Elke zaterdagavond om Half 7 op Radio 1",
				Date:     time.Date(2007, time.September, 8, 18, 32, 0, 0, l),
				Length:   1500000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk",
					Description: "De zomerspeciaal van deze week wordt ge-evalueerd door Joke Leermakers.\nOverblijfkrachtcentrale zoekt ex-TBS er die zijn talent wil inzetten.\nHenk en ...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk/01-09-2007/POMS_VPRO_396274",
				},
				SubTitle: "Elke zaterdagavond om Half 7 op Radio 1",
				Date:     time.Date(2007, time.September, 1, 18, 32, 0, 0, l),
				Length:   1486000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Geen Radio Bergeijk",
					Description: "",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/geen-radio-bergeijk/25-08-2007/POMS_VPRO_396273",
				},
				SubTitle: "NOS verslag EK finale dameshockey",
				Date:     time.Date(2007, time.August, 25, 18, 32, 0, 0, l),
				Length:   1680000000000,
			},
			&BroadcastProxy{
				MediaItem: MediaItem{
					Title:       "Radio Bergeijk",
					Description: "De zomerspeciaal van deze week wordt ge-evalueerd door Klaas Groot.\nIn het negerrubriek De Tamtam negernieuws over het negerboek over de legendarische ...",
					ImageURLs: []string{
						"http://images.poms.omroep.nl/image/s174/c174x98/215303.png",
					},
					URL: "http://www.npo.nl/radio-bergeijk/18-08-2007/POMS_VPRO_396272",
				},
				SubTitle: "Elke zaterdagavond om Half 7 op Radio 1",
				Date:     time.Date(2007, time.August, 18, 18, 32, 0, 0, l),
				Length:   1680000000000,
			},
		},
	}

	p, err := ParseProgram(r)
	assert.NoError(err)
	assert.Equal(_p.Title, p.Title, "title not equal")
	assert.Equal(_p.Description, p.Description, "description not equal")
	assert.Equal(_p.ImageURLs, p.ImageURLs, "image URLs not equal")
	assert.Equal(_p.URL, p.URL, "program URL not equal")
	assert.Len(p.bs, len(_p.bs), "broadcast proxy list length not correct")
	for i := 0; i < len(_p.bs); i++ {
		_bp, bp := _p.bs[i], p.bs[i]

		assert.Equal(_bp.Title, bp.Title, "broadcast proxy %d title not equal", i)
		assert.Equal(_bp.Description, bp.Description, "broadcast proxy %d description not equal", i)
		assert.Equal(_bp.ImageURLs, bp.ImageURLs, "broadcast proxy %d image URLs not equal", i)
		assert.Equal(_bp.URL, bp.URL, "broadcast proxy %d URL not equal", i)
		assert.True(_bp.Date.Equal(bp.Date), "broadcast proxy %d date not equal (%v != %v)", i, _bp.Date, bp.Date)
		assert.Equal(_bp.SubTitle, bp.SubTitle, "broadcast proxy %d sub title not equal", i)
		assert.Equal(_bp.Length, bp.Length, "broadcast proxy %d length not equal", i)
	}
}

var testDataProgramBroadcast = `<!DOCTYPE html>
<html>
<head>
<title>
Alles van Radio Bergeijk kijk je op npo.nl
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
<meta content='Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een keur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.&#x000A; ' name='description'>
<script id='signup-popover' type='text/x-jquery-tmpl'><div class='signup-popover'>
<h3 class='omroep-bold'>Wil je dit programma bewaren als favoriet?</h3>
<p>Maak een eigen account aan en bewaar je favoriete programma’s. Een account zorgt ook voor tips op maat.</p>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280%3Fmedia_type%3Dbroadcast" class="button action-button">Inloggen</a>
<a href="https://mijn.npo.nl/account_aanmaken" class="button action-button">Account aanmaken</a>
</div>
</script>
<meta content='true' name='guest'>



<script type="text/javascript">
//<![CDATA[
npo_cookies.init();
//]]>
</script>

<meta content="authenticity_token" name="csrf-param" />
<meta content="OVLm8/U81apbgHzmBNC8poUAWMcQE3F8TPO/XIkc/cs=" name="csrf-token" />
<meta content="Radio Bergeijk" name="og:title" />
<meta content="http://www.npo.nl/radio-bergeijk/POMS_S_VPRO_396280" name="og:url" />
<meta content="http://images.poms.omroep.nl/image/215303.png" name="og:image" />
<meta content="http://images.poms.omroep.nl/image/193478.png" name="og:image" />
<meta content="http://images.poms.omroep.nl/image/215304.png" name="og:image" />
<meta content="Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een keur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.
 " name="og:description" />
</head>
<body>
<!-- Begin comScore Inline Tag 1.1302.13 -->
<script type="text/plain" class="npo_cc_inline_analytics">
// <![CDATA[
function udm_(e){var t="comScore=",n=document,r=n.cookie,i="",s="indexOf",o="substring",u="length",a=2048,f,l="&ns_",c="&",h,p,d,v,m=window,g=m.encodeURIComponent||escape;if(r[s](t)+1)for(d=0,p=r.split(";"),v=p[u];d<v;d++)h=p[d][s](t),h+1&&(i=c+unescape(p[d][o](h+t[u])));e+=l+"_t="+ +(new Date)+l+"c="+(n.characterSet||n.defaultCharset||"")+"&c8="+g(n.title)+i+"&c7="+g(n.URL)+"&c9="+g(n.referrer),e[u]>a&&e[s](c)>0&&(f=e[o](0,a-8).lastIndexOf(c),e=(e[o](0,f)+l+"cut="+g(e[o](f+1)))[o](0,a)),n.images?(h=new Image,m.ns_p||(ns_p=h),h.src=e):n.write("<","p","><",'img src="',e,'" height="1" width="1" alt="*"',"><","/p",">")};
udm_('http'+(document.location.href.charAt(4)=='s'?'s://sb':'://b')+'.scorecardresearch.com/b?c1=2&c2=17827132&name=npo.programmas.radio-bergeijk.home&npo_ingelogd=no&npo_login_id=geen&ns_site=po-totaal&potag1=npoportal&potag2=programmas&potag3=npo&potag4=npo&potag5=zenderportal&potag6=video&potag7=geen&potag8=site&potag9=site');
// ]]>
</script>
<noscript><p><img src="http://b.scorecardresearch.com/p?c1=2&amp;c2=17827132&amp;name=npo.programmas.radio-bergeijk.home&amp;npo_ingelogd=no&amp;npo_login_id=geen&amp;ns_site=po-totaal&amp;potag1=npoportal&amp;potag2=programmas&amp;potag3=npo&amp;potag4=npo&amp;potag5=zenderportal&amp;potag6=video&amp;potag7=geen&amp;potag8=site&amp;potag9=site" height="1" width="1" alt="*"></p></noscript>
<!-- End comScore Inline Tag -->


<div class="hidden-item-props"><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/" itemprop="url"><span itemprop="title">NPO</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="http://www.npo.nl/a-z" itemprop="url"><span itemprop="title">Series</span></a></div><div itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="http://www.npo.nl/POMS_S_VPRO_396280/POMS_S_VPRO_396280" itemprop="url"><span itemprop="title">Radio Bergeijk</span></a></div></div>
<div class='popup-overlay hide' id='popup-overlay'></div>
<div class='popup-alerts' id='popup-alerts'>
</div>
<div class='page-container shows'>
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
<li><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li class="active"><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
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
<li><a href="http://www.npo.nl/uitzending-gemist" data-show-sub-navigation="programs">Gemist</a></li>
<li><a href="http://www.npo.nl/gids" data-show-sub-navigation="guides">Gids</a></li>
<li class="active"><a href="http://www.npo.nl/a-z" data-show-sub-navigation="shows">A-Z</a></li>
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
 </span></div></div>
<div class='share-modal modal hide fade' data-id='POMS_S_VPRO_396280' id='share-modal' tabindex='-1'>
<div class='modal-header'>
<div class='share-close' data-dismiss='modal' type='button'>
<i class="npo-icon-close close"></i>
</div>
<h2 class='omroep-bold'>Radio Bergeijk</h2>
</div>
<div class='modal-body'><div class='row-fluid'>
<div class='span4 social-buttons'>
<h3>Deel via social media</h3>
<a href="/radio-bergeijk/POMS_S_VPRO_396280/delen/facebook?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280" class="button facebook" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.facebook.radio-bergeijk.POMS_S_VPRO_396280&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph facebook"></span>Facebook</a>
<a href="/radio-bergeijk/POMS_S_VPRO_396280/delen/google?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280" class="button google" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.google.radio-bergeijk.POMS_S_VPRO_396280&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph google"></span>Google</a>
<a href="/radio-bergeijk/POMS_S_VPRO_396280/delen/twitter?url=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280" class="button twitter" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.twitter.radio-bergeijk.POMS_S_VPRO_396280&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph twitter"></span>Twitter</a>
<a href="whatsapp://send?text=Bekijk%20Radio%20Bergeijk:%20http://www.npo.nl/radio-bergeijk/POMS_S_VPRO_396280" class="button whatsapp hide" data-scorecard="{&quot;name&quot;:&quot;programmas.delen.whatsapp.radio-bergeijk.POMS_S_VPRO_396280&quot;}" rel="nofollow" target="_blank"><span class="npo-glyph whatsapp"></span>Whatsapp</a>
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
</div>
<div class='share-path'>
<label>Directe link</label>
<div class='share-action-wrapper' id='share-modal-url-copy-wrapper'>
<div class='share-action-link' id='share-modal-url-copy'>Kopieer link</div>
</div>
<input id="share_url" name="share_url" type="text" value="http://www.npo.nl/radio-bergeijk/POMS_S_VPRO_396280" />
</div>
</div>
</div>
</div>
<div class='email-friend email-block switch-list'>
<h3>E-mail vriend</h3>
<form accept-charset="UTF-8" action="/shares" class="simple_form new_share" id="new_share_" method="post" novalidate="novalidate"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="OVLm8/U81apbgHzmBNC8poUAWMcQE3F8TPO/XIkc/cs=" /></div><div class='row-fluid'>
<div class='span5'>
<input id="id" name="id" type="hidden" value="POMS_S_VPRO_396280" />
<input id="type" name="type" type="hidden" value="show" />
<input id="mid" name="mid" type="hidden" />
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
<a href="/radio-bergeijk/POMS_S_VPRO_396280" class="button gray right" data-dismiss="modal">Annuleren</a>
</div>
</div>
</form>


</div>
</div>
</div>
</div>
<div class='modal-footer'></div>
</div>

<div class='showcase showcase-fluid'>
<div class='showcase-background'>
<div class='header-gradient no-gradient'>
<div class='info no-overlay-for-small-screen'>
<div class='container'><div id='header'>
<div class='row-fluid visible-no-grid-collapse'>
<div class='span-main'>
<div class='row-fluid'>
<div class='broadcaster-logo'><a href="http://www.vpro.nl" target="_blank" title="VPRO"><img alt="Regular_regular_vpro_alt2" class="no-shadow" src="//www-assets.npo.nl/uploads/broadcaster/40/logo/regular_regular_VPRO_alt2.png" /></a></div>
<h1>Radio Bergeijk</h1>
</div>
<div class='row-fluid'>
<div class='span12'>
</div>
</div>
<div class='row-fluid'>
<div class='span4'><img alt="Afbeelding van Radio Bergeijk" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" /></div>
<div class='span8'>
<div class='metadata overflow-description'><span>Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een ke</span><span class="omission">...<a href="#" class="overflow-toggle">Meer</a></span><span class="omission-overflow hide">ur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.
 <a href="#" class="overflow-toggle less">Minder</a></span></div>
</div>
</div>
</div>
<div class='span-sidebar'>
<div class='actions'><div class='action-buttons for-guest'>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280%3Fmedia_type%3Dbroadcast" class="button favorite-button guest-favorite-button"><span class="npo-glyph heart contrast"></span> <span class="glyph-label">Favoriet</span></a>
<a href="/radio-bergeijk/POMS_S_VPRO_396280/delen" class="last button share-modal-button" data-id="POMS_S_VPRO_396280" rel="nofollow"><span class="npo-glyph share contrast"></span> <span class="glyph-label">Delen</span></a>
</div>
</div>
</div>
</div>
<div class='row-fluid visible-with-grid-collapse non-responsive'>
<div class='span8'>
<h1>
Radio Bergeijk
<span class='inactive'>(<span id="broadcaster">VPRO</span>)</span>
</h1>
<h2>Zaterdag om 18:32</h2>
<div class='metadata visible-no-grid-collapse' title='Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een keur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.&#x000A; '>
<p><span>Radio Bergeijk, het radiostation voor Bergeijk, wordt gepresenteerd door ankerman Toon Spoorenberg en zijn beste kennis Peer van Eersel. Met </span><span class="omission">...<a href="#" class="overflow-toggle">Meer</a></span><span class="omission-overflow hide">facilitaire ondersteuning van technicus Tedje van Lieshout ontvangen de radiomannen een keur aan plaatselijke gasten. Of doen verslag van bijzondere culturele gebeurtenissen of particuliere initiatieven. Radio Bergeijk zond uit van 2001 t/m 2007.
 <a href="#" class="overflow-toggle less">Minder</a></span></p>
</div>
<div class='actions visible-no-grid-collapse'><div class='action-buttons for-guest'>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280%3Fmedia_type%3Dbroadcast" class="button favorite-button guest-favorite-button"><span class="npo-glyph heart contrast"></span> <span class="glyph-label">Favoriet</span></a>
<a href="/radio-bergeijk/POMS_S_VPRO_396280/delen" class="last button share-modal-button" data-id="POMS_S_VPRO_396280" rel="nofollow"><span class="npo-glyph share contrast"></span> <span class="glyph-label">Delen</span></a>
</div>
</div>
</div>
<div class='span4'><div class='mobile-buttons single-button'>
<a href="#" class="info-button"><i class="npo-icon-information"></i></a>
<a href="/radio-bergeijk/POMS_S_VPRO_396280/delen" class="last  share-modal-button" data-id="POMS_S_VPRO_396280" rel="nofollow"><span class="npo-glyph share"></span></a>
<a href="https://mijn.npo.nl/inloggen?redirect=http%3A%2F%2Fwww.npo.nl%2Fradio-bergeijk%2FPOMS_S_VPRO_396280%3Fmedia_type%3Dbroadcast" class=" favorite-button guest-favorite-button"><span class="npo-glyph heart"></span></a>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
<div class="page-content"><div class="shadow"></div><div class="container"><div class='row-fluid'>
<div class='span-main show-container' data-source='/radio-bergeijk/POMS_S_VPRO_396280/search'>
<form accept-charset="UTF-8" action="/search" id="program-search-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
<div class='filter-block'>
<div class="broadcast-block switch-block filter "><a href="#" class="omroep-light " data-media-type="broadcast" data-scorecard="{&quot;name&quot;:&quot;programmas.content.afleveringen&quot;,&quot;soft&quot;:true}">Afleveringen <span class="num-found"></span></a></div>
<div class="segment-block switch-block filter last"><a href="#" class="omroep-light inactive" data-media-type="segment" data-scorecard="{&quot;name&quot;:&quot;programmas.content.fragmenten&quot;,&quot;soft&quot;:true}">Fragmenten <span class="num-found"></span></a></div>

<div class='search-input filter with-icon'>
<div class="btn-group filter-dropdown search-dropdown"><a href="#" class="dropdown-toggle" data-target="search-dropdown" data-toggle="dropdown"><i class="npo-icon-search gray"></i></a><ul class="pull- dropdown-menu" filter_type="search" icon="search gray" id="search-dropdown"><li class='search-box'>
<span class="npo-glyph search"></span>
<input class="search-query-box" id="filter-programs" name="filter-programs" placeholder="Zoeken..." type="text" />
</li>
</ul></div></div>
<div class='calendar-input filter with-icon'>
<div class="btn-group filter-dropdown calendar-dropdown"><a href="#" class="dropdown-toggle" data-target="calendar-dropdown" data-toggle="dropdown"><i class="npo-icon-calendar"></i></a><ul class="pull- dropdown-menu" filter_type="calendar" id="calendar-dropdown"><li>
<label for="Van">Van</label>
<input id="program-start" name="program-start" type="text" />
</li>
<li>
<label for="Tot">Tot</label>
<input id="program-end" name="program-end" type="text" />
</li>
</ul></div></div>
</div>
<div class='row-fluid'>
<div class='span12'>
<div class='broadcast-block video-container-block' data-media-type='broadcast' id='broadcasts-block'>
<span class='filter-info'></span>
<div class='content'><div class='search-results' data-num-found='693' data-rows='8' data-start='0'>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk-de-allerlaatste/06-10-2007/POMS_VPRO_396279"><img alt="Afbeelding van Radio Bergeijk : De allerlaatste !" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 25:02</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk-de-allerlaatste/06-10-2007/POMS_VPRO_396279"><h4>
Radio Bergeijk : De allerlaatste !
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Zaterdagavond 6 okt om Half 7 op Radio 1 · Za 6 okt 2007 18:32 · 25 min</h5>
<p class='without-title visible-for-desktop'>Aan het begin van de uitzending wordt Tedje van Lieshout dood aantroffen in zijn werkhok. Een groots gevoel van machteloze moedeloosheid overvalt presen...</p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk-de-een-na-laatste/29-09-2007/POMS_VPRO_396278"><img alt="Afbeelding van Radio Bergeijk: De één na laatste!" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 26:40</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk-de-een-na-laatste/29-09-2007/POMS_VPRO_396278"><h4>
Radio Bergeijk: De één na laatste!
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Zaterdagavond 29 sept om Half 7 op Radio 1 · Za 29 sep 2007 18:32 · 26 min</h5>
<p class='without-title visible-for-desktop'>Hoe kan het nu opeens zo snel gaan! Met de acties tégen Radio Bergeijk.
Toon leent 17 euro van Peer. Krijgt hij volgende week weer terug.
De vader van...</p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/22-09-2007/POMS_VPRO_396277"><img alt="Afbeelding van Radio Bergeijk" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 25:22</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/22-09-2007/POMS_VPRO_396277"><h4>
Radio Bergeijk
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Elke zaterdagavond om Half 7 op Radio 1 · Za 22 sep 2007 18:32 · 25 min</h5>
<p class='without-title visible-for-desktop'>Er hebben zich inmiddels veel meer dan 200 inwoners aangesloten bij het gemeentelijk actiecomité tégen Radio Bergeijk. Toon heeft daarover een 18 pagina...</p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/15-09-2007/POMS_VPRO_396276"><img alt="Afbeelding van Radio Bergeijk" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 22:31</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/15-09-2007/POMS_VPRO_396276"><h4>
Radio Bergeijk
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Elke zaterdagavond om Half 7 op Radio 1 · Za 15 sep 2007 18:32 · 22 min</h5>
<p class='without-title visible-for-desktop'>Toon en Peer zijn vorige week uit GrandCafé Beeks gezet en zijn nu teruggekeerd
in hun oude loods op het industrieterrein. En alsof er niets gebeurd is...</p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/08-09-2007/POMS_VPRO_396275"><img alt="Afbeelding van Radio Bergeijk" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 25:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/08-09-2007/POMS_VPRO_396275"><h4>
Radio Bergeijk
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Elke zaterdagavond om Half 7 op Radio 1 · Za 8 sep 2007 18:32 · 25 min</h5>
<p class='without-title visible-for-desktop'>Deze laatste zomerspeciaal van 2007 wordt ge-evalueerd door cafébaas Jan Beeks.
In de rubriek vacaturebank zoekt uitvaartverzorging van der Stappen een...</p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/01-09-2007/POMS_VPRO_396274"><img alt="Afbeelding van Radio Bergeijk" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 24:46</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/01-09-2007/POMS_VPRO_396274"><h4>
Radio Bergeijk
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Elke zaterdagavond om Half 7 op Radio 1 · Za 1 sep 2007 18:32 · 24 min</h5>
<p class='without-title visible-for-desktop'>De zomerspeciaal van deze week wordt ge-evalueerd door Joke Leermakers.
Overblijfkrachtcentrale zoekt ex-TBS er die zijn talent wil inzetten.
Henk en ...</p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/geen-radio-bergeijk/25-08-2007/POMS_VPRO_396273"><img alt="Afbeelding van Geen Radio Bergeijk" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 28:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/geen-radio-bergeijk/25-08-2007/POMS_VPRO_396273"><h4>
Geen Radio Bergeijk
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>NOS verslag EK finale dameshockey · Za 25 aug 2007 18:32 · 28 min</h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/18-08-2007/POMS_VPRO_396272"><img alt="Afbeelding van Radio Bergeijk" class="preview program-image" data-images="[&quot;http://images.poms.omroep.nl/image/s174/c174x98/215303.png&quot;]" data-toggle="image-skimmer" src="http://images.poms.omroep.nl/image/s174/c174x98/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 28:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/18-08-2007/POMS_VPRO_396272"><h4>
Radio Bergeijk
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Elke zaterdagavond om Half 7 op Radio 1 · Za 18 aug 2007 18:32 · 28 min</h5>
<p class='without-title visible-for-desktop'>De zomerspeciaal van deze week wordt ge-evalueerd door Klaas Groot.
In het negerrubriek De Tamtam negernieuws over het negerboek over de legendarische ...</p>
</a></div>
</div>

</div>
</div>
<div class="more-button"><a href="/radio-bergeijk/POMS_S_VPRO_396280?media_type=broadcast" class="dark-well dark-block with-icon left"><span class="npo-glyph plus"></span><span>Meer afleveringen</span></a></div>
</div>
<div class='segment-block video-container-block' data-media-type='segment' id='segments-block'>
<span class='filter-info'></span>
<div class='content'><div class='search-results' data-num-found='139' data-rows='8' data-start='0'>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/28-12-2001/POMS_VPRO_396637/POMS_VPRO_396638"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 28 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/28-12-2001/POMS_VPRO_396637/POMS_VPRO_396638"><h4>
Radio Bergeijk 00:45 uur 28 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/28-12-2001/POMS_VPRO_396637">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/27-12-2001/POMS_VPRO_396635/POMS_VPRO_396636"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 27 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/27-12-2001/POMS_VPRO_396635/POMS_VPRO_396636"><h4>
Radio Bergeijk 00:45 uur 27 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/27-12-2001/POMS_VPRO_396635">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/26-12-2001/POMS_VPRO_396633/POMS_VPRO_396634"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 26 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/552197.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/26-12-2001/POMS_VPRO_396633/POMS_VPRO_396634"><h4>
Radio Bergeijk 00:45 uur 26 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/26-12-2001/POMS_VPRO_396633">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/25-12-2001/POMS_VPRO_396631/POMS_VPRO_396632"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 25 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/25-12-2001/POMS_VPRO_396631/POMS_VPRO_396632"><h4>
Radio Bergeijk 00:45 uur 25 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/25-12-2001/POMS_VPRO_396631">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/22-12-2001/POMS_VPRO_396629/POMS_VPRO_396630"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 22 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/22-12-2001/POMS_VPRO_396629/POMS_VPRO_396630"><h4>
Radio Bergeijk 00:45 uur 22 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/22-12-2001/POMS_VPRO_396629">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/21-12-2001/POMS_VPRO_396627/POMS_VPRO_396628"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 21 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/21-12-2001/POMS_VPRO_396627/POMS_VPRO_396628"><h4>
Radio Bergeijk 00:45 uur 21 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/21-12-2001/POMS_VPRO_396627">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/20-12-2001/POMS_VPRO_396625/POMS_VPRO_396626"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 20 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/20-12-2001/POMS_VPRO_396625/POMS_VPRO_396626"><h4>
Radio Bergeijk 00:45 uur 20 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/20-12-2001/POMS_VPRO_396625">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>
<div class='list-item non-responsive row-fluid'>
<div class='span4'>
<div class='image-container'>
<a href="/radio-bergeijk/19-12-2001/POMS_VPRO_396623/POMS_VPRO_396624"><img alt="Afbeelding van Radio Bergeijk 00:45 uur 19 Dec 2001" class="segment-image" render_overlay="true" src="http://images.poms.omroep.nl/image/s265/c265x150/215303.png" />
<div class="overlay-icon"><span class="npo-glyph speaker"></span> 60:00</div>
</a></div>
</div>
<div class='span8'>
<a href="/radio-bergeijk/19-12-2001/POMS_VPRO_396623/POMS_VPRO_396624"><h4>
Radio Bergeijk 00:45 uur 19 Dec 2001
<span class='inactive'>(VPRO)</span>
<span class='av-icon'></span>
</h4>
<h5>Fragment uit <a href="/radio-bergeijk/19-12-2001/POMS_VPRO_396623">Radio bergeijk</a></h5>
<p class='without-title visible-for-desktop'></p>
</a></div>
</div>

</div>
</div>
<div class="more-button"><a href="/radio-bergeijk/POMS_S_VPRO_396280?media_type=segment" class="dark-well dark-block with-icon left"><span class="npo-glyph plus"></span><span>Meer fragmenten</span></a></div>
</div>
</div>
</div>
</form>

</div>
</div>
</div></div>
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
