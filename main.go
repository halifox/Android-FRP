package frp

import (
	"context"
	"fmt"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
	"os"
)

func RunClient(cfgFile string) {
	cfg, proxyCfgs, visitorCfgs, isLegacyFormat, err := config.LoadClientConfig(cfgFile, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isLegacyFormat {
		fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
			"please use yaml/json/toml format instead!\n")
	}

	warning, err := validation.ValidateAllClientConfig(cfg, proxyCfgs, visitorCfgs)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)

	if cfgFile != "" {
		log.Infof("start frpc service for config file [%s]", cfgFile)
		defer log.Infof("frpc service for config file [%s] stopped", cfgFile)
	}
	svr, err := client.NewService(client.ServiceOptions{
		Common:         cfg,
		ProxyCfgs:      proxyCfgs,
		VisitorCfgs:    visitorCfgs,
		ConfigFilePath: cfgFile,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	svr.Run(context.Background())
}

func RunServer(cfgFile string) {
	cfg, isLegacyFormat, err := config.LoadServerConfig(cfgFile, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isLegacyFormat {
		fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
			"please use yaml/json/toml format instead!\n")
	}

	warning, err := validation.ValidateServerConfig(cfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)
	svr, err := server.NewService(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Infof("frps started successfully")
	svr.Run(context.Background())
}
