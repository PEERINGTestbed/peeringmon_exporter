package main

var prefixes = map[string]string{
	"2804:269c:fe41::/48": "seattle01",
	"2804:269c:fe42::/48": "isi01",
	"2804:269c:fe45::/48": "amsterdam01",
	"2804:269c:fe46::/48": "gatech01",
	"2804:269c:fe47::/48": "ufmg01",
	"2804:269c:fe49::/48": "grnet01",
	"2804:269c:fe50::/48": "uw01",
	"2804:269c:fe51::/48": "wisc01",
	"2804:269c:fe54::/48": "neu01",
	"2804:269c:fe56::/48": "clemson01",
	"2804:269c:fe57::/48": "utah01",
	"2804:269c:fe59::/48": "saopaulo01",
	"2804:269c:fe63::/48": "vtrmiami",
	"2804:269c:fe64::/48": "vtratlanta",
	"2804:269c:fe65::/48": "vtramsterdam",
	"2804:269c:fe66::/48": "vtrtokyo",
	"2804:269c:fe67::/48": "vtrsydney",
	"2804:269c:fe68::/48": "vtrfrankfurt",
	"2804:269c:fe69::/48": "vtrseattle",
	"2804:269c:fe70::/48": "vtrchicago",
	"2804:269c:fe71::/48": "vtrparis",
	"2804:269c:fe72::/48": "vtrsingapore",
	"2804:269c:fe73::/48": "vtrwarsaw",
	"2804:269c:fe74::/48": "vtrnewyork",
	"2804:269c:fe75::/48": "vtrdallas",
	"2804:269c:fe76::/48": "vtrmexico",
	"2804:269c:fe77::/48": "vtrtoronto",
	"2804:269c:fe78::/48": "vtrmadrid",
	"2804:269c:fe79::/48": "vtrstockholm",
	"2804:269c:fe80::/48": "vtrbangalore",
	"2804:269c:fe81::/48": "vtrdelhi",
	"2804:269c:fe82::/48": "vtrlosangelas",
	"2804:269c:fe83::/48": "vtrsilicon",
	"2804:269c:fe84::/48": "vtrlondon",
	"2804:269c:fe85::/48": "vtrmumbai",
	"2804:269c:fe86::/48": "vtrseoul",
	"2804:269c:fe87::/48": "vtrmelbourne",
	"2804:269c:fe88::/48": "vtrsaopaulo",
	"2804:269c:fe89::/48": "vtrjohannesburg",
	"2804:269c:fe90::/48": "vtrosaka",
	"2804:269c:fe91::/48": "vtrsantiago",
	"2804:269c:fe92::/48": "vtrmanchester",
	"2804:269c:fe93::/48": "vtrtelaviv",
	"2804:269c:fe94::/48": "vtrhonolulu",

	//"104.17.224.0/20": "cf valid4",
	//"2606:4700::/44":  "cf valid6",

	//"103.21.244.0/24":     "cf invalid4",
	//"2606:4700:7000::/48": "cf invalid6",
}
