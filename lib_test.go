package lib_test

import (
	"get-tweets/lib"
	"github.com/ChimeraCoder/anaconda"
	"testing"
)

func Test_removeDuplicate(t *testing.T) {
	// list := []anaconda.Tweet{, , , }
	// expected := {"Go":struct {}{}, "Java":struct {}{}, "Python":struct {}{}, "Ruby":struct {}{}}
	output := lib.removeDuplicate(list)
	if expected != output {
		t.Errorf("")
	}
}
