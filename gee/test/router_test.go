package gee

import (
	"gee"
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(gee.ParsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(gee.ParsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(gee.ParsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}
