package load

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cybertec-postgresql/pg_timetable/embed/tomlconfig"
	"github.com/cybertec-postgresql/pg_timetable/internal/config"
)

func TestLoad(t *testing.T) {
	c := &config.CmdOptions{}
	c.Version = true
	c.ClientName = "pgt"
	c.Connection.Host = "127.0.0.1"
	c.Connection.Port = 5433
	c.Connection.Password = "postgresPWD"
	c.Connection.User = "postgrest"
	c.Connection.SSLMode = "disable"
	c.Connection.Timeout = 45
	c.Resource.CronWorkers = 16
	c.RestApi.Port = 8008
	c.Logging.LogLevel = "info"
	c.Logging.LogDBLevel = "info"
	c.Logging.LogFile = "pgt_session.log"
	c.Logging.LogFileFormat = "json"

	f := "testdata/pgt_config.toml"
	err := tomlconfig.SaveTOML(c, f)
	assert.NoError(t, err)

	cfg, er1 := Load(f)
	assert.NoError(t, er1)
	assert.Equal(t, cfg.Connection.Host, "127.0.0.1")
}
