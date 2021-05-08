package installer

import (
	"os"
	"reflect"
	"testing"
)

func Test_vim_installed(t *testing.T) {
	vim_install("$HOME/vim")
	if f, err := os.Stat("$HOME/vim"); os.IsNotExist(err) || !f.IsDir() {
		t.Errorf("Vim's dir is not found")
	} else {
		t.Log("Vim's dir is found")
	}
}

func Test_cmd_exist(t *testing.T) {
	actual := true
	expect := cmd_pathCheck("git")

	if reflect.DeepEqual(actual, expect) {
		t.Log("git has found")
	} else {
		t.Errorf("git not fount")
	}
}
