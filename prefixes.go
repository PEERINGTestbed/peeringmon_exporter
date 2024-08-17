package main

type Prefix struct {
	prefix    string
	pop       string
	available bool
	origin    int
}

type Upstream struct {
	name string
	asn  int
}

var dbUpstreams = []Upstream{
	Upstream{"Northeastern University", 156},
	Upstream{"FABRIC Testbed", 398900},
	Upstream{"GRNet", 5408},
	Upstream{"Bit BV", 12859},
	Upstream{"Netwerkvereniging Coloclue", 8283},
	Upstream{"Stony Brook University", 5719},
	Upstream{"Clemson University", 12148},
	Upstream{"Utah Education Network", 210},
	Upstream{"Georgia Institute of Technology", 2637},
	Upstream{"University of Wisconsin - Madison", 3128},
	Upstream{"Rede Nacional de Ensino e Pesquisa (RNP)", 1916},
	Upstream{"Cornell University", 26},
	Upstream{"psg.com RGNet", 3130},
	Upstream{"Los Nettos Regional Network", 226},
	Upstream{"UW at PNW GigaPoP", 101},
	Upstream{"vultr", 20473},
	Upstream{"HE", 6939},
	Upstream{"PEERING", 47065},
	Upstream{"1299 cf", 1299},
}

var monitorState = []Prefix{
	Prefix{prefix: "2804:269c:fe01::/48", pop: "seattle01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe02::/48", pop: "isi01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe03::/48", pop: "cornell01", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe04::/48", pop: "phoenix01", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe05::/48", pop: "amsterdam01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe06::/48", pop: "gatech01", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe07::/48", pop: "ufmg01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe09::/48", pop: "grnet01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe0a::/48", pop: "uw01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe0b::/48", pop: "wisc01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe0c::/48", pop: "usc01", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe0d::/48", pop: "dallas01", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe0e::/48", pop: "neu01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe0f::/48", pop: "sbu01", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe10::/48", pop: "clemson01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe11::/48", pop: "utah01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe13::/48", pop: "saopaulo01", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe14::/48", pop: "fabwash", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe16::/48", pop: "fabstar", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe17::/48", pop: "vtrmiami", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe18::/48", pop: "vtratlanta", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe19::/48", pop: "vtramsterdam", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe1a::/48", pop: "vtrtokyo", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe1b::/48", pop: "vtrsydney", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe1c::/48", pop: "vtrfrankfurt", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe1d::/48", pop: "vtrseattle", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe1e::/48", pop: "vtrchicago", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe1f::/48", pop: "vtrparis", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe20::/48", pop: "vtrsingapore", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe21::/48", pop: "vtrwarsaw", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe22::/48", pop: "vtrnewyork", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe23::/48", pop: "vtrdallas", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe24::/48", pop: "vtrmexico", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe25::/48", pop: "vtrtoronto", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe26::/48", pop: "vtrmadrid", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe27::/48", pop: "vtrstockholm", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe28::/48", pop: "vtrbangalore", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe29::/48", pop: "vtrdelhi", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe2a::/48", pop: "vtrlosangelas", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe2b::/48", pop: "vtrsilicon", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe2c::/48", pop: "vtrlondon", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe2d::/48", pop: "vtrmumbai", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe2e::/48", pop: "vtrseoul", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe2f::/48", pop: "vtrmelbourne", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe30::/48", pop: "vtrsaopaulo", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe31::/48", pop: "vtrjohannesburg", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe32::/48", pop: "vtrosaka", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe33::/48", pop: "vtrsantiago", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe34::/48", pop: "vtrmanchester", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe35::/48", pop: "vtrtelaviv", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe36::/48", pop: "vtrhonolulu", available: true, origin: 47065},
	Prefix{prefix: "2804:269c:fe37::/48", pop: "cfuseast1", available: false, origin: 47065},
	Prefix{prefix: "2804:269c:fe38::/48", pop: "vtrezri1", available: false, origin: 47065},

	Prefix{prefix: "2804:269c:fe41::/48", pop: "seattle01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe42::/48", pop: "isi01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe43::/48", pop: "cornell01", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe44::/48", pop: "phoenix01", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe45::/48", pop: "amsterdam01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe46::/48", pop: "gatech01", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe47::/48", pop: "ufmg01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe49::/48", pop: "grnet01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe4a::/48", pop: "uw01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe4b::/48", pop: "wisc01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe4c::/48", pop: "usc01", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe4d::/48", pop: "dallas01", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe4e::/48", pop: "neu01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe4f::/48", pop: "sbu01", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe50::/48", pop: "clemson01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe51::/48", pop: "utah01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe53::/48", pop: "saopaulo01", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe54::/48", pop: "fabwash", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe56::/48", pop: "fabstar", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe57::/48", pop: "vtrmiami", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe58::/48", pop: "vtratlanta", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe59::/48", pop: "vtramsterdam", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe5a::/48", pop: "vtrtokyo", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe5b::/48", pop: "vtrsydney", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe5c::/48", pop: "vtrfrankfurt", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe5d::/48", pop: "vtrseattle", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe5e::/48", pop: "vtrchicago", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe5f::/48", pop: "vtrparis", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe60::/48", pop: "vtrsingapore", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe61::/48", pop: "vtrwarsaw", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe62::/48", pop: "vtrnewyork", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe63::/48", pop: "vtrdallas", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe64::/48", pop: "vtrmexico", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe65::/48", pop: "vtrtoronto", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe66::/48", pop: "vtrmadrid", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe67::/48", pop: "vtrstockholm", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe68::/48", pop: "vtrbangalore", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe69::/48", pop: "vtrdelhi", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe6a::/48", pop: "vtrlosangelas", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe6b::/48", pop: "vtrsilicon", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe6c::/48", pop: "vtrlondon", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe6d::/48", pop: "vtrmumbai", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe6e::/48", pop: "vtrseoul", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe6f::/48", pop: "vtrmelbourne", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe70::/48", pop: "vtrsaopaulo", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe71::/48", pop: "vtrjohannesburg", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe72::/48", pop: "vtrosaka", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe73::/48", pop: "vtrsantiago", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe74::/48", pop: "vtrmanchester", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe75::/48", pop: "vtrtelaviv", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe76::/48", pop: "vtrhonolulu", available: true, origin: 61574},
	Prefix{prefix: "2804:269c:fe77::/48", pop: "cfuseast1", available: false, origin: 61574},
	Prefix{prefix: "2804:269c:fe78::/48", pop: "vtrezri1", available: false, origin: 61574},

	Prefix{prefix: "104.17.224.0/20", pop: "cf valid4", available: true, origin: 13335},
	Prefix{prefix: "2606:4700::/44", pop: "cf valid6", available: true, origin: 13335},

	Prefix{prefix: "103.21.244.0/24", pop: "cf invalid4", available: false, origin: 13335},
	Prefix{prefix: "2606:4700:7000::/48", pop: "cf invalid6", available: false, origin: 13335},
}