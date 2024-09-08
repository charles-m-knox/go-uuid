package uuid_test

import (
	"regexp"
	"testing"

	uuid "github.com/charles-m-knox/go-uuid"
)

var r = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

func TestNew(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10000; i++ {
		u := uuid.New()

		if !r.MatchString(u) {
			t.Logf("test %v failed: regular expression did not match %v", i, u)
			t.FailNow()
		}
	}
}
