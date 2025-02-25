package main

type RIPEStatVisibilityResp struct {
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

type RIPEStatLookingGlassResp struct {
	Messages       []interface{} `json:"messages"`
	SeeAlso        []interface{} `json:"see_also"`
	Version        string        `json:"version"`
	DataCallName   string        `json:"data_call_name"`
	DataCallStatus string        `json:"data_call_status"`
	Cached         bool          `json:"cached"`
	Data           struct {
		Rrcs []struct {
			Rrc      string `json:"rrc"`
			Location string `json:"location"`
			Peers    []struct {
				AsnOrigin   string `json:"asn_origin"`
				AsPath      string `json:"as_path"`
				Community   string `json:"community"`
				LastUpdated string `json:"last_updated"`
				Prefix      string `json:"prefix"`
				Peer        string `json:"peer"`
				Origin      string `json:"origin"`
				NextHop     string `json:"next_hop"`
				LatestTime  string `json:"latest_time"`
			} `json:"peers"`
		} `json:"rrcs"`
		QueryTime  string `json:"query_time"`
		LatestTime string `json:"latest_time"`
		Parameters struct {
			Resource      string      `json:"resource"`
			LookBackLimit int         `json:"look_back_limit"`
			Cache         interface{} `json:"cache"`
		} `json:"parameters"`
	} `json:"data"`
	QueryID      string `json:"query_id"`
	ProcessTime  int    `json:"process_time"`
	ServerID     string `json:"server_id"`
	BuildVersion string `json:"build_version"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	Time         string `json:"time"`
}

type risResponse struct {
	Messages [][]string `json:"messages"`
	SeeAlso  []string   `json:"see_also"`
	Version  string     `json:"version"`
	DataCall struct {
		Name   string `json:"data_call_name"`
		Status string `json:"data_call_status"`
	} `json:"data_call"`
	Cached bool `json:"cached"`
	Data   struct {
		Peers struct {
			RRC00 []risPeer `json:"rrc00"`
			RRC01 []risPeer `json:"rrc01"`
			RRC03 []risPeer `json:"rrc03"`
			RRC04 []risPeer `json:"rrc04"`
			RRC05 []risPeer `json:"rrc05"`
			RRC06 []risPeer `json:"rrc06"`
			RRC07 []risPeer `json:"rrc07"`
			RRC10 []risPeer `json:"rrc10"`
			RRC11 []risPeer `json:"rrc11"`
			RRC13 []risPeer `json:"rrc13"`
			RRC14 []risPeer `json:"rrc14"`
			RRC15 []risPeer `json:"rrc15"`
			RRC16 []risPeer `json:"rrc16"`
			RRC18 []risPeer `json:"rrc18"`
			RRC19 []risPeer `json:"rrc19"`
			RRC20 []risPeer `json:"rrc20"`
			RRC21 []risPeer `json:"rrc21"`
			RRC22 []risPeer `json:"rrc22"`
			RRC23 []risPeer `json:"rrc23"`
			RRC24 []risPeer `json:"rrc24"`
			RRC25 []risPeer `json:"rrc25"`
			RRC26 []risPeer `json:"rrc26"`
		} `json:"peers"`
		LatestTime   string        `json:"latest_time"`
		EarliestTime string        `json:"earliest_time"`
		Parameters   risParameters `json:"parameters"`
	} `json:"data"`
	QueryID      string `json:"query_id"`
	ProcessTime  int    `json:"process_time"`
	ServerID     string `json:"server_id"`
	BuildVersion string `json:"build_version"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	Time         string `json:"time"`
}

// Peer represents an individual peer in the RIS system
type risPeer struct {
	ASN           string `json:"asn"`
	IP            string `json:"ip"`
	V4PrefixCount int    `json:"v4_prefix_count"`
	V6PrefixCount int    `json:"v6_prefix_count"`
}

// Parameters represents the query parameters
type risParameters struct {
	QueryTime string      `json:"query_time"`
	Cache     interface{} `json:"cache"`
}
