package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port int
var appId string
var debug bool
var jsonLog bool

const ripestatBase = "https://stat.ripe.net"

func updateStates() {
	log.Debug().Msg("Updating Prefixes")

	upstreamsGauge.Reset()
	upstreams2Gauge.Reset()
	bgpCommunitiesGauge.Reset()

	for _, prefix := range monitorState {
		prefix.checkVisState()
		prefix.checkLGState()
	}
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	flag.StringVar(&appId, "appid", "exporter", "provide a unique identifier to every data call")
	flag.IntVar(&port, "port", 2112, "port")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.BoolVar(&jsonLog, "json", false, "json logging")
}

func main() {
	flag.Parse()

	if !jsonLog {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug log enabled")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Info().
		Str("appID", appId).
		Str("Data Source", "RIPE RIS via RIPEstat API").
		Msg("Starting PEERINGMON Exporter")

	setUpstreamGauge()
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
