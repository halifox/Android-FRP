package frpandroid

import (
	"context"
	"fmt"
	"sync"

	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/config/source"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	frplog "github.com/fatedier/frp/pkg/util/log"
)

var defaultController = &frpcController{}

func Start(cfg *FrpcConfig) error {
	return defaultController.start(cfg)
}

func Stop() error {
	return defaultController.stop()
}

func Reload(cfg *FrpcConfig) error {
	return defaultController.reload(cfg)
}

type frpcController struct {
	mu sync.Mutex

	service *client.Service
	cancel  context.CancelFunc
	done    chan error

	running bool
}

func (c *frpcController) start(cfg *FrpcConfig) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.running {
		return fmt.Errorf("frpc is already running")
	}

	service, err := buildService(cfg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)

	c.service = service
	c.cancel = cancel
	c.done = done
	c.running = true

	go c.run(service, ctx, done)
	return nil
}

func (c *frpcController) stop() error {
	c.mu.Lock()
	if !c.running {
		c.mu.Unlock()
		return nil
	}

	cancel := c.cancel
	done := c.done
	c.mu.Unlock()

	cancel()
	return <-done
}

func (c *frpcController) reload(cfg *FrpcConfig) error {
	if err := c.stop(); err != nil {
		return err
	}
	return c.start(cfg)
}

func (c *frpcController) run(service *client.Service, ctx context.Context, done chan error) {
	err := service.Run(ctx)

	c.mu.Lock()
	if c.service == service {
		c.service = nil
		c.cancel = nil
		c.done = nil
		c.running = false
	}
	c.mu.Unlock()

	done <- err
	close(done)
}

func buildService(cfg *FrpcConfig) (*client.Service, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	fullConfig := cfg.ClientConfig
	common := fullConfig.ClientCommonConfig
	if err := common.Complete(); err != nil {
		return nil, err
	}
	if common.ServerAddr == "" {
		return nil, fmt.Errorf("ServerAddr is required")
	}
	if common.ServerPort <= 0 {
		return nil, fmt.Errorf("ServerPort is required")
	}

	frplog.InitLogger(common.Log.To, common.Log.Level, int(common.Log.MaxDays), common.Log.DisablePrintColor)

	proxies := make([]v1.ProxyConfigurer, 0, len(fullConfig.Proxies))
	for _, item := range fullConfig.Proxies {
		if item.ProxyConfigurer == nil {
			return nil, fmt.Errorf("proxy config is nil")
		}
		proxies = append(proxies, item.ProxyConfigurer.Clone())
	}

	visitors := make([]v1.VisitorConfigurer, 0, len(fullConfig.Visitors))
	for _, item := range fullConfig.Visitors {
		if item.VisitorConfigurer == nil {
			return nil, fmt.Errorf("visitor config is nil")
		}
		visitors = append(visitors, item.VisitorConfigurer.Clone())
	}

	if len(proxies) == 0 && len(visitors) == 0 {
		return nil, fmt.Errorf("at least one proxy or visitor is required")
	}

	proxies, visitors = config.FilterClientConfigurers(&common, proxies, visitors)
	proxies = config.CompleteProxyConfigurers(proxies)
	visitors = config.CompleteVisitorConfigurers(visitors)

	configSource := source.NewConfigSource()
	if err := configSource.ReplaceAll(proxies, visitors); err != nil {
		return nil, err
	}

	return client.NewService(client.ServiceOptions{
		Common:                 &common,
		ConfigSourceAggregator: source.NewAggregator(configSource),
	})
}
