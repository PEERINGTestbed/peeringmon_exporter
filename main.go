package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	appID = "PEERINGMON-DEV"
	port  = ":2112"
)

var (
	prefixStateGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "prefix_visibility",
		Help: "Visibility of the prefix",
	}, []string{"prefix", "city"})
)

type PrefixState struct {
	Prefix string
	State  map[string]float32
	Mu     sync.Mutex
}

func (p *PrefixState) checkState() {
	log.Trace().Str("Prefix", p.Prefix).Msg("checking prefix state")
	url := "https://stat.ripe.net/data/visibility/data.json?data_overload_limit=ignore&include=peers_seeing&resource=" + p.Prefix + "&sourceapp=" + appID
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Fetching ripestat")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("reading ripestat resp")
	}
	defer resp.Body.Close()

	var ripeStatResp RIPEStatResp
	json.Unmarshal(body, &ripeStatResp)

	ipv6 := strings.Contains(p.Prefix, ":")

	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, probe := range ripeStatResp.Data.Visibilities {
		var vis float32
		if ipv6 {
			vis = float32(len(probe.Ipv6FullTablePeersSeeing)) /
				float32(probe.Ipv6FullTablePeerCount)
		} else {
			vis = float32(len(probe.Ipv4FullTablePeersSeeing)) /
				float32(probe.Ipv4FullTablePeerCount)

		}
		p.State[probe.Probe.City] = vis
		prefixStateGauge.WithLabelValues(p.Prefix, probe.Probe.City).Set(float64(vis))
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Info().Msg("Starting PEERINGMON Exporter")

	prefixes := []string{
		"2804:269c:fe53::/48",
		"2804:269c:fe56::/48",
		"2606:4700:7000::/48",
		"198.8.58.0/24",
	}

	prefixStates := make([]*PrefixState, len(prefixes))
	for i, prefix := range prefixes {
		prefixStates[i] = &PrefixState{Prefix: prefix, State: make(map[string]float32)}
	}

	for _, ps := range prefixStates {
		ps.checkState()
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			for _, ps := range prefixStates {
				ps.checkState()
			}
		}
	}()

	log.Info().Msg("Starting exporter on port " + port)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Error().Err(err).Msg("Failed on http listening")
	}
}
