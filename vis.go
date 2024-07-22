package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

var (
	prefixStateGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefix_visibility",
		Help: "Visibility of the prefix",
	}, []string{"prefix", "city", "mux"})
)

func (p *PrefixState) checkVisState() {
	log.Trace().Str("Prefix", p.Prefix).Msg("checking prefix state")
	url := ripestatBase + "/data/visibility/data.json?data_overload_limit=ignore&include=peers_seeing&resource=" + p.Prefix + "&sourceapp=" + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Fetching ripestat")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("reading ripestat resp")
	}
	defer resp.Body.Close()

	var ripeStatVisibilityResp RIPEStatVisibilityResp
	json.Unmarshal(body, &ripeStatVisibilityResp)

	if statusCode := ripeStatVisibilityResp.StatusCode; statusCode != 200 {
		log.Error().Int("status code", statusCode).
			Str("status", ripeStatVisibilityResp.Status).
			Msg("ripestat(vis) resp status code != 200")
	}

	ipv6 := strings.Contains(p.Prefix, ":")

	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, probe := range ripeStatVisibilityResp.Data.Visibilities {
		var vis float32
		if ipv6 {
			vis = float32(len(probe.Ipv6FullTablePeersSeeing)) /
				float32(probe.Ipv6FullTablePeerCount)
		} else {
			vis = float32(len(probe.Ipv4FullTablePeersSeeing)) /
				float32(probe.Ipv4FullTablePeerCount)

		}
		p.State[probe.Probe.City] = vis
		prefixStateGauge.WithLabelValues(
			p.Prefix,
			probe.Probe.City,
			prefixes[p.Prefix],
		).Set(float64(vis))
	}
}
