// PicoClawQuant - Ultra-lightweight personal AI agent
// License: MIT
//
// Copyright (c) 2026 PicoClawQuant contributors

package config

import (
	"os"
	"path/filepath"

	"github.com/BlackOceanX/picoclawquant/pkg"
)

// Runtime environment variable keys for the picoclawquant process.
// All keys use the PICOCLAWQUANT_ prefix.
const (
	// EnvHome overrides the base directory for all picoclawquant data
	// (config, workspace, skills, auth store, …).
	// Default: ~/.picoclawquant
	EnvHome = "PICOCLAWQUANT_HOME"

	// EnvConfig overrides the full path to the JSON config file.
	// Default: $PICOCLAWQUANT_HOME/config.json
	EnvConfig = "PICOCLAWQUANT_CONFIG"

	// EnvBuiltinSkills overrides the directory from which built-in
	// skills are loaded.
	// Default: <cwd>/skills
	EnvBuiltinSkills = "PICOCLAWQUANT_BUILTIN_SKILLS"

	// EnvBinary overrides the path to the picoclawquant executable.
	// Used by the web launcher when spawning the gateway subprocess.
	// Default: resolved from the same directory as the current executable.
	EnvBinary = "PICOCLAWQUANT_BINARY"

	// EnvGatewayHost overrides the host address for the gateway server.
	// Default: "localhost"
	EnvGatewayHost = "PICOCLAWQUANT_GATEWAY_HOST"
)

func GetHome() string {
	homePath, _ := os.UserHomeDir()
	if picoclawquantHome := os.Getenv(EnvHome); picoclawquantHome != "" {
		homePath = picoclawquantHome
	} else if homePath != "" {
		homePath = filepath.Join(homePath, pkg.DefaultPicoClawQuantHome)
	}
	if homePath == "" {
		homePath = "."
	}
	return homePath
}