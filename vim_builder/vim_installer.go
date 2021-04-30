package vim_builder

import (
	"os/exec"
)

func install() string{
    cmd, _ := exec.Command("ls", "-la").Output()
    return string(cmd)
}
