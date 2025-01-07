package mock

import (
	"io/fs"
	"time"
)

type FileInfo struct {
	FiModTime time.Time
	FiSys     any
	FiName    string
	FiSize    int64
	FiMode    fs.FileMode
	FiIsDir   bool
}

func (f FileInfo) Name() string       { return f.FiName }
func (f FileInfo) Size() int64        { return f.FiSize }
func (f FileInfo) Mode() fs.FileMode  { return f.FiMode }
func (f FileInfo) ModTime() time.Time { return f.FiModTime }
func (f FileInfo) IsDir() bool        { return f.FiIsDir }
func (f FileInfo) Sys() any           { return f.FiSys }
