package gobblr

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	gsl "github.com/mrsimonemms/gin-structured-logger"
	"github.com/mrsimonemms/gobblr/pkg/drivers"
)

var runCount = 0

func Serve(dataPath string, db drivers.Driver, retries uint64, port int) error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(
		requestid.New(),
		gsl.New(),
		gin.Recovery(),
		func(ctx *gin.Context) {
			gsl.Get(ctx).Debug().Str("path", ctx.Request.URL.Path).Msg("New HTTP call")
		},
	)

	h := handler{
		DataPath: dataPath,
		Driver:   db,
		Retries:  retries,
	}

	// Register the routes
	r.POST("/data/reset", h.ResetData)

	logger.Info().Int("port", port).Msg("Starting web server")

	return r.Run(fmt.Sprintf(":%d", port))
}

type handler struct {
	DataPath string
	Driver   drivers.Driver
	Retries  uint64
}

// ResetData runs the Execute command whenever it receives a call
func (h handler) ResetData(c *gin.Context) {
	// Increment the run count
	runCount++

	log := gsl.Get(c).With().Logger()

	inserted, err := Execute(h.DataPath, h.Driver, h.Retries)
	if err != nil {
		log.Error().Err(err).Msg("Failed to ingest data to database")

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   http.StatusText(http.StatusServiceUnavailable),
			"message": err.Error(),
		})

		return
	}

	// Log the result
	log.Info().Fields(map[string]interface{}{
		"inserted": inserted,
	}).Int("run count", runCount).Msg("Successfully inserted data")

	c.JSON(http.StatusOK, inserted)
}
