package httpreq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whale-clouds/goutil/dump"
	"github.com/whale-clouds/goutil/jsonutil"
	"github.com/whale-clouds/goutil/netutil/httpctype"
	"github.com/whale-clouds/goutil/netutil/httpreq"
)

func TestHttpReq_Send(t *testing.T) {
	resp, err := httpreq.New("https://httpbin.org").
		StringBody("hi").
		ContentType(httpctype.JSON).
		Send("/get")

	assert.NoError(t, err)
	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
	assert.False(t, httpreq.IsRedirect(sc))
	assert.False(t, httpreq.IsForbidden(sc))
	assert.False(t, httpreq.IsNotFound(sc))
	assert.False(t, httpreq.IsClientError(sc))
	assert.False(t, httpreq.IsServerError(sc))

	retMp := make(map[string]interface{})
	err = jsonutil.DecodeReader(resp.Body, &retMp)
	assert.NoError(t, err)
	dump.P(retMp)
}

func TestHttpReq_MustSend(t *testing.T) {
	resp := httpreq.New("https://httpbin.org").
		BytesBody([]byte("hi")).
		Method("POST").
		MustSend("/post")

	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
}
