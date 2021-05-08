package builder

import "os/exec"

func vim_build(buildFlags []string, vimPath string) bool {
	err := exec.Command("cd", vimPath)
	if err == nil {
		return false
	}
	return true
}
