package load

import (
	"github.com/BurntSushi/toml"

	"github.com/cybertec-postgresql/pg_timetable/internal/config"
)

func Load(f string) (*config.CmdOptions, error) {
	conf := &config.CmdOptions{}

	_, err := toml.DecodeFile(f, &conf)

	return conf, err
}
