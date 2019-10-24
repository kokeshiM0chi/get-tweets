package lib_test

import (
	"testing"
	"github.com/ChimeraCoder/anaconda"
	"get-tweets/lib"
)

func Test_removeDuplicate(t *testing.T) {
	// list := []anaconda.Tweet{, , , }
	// expected := {"Go":struct {}{}, "Java":struct {}{}, "Python":struct {}{}, "Ruby":struct {}{}}
	output := lib.removeDuplicate(list)
	if expected != output {
		t.Errorf("")
	}
}