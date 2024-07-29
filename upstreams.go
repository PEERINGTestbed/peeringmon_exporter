package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

var (
	upstreamGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "defined_upstreams",
		Help: "defined upstreams pulled from the db",
	}, []string{"asn", "name"})
)

func setUpstreamGauge() {
	log.Trace().Msg("setting upstreeams gauge")
	for _, dbUpstream := range dbUpstreams {
		upstreamGauge.WithLabelValues(
			dbUpstream.asn,
			dbUpstream.name,
		).Set(1)
	}
}
