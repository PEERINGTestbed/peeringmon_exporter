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
)

var (
	prefixStateGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefix_visibility",
		Help: "Visibility of the prefix",
	}, []string{"prefix", "city", "mux", "available", "origin"})
	ripeStatVisibilityErr = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ripestatvis_err",
		Help: "error count for ripestat vis endpoint",
	})
)

func (p *Prefix) checkVisState() {
	log.Trace().Str("Prefix", p.prefix).Msg("checking prefix state")
	prefixStateGauge.Reset()
	url := ripestatBase + "/data/visibility/data.json?data_overload_limit=ignore&include=peers_seeing&resource=" + p.prefix + "&sourceapp=" + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Fetching ripestat")
		ripeStatVisibilityErr.Inc()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("reading ripestat resp")
		return
	}
	defer resp.Body.Close()

	var ripeStatVisibilityResp RIPEStatVisibilityResp
	json.Unmarshal(body, &ripeStatVisibilityResp)

	if statusCode := ripeStatVisibilityResp.StatusCode; statusCode != 200 {
		log.Error().Int("status code", statusCode).
			Str("status", ripeStatVisibilityResp.Status).
			Msg("ripestat(vis) resp status code != 200")
		ripeStatVisibilityErr.Inc()
		return
	}

	ipv6 := strings.Contains(p.prefix, ":")

	availableStr := "y"
	if !p.available {
		availableStr = "n"
	}
	origin := strconv.Itoa(p.origin)

	for _, probe := range ripeStatVisibilityResp.Data.Visibilities {
		var vis float32
		if ipv6 {
			vis = float32(len(probe.Ipv6FullTablePeersSeeing)) /
				float32(probe.Ipv6FullTablePeerCount)
		} else {
			vis = float32(len(probe.Ipv4FullTablePeersSeeing)) /
				float32(probe.Ipv4FullTablePeerCount)

		}
		prefixStateGauge.WithLabelValues(
			p.prefix,
			probe.Probe.City,
			p.pop,
			availableStr,
			origin,
		).Set(float64(vis))
	}
}
