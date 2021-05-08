package builder

import "os/exec"

func vim_build(buildFlags []string, vimPath string) bool {
	vimPath += "/src"
	err := exec.Command("cd", vimPath).Run()
	if err != nil {
		return false
	}
	return true
}

func run_configure(buildFlags []string) bool {
	err := exec.Command("./configure").Run()
	if err != nil {
		return false
	}
	return true
}

func run_make() bool {
	err := exec.Command("make")
	if err != nil {
		return false
	}
	return true
}

func run_make_install() bool {
	err := exec.Command("make", "install")
	if err != nil {
		return false
	}
	return true
}
