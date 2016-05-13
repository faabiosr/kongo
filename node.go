package kongo

import (
	"net/http"
)

type NodeService interface {
	Info() (*Info, *http.Response, error)
	Status() (*Status, *http.Response, error)
}

type NodeServiceOp struct {
	client *Kongo
}

type Info struct {
	Configuration *Configuration `json:"configuration, omitempty"`
	Hostname      string         `json:"hostname, omitempty"`
	LuaVersion    string         `json:"lua_version, omitempty"`
	Plugins       *Plugins       `json:"plugins, omitempty"`
	Tagline       string         `json:"tagline, omitempty"`
	Timers        *Timers        `json:"timers, omitempty"`
	Version       string         `json:"version, omitempty"`
}

type Configuration struct {
	AdminApiListen        string                 `json:"admin_api_listen, omitempty"`
	Cassandra             *Cassandra             `json:"cassandra, omitempty"`
	Cluster               *Cluster               `json:"cluster, omitempty"`
	ClusterListen         string                 `json:"cluster_listen, omitempty"`
	ClusterListenRpc      string                 `json:"cluster_listen_rpc, omitempty"`
	CustomPlugins         *CustomPlugins         `json:"custom_plugins, omitempty"`
	DaoConfig             *DaoConfig             `json:"dao_config, omitempty"`
	Database              string                 `json:"database, omitempty"`
	DnsResolver           *DnsResolver           `json:"dns_resolver, omitempty"`
	DnsResolversAvailable *DnsResolversAvailable `json:"dns_resolvers_available, omitempty"`
	MemoryCacheSize       int                    `json:"memory_cache_size, omitempty"`
	Nginx                 string                 `json:"nginx, omitempty"`
	NginxWorkingDir       string                 `json:"nginx_working_dir, omitempty"`
	Pidfile               string                 `json:"pid_file, omitempty"`
	Plugins               []string               `json:"plugins, omitempty"`
	Postgres              *Postgres              `json:"postgres, omitempty"`
	ProxyListen           string                 `json:"proxy_listen, omitempty"`
	ProxyListenSsl        string                 `json:"proxy_listen_ssl, omitempty"`
	SendAnonymousReports  bool                   `json:"send_anonymous_reports, omitempty"`
}

type Cassandra struct {
	Consistency         string       `json:"consistency, omitempty"`
	ContactPoints       []string     `json:"contact_points, omitempty"`
	DataCenters         *DataCenters `json:"data_centers, omitempty"`
	Keyspace            string       `json:"keyspace, omitempty"`
	Port                int          `json:"port, omitempty"`
	ReplicationFactor   int          `json:"replication_factor, omitempty"`
	ReplicationStrategy string       `json:"replication_strategy, omitempty"`
	Ssl                 *Ssl         `json:"ssl, omitempty"`
	Timeout             int          `json:"timeout, omitempty"`
}

type DataCenters struct {
}

type Ssl struct {
	Enabled bool `json:"enabled, omitempty"`
	Verify  bool `json:"verify, omitempty"`
}

type Cluster struct {
	AutoJoin     bool   `json:"auto-join, omitempty"`
	Profile      string `json:"profile, omitempty"`
	TtlOnFailure int    `json:"ttl_on_failure, omitempty"`
}

type CustomPlugins struct {
}

type DaoConfig struct {
	Database string `json:"database, omitempty"`
	Host     string `json:"host, omitempty"`
	Port     int    `json:"port, omitempty"`
	User     string `json:"user, omitempty"`
}

type DnsResolver struct {
	Address string `json:"address, omitempty"`
	DnsMasq bool   `json:"dnsmasq, omitempty"`
	Port    int    `json:"port, omitempty"`
}

type DnsResolversAvailable struct {
	DnsMasq *DnsResolversAvailableDnsMasq `json:"dnsmasq, omitempty"`
	Server  *DnsResolversAvailableServer  `json:"server, omitempty"`
}

type DnsResolversAvailableDnsMasq struct {
	Port int `json:"port, omitempty"`
}

type DnsResolversAvailableServer struct {
	Address string `json:"address, omitempty"`
}

type Postgres struct {
	Database string `json:"database, omitempty"`
	Host     string `json:"host, omitempty"`
	Port     int    `json:"port, omitempty"`
	User     string `json:"user, omitempty"`
}

type Plugins struct {
	AvailableOnServer []string         `json:"available_on_server, omitempty"`
	EnableInCluster   *EnableInCluster `json:"enable_in_cluster, omitempty"`
}

type EnableInCluster struct {
}

type Timers struct {
	Pending int `json:"pending, omitempty"`
	Running int `json:"running, omitempty"`
}

type Status struct {
	Server   *StatusServer   `json:"server, omitempty"`
	Database *StatusDatabase `json:"database, omitempty"`
}

type StatusDatabase struct {
	Acls                        int `json:"acls, omitempty"`
	Apis                        int `json:"apis, omitempty"`
	BasicAuthCredentials        int `json:"basicauth_credentials, omitempty"`
	Consumers                   int `json:"consumers, omitempty"`
	HmacAuthCredentials         int `json:"hmacauth_credentials, omitempty"`
	JwtSecrets                  int `json:"jwt_secrets, omitempty"`
	KeyAuthCredentials          int `json:"keyauth_credentials, omitempty"`
	Nodes                       int `json:"nodes, omitempty"`
	Oauth2AuthorizationCodes    int `json:"oauth2_authorization_codes, omitempty"`
	Oauth2Credentials           int `json:"oauth2_credentials, omitempty"`
	Oauth2Tokens                int `json:"oauth2_tokens, omitempty"`
	Plugins                     int `json:"plugins, omitempty"`
	RateLimitingMetrics         int `json:"ratelimiting_metrics, omitempty"`
	ResponseRateLimitingMetrics int `json:"response_ratelimiting_metrics, omitempty"`
}

type StatusServer struct {
	ConnectionsAccepted int `json:"connections_accepted, omitempty"`
	ConnectionsActive   int `json:"connections_active, omitempty"`
	ConnectionsHandled  int `json:"connections_handled, omitempty"`
	ConnectionsReading  int `json:"connections_reading, omitempty"`
	ConnectionsWaiting  int `json:"connections_waiting, omitempty"`
	ConnectionsWriting  int `json:"connections_writing, omitempty"`
	TotalRequests       int `json:"total_requests, omitempty"`
}

func (n *NodeServiceOp) Info() (*Info, *http.Response, error) {
	resource := "/"

	req, err := n.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	info := new(Info)

	res, err := n.client.Do(req, info)

	if err != nil {
		return nil, res, err
	}

	return info, res, nil
}

func (n *NodeServiceOp) Status() (*Status, *http.Response, error) {
	resource := "/status"

	req, err := n.client.NewRequest("GET", resource, nil)

	if err != nil {
		return nil, nil, err
	}

	status := new(Status)

	res, err := n.client.Do(req, status)

	if err != nil {
		return nil, res, err
	}

	return status, res, nil
}
