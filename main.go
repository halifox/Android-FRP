package frp

import (
	"context"
	"os"
	"path/filepath"

	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/config/source"
	"github.com/fatedier/frp/pkg/policy/featuregate"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
)

func RunClient(cfgFile string) {
	result, err := config.LoadClientConfigResult(cfgFile, true)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if len(result.Common.FeatureGates) > 0 {
		if err := featuregate.SetFromMap(result.Common.FeatureGates); err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}

	configSource := source.NewConfigSource()
	if err := configSource.ReplaceAll(result.Proxies, result.Visitors); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	var storeSource *source.StoreSource
	if result.Common.Store.IsEnabled() {
		storePath := result.Common.Store.Path
		if storePath != "" && cfgFile != "" && !filepath.IsAbs(storePath) {
			storePath = filepath.Join(filepath.Dir(cfgFile), storePath)
		}

		storeSource, err = source.NewStoreSource(source.StoreSourceConfig{
			Path: storePath,
		})
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}

	aggregator := source.NewAggregator(configSource)
	if storeSource != nil {
		aggregator.SetStoreSource(storeSource)
	}

	proxyCfgs, visitorCfgs, err := aggregator.Load()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	proxyCfgs, visitorCfgs = config.FilterClientConfigurers(result.Common, proxyCfgs, visitorCfgs)
	proxyCfgs = config.CompleteProxyConfigurers(proxyCfgs)
	visitorCfgs = config.CompleteVisitorConfigurers(visitorCfgs)
	log.InitLogger(result.Common.Log.To, result.Common.Log.Level, int(result.Common.Log.MaxDays), result.Common.Log.DisablePrintColor)

	svr, err := client.NewService(client.ServiceOptions{
		Common:                 result.Common,
		ConfigSourceAggregator: aggregator,
		ConfigFilePath:         cfgFile,
	})
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	svr.Run(context.Background())
}

func RunServer(cfgFile string) {
	cfg, _, err := config.LoadServerConfig(cfgFile, true)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)
	svr, err := server.NewService(cfg)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	svr.Run(context.Background())
}
