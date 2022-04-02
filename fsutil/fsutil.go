package fsutil

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const (
	// MimeSniffLen sniff Length, use for detect file mime type
	MimeSniffLen = 512
)

// DiscardReader anything from the reader
func DiscardReader(src io.Reader) {
	_, _ = io.Copy(ioutil.Discard, src)
}

// OSTempFile create a temp file on os.TempDir()
//
// Usage:
// 	fsutil.OSTempFile("example.*.txt")
func OSTempFile(pattern string) (*os.File, error) {
	return ioutil.TempFile(os.TempDir(), pattern)
}

// TempFile is like ioutil.TempFile, but can custom temp dir.
//
// Usage:
// 	fsutil.TempFile("", "example.*.txt")
func TempFile(dir, pattern string) (*os.File, error) {
	return ioutil.TempFile(dir, pattern)
}

// OSTempDir creates a new temp dir on os.TempDir and return the temp dir path
//
// Usage:
// 	fsutil.OSTempDir("example.*")
func OSTempDir(pattern string) (string, error) {
	return ioutil.TempDir(os.TempDir(), pattern)
}

// TempDir creates a new temp dir and return the temp dir path
//
// Usage:
// 	fsutil.TempDir("", "example.*")
func TempDir(dir, pattern string) (string, error) {
	return ioutil.TempDir(dir, pattern)
}

// ExpandPath will parse `~` as user home dir path.
func ExpandPath(path string) string {
	path, _ = homedir.Expand(path)
	return path
}

// Realpath parse and get
func Realpath(path string) string {
	// TODO
	return path
}

// MimeType get File Mime Type name. eg "image/png"
func MimeType(path string) (mime string) {
	if path == "" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}

	return ReaderMimeType(file)
}

// ReaderMimeType get the io.Reader mimeType
//
// Usage:
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return
// 	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) {
	var buf [MimeSniffLen]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}

	return http.DetectContentType(buf[:n])
}

// RealWorkerFilePath 获取真实工作路径
func RealWorkerFilePath() string {
	executable, err := os.Executable()
	if err != nil {
		return ""
	}
	exPath := filepath.Dir(executable)
	dir, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		return ""
	}
	return dir
}
