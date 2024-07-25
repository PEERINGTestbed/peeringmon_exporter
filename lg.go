package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
)

var (
	upstreamsGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "direct_upstreams",
		Help: "upstreams",
	},
		[]string{"prefix", "city", "mux", "upstreams", "available", "origin"},
	)
	upstreams2Gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "indirect_upstreams",
		Help: "upstreams",
	},
		[]string{"prefix", "city", "mux", "upstreams", "available", "origin"},
	)
	//bgpCommunitiesGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
	//	Name: "bgp_communities",
	//	Help: "BGP Communities",
	//},
	//	[]string{"prefix", "city", "mux", "communities"},
	//)
)

func (p *Prefix) checkLGState() {
	log.Trace().Str("Prefix", p.prefix).Msg("checking prefix state")
	url := ripestatBase + "/data/looking-glass/data.json?resource=" + p.prefix + "&sourceapp=" + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Fetching ripestat")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("reading ripestat resp")
	}
	defer resp.Body.Close()

	var ripeStatLookingGlassResp RIPEStatLookingGlassResp
	json.Unmarshal(body, &ripeStatLookingGlassResp)

	if statusCode := ripeStatLookingGlassResp.StatusCode; statusCode != 200 {
		log.Error().Int("status code", statusCode).
			Str("status", ripeStatLookingGlassResp.Status).
			Msg("ripestat(lg) resp status code != 200")
	}

	availableStr := "y"
	if !p.available {
		availableStr = "n"
	}

	origin := strconv.Itoa(p.origin)

	for _, rrc := range ripeStatLookingGlassResp.Data.Rrcs {
		upstreams := []string{}
		upstreams2 := []string{}
		//communities := []string{}

		for _, peer := range rrc.Peers {
			asPathSplit := strings.Split(peer.AsPath, " ")
			upstream := ""
			upstream2 := ""
			if len(asPathSplit) >= 4 {
				upstream = asPathSplit[len(asPathSplit)-4]
				if err != nil {
					log.Err(err).Msg("atoi fail")
				}
			}
			if len(asPathSplit) >= 5 {
				upstream2 = asPathSplit[len(asPathSplit)-5]
				if err != nil {
					log.Err(err).Msg("atoi fail")
				}
			}
			if !slices.Contains(upstreams, upstream) {
				upstreams = append(upstreams, upstream)
			}
			if !slices.Contains(upstreams2, upstream2) {
				upstreams2 = append(upstreams2, upstream2)
			}
			//communities = append(communities, peer.Community)
		}

		upstreamsGauge.WithLabelValues(
			p.prefix,
			rrc.Location,
			p.pop,
			strings.Join(upstreams, " "),
			availableStr,
			origin,
		).Set(float64(len(upstreams)))

		upstreams2Gauge.WithLabelValues(
			p.prefix,
			rrc.Location,
			p.pop,
			strings.Join(upstreams2, " "),
			availableStr,
			origin,
		).Set(float64(len(upstreams2)))

		//communities = slices.Compact(communities)
		//for _, e := range communities {
		//	bgpCommunitiesGauge.WithLabelValues(
		//		p.prefix,
		//		rrc.Location,
		//		prefixes[p.prefix],
		//		e,
		//	).Set(1)
		//}
	}
}
