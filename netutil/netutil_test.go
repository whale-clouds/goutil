package netutil_test

import (
	"net/url"
	"testing"

	"github.com/gookit/goutil/netutil"
	"github.com/stretchr/testify/assert"
)

func TestInternalIP(t *testing.T) {
	assert.NotEmpty(t, netutil.InternalIP())
}

func TestHttpQueryFromValues(t *testing.T) {
	type args struct {
		v      url.Values
		isSort bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				v: map[string][]string{
					"test": {"123"},
				},
				isSort: false,
			},
			want: "test=123",
		},
		{
			name: "testSort",
			args: args{
				v: map[string][]string{
					"b": {"123"},
					"a": {"456"},
				},
				isSort: true,
			},
			want: "a=456&b=123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, netutil.HttpQueryFromValues(tt.args.v, tt.args.isSort), "HttpQueryFromValues(%v)", tt.args.v)
		})
	}
}
