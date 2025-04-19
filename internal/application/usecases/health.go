package usecase

import (
	"dms/config"
	"time"

	"github.com/hellofresh/health-go/v5"
	healthHttp "github.com/hellofresh/health-go/v5/checks/http"
	healthInflux "github.com/hellofresh/health-go/v5/checks/influxdb"
	healthMongo "github.com/hellofresh/health-go/v5/checks/mongo"
	healthMySql "github.com/hellofresh/health-go/v5/checks/mysql"
)

func NewHealthUseCase(h *health.Health) (*health.Health, error) {
	cfg := config.LoadConfig()

	// Connection HTTP Check
	h.Register(health.Config{
		Name:      "http-check",
		Timeout:   5 * time.Second,
		SkipOnErr: true,
		Check: healthHttp.New(healthHttp.Config{
			URL: `http://example.com`,
		}),
	})

	// MySQL Check
	h.Register(health.Config{
		Name:      "mysql-check",
		Timeout:   5 * time.Second,
		SkipOnErr: true,
		Check: healthMySql.New(healthMySql.Config{
			DSN: `test:test@tcp(localhost:3306)/test?charset=utf8`,
		}),
	})

	// MongoDB Check
	h.Register(health.Config{
		Name:      "mongo-check",
		Timeout:   5 * time.Second,
		SkipOnErr: true,
		Check: healthMongo.New(healthMongo.Config{
			DSN: cfg.GetString("MONGODB_URL"),
		}),
	})

	// InfluxDB Check
	h.Register(health.Config{
		Name:      "influxdb-check",
		Timeout:   5 * time.Second,
		SkipOnErr: true,
		Check: healthInflux.New(healthInflux.Config{
			URL: "http://localhost:8086/health",
		}),
	})

	return h, nil
}
