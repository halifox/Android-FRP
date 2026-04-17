package frpandroid

import (
	"fmt"

	"github.com/fatedier/frp/pkg/config/types"
	v1 "github.com/fatedier/frp/pkg/config/v1"
)

type FrpcConfig struct {
	v1.ClientConfig
}

func NewFrpcConfig() *FrpcConfig {
	return &FrpcConfig{}
}

func (c *FrpcConfig) SetServerAddr(value string) {
	c.ServerAddr = value
}

func (c *FrpcConfig) SetServerPort(value int) {
	c.ServerPort = value
}

func (c *FrpcConfig) SetUser(value string) {
	c.User = value
}

func (c *FrpcConfig) SetClientID(value string) {
	c.ClientID = value
}

func (c *FrpcConfig) SetNatHoleSTUNServer(value string) {
	c.NatHoleSTUNServer = value
}

func (c *FrpcConfig) SetDNSServer(value string) {
	c.DNSServer = value
}

func (c *FrpcConfig) SetLoginFailExit(value bool) {
	c.LoginFailExit = &value
}

func (c *FrpcConfig) AddStart(value string) {
	c.Start = append(c.Start, value)
}

func (c *FrpcConfig) ClearStart() {
	c.Start = nil
}

func (c *FrpcConfig) SetAuthMethod(value string) {
	c.Auth.Method = v1.AuthMethod(value)
}

func (c *FrpcConfig) SetAuthToken(value string) {
	c.Auth.Token = value
}

func (c *FrpcConfig) AddAuthAdditionalScope(value string) {
	c.Auth.AdditionalScopes = append(c.Auth.AdditionalScopes, v1.AuthScope(value))
}

func (c *FrpcConfig) ClearAuthAdditionalScopes() {
	c.Auth.AdditionalScopes = nil
}

func (c *FrpcConfig) SetOIDCClientID(value string) {
	c.Auth.OIDC.ClientID = value
}

func (c *FrpcConfig) SetOIDCClientSecret(value string) {
	c.Auth.OIDC.ClientSecret = value
}

func (c *FrpcConfig) SetOIDCAudience(value string) {
	c.Auth.OIDC.Audience = value
}

func (c *FrpcConfig) SetOIDCScope(value string) {
	c.Auth.OIDC.Scope = value
}

func (c *FrpcConfig) SetOIDCTokenEndpointURL(value string) {
	c.Auth.OIDC.TokenEndpointURL = value
}

func (c *FrpcConfig) SetOIDCTrustedCaFile(value string) {
	c.Auth.OIDC.TrustedCaFile = value
}

func (c *FrpcConfig) SetOIDCInsecureSkipVerify(value bool) {
	c.Auth.OIDC.InsecureSkipVerify = value
}

func (c *FrpcConfig) SetOIDCProxyURL(value string) {
	c.Auth.OIDC.ProxyURL = value
}

func (c *FrpcConfig) PutOIDCAdditionalEndpointParam(key string, value string) {
	if c.Auth.OIDC.AdditionalEndpointParams == nil {
		c.Auth.OIDC.AdditionalEndpointParams = map[string]string{}
	}
	c.Auth.OIDC.AdditionalEndpointParams[key] = value
}

func (c *FrpcConfig) SetLogTo(value string) {
	c.Log.To = value
}

func (c *FrpcConfig) SetLogLevel(value string) {
	c.Log.Level = value
}

func (c *FrpcConfig) SetLogMaxDays(value int64) {
	c.Log.MaxDays = value
}

func (c *FrpcConfig) SetLogDisablePrintColor(value bool) {
	c.Log.DisablePrintColor = value
}

func (c *FrpcConfig) SetWebServerAddr(value string) {
	c.WebServer.Addr = value
}

func (c *FrpcConfig) SetWebServerPort(value int) {
	c.WebServer.Port = value
}

func (c *FrpcConfig) SetWebServerUser(value string) {
	c.WebServer.User = value
}

func (c *FrpcConfig) SetWebServerPassword(value string) {
	c.WebServer.Password = value
}

func (c *FrpcConfig) SetWebServerAssetsDir(value string) {
	c.WebServer.AssetsDir = value
}

func (c *FrpcConfig) SetWebServerPprofEnable(value bool) {
	c.WebServer.PprofEnable = value
}

func (c *FrpcConfig) SetWebServerTLSCertFile(value string) {
	c.ensureWebServerTLS()
	c.WebServer.TLS.CertFile = value
}

func (c *FrpcConfig) SetWebServerTLSKeyFile(value string) {
	c.ensureWebServerTLS()
	c.WebServer.TLS.KeyFile = value
}

func (c *FrpcConfig) SetWebServerTLSTrustedCaFile(value string) {
	c.ensureWebServerTLS()
	c.WebServer.TLS.TrustedCaFile = value
}

func (c *FrpcConfig) SetWebServerTLSServerName(value string) {
	c.ensureWebServerTLS()
	c.WebServer.TLS.ServerName = value
}

func (c *FrpcConfig) SetTransportProtocol(value string) {
	c.Transport.Protocol = value
}

func (c *FrpcConfig) SetTransportDialServerTimeout(value int64) {
	c.Transport.DialServerTimeout = value
}

func (c *FrpcConfig) SetTransportDialServerKeepAlive(value int64) {
	c.Transport.DialServerKeepAlive = value
}

func (c *FrpcConfig) SetTransportConnectServerLocalIP(value string) {
	c.Transport.ConnectServerLocalIP = value
}

func (c *FrpcConfig) SetTransportProxyURL(value string) {
	c.Transport.ProxyURL = value
}

func (c *FrpcConfig) SetTransportPoolCount(value int) {
	c.Transport.PoolCount = value
}

func (c *FrpcConfig) SetTransportTCPMux(value bool) {
	c.Transport.TCPMux = &value
}

func (c *FrpcConfig) SetTransportTCPMuxKeepaliveInterval(value int64) {
	c.Transport.TCPMuxKeepaliveInterval = value
}

func (c *FrpcConfig) SetTransportHeartbeatInterval(value int64) {
	c.Transport.HeartbeatInterval = value
}

func (c *FrpcConfig) SetTransportHeartbeatTimeout(value int64) {
	c.Transport.HeartbeatTimeout = value
}

func (c *FrpcConfig) SetTransportQUICKeepalivePeriod(value int) {
	c.Transport.QUIC.KeepalivePeriod = value
}

func (c *FrpcConfig) SetTransportQUICMaxIdleTimeout(value int) {
	c.Transport.QUIC.MaxIdleTimeout = value
}

func (c *FrpcConfig) SetTransportQUICMaxIncomingStreams(value int) {
	c.Transport.QUIC.MaxIncomingStreams = value
}

func (c *FrpcConfig) SetTransportTLSEnable(value bool) {
	c.Transport.TLS.Enable = &value
}

func (c *FrpcConfig) SetTransportTLSDisableCustomTLSFirstByte(value bool) {
	c.Transport.TLS.DisableCustomTLSFirstByte = &value
}

func (c *FrpcConfig) SetTransportTLSCertFile(value string) {
	c.Transport.TLS.CertFile = value
}

func (c *FrpcConfig) SetTransportTLSKeyFile(value string) {
	c.Transport.TLS.KeyFile = value
}

func (c *FrpcConfig) SetTransportTLSTrustedCaFile(value string) {
	c.Transport.TLS.TrustedCaFile = value
}

func (c *FrpcConfig) SetTransportTLSServerName(value string) {
	c.Transport.TLS.ServerName = value
}

func (c *FrpcConfig) SetUDPPacketSize(value int64) {
	c.UDPPacketSize = value
}

func (c *FrpcConfig) PutMetadata(key string, value string) {
	if c.Metadatas == nil {
		c.Metadatas = map[string]string{}
	}
	c.Metadatas[key] = value
}

func (c *FrpcConfig) PutFeatureGate(key string, value bool) {
	if c.FeatureGates == nil {
		c.FeatureGates = map[string]bool{}
	}
	c.FeatureGates[key] = value
}

func (c *FrpcConfig) AddProxy(proxy *FrpcProxyConfig) {
	if proxy == nil || proxy.ProxyConfigurer == nil {
		return
	}
	c.Proxies = append(c.Proxies, proxy.TypedProxyConfig)
}

func (c *FrpcConfig) AddVisitor(visitor *FrpcVisitorConfig) {
	if visitor == nil || visitor.VisitorConfigurer == nil {
		return
	}
	c.Visitors = append(c.Visitors, visitor.TypedVisitorConfig)
}

func (c *FrpcConfig) ClearProxies() {
	c.Proxies = nil
}

func (c *FrpcConfig) ClearVisitors() {
	c.Visitors = nil
}

type FrpcProxyConfig struct {
	v1.TypedProxyConfig
}

func NewFrpcProxyConfig(proxyType string) *FrpcProxyConfig {
	cfg := v1.NewProxyConfigurerByType(v1.ProxyType(proxyType))
	if cfg == nil {
		return nil
	}
	return &FrpcProxyConfig{
		TypedProxyConfig: v1.TypedProxyConfig{
			Type:            cfg.GetBaseConfig().Type,
			ProxyConfigurer: cfg,
		},
	}
}

func (p *FrpcProxyConfig) SetName(value string) {
	if base := p.baseConfig(); base != nil {
		base.Name = value
	}
}

func (p *FrpcProxyConfig) SetLocalIP(value string) {
	if base := p.baseConfig(); base != nil {
		base.LocalIP = value
	}
}

func (p *FrpcProxyConfig) SetLocalPort(value int) {
	if base := p.baseConfig(); base != nil {
		base.LocalPort = value
	}
}

func (p *FrpcProxyConfig) SetUseEncryption(value bool) {
	if base := p.baseConfig(); base != nil {
		base.Transport.UseEncryption = value
	}
}

func (p *FrpcProxyConfig) SetUseCompression(value bool) {
	if base := p.baseConfig(); base != nil {
		base.Transport.UseCompression = value
	}
}

func (p *FrpcProxyConfig) SetBandwidthLimit(value string) error {
	base := p.baseConfig()
	if base == nil {
		return fmt.Errorf("proxy config is nil")
	}
	q, err := types.NewBandwidthQuantity(value)
	if err != nil {
		return err
	}
	base.Transport.BandwidthLimit = q
	return nil
}

func (p *FrpcProxyConfig) SetBandwidthLimitMode(value string) {
	if base := p.baseConfig(); base != nil {
		base.Transport.BandwidthLimitMode = value
	}
}

func (p *FrpcProxyConfig) SetProxyProtocolVersion(value string) {
	if base := p.baseConfig(); base != nil {
		base.Transport.ProxyProtocolVersion = value
	}
}

func (p *FrpcProxyConfig) SetGroup(value string) {
	if base := p.baseConfig(); base != nil {
		base.LoadBalancer.Group = value
	}
}

func (p *FrpcProxyConfig) SetGroupKey(value string) {
	if base := p.baseConfig(); base != nil {
		base.LoadBalancer.GroupKey = value
	}
}

func (p *FrpcProxyConfig) SetEnabled(value bool) {
	if base := p.baseConfig(); base != nil {
		base.Enabled = &value
	}
}

func (p *FrpcProxyConfig) PutMetadata(key string, value string) {
	if base := p.baseConfig(); base != nil {
		if base.Metadatas == nil {
			base.Metadatas = map[string]string{}
		}
		base.Metadatas[key] = value
	}
}

func (p *FrpcProxyConfig) PutAnnotation(key string, value string) {
	if base := p.baseConfig(); base != nil {
		if base.Annotations == nil {
			base.Annotations = map[string]string{}
		}
		base.Annotations[key] = value
	}
}

func (p *FrpcProxyConfig) SetHealthCheckType(value string) {
	if base := p.baseConfig(); base != nil {
		base.HealthCheck.Type = value
	}
}

func (p *FrpcProxyConfig) SetHealthCheckTimeoutSeconds(value int) {
	if base := p.baseConfig(); base != nil {
		base.HealthCheck.TimeoutSeconds = value
	}
}

func (p *FrpcProxyConfig) SetHealthCheckMaxFailed(value int) {
	if base := p.baseConfig(); base != nil {
		base.HealthCheck.MaxFailed = value
	}
}

func (p *FrpcProxyConfig) SetHealthCheckIntervalSeconds(value int) {
	if base := p.baseConfig(); base != nil {
		base.HealthCheck.IntervalSeconds = value
	}
}

func (p *FrpcProxyConfig) SetHealthCheckPath(value string) {
	if base := p.baseConfig(); base != nil {
		base.HealthCheck.Path = value
	}
}

func (p *FrpcProxyConfig) AddHealthCheckHeader(name string, value string) {
	if base := p.baseConfig(); base != nil {
		base.HealthCheck.HTTPHeaders = append(base.HealthCheck.HTTPHeaders, v1.HTTPHeader{Name: name, Value: value})
	}
}

func (p *FrpcProxyConfig) SetRemotePort(value int) error {
	switch cfg := p.configurer().(type) {
	case *v1.TCPProxyConfig:
		cfg.RemotePort = value
	case *v1.UDPProxyConfig:
		cfg.RemotePort = value
	default:
		return fmt.Errorf("RemotePort is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) SetSecretKey(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.STCPProxyConfig:
		cfg.Secretkey = value
	case *v1.XTCPProxyConfig:
		cfg.Secretkey = value
	case *v1.SUDPProxyConfig:
		cfg.Secretkey = value
	default:
		return fmt.Errorf("SecretKey is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) AddAllowUser(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.STCPProxyConfig:
		cfg.AllowUsers = append(cfg.AllowUsers, value)
	case *v1.XTCPProxyConfig:
		cfg.AllowUsers = append(cfg.AllowUsers, value)
	case *v1.SUDPProxyConfig:
		cfg.AllowUsers = append(cfg.AllowUsers, value)
	default:
		return fmt.Errorf("AllowUsers is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) SetSubDomain(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.HTTPProxyConfig:
		cfg.SubDomain = value
	case *v1.HTTPSProxyConfig:
		cfg.SubDomain = value
	case *v1.TCPMuxProxyConfig:
		cfg.SubDomain = value
	default:
		return fmt.Errorf("SubDomain is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) AddCustomDomain(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.HTTPProxyConfig:
		cfg.CustomDomains = append(cfg.CustomDomains, value)
	case *v1.HTTPSProxyConfig:
		cfg.CustomDomains = append(cfg.CustomDomains, value)
	case *v1.TCPMuxProxyConfig:
		cfg.CustomDomains = append(cfg.CustomDomains, value)
	default:
		return fmt.Errorf("CustomDomains is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) AddLocation(value string) error {
	cfg, ok := p.configurer().(*v1.HTTPProxyConfig)
	if !ok {
		return fmt.Errorf("Locations is not supported for this proxy type")
	}
	cfg.Locations = append(cfg.Locations, value)
	return nil
}

func (p *FrpcProxyConfig) SetHTTPUser(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.HTTPProxyConfig:
		cfg.HTTPUser = value
	case *v1.TCPMuxProxyConfig:
		cfg.HTTPUser = value
	default:
		return fmt.Errorf("HTTPUser is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) SetHTTPPassword(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.HTTPProxyConfig:
		cfg.HTTPPassword = value
	case *v1.TCPMuxProxyConfig:
		cfg.HTTPPassword = value
	default:
		return fmt.Errorf("HTTPPassword is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) SetHostHeaderRewrite(value string) error {
	cfg, ok := p.configurer().(*v1.HTTPProxyConfig)
	if !ok {
		return fmt.Errorf("HostHeaderRewrite is not supported for this proxy type")
	}
	cfg.HostHeaderRewrite = value
	return nil
}

func (p *FrpcProxyConfig) PutRequestHeader(key string, value string) error {
	cfg, ok := p.configurer().(*v1.HTTPProxyConfig)
	if !ok {
		return fmt.Errorf("RequestHeaders is not supported for this proxy type")
	}
	if cfg.RequestHeaders.Set == nil {
		cfg.RequestHeaders.Set = map[string]string{}
	}
	cfg.RequestHeaders.Set[key] = value
	return nil
}

func (p *FrpcProxyConfig) PutResponseHeader(key string, value string) error {
	cfg, ok := p.configurer().(*v1.HTTPProxyConfig)
	if !ok {
		return fmt.Errorf("ResponseHeaders is not supported for this proxy type")
	}
	if cfg.ResponseHeaders.Set == nil {
		cfg.ResponseHeaders.Set = map[string]string{}
	}
	cfg.ResponseHeaders.Set[key] = value
	return nil
}

func (p *FrpcProxyConfig) SetRouteByHTTPUser(value string) error {
	switch cfg := p.configurer().(type) {
	case *v1.HTTPProxyConfig:
		cfg.RouteByHTTPUser = value
	case *v1.TCPMuxProxyConfig:
		cfg.RouteByHTTPUser = value
	default:
		return fmt.Errorf("RouteByHTTPUser is not supported for this proxy type")
	}
	return nil
}

func (p *FrpcProxyConfig) SetMultiplexer(value string) error {
	cfg, ok := p.configurer().(*v1.TCPMuxProxyConfig)
	if !ok {
		return fmt.Errorf("Multiplexer is not supported for this proxy type")
	}
	cfg.Multiplexer = value
	return nil
}

type FrpcVisitorConfig struct {
	v1.TypedVisitorConfig
}

func NewFrpcVisitorConfig(visitorType string) *FrpcVisitorConfig {
	cfg := v1.NewVisitorConfigurerByType(v1.VisitorType(visitorType))
	if cfg == nil {
		return nil
	}
	return &FrpcVisitorConfig{
		TypedVisitorConfig: v1.TypedVisitorConfig{
			Type:              cfg.GetBaseConfig().Type,
			VisitorConfigurer: cfg,
		},
	}
}

func (v *FrpcVisitorConfig) SetName(value string) {
	if base := v.baseConfig(); base != nil {
		base.Name = value
	}
}

func (v *FrpcVisitorConfig) SetEnabled(value bool) {
	if base := v.baseConfig(); base != nil {
		base.Enabled = &value
	}
}

func (v *FrpcVisitorConfig) SetSecretKey(value string) {
	if base := v.baseConfig(); base != nil {
		base.SecretKey = value
	}
}

func (v *FrpcVisitorConfig) SetServerUser(value string) {
	if base := v.baseConfig(); base != nil {
		base.ServerUser = value
	}
}

func (v *FrpcVisitorConfig) SetServerName(value string) {
	if base := v.baseConfig(); base != nil {
		base.ServerName = value
	}
}

func (v *FrpcVisitorConfig) SetBindAddr(value string) {
	if base := v.baseConfig(); base != nil {
		base.BindAddr = value
	}
}

func (v *FrpcVisitorConfig) SetBindPort(value int) {
	if base := v.baseConfig(); base != nil {
		base.BindPort = value
	}
}

func (v *FrpcVisitorConfig) SetUseEncryption(value bool) {
	if base := v.baseConfig(); base != nil {
		base.Transport.UseEncryption = value
	}
}

func (v *FrpcVisitorConfig) SetUseCompression(value bool) {
	if base := v.baseConfig(); base != nil {
		base.Transport.UseCompression = value
	}
}

func (v *FrpcVisitorConfig) SetProtocol(value string) error {
	cfg, ok := v.configurer().(*v1.XTCPVisitorConfig)
	if !ok {
		return fmt.Errorf("Protocol is not supported for this visitor type")
	}
	cfg.Protocol = value
	return nil
}

func (v *FrpcVisitorConfig) SetKeepTunnelOpen(value bool) error {
	cfg, ok := v.configurer().(*v1.XTCPVisitorConfig)
	if !ok {
		return fmt.Errorf("KeepTunnelOpen is not supported for this visitor type")
	}
	cfg.KeepTunnelOpen = value
	return nil
}

func (v *FrpcVisitorConfig) SetMaxRetriesAnHour(value int) error {
	cfg, ok := v.configurer().(*v1.XTCPVisitorConfig)
	if !ok {
		return fmt.Errorf("MaxRetriesAnHour is not supported for this visitor type")
	}
	cfg.MaxRetriesAnHour = value
	return nil
}

func (v *FrpcVisitorConfig) SetMinRetryInterval(value int) error {
	cfg, ok := v.configurer().(*v1.XTCPVisitorConfig)
	if !ok {
		return fmt.Errorf("MinRetryInterval is not supported for this visitor type")
	}
	cfg.MinRetryInterval = value
	return nil
}

func (v *FrpcVisitorConfig) SetFallbackTo(value string) error {
	cfg, ok := v.configurer().(*v1.XTCPVisitorConfig)
	if !ok {
		return fmt.Errorf("FallbackTo is not supported for this visitor type")
	}
	cfg.FallbackTo = value
	return nil
}

func (v *FrpcVisitorConfig) SetFallbackTimeoutMs(value int) error {
	cfg, ok := v.configurer().(*v1.XTCPVisitorConfig)
	if !ok {
		return fmt.Errorf("FallbackTimeoutMs is not supported for this visitor type")
	}
	cfg.FallbackTimeoutMs = value
	return nil
}

func (c *FrpcConfig) ensureWebServerTLS() {
	if c.WebServer.TLS == nil {
		c.WebServer.TLS = &v1.TLSConfig{}
	}
}

func (p *FrpcProxyConfig) baseConfig() *v1.ProxyBaseConfig {
	cfg := p.configurer()
	if cfg == nil {
		return nil
	}
	return cfg.GetBaseConfig()
}

func (v *FrpcVisitorConfig) baseConfig() *v1.VisitorBaseConfig {
	cfg := v.configurer()
	if cfg == nil {
		return nil
	}
	return cfg.GetBaseConfig()
}

func (p *FrpcProxyConfig) configurer() v1.ProxyConfigurer {
	if p == nil {
		return nil
	}
	return p.ProxyConfigurer
}

func (v *FrpcVisitorConfig) configurer() v1.VisitorConfigurer {
	if v == nil {
		return nil
	}
	return v.VisitorConfigurer
}
