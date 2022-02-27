package embed

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/cybertec-postgresql/pg_timetable/internal/config"
)

func LoadConfig(f string) (*config.CmdOptions, error) {
	v := viper.New()
	v.SetConfigFile(f)
	conf := &config.CmdOptions{}
	if err := v.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("Fatal error unmarshalling load file: %w", err)
	}
	return conf, nil
}

// GetCurrentExecDir get exec dir
func GetCurrentExecDir() (dir string, err error) {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		// fmt.Printf("exec.LookPath(%s), err: %s\n", os.Args[0], err)
		return "", err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		// fmt.Printf("filepath.Abs(%s), err: %s\n", path, err)
		return "", err
	}
	dir = filepath.Dir(absPath)
	return dir, nil
}
