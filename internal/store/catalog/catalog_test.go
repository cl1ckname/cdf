package catalog_test

import (
	"io/fs"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/cl1ckname/cdf/internal/store/catalog"
)

func TestPaths(t *testing.T) {
	t.Parallel()

	root := "/home"
	c := catalog.New(root, nil)
	if filename := c.Root(); filename != root {
		t.Errorf("invalid root path: %s", filename)
	}
	if filename := c.Marks(); filename != filepath.Join(root, catalog.MarksFilename) {
		t.Errorf("invalid filename: %s", filename)
	}
}

func TestEnsureRoot(t *testing.T) {
	t.Parallel()

	home := "home"
	tests := []struct {
		name  string
		fs    fstest.MapFS
		root  string
		mkdir *string
		err   bool
	}{
		{
			name:  "already exists",
			root:  home,
			fs:    fstest.MapFS{home: &fstest.MapFile{Mode: fs.ModeDir}},
			mkdir: nil,
			err:   false,
		},
		{
			name:  "should be created",
			root:  home,
			fs:    fstest.MapFS{},
			mkdir: &home,
			err:   false,
		},
		{
			name:  "exists but file",
			root:  home,
			fs:    fstest.MapFS{home: &fstest.MapFile{Mode: fs.ModeTemporary}},
			mkdir: nil,
			err:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := new(fsmock)
			f.MapFS = test.fs
			c := catalog.New(test.root, f)
			err := c.EnsureRoot()
			if (err != nil) != test.err {
				t.Fatalf("should be error")
			}
			if mkd := test.mkdir; mkd != nil {
				if len(f.mkdirs) == 0 {
					t.Fatalf("no mkdirs but expected")
				}
				if *mkd != f.mkdirs[0] {
					t.Fatalf("wrong mkdir: %s vs %s", *mkd, f.mkdirs[0])
				}
			} else if len(f.mkdirs) != 0 {
				t.Fatalf("mkdir should not be caused")
			}
		})
	}
}

func TestEnsureMarks(t *testing.T) {
	t.Parallel()

	home := "home"
	marks := filepath.Join(home, catalog.MarksFilename)

	tests := []struct {
		name  string
		fs    fstest.MapFS
		root  string
		touch *string
		err   bool
	}{
		{
			name:  "already exists",
			root:  home,
			fs:    fstest.MapFS{marks: &fstest.MapFile{}},
			touch: nil,
			err:   false,
		},
		{
			name:  "should be created",
			root:  home,
			fs:    fstest.MapFS{},
			touch: &marks,
			err:   false,
		},
		{
			name:  "exists but folder",
			root:  home,
			fs:    fstest.MapFS{marks: &fstest.MapFile{Mode: fs.ModeDir}},
			touch: nil,
			err:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := new(fsmock)
			f.MapFS = test.fs
			c := catalog.New(test.root, f)
			err := c.EnsureMarks()
			if (err != nil) != test.err {
				t.Fatalf("should be error")
			}
			if mkd := test.touch; mkd != nil {
				if len(f.touchs) == 0 {
					t.Fatalf("no touchs but expected")
				}
				if *mkd != f.touchs[0] {
					t.Fatalf("wrong mkdir: %s vs %s", *mkd, f.touchs[0])
				}
			} else if len(f.touchs) != 0 {
				t.Fatalf("mkdir should not be caused")
			}
		})
	}
}

type fsmock struct {
	fstest.MapFS
	mkdirs []string
	touchs []string
}

func (f *fsmock) Touch(path string, _ fs.FileMode) error {
	f.touchs = append(f.touchs, path)
	return nil
}

func (f *fsmock) Mkdir(path string, _ fs.FileMode) error {
	f.mkdirs = append(f.mkdirs, path)
	return nil
}
