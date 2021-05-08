package vim_builder

import (
	"fmt"
	"log"
	"os/exec"
)

func vim_install(vimPath string) {
	if !cmd_pathCheck("git") {
		return
	}
	err := exec.Command("git", "clone", "https://github.com/vim/vim", vimPath).Run()
	if err == nil {
		fmt.Println("error")
	} else {
		fmt.Println("Vim is installed")
	}
}

func cmd_pathCheck(cmd string) bool {
	if _, err := exec.LookPath(cmd); err != nil {
		log.Print(err)
		return false
	}
	return true
}
