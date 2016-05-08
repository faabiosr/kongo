package kongo

import (
	"net/http"
)

type NodeService interface {
	Info() (*NodeInfo, *http.Response, error)
	Status() (*NodeStatus, *http.Response, error)
}

type NodeServiceOp struct {
	client *Kongo
}

type NodeInfo struct {
	Configuration *NodeInfoConfiguration `json:"configuration"`
	Hostname      string                 `json:"hostname"`
	LuaVersion    string                 `json:"lua_version"`
	Plugins       *NodeInfoPlugins       `json:"plugins"`
	Tagline       string                 `json:"tagline"`
	Timers        *NodeInfoTimers        `json:"timers"`
	Version       string                 `json:"version"`
}

type NodeInfoConfiguration struct {
	AdminApiListen        string                                      `json:"admin_api_listen"`
	Cassandra             *NodeInfoConfigurationCassandra             `json:"cassandra"`
	Cluster               *NodeInfoConfigurationCluster               `json:"cluster"`
	ClusterListen         string                                      `json:"cluster_listen"`
	ClusterListenRpc      string                                      `json:"cluster_listen_rpc"`
	CustomPlugins         *NodeInfoConfigurationCustomPlugins         `json:"custom_plugins"`
	DaoConfig             *NodeInfoConfigurationDaoConfig             `json:"dao_config"`
	Database              string                                      `json:"database"`
	DnsResolver           *NodeInfoConfigurationDnsResolver           `json:"dns_resolver"`
	DnsResolversAvailable *NodeInfoConfigurationDnsResolversAvailable `json:"dns_resolvers_available"`
	MemoryCacheSize       int                                         `json:"memory_cache_size"`
	Nginx                 string                                      `json:"nginx"`
	NginxWorkingDir       string                                      `json:"nginx_working_dir"`
	Pidfile               string                                      `json:"pid_file"`
	Plugins               []string                                    `json:"plugins"`
	Postgres              *NodeInfoConfigurationPostgres              `json:"postgres"`
	ProxyListen           string                                      `json:"proxy_listen"`
	ProxyListenSsl        string                                      `json:"proxy_listen_ssl"`
	SendAnonymousReports  bool                                        `json:"send_anonymous_reports"`
}

type NodeInfoConfigurationCassandra struct {
	Consistency         string                                     `json:"consistency"`
	ContactPoints       []string                                   `json:"contact_points"`
	DataCenters         *NodeInfoConfigurationCassandraDataCenters `json:"data_centers"`
	Keyspace            string                                     `json:"keyspace"`
	Port                int                                        `json:"port"`
	ReplicationFactor   int                                        `json:"replication_factor"`
	ReplicationStrategy string                                     `json:"replication_strategy"`
	Ssl                 *NodeInfoConfigurationCassandraSsl         `json:"ssl"`
	Timeout             int                                        `json:"timeout"`
}

type NodeInfoConfigurationCassandraDataCenters struct {
}

type NodeInfoConfigurationCassandraSsl struct {
	Enabled bool `json:"enabled"`
	Verify  bool `json:"verify"`
}

type NodeInfoConfigurationCluster struct {
	AutoJoin     bool   `json:"auto-join"`
	Profile      string `json:"profile"`
	TtlOnFailure int    `json:"ttl_on_failure"`
}

type NodeInfoConfigurationCustomPlugins struct {
}

type NodeInfoConfigurationDaoConfig struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
}

type NodeInfoConfigurationDnsResolver struct {
	Address string `json:"address"`
	DnsMasq bool   `json:"dnsmasq"`
	Port    int    `json:"port"`
}

type NodeInfoConfigurationDnsResolversAvailable struct {
	DnsMasq *NodeInfoConfigurationDnsResolversAvailableDnsMasq `json:"dnsmasq"`
	Server  *NodeInfoConfigurationDnsResolversAvailableServer  `json:"server"`
}

type NodeInfoConfigurationDnsResolversAvailableDnsMasq struct {
	Port int `json:"port"`
}

type NodeInfoConfigurationDnsResolversAvailableServer struct {
	Address string `json:"address"`
}

type NodeInfoConfigurationPostgres struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
}

type NodeInfoPlugins struct {
	AvailableOnServer []string                        `json:"available_on_server"`
	EnableInCluster   *NodeInfoPluginsEnableInCluster `json:"enable_in_cluster"`
}

type NodeInfoPluginsEnableInCluster struct {
}

type NodeInfoTimers struct {
	Pending int `json:"pending"`
	Running int `json:"running"`
}

type NodeStatus struct {
	Server   *NodeStatusServer   `json:"server"`
	Database *NodeStatusDatabase `json:"database"`
}

type NodeStatusDatabase struct {
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

type NodeStatusServer struct {
	ConnectionsAccepted int `json:"connections_accepted, omitempty"`
	ConnectionsActive   int `json:"connections_active, omitempty"`
	ConnectionsHandled  int `json:"connections_handled, omitempty"`
	ConnectionsReading  int `json:"connections_reading, omitempty"`
	ConnectionsWaiting  int `json:"connections_waiting, omitempty"`
	ConnectionsWriting  int `json:"connections_writing, omitempty"`
	TotalRequests       int `json:"total_requests, omitempty"`
}

func (n *NodeServiceOp) Info() (*NodeInfo, *http.Response, error) {
	resource := "/"

	req, err := n.client.NewRequest("GET", resource)

	if err != nil {
		return nil, nil, err
	}

	nodeInfo := new(NodeInfo)

	res, err := n.client.Do(req, nodeInfo)

	if err != nil {
		return nil, res, err
	}

	return nodeInfo, res, nil
}

func (n *NodeServiceOp) Status() (*NodeStatus, *http.Response, error) {
	resource := "/status"

	req, err := n.client.NewRequest("GET", resource)

	if err != nil {
		return nil, nil, err
	}

	nodeStatus := new(NodeStatus)

	res, err := n.client.Do(req, nodeStatus)

	if err != nil {
		return nil, res, err
	}

	return nodeStatus, res, nil
}
