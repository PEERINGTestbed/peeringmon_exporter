package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

var (
	ripeRISPeerRouteCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "rispeer",
		Help: "RIPE RIS Peer Route Count",
	}, []string{"ip", "asn", "rrc", "stack"})
	ripeRISPeerErr = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rispeer_err",
		Help: "error count for ripe ris peer endpoint",
	})
)

func fetchRisPeer() {
	log.Trace().Msg("checking ris peers state")
	ripeRISPeerRouteCount.Reset()
	url := ripestatBase + "/data/ris-peers/data.json?sourceapp=" + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Fetching ripestat")
		ripeRISPeerErr.Inc()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("reading ripestat resp")
		return
	}
	defer resp.Body.Close()

	var response risResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Error().Err(err).Msg("err unmarhsalling resp")
		return
	}

	// Use reflection to iterate through all RRCs
	peersValue := reflect.ValueOf(response.Data.Peers)
	peersType := peersValue.Type()

	for i := 0; i < peersValue.NumField(); i++ {
		field := peersType.Field(i)
		rrc := strings.ToLower(field.Name) // Get RRC name from struct field
		peers := peersValue.Field(i).Interface().([]risPeer)

		// Process peers for this RRC
		for _, peer := range peers {
			// Set IPv4 prefix count
			if peer.V4PrefixCount > 0 {
				ripeRISPeerRouteCount.With(prometheus.Labels{
					"ip":    peer.IP,
					"asn":   peer.ASN,
					"rrc":   rrc,
					"stack": "ipv4",
				}).Set(float64(peer.V4PrefixCount))
			}

			// Set IPv6 prefix count
			if peer.V6PrefixCount > 0 {
				ripeRISPeerRouteCount.With(prometheus.Labels{
					"ip":    peer.IP,
					"asn":   peer.ASN,
					"rrc":   rrc,
					"stack": "ipv6",
				}).Set(float64(peer.V6PrefixCount))
			}
		}
	}
}
