package fsutil_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whale-clouds/goutil/fsutil"
)

func TestMimeType(t *testing.T) {
	assert.Equal(t, "", fsutil.MimeType(""))
	assert.Equal(t, "", fsutil.MimeType("not-exist"))
	assert.Equal(t, "image/jpeg", fsutil.MimeType("testdata/test.jpg"))

	buf := new(bytes.Buffer)
	buf.Write([]byte("\xFF\xD8\xFF"))
	assert.Equal(t, "image/jpeg", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte("text"))
	assert.Equal(t, "text/plain; charset=utf-8", fsutil.ReaderMimeType(buf))
	buf.Reset()

	buf.Write([]byte(""))
	assert.Equal(t, "", fsutil.ReaderMimeType(buf))
	buf.Reset()

	assert.True(t, fsutil.IsImageFile("testdata/test.jpg"))
}

func TestDiscardReader(t *testing.T) {
	sr := strings.NewReader("hello")
	fsutil.DiscardReader(sr)

	assert.Empty(t, fsutil.MustReadReader(sr))
}

// func TestDir(t *testing.T) {
//
// }

func TestTempDir(t *testing.T) {

}
