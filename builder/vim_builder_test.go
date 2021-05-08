package builder

import (
	"testing"
)

func Test_vim_build(t *testing.T) {
	if !vim_build([]string{"hoge"}, "$HOME/vim") {
		t.Error("cd失敗")
		return
	}
	t.Log("Success")
}

func Test_run_configure(t *testing.T) {
	if !run_configure([]string{}) {
		t.Error("make失敗")
		return
	}
	t.Log("Success")
}
