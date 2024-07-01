package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var prefixes = []string{
	// PEERING v6
	"2804:269c:fe41::/48",
	"2804:269c:fe42::/48",
	"2804:269c:fe44::/48",
	"2804:269c:fe45::/48",
	"2804:269c:fe47::/48",
	"2804:269c:fe50::/48",
	"2804:269c:fe51::/48",
	"2804:269c:fe53::/48",
	"2804:269c:fe56::/48",
	"2804:269c:fe57::/48",
	"2804:269c:fe58::/48",
	"2804:269c:fe59::/48",
	"2804:269c:fe5a::/48",
	"2804:269c:fe5b::/48",
	"2804:269c:fe5c::/48",
	"2804:269c:fe5d::/48",
	"2804:269c:fe5e::/48",
	"2804:269c:fe5f::/48",
	"2804:269c:fe60::/48",
	"2804:269c:fe61::/48",
	"2804:269c:fe62::/48",
	"2804:269c:fe63::/48",
	"2804:269c:fe64::/48",
	"2804:269c:fe65::/48",
	"2804:269c:fe66::/48",
	"2804:269c:fe67::/48",
	"2804:269c:fe68::/48",
	"2804:269c:fe69::/48",
	"2804:269c:fe6a::/48",
	"2804:269c:fe6b::/48",
	"2804:269c:fe6c::/48",
	"2804:269c:fe6d::/48",
	"2804:269c:fe6e::/48",
	"2804:269c:fe6f::/48",
	"2804:269c:fe70::/48",
	"2804:269c:fe71::/48",
	"2804:269c:fe72::/48",
	"2804:269c:fe73::/48",
	"2804:269c:fe74::/48",
	"2804:269c:fe76::/48",

	// isbgpsafeyet.com valid
	"104.17.224.0/20",
	"2606:4700::/44",

	// isbgpsafeyet.com invalid
	"103.21.244.0/24",
	"2606:4700:7000::/48",
}

var prefixStates = []*PrefixState{}

var port int
var appId string

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
	url := "https://stat.ripe.net/data/visibility/data.json?data_overload_limit=ignore&include=peers_seeing&resource=" + p.Prefix + "&sourceapp=" + appId
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

func updateStates() {
	log.Debug().Msg("Updating Prefixes")
	for _, ps := range prefixStates {
		go ps.checkState()
	}
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	flag.StringVar(&appId, "appid", "exporter", "provide a unique identifier to every data call")
	flag.IntVar(&port, "port", 2112, "port")
}

func main() {
	flag.Parse()

	log.Info().
		Str("appID", appId).
		Str("Data Source", "RIPE RIS via RIPEstat API").
		Msg("Starting PEERINGMON Exporter")

	for _, prefix := range prefixes {
		prefixStates = append(prefixStates, &PrefixState{
			Prefix: prefix,
			State:  make(map[string]float32),
		})
	}

	updateStates()

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			updateStates()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr: ":" + strconv.Itoa(port),
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()
	log.Info().Int("port", port).Msg("Started exporter")

	<-done
	log.Info().Msg("Stopping")
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Err(err).Msg("Failed to gracefully stop server")
	}
	log.Info().Msg("Graceful Shutdown Successful, bye")
}
