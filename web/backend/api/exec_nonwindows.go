//go:build !windows

package api

import (
	"os/exec"

	"github.com/BlackOceanX/picoclawquant/web/backend/utils"
)

func launcherExecCommand(name string, args ...string) *exec.Cmd {
	return utils.LauncherExecCommand(name, args...)
}

func applyLauncherProcAttrs(cmd *exec.Cmd) {
	utils.ApplyLauncherProcAttrs(cmd)
}
