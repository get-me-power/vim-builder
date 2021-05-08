package vim_builder

import (
	"testing"
)

func Test_vim_build(t *testing.T) {
	if !vim_build([]string{"hoge"}, "$HOME/vim") {
		t.Error("あほしね")
		return
	}
	t.Log("Success")
}
