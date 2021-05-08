package builder

import (
	"testing"

	"github.com/kazukazuinaina/vim-builder/installer"
)

func Test_vim_build(t *testing.T) {
	installer.Vim_install("$HOME/vim")
	if !vim_build([]string{"hoge"}, "$HOME/vim") {
		t.Error("cd失敗")
		return
	}
	t.Log("Success")
}

// func Test_run_configure(t *testing.T) {
// 	if !run_configure([]string{}) {
// 		t.Error("make失敗")
// 		return
// 	}
// 	t.Log("Success")
// }
