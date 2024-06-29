package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const appID = "PEERINGMON-DEV"

type PrefixState struct {
	Prefix string
	State  map[string]int
	Mu     sync.Mutex
}

func (p *PrefixState) checkState() {
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

	// compare values
	ipv6 := strings.Contains(p.Prefix, ":")

	p.Mu.Lock()
	defer p.Mu.Unlock()

	for _, probe := range ripeStatResp.Data.Visibilities {
		var vis int
		if ipv6 {
			vis = probe.Ipv6FullTablePeerCount /
				len(probe.Ipv6FullTablePeersSeeing)
		} else {
			vis = probe.Ipv4FullTablePeerCount /
				len(probe.Ipv4FullTablePeersSeeing)
		}
		p.State[probe.Probe.City] = vis
		fmt.Println(probe.Probe.City)
		fmt.Println(vis)
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	a := &PrefixState{Prefix: "2804:269c:fe53::/48", State: make(map[string]int)}
	b := &PrefixState{Prefix: "2804:269c:fe56::/48", State: make(map[string]int)}
	a.checkState()
	b.checkState()

	// init exporter
	// init struct with prefilled ranges
}
