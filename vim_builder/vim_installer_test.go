package vim_builder

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_cmd(t *testing.T) {
    actual := install()
    expect := "hogee"

    if reflect.DeepEqual(actual, expect) {
        fmt.Println("Success")
    } else {
        t.Errorf("Expect %s,\n but got %s", expect, actual)
    }

}
