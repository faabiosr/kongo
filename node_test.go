package kongo

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type NodeTestSuite struct {
	BaseTestSuite
}

func (s *NodeTestSuite) TestInfoReturnsHttpError() {
	s.mux.HandleFunc(nodeInfoResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.Node.Info()

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *NodeTestSuite) TestInfo() {
	s.mux.HandleFunc(nodeInfoResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/node_info_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	info, res, err := s.client.Node.Info()

	s.assert.IsType(&NodeInfo{}, info)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(info.Plugins.AvailableOnServer)
	s.assert.NotZero(info.Plugins.EnabledInCluster)
	s.assert.NotEmpty(info.Tagline)
	s.assert.NotZero(info.Configuration.Plugins)
	s.assert.NotZero(info.Configuration.AdminListen)
	s.assert.NotZero(info.Configuration.LuaSSLVerifyDepth)
	s.assert.NotZero(info.Configuration.TrustedIps)
	s.assert.NotEmpty(info.Configuration.Prefix)
	s.assert.NotEmpty(info.Configuration.NginxConf)
	s.assert.NotEmpty(info.Configuration.CassandraUsername)
	s.assert.NotEmpty(info.Configuration.AdminSSLCertificateCsrDefault)
	s.assert.NotEmpty(info.Configuration.SSLCertificateKey)
	s.assert.NotZero(info.Configuration.DNSResolver)
	s.assert.NotEmpty(info.Configuration.PostgresUsername)
	s.assert.NotEmpty(info.Configuration.MemoryCacheSize)
	s.assert.NotEmpty(info.Configuration.SSLCiphers)
	s.assert.NotZero(info.Configuration.CustomPlugins)
	s.assert.NotEmpty(info.Configuration.PostgresHost)
	s.assert.NotEmpty(info.Configuration.NginxAccessLogs)
	s.assert.NotZero(info.Configuration.ProxyListen)
	s.assert.NotEmpty(info.Configuration.ClientSSLCertificateDefault)
	s.assert.NotEmpty(info.Configuration.SSLCertificateDefaultKey)
	s.assert.NotZero(info.Configuration.DatabaseUpdateFrequency)
	s.assert.Zero(info.Configuration.DatabaseUpdatePropagation)
	s.assert.NotEmpty(info.Configuration.NginxErrorLogs)
	s.assert.NotZero(info.Configuration.CassandraPort)
	s.assert.NotZero(info.Configuration.DNSOrder)
	s.assert.NotZero(info.Configuration.DNSErrorTTL)
	s.assert.NotEmpty(info.Configuration.CassandraLBPolicy)
	s.assert.True(info.Configuration.NginxOptimizations)
	s.assert.NotEmpty(info.Configuration.Database)
	s.assert.NotEmpty(info.Configuration.PostgresDatabase)
	s.assert.NotEmpty(info.Configuration.NginxWorkerProcesses)
	s.assert.Empty(info.Configuration.LuaPackageCPath)
	s.assert.NotEmpty(info.Configuration.LuaPackagePath)
	s.assert.NotEmpty(info.Configuration.NginxPID)
	s.assert.NotZero(info.Configuration.UpstreamKeepAlive)
	s.assert.NotEmpty(info.Configuration.AdminAcessLog)
	s.assert.NotEmpty(info.Configuration.ClientSSLCertificateCsrDefault)
	s.assert.NotZero(info.Configuration.ProxyListeners, 2)
	s.assert.False(info.Configuration.ProxyListeners[0].SSL)
	s.assert.NotEmpty(info.Configuration.ProxyListeners[0].Ip)
	s.assert.False(info.Configuration.ProxyListeners[0].Protocol)
	s.assert.NotZero(info.Configuration.ProxyListeners[0].Port)
	s.assert.False(info.Configuration.ProxyListeners[0].Http2)
	s.assert.NotEmpty(info.Configuration.ProxyListeners[0].Listener)
	s.assert.True(info.Configuration.ProxySSLEnabled)
	s.assert.NotZero(info.Configuration.LuaSocketPoolSize)
	s.assert.NotEmpty(info.Configuration.ErrorDefaultType)
	s.assert.NotEmpty(info.Configuration.ProxyAccessLog)
	s.assert.False(info.Configuration.CassandraSSL)
	s.assert.NotEmpty(info.Configuration.CassandraConsistency)
	s.assert.NotEmpty(info.Configuration.ClientMaxBodySize)
	s.assert.NotEmpty("logs/error.log", info.Configuration.AdminErrorLog)
	s.assert.NotEmpty("/usr/local/kong/ssl/admin-kong-default.crt", info.Configuration.AdminSSLCertificateDefault)
	s.assert.NotZero(info.Configuration.DNSNotFoundTTL)
	s.assert.False(info.Configuration.PostgresSSL)
	s.assert.NotEmpty("notice", info.Configuration.LogLevel)
	s.assert.NotZero(info.Configuration.CassandraReplicationFactor)
	s.assert.NotEmpty("SimpleStrategy", info.Configuration.CassandraReplicationStrategy)
	s.assert.True(info.Configuration.LatencyTokens)
	s.assert.NotZero(info.Configuration.CassandraDataCenters)
	s.assert.NotEmpty("X-Real-IP", info.Configuration.RealIpHeader)
	s.assert.NotEmpty("/usr/local/kong/ssl/admin-kong-default.key", info.Configuration.AdminSSLCertificateKeyDefault)
	s.assert.NotEmpty("/usr/local/kong/.kong_env", info.Configuration.KongEnv)
	s.assert.NotZero(info.Configuration.CassandraSchemaConsensusTimeout)
	s.assert.NotEmpty("/etc/hosts", info.Configuration.DNSHostsFile)
	s.assert.NotZero(info.Configuration.AdminListeners)
	s.assert.False(info.Configuration.AdminListeners[0].SSL)
	s.assert.NotEmpty("0.0.0.0", info.Configuration.AdminListeners[0].Ip)
	s.assert.False(info.Configuration.AdminListeners[0].Protocol)
	s.assert.NotZero(info.Configuration.AdminListeners[0].Port)
	s.assert.False(info.Configuration.AdminListeners[0].Http2)
	s.assert.NotEmpty("0.0.0.0:8001", info.Configuration.AdminListeners[0].Listener)
	s.assert.False(info.Configuration.DNSNoSync)
	s.assert.NotEmpty("/usr/local/kong/ssl/kong-default.crt", info.Configuration.SSLCertificate)
	s.assert.False(info.Configuration.ClientSSL)
	s.assert.NotZero(info.Configuration.CassandraTimeout)
	s.assert.False(info.Configuration.CassandraSSLVerify)
	s.assert.NotZero(info.Configuration.CassandraContactPoints)
	s.assert.True(info.Configuration.ServerTokens)
	s.assert.NotEmpty("off", info.Configuration.RealIpRecursive)
	s.assert.NotEmpty("logs/error.log", info.Configuration.ProxyErrorLog)
	s.assert.NotEmpty("/usr/local/kong/ssl/kong-default.key", info.Configuration.ClientSSLCertificateKeyDefault)
	s.assert.NotEmpty("off", info.Configuration.NginxDaemon)
	s.assert.True(info.Configuration.AnonymousReports)
	s.assert.NotEmpty("modern", info.Configuration.SSLCipherSuite)
	s.assert.NotZero(info.Configuration.DNSStaleTTL)
	s.assert.NotZero(info.Configuration.PostgresPort)
	s.assert.NotEmpty("/usr/local/kong/nginx-kong.conf", info.Configuration.NginxKongConf)
	s.assert.NotEmpty("8k", info.Configuration.ClientBodyBufferSize)
	s.assert.NotZero(info.Configuration.DatabaseCacheTTL)
	s.assert.False(info.Configuration.PostgresSSLVerify)
	s.assert.NotEmpty("/usr/local/kong/logs/admin_access.log", info.Configuration.NginxAdminAccessLog)
	s.assert.NotEmpty("kong", info.Configuration.CassandraKeyspace)
	s.assert.NotEmpty("/usr/local/kong/ssl/kong-default.crt", info.Configuration.SSLCertificateDefault)
	s.assert.NotEmpty("/usr/local/kong/ssl/kong-default.csr", info.Configuration.SSLCertificateCsrDefault)
	s.assert.False(info.Configuration.AdminSSLEnabled)

	s.assert.NotEmpty(info.Version)
	s.assert.NotEmpty(info.LuaVersion)
	s.assert.NotZero(info.PrngSeeds)
	s.assert.NotZero(info.Timers.Pending)
	s.assert.Zero(info.Timers.Running)
	s.assert.NotEmpty(info.Hostname)
}

func (s *NodeTestSuite) TestStatusReturnsHttpError() {
	s.mux.HandleFunc(nodeStatusResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "")
	})

	client, _ := New(nil, s.server.URL)
	_, res, err := client.Node.Status()

	s.assert.IsType(&http.Response{}, res)
	s.assert.Error(err)
}

func (s *NodeTestSuite) TestStatus() {
	s.mux.HandleFunc(nodeStatusResourcePath, func(w http.ResponseWriter, r *http.Request) {
		s.assert.Equal(http.MethodGet, r.Method)

		file, _ := s.LoadFixture("fixtures/node_status_payload.json")

		io.Copy(w, file)

		defer file.Close()
	})

	status, res, err := s.client.Node.Status()

	s.assert.IsType(&NodeStatus{}, status)
	s.assert.IsType(&http.Response{}, res)
	s.assert.Nil(err)

	s.assert.NotZero(status.Server.ConnectionsAccepted)
	s.assert.NotZero(status.Server.ConnectionsActive)
	s.assert.NotZero(status.Server.ConnectionsHandled)
	s.assert.Zero(status.Server.ConnectionsReading)
	s.assert.Zero(status.Server.ConnectionsWaiting)
	s.assert.NotZero(status.Server.ConnectionsWriting)
	s.assert.NotZero(status.Server.TotalRequests)

	s.assert.True(status.Database.Reachable)
}

func TestNodeTestSuite(t *testing.T) {
	suite.Run(t, new(NodeTestSuite))
}
