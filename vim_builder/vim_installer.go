package vim_builder

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func vim_install(vimPath string) {
	if !cmd_pathCheck("git") || vimpath_exists(vimPath) {
		return
	}

	err := exec.Command("git", "clone", "https://github.com/vim/vim", vimPath).Run()
	if err != nil {
		fmt.Printf("Error, %v", err)
		fmt.Println("Failed to clone the Vim repository.")
		return
	}
	fmt.Println("Vim is installed")
	return
}

func cmd_pathCheck(cmd string) bool {
	if _, err := exec.LookPath(cmd); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func vimpath_exists(vimPath string) bool {
	if f, err := os.Stat(vimPath); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		fmt.Println("Vim dir has alrady exists")
		return true
	}
}
