package conf

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// MySSHCfg config_myssh
type MySSHCfg struct {
	Servers []ServerCfg `toml:"server"`
}

// ServerCfg server section
type ServerCfg struct {
	Host string `toml:"host"`
	Desc string `toml:"description"`
	Cred string `toml:"credential"`
}

// LoadMySSHConfig load config from file
func LoadMySSHConfig() *MySSHCfg {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(dir, ".ssh/config_myssh")
	// fmt.Println(filePath)

	var cfg MySSHCfg
	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
