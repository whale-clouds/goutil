package optional_test

import (
	"testing"

	"github.com/whale-clouds/goutil/dump"
	"github.com/whale-clouds/goutil/stdutil/optional"
)

func TestOptFactory_of(t *testing.T) {
	opt := optional.Of(nil)

	dump.P(opt.OrElseGet(34))
}
