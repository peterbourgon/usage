package usage_test

import (
	"bytes"
	"flag"
	"strings"
	"testing"

	"github.com/peterbourgon/usage"
)

func TestEllipsis(t *testing.T) {
	var (
		fs  = flag.NewFlagSet("TestEllipsis", flag.ExitOnError)
		_   = fs.String("s", "", "string")
		buf bytes.Buffer
	)

	usage.ForWriter(&buf, fs, "TestEllipsis [flags]")()

	if needle, haystack := `-s ...`, buf.String(); !strings.Contains(haystack, needle) {
		t.Errorf("couldn't find needle %q in haystack:\n%s\n", needle, haystack)
	}
}
