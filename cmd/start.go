package cmd

import (
	"fmt"
	"github.com/cksidharthan/sih-server/pkg/config"
	"github.com/cksidharthan/sih-server/pkg/logger"
	"github.com/cksidharthan/sih-server/pkg/router"
	filesCtrl "github.com/cksidharthan/sih-server/pkg/v1/files/controller"
	filesSvc "github.com/cksidharthan/sih-server/pkg/v1/files/service"
	"go.uber.org/fx"
	"net"
)

func Start() {
	fmt.Println("Starting Server...")

	// creating a http listener to serve the api
	httpListener := func(envCfg *config.Config) (net.Listener, error) {
		return net.Listen("tcp", fmt.Sprintf(":%s", envCfg.Port))
	}

	app := fx.New(
		fx.Provide(
			config.New,
			logger.New,
			httpListener,
			router.New,

			filesSvc.New,
		),
		fx.Invoke(
			router.DefaultEndpoints,
			filesCtrl.New,
		),
	)

	app.Run()
}
