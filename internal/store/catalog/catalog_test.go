package catalog_test

import (
	"io/fs"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/cl1ckname/cdf/internal/store/catalog"
)

func TestEnsureFolder(t *testing.T) {
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
			_, err := catalog.InitInFolder(test.root, f)
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

func TestEnsureFile(t *testing.T) {
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
			name:  "marks exists but folder",
			root:  marks,
			fs:    fstest.MapFS{marks: &fstest.MapFile{Mode: fs.ModeDir}},
			touch: nil,
			err:   true,
		},
		{
			name:  "marks already exists",
			root:  marks,
			fs:    fstest.MapFS{marks: &fstest.MapFile{}},
			touch: nil,
			err:   false,
		},
		{
			name:  "save file should be created",
			root:  marks,
			fs:    fstest.MapFS{},
			touch: &marks,
			err:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := new(fsmock)
			f.MapFS = test.fs
			err := catalog.EnsureFile(test.root, f)
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
