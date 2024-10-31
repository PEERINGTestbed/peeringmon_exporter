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
		[]string{"prefix", "mux", "upstreams", "available", "origin"},
	)
	upstreams2Gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "indirect_upstreams",
		Help: "upstreams",
	},
		[]string{"prefix", "mux", "upstreams", "available", "origin"},
	)
	ripeStatLGErr = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ripestatlg_err",
		Help: "error count for ripestat lg endpoint",
	})
	possibleHijack = promauto.NewCounter(prometheus.CounterOpts{
		Name: "possible_hijack",
		Help: "upstream mismatch, possible hijack",
	})
	bgpCommunitiesGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "bgp_communities",
		Help: "BGP Communities",
	},
		[]string{"prefix", "city", "mux", "communities"},
	)
)

func (p *Prefix) checkLGState() {
	log.Trace().Str("Prefix", p.prefix).Msg("checking prefix state")
	url := ripestatBase + "/data/looking-glass/data.json?resource=" + p.prefix + "&sourceapp=" + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Fetching ripestat")
		ripeStatLGErr.Inc()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("reading ripestat resp")
		return
	}
	defer resp.Body.Close()

	var ripeStatLookingGlassResp RIPEStatLookingGlassResp
	json.Unmarshal(body, &ripeStatLookingGlassResp)

	if statusCode := ripeStatLookingGlassResp.StatusCode; statusCode != 200 {
		log.Error().Int("status code", statusCode).
			Str("status", ripeStatLookingGlassResp.Status).
			Msg("ripestat(lg) resp status code != 200")
		ripeStatLGErr.Inc()
		return
	}

	availableStr := "y"
	if !p.available {
		availableStr = "n"
	}

	origin := strconv.Itoa(p.origin)

	for _, rrc := range ripeStatLookingGlassResp.Data.Rrcs {
		upstreams := []string{}
		upstreams2 := []string{}
		communities := []string{}

		for _, peer := range rrc.Peers {
			asPathSplit := strings.Split(peer.AsPath, " ")
			upstream := ""
			upstream2 := ""
			offset := 2
			if len(asPathSplit) < offset+1 {
				continue
			}
			pos := 0
			for i, asn := range asPathSplit {
				if asn == origin {
					pos = i - 1
				}
				if pos < 0 {
					continue
				}
			}
			if pos == 0 {
				log.Info().
					Str("prefix", p.prefix).
					Str("asPath", peer.AsPath).
					Str("expected origin", origin).
					Msg("correct origin not found, hijack possible")
				possibleHijack.Inc()
				continue
			}
			upstream = asPathSplit[pos]
			if err != nil {
				log.Error().Err(err).Msg("atoi fail")
				continue
			}
			matched := false
			for _, dbUpstream := range dbUpstreams {
				if strconv.Itoa(dbUpstream.asn) == upstream {
					matched = true
					break
				}
			}
			if !matched {
				//change this back to info after stabalizes
				log.Debug().
					Str("prefix", p.prefix).
					Str("upstream", upstream).
					Str("path", peer.AsPath).
					Int("pos", pos).
					Msg("expected upstream not found, hijack possible")
				possibleHijack.Inc()
				continue
			}
			if !slices.Contains(upstreams, upstream) {
				upstreams = append(upstreams, upstream)
			}

			// second upstream
			if len(asPathSplit) < offset+2 {
				continue
			}
			upstream2 = asPathSplit[len(asPathSplit)-offset-1]
			if err != nil {
				log.Error().Err(err).Msg("atoi fail")
				continue
			}
			if !slices.Contains(upstreams2, upstream2) {
				upstreams2 = append(upstreams2, upstream2)
			}
			communities = append(communities, peer.Community)
		}

		upstreamsGauge.WithLabelValues(
			p.prefix,
			p.pop,
			strings.Join(upstreams, " "),
			availableStr,
			origin,
		).Set(float64(len(upstreams)))

		upstreams2Gauge.WithLabelValues(
			p.prefix,
			p.pop,
			strings.Join(upstreams2, " "),
			availableStr,
			origin,
		).Set(float64(len(upstreams2)))

		communities = slices.Compact(communities)
		for _, e := range communities {
			bgpCommunitiesGauge.WithLabelValues(
				p.prefix,
				rrc.Location,
				p.pop,
				e,
			).Set(1)
		}
	}
}
