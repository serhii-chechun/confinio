package core

type (
	// Configuration defines all parameters used by the application
	Configuration struct {
		Core struct {
			HTTPEngines []struct {
				Alias            string `json:"alias"`
				ListenAddress    string `json:"listen"`
				ListenAddressTLS string `json:"tls_listen"`
				CertFile         string `json:"tls_cert"`
				KeyFile          string `json:"tls_key"`
				ServerName       string `json:"server_name"`
			} `json:"http_engines"`
		} `json:"core"`
		Servers []struct {
			HTTPEngineAlias   string   `json:"http_engine_alias"`
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
		} `json:"servers"`
		Upstreams []struct {
			Alias       string `json:"alias"`
			DialAddress string `json:"dial_address"`
		}
	}
)
