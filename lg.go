package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
		[]string{"prefix", "city", "mux", "upstreams"},
	)
	upstreams2Gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "indirect_upstreams",
		Help: "upstreams",
	},
		[]string{"prefix", "city", "mux", "upstreams"},
	)
	//bgpCommunitiesGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
	//	Name: "bgp_communities",
	//	Help: "BGP Communities",
	//},
	//	[]string{"prefix", "city", "mux", "communities"},
	//)
)

func (p *PrefixState) checkLGState() {
	log.Trace().Str("Prefix", p.Prefix).Msg("checking prefix state")
	url := ripestatBase + "/data/looking-glass/data.json?resource=" + p.Prefix + "&sourceapp=" + appId
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

	p.Mu.Lock()
	defer p.Mu.Unlock()

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
			upstreams = append(upstreams, upstream)
			if len(asPathSplit) >= 5 {
				upstream2 = asPathSplit[len(asPathSplit)-5]
				if err != nil {
					log.Err(err).Msg("atoi fail")
				}
			}
			upstreams2 = append(upstreams2, upstream2)
			//communities = append(communities, peer.Community)
		}

		upstreams = slices.Compact(upstreams)
		upstreamsGauge.WithLabelValues(
			p.Prefix,
			rrc.Location,
			prefixes[p.Prefix],
			strings.Join(upstreams, " "),
		).Set(float64(len(upstreams)))

		upstreams2 = slices.Compact(upstreams2)
		upstreams2Gauge.WithLabelValues(
			p.Prefix,
			rrc.Location,
			prefixes[p.Prefix],
			strings.Join(upstreams2, " "),
		).Set(float64(len(upstreams2)))

		//communities = slices.Compact(communities)
		//for _, e := range communities {
		//	bgpCommunitiesGauge.WithLabelValues(
		//		p.Prefix,
		//		rrc.Location,
		//		prefixes[p.Prefix],
		//		e,
		//	).Set(1)
		//}
	}
}
