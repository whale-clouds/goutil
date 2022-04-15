package netutil_test

import (
	"fmt"
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
			assert.Equalf(t, tt.want, netutil.HttpQueryStringFromValues(tt.args.v, tt.args.isSort), "HttpQueryFromValues(%v)", tt.args.v)
		})
	}
}

func TestHttpQueryToValues(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    url.Values
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{s: "https://localhost?a=123&b=test"},
			want: map[string][]string{
				"a": {"123"},
				"b": {"test"},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := netutil.HttpQueryToValues(tt.args.s)
			if !tt.wantErr(t, err, fmt.Sprintf("HttpQueryToValues(%v)", tt.args.s)) {
				return
			}
			assert.Equalf(t, tt.want, got, "HttpQueryToValues(%v)", tt.args.s)
		})
	}
}
