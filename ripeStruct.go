package main

type RIPEStatResp struct {
	Messages       [][]string    `json:"messages"`
	SeeAlso        []interface{} `json:"see_also"`
	Version        string        `json:"version"`
	DataCallName   string        `json:"data_call_name"`
	DataCallStatus string        `json:"data_call_status"`
	Cached         bool          `json:"cached"`
	Data           struct {
		Visibilities []struct {
			Probe struct {
				City          string  `json:"city"`
				Country       string  `json:"country"`
				Longitude     float64 `json:"longitude"`
				Latitude      float64 `json:"latitude"`
				Name          string  `json:"name"`
				Ipv4PeerCount int     `json:"ipv4_peer_count"`
				Ipv6PeerCount int     `json:"ipv6_peer_count"`
				Ixp           string  `json:"ixp"`
			} `json:"probe"`
			Ipv4FullTablePeersNotSeeing []interface{} `json:"ipv4_full_table_peers_not_seeing"`
			Ipv6FullTablePeersNotSeeing []interface{} `json:"ipv6_full_table_peers_not_seeing"`
			Ipv4FullTablePeerCount      int           `json:"ipv4_full_table_peer_count"`
			Ipv6FullTablePeerCount      int           `json:"ipv6_full_table_peer_count"`
			Ipv4FullTablePeersSeeing    []struct {
				Asn         int    `json:"asn"`
				IP          string `json:"ip"`
				PrefixCount int    `json:"prefix_count"`
			} `json:"ipv4_full_table_peers_seeing"`
			Ipv6FullTablePeersSeeing []struct {
				Asn         int    `json:"asn"`
				IP          string `json:"ip"`
				PrefixCount int    `json:"prefix_count"`
			} `json:"ipv6_full_table_peers_seeing"`
		} `json:"visibilities"`
		Resource        string        `json:"resource"`
		RelatedPrefixes []interface{} `json:"related_prefixes"`
		QueryTime       string        `json:"query_time"`
		LatestTime      string        `json:"latest_time"`
		Include         []string      `json:"include"`
	} `json:"data"`
	QueryID      string `json:"query_id"`
	ProcessTime  int    `json:"process_time"`
	ServerID     string `json:"server_id"`
	BuildVersion string `json:"build_version"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	Time         string `json:"time"`
}
