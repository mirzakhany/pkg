package status

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoas/stats"
	"github.com/mirzakhany/pkg/logger"
	"github.com/mirzakhany/pkg/status/leveldb"
	"github.com/mirzakhany/pkg/status/boltdb"
	"github.com/mirzakhany/pkg/status/redis"
)

// StatStorage implements the storage interface
var StatStorage KVStorage

// Stats provide response time, status code count, etc.
var Stats = stats.New()

// AppStatus is app status structure
type AppStatus struct {
	Version    string     `json:"version"`
}

// InitAppStatus for initialize app status
func InitAppStatus(settings ConfStatus) error {
	logger.Infof("Init App Status Engine as %s", settings.Engine)
	switch settings.Engine {
	case "redis":
		StatStorage = redis.New(settings)
	case "boltdb":
		StatStorage = boltdb.New(settings)
	case "leveldb":
		StatStorage = leveldb.New(settings)
	default:
		logger.LogError.Error("storage error: can't find storage driver")
		return errors.New("can't find storage driver")
	}

	if err := StatStorage.Init(); err != nil {
		logger.LogError.Error("storage error: " + err.Error())
		return err
	}

	return nil
}

// SysStatsHandler return system status
func SysStatsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Stats.Data())
}

// StatMiddleware response time, status code count, etc.
func StatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginning, recorder := Stats.Begin(c.Writer)
		c.Next()
		Stats.End(beginning, recorder)
	}
}
