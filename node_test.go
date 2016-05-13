package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type NodeTestSuite struct {
	KongoTestSuite
}

func (s *NodeTestSuite) TestInfoShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	node := &NodeServiceOp{client}

	info, res, err := node.Info()

	s.assert.Nil(info)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *NodeTestSuite) TestInfoShouldRetrieveErrorWhenRequest() {
	info, res, err := s.client.Node.Info()

	s.assert.Nil(info)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *NodeTestSuite) TestInfo() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		{
            "configuration": {
                "admin_api_listen": "0.0.0.0:8001",
                "cassandra": {
                    "consistency": "ONE",
                    "contact_points": [
                        "kong-database:9042"
                    ],
                    "data_centers": {},
                    "keyspace": "kong",
                    "port": 9042,
                    "replication_factor": 1,
                    "replication_strategy": "SimpleStrategy",
                    "ssl": {
                        "enabled": false,
                        "verify": false
                    },
                    "timeout": 5000
                },
                "cluster": {
                    "auto-join": true,
                    "profile": "wan",
                    "ttl_on_failure": 3600
                },
                "cluster_listen": "0.0.0.0:7946",
                "cluster_listen_rpc": "127.0.0.1:7373",
                "custom_plugins": {},
                "dao_config": {
                    "database": "kong",
                    "host": "kong-database",
                    "port": 5432,
                    "user": "kong"
                },
                "database": "postgres",
                "dns_resolver": {
                    "address": "127.0.0.1:8053",
                    "dnsmasq": true,
                    "port": 8053
                },
                "dns_resolvers_available": {
                    "dnsmasq": {
                        "port": 8053
                    },
                    "server": {
                        "address": "8.8.8.8"
                    }
                },
                "memory_cache_size": 128,
                "nginx": "NGINX CONFIGURATION",
                "nginx_working_dir": "/usr/local/kong",
                "pid_file": "/usr/local/kong",
                "plugins": [
                    "ssl"
                ],
                "postgres": {
                    "database": "kong",
                    "host": "kong-database",
                    "port": 5432,
                    "user": "kong"
                },
                "proxy_listen": "0.0.0.0:8000",
                "proxy_listen_ssl": "0.0.0.0:8443",
                "send_anonymous_reports": true
            },
            "hostname": "dd90b6072768",
            "lua_version": "LuaJIT 2.1.0-beta1",
            "plugins": {
                "available_on_server": [
                    "rate-limiting"
                ],
                "enabled_in_cluster": {}
            },
            "tagline": "Welcome to kong",
            "timers": {
                "pending": 4,
                "running": 0
            },
            "version": "0.8.1"
		}`

		fmt.Fprint(w, response)
	})

	info, res, err := s.client.Node.Info()

	s.assert.IsType(&Info{}, info)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Equal("0.0.0.0:8001", info.Configuration.AdminApiListen)
	s.assert.Equal("ONE", info.Configuration.Cassandra.Consistency)
	s.assert.Equal([]string{"kong-database:9042"}, info.Configuration.Cassandra.ContactPoints)
	s.assert.IsType(&DataCenters{}, info.Configuration.Cassandra.DataCenters)
	s.assert.Equal("kong", info.Configuration.Cassandra.Keyspace)
	s.assert.Equal(9042, info.Configuration.Cassandra.Port)
	s.assert.Equal(1, info.Configuration.Cassandra.ReplicationFactor)
	s.assert.Equal("SimpleStrategy", info.Configuration.Cassandra.ReplicationStrategy)
	s.assert.Equal(false, info.Configuration.Cassandra.Ssl.Enabled)
	s.assert.Equal(false, info.Configuration.Cassandra.Ssl.Verify)
	s.assert.Equal(5000, info.Configuration.Cassandra.Timeout)
	s.assert.Equal(true, info.Configuration.Cluster.AutoJoin)
	s.assert.Equal("wan", info.Configuration.Cluster.Profile)
	s.assert.Equal(3600, info.Configuration.Cluster.TtlOnFailure)
	s.assert.Equal("0.0.0.0:7946", info.Configuration.ClusterListen)
	s.assert.Equal("127.0.0.1:7373", info.Configuration.ClusterListenRpc)
	s.assert.IsType(&CustomPlugins{}, info.Configuration.CustomPlugins)
	s.assert.Equal("kong", info.Configuration.DaoConfig.Database)
	s.assert.Equal("kong-database", info.Configuration.DaoConfig.Host)
	s.assert.Equal(5432, info.Configuration.DaoConfig.Port)
	s.assert.Equal("kong", info.Configuration.DaoConfig.User)
	s.assert.Equal("postgres", info.Configuration.Database)
	s.assert.Equal("127.0.0.1:8053", info.Configuration.DnsResolver.Address)
	s.assert.Equal(true, info.Configuration.DnsResolver.DnsMasq)
	s.assert.Equal(8053, info.Configuration.DnsResolver.Port)
	s.assert.Equal(8053, info.Configuration.DnsResolversAvailable.DnsMasq.Port)
	s.assert.Equal("8.8.8.8", info.Configuration.DnsResolversAvailable.Server.Address)
	s.assert.Equal(128, info.Configuration.MemoryCacheSize)
	s.assert.Equal("NGINX CONFIGURATION", info.Configuration.Nginx)
	s.assert.Equal("/usr/local/kong", info.Configuration.NginxWorkingDir)
	s.assert.Equal("/usr/local/kong", info.Configuration.Pidfile)
	s.assert.Equal([]string{"ssl"}, info.Configuration.Plugins)
	s.assert.Equal("kong", info.Configuration.Postgres.Database)
	s.assert.Equal("kong-database", info.Configuration.Postgres.Host)
	s.assert.Equal(5432, info.Configuration.Postgres.Port)
	s.assert.Equal("kong", info.Configuration.Postgres.User)
	s.assert.Equal("0.0.0.0:8000", info.Configuration.ProxyListen)
	s.assert.Equal("0.0.0.0:8443", info.Configuration.ProxyListenSsl)
	s.assert.Equal(true, info.Configuration.SendAnonymousReports)
	s.assert.Equal("dd90b6072768", info.Hostname)
	s.assert.Equal("LuaJIT 2.1.0-beta1", info.LuaVersion)
	s.assert.Equal([]string{"rate-limiting"}, info.Plugins.AvailableOnServer)
	s.assert.IsType(&EnableInCluster{}, info.Plugins.EnableInCluster)
	s.assert.Equal("Welcome to kong", info.Tagline)
	s.assert.Equal(4, info.Timers.Pending)
	s.assert.Equal(0, info.Timers.Running)
	s.assert.Equal("0.8.1", info.Version)
}

func (s *NodeTestSuite) TestStatusShouldRetrieveErrorWhenCreateRequest() {
	client := &Kongo{baseUrl: "%a"}
	node := &NodeServiceOp{client}

	status, res, err := node.Status()

	s.assert.Nil(status)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *NodeTestSuite) TestStatusShouldRetrieveErrorWhenRequest() {
	status, res, err := s.client.Node.Status()

	s.assert.Nil(status)
	s.assert.Nil(res)
	s.assert.Error(err)
}

func (s *NodeTestSuite) TestStatus() {
	s.mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal("GET", r.Method)

		response := `
		{
            "database": {
                "acls": 0,
                "apis": 0,
                "basicauth_credentials": 0,
                "consumers": 0,
                "hmacauth_credentials": 0,
                "jwt_secrets": 0,
                "keyauth_credentials": 0,
                "nodes": 1,
                "oauth2_authorization_codes": 0,
                "oauth2_credentials": 0,
                "oauth2_tokens": 0,
                "plugins": 0,
                "ratelimiting_metrics": 0,
                "response_ratelimiting_metrics": 0
            },
            "server": {
                "connections_accepted": 532,
                "connections_active": 1,
                "connections_handled": 532,
                "connections_reading": 0,
                "connections_waiting": 0,
                "connections_writing": 1,
                "total_requests": 532
            }
		}`

		fmt.Fprint(w, response)
	})

	status, res, err := s.client.Node.Status()

	s.assert.IsType(&Status{}, status)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.Equal(532, status.Server.ConnectionsAccepted)
	s.assert.Equal(1, status.Server.ConnectionsActive)
	s.assert.Equal(532, status.Server.ConnectionsHandled)
	s.assert.Equal(0, status.Server.ConnectionsReading)
	s.assert.Equal(0, status.Server.ConnectionsWaiting)
	s.assert.Equal(1, status.Server.ConnectionsWriting)
	s.assert.Equal(532, status.Server.TotalRequests)

	s.assert.Equal(0, status.Database.Acls)
	s.assert.Equal(0, status.Database.Apis)
	s.assert.Equal(0, status.Database.BasicAuthCredentials)
	s.assert.Equal(0, status.Database.Consumers)
	s.assert.Equal(0, status.Database.HmacAuthCredentials)
	s.assert.Equal(0, status.Database.JwtSecrets)
	s.assert.Equal(0, status.Database.KeyAuthCredentials)
	s.assert.Equal(1, status.Database.Nodes)
	s.assert.Equal(0, status.Database.Oauth2AuthorizationCodes)
	s.assert.Equal(0, status.Database.Oauth2Credentials)
	s.assert.Equal(0, status.Database.Oauth2Tokens)
	s.assert.Equal(0, status.Database.Plugins)
	s.assert.Equal(0, status.Database.RateLimitingMetrics)
	s.assert.Equal(0, status.Database.ResponseRateLimitingMetrics)
}

func TestNodeTestSuite(t *testing.T) {
	suite.Run(t, new(NodeTestSuite))
}
