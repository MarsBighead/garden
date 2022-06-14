package config

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestDynamicOptionChange(t *testing.T) {
	re := require.New(t)
	env, err := GetEnvironment("")
	re.NotNil(t, err)
	env, err = GetEnvironment("config.yml")
	re.NotNil(t, err)
	env, err = GetEnvironment("config.yaml")
	re.Equal(env.Version, "v2.0.1")
	env, err = GetEnvironment("config.json")
	re.Equal(env.Version, "")
	_, err = GetEnvironment("env.go")
	re.EqualError(err, "Unsupport file format.")
}
