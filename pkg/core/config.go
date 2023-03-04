package core

type (
	// Configuration defines all parameters used by the application
	Configuration struct {
		Servers []struct {
			DomainNames       []string `json:"domain_names"`
			ServerRoot        *string  `json:"server_root"`
			ResourceLocations []struct {
				PatternURL   string  `json:"url_pattern"`
				LocationRoot *string `json:"location_root"`
				Backends     []struct {
					DialAddress       *string `json:"dial_address"`
					UpstreamAlias     *string `json:"upstream_alias"`
					TrafficPercentage *int    `json:"traffic_percentage"`
				} `json:"backends"`
			} `json:"resource_locations"`
			HTTPEngine struct {
				ServerName       string `json:"server_name"`
				ListenAddress    string `json:"listen"`
				ListenAddressTLS string `json:"tls_listen"`
				CertFile         string `json:"tls_cert"`
				KeyFile          string `json:"tls_key"`
			} `json:"http_engine"`
		} `json:"servers"`
		Upstreams []struct {
			Alias       string `json:"alias"`
			DialAddress string `json:"dial_address"`
		}
	}
)
