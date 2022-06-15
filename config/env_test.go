package config

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestDynamicOptionChange(t *testing.T) {
	re := require.New(t)
	cfg := &Config{configFile: ""}
	env, err := cfg.NewEnvironment()
	re.NotNil(t, err)
	cfg.configFile = "config.yml"
	env, err = cfg.NewEnvironment()
	re.NotNil(t, err)
	cfg.configFile = "config.yaml"
	env, err = cfg.NewEnvironment()
	re.Equal(env.Version, "v2.0.1")
	cfg.configFile = "config.json"
	env, err = cfg.NewEnvironment()
	re.Equal(env.Version, "")
	cfg.configFile = "config.go"
	_, err = cfg.NewEnvironment()
	re.EqualError(err, "Unsupport file format.")
}
