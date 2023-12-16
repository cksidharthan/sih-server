package router

import (
	"context"
	"errors"
	"github.com/cksidharthan/sih-server/pkg/config"
	"net"
	"net/http"
	"sync"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var bootstrapOnly sync.Once

type Router struct {
	fx.Out

	Engine *gin.Engine
	Server *http.Server
	V1     *gin.RouterGroup `name:"v1"`
}

// New - creates a new router object that will be used by the application - contains the default routes and groups
func New(lc fx.Lifecycle, l net.Listener, envCfg *config.Config, logger *logrus.Logger) Router {
	// These calls mutate a global state
	// using sync.Once here prevents data races
	bootstrapOnly.Do(func() {
		gin.SetMode(gin.ReleaseMode)
		// Setup validator to ignore fields with no `validate` tag
		validator.SetFieldsRequiredByDefault(false)
	})

	engine := gin.Default()
	engine.Use(gin.Recovery())

	server := &http.Server{
		Handler: engine,
	}

	engine.RedirectTrailingSlash = false

	// Appends the on start and on stop lifecycle to the uberfx framework
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			errChan := make(chan error)

			go func(errChan chan error) {
				if err := server.Serve(l); err != nil && !errors.Is(err, http.ErrServerClosed) {
					errChan <- err
				}
			}(errChan)

			select {
			case err := <-errChan:
				return err
			case <-time.After(1 * time.Second):
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return Router{
		Engine: engine,
		Server: server,
		V1:     engine.Group("/api/v1"),
	}
}
