package walker_test

import (
	"os"
	"sort"
	"testing"

	"github.com/minodisk/go-walker"
)

func TestWalkWithDirectory(t *testing.T) {
	expectedFiles := []string{"fixtures/bar/baz", "fixtures/bar/qux/quux", "fixtures/foo"}
	expectedDirs := []string{"fixtures", "fixtures/bar", "fixtures/bar/qux"}
	var actualFiles, actualDirs []string
	err := walker.Walk("fixtures", func(name string, fi os.FileInfo) (bool, error) {
		if fi.IsDir() {
			actualDirs = append(actualDirs, name)
		} else {
			actualFiles = append(actualFiles, name)
		}
		return true, nil
	})
	if err != nil {
		t.Errorf("should not return error: %v", err)
	}
	if !equalStrings(expectedFiles, actualFiles) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedFiles, actualFiles)
	}
	if !equalStrings(expectedDirs, actualDirs) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedDirs, actualDirs)
	}
}

func TestWalkWithFile(t *testing.T) {
	expectedFiles := []string{"fixtures/foo"}
	expectedDirs := []string{}
	var actualFiles, actualDirs []string
	err := walker.Walk("fixtures/foo", func(name string, fi os.FileInfo) (bool, error) {
		if fi.IsDir() {
			actualDirs = append(actualDirs, name)
		} else {
			actualFiles = append(actualFiles, name)
		}
		return true, nil
	})
	if err != nil {
		t.Errorf("should not return error: %v", err)
	}
	if !equalStrings(expectedFiles, actualFiles) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedFiles, actualFiles)
	}
	if !equalStrings(expectedDirs, actualDirs) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedDirs, actualDirs)
	}
}

func TestWalkUnderWithDirectory(t *testing.T) {
	expectedFiles := []string{"fixtures/bar/baz", "fixtures/bar/qux/quux", "fixtures/foo"}
	expectedDirs := []string{"fixtures/bar", "fixtures/bar/qux"}
	var actualFiles, actualDirs []string
	err := walker.WalkUnder("fixtures", func(name string, fi os.FileInfo) (bool, error) {
		if fi.IsDir() {
			actualDirs = append(actualDirs, name)
		} else {
			actualFiles = append(actualFiles, name)
		}
		return true, nil
	})
	if err != nil {
		t.Errorf("should not return error: %v", err)
	}
	if !equalStrings(expectedFiles, actualFiles) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedFiles, actualFiles)
	}
	if !equalStrings(expectedDirs, actualDirs) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedDirs, actualDirs)
	}
}

func TestWalkUnderWithFile(t *testing.T) {
	expectedFiles := []string{}
	expectedDirs := []string{}
	var actualFiles, actualDirs []string
	walker.WalkUnder("fixtures/foo", func(name string, fi os.FileInfo) (bool, error) {
		if fi.IsDir() {
			actualDirs = append(actualDirs, name)
		} else {
			actualFiles = append(actualFiles, name)
		}
		return true, nil
	})
	// Should return error in Linux, but doesn't in Mac OS X, in go 1.4.
	// See https://github.com/golang/go/issues/9789
	// if err == nil {
	// 	t.Errorf("should return error: %v", err)
	// }
	if !equalStrings(expectedFiles, actualFiles) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedFiles, actualFiles)
	}
	if !equalStrings(expectedDirs, actualDirs) {
		t.Errorf("wrong files: expected %v, but actual %v", expectedDirs, actualDirs)
	}
}

func TestFindDirs(t *testing.T) {
	expected := []string{"fixtures", "fixtures/bar", "fixtures/bar/qux"}
	actual := walker.FindDirs("fixtures")
	if !equalStrings(expected, actual) {
		t.Errorf("wrong files: expected %v, but actual %v", expected, actual)
	}
}

func TestFindFiles(t *testing.T) {
	expected := []string{"fixtures/bar/baz", "fixtures/bar/qux/quux", "fixtures/foo"}
	actual := walker.FindFiles("fixtures")
	if !equalStrings(expected, actual) {
		t.Errorf("wrong files: expected %v, but actual %v", expected, actual)
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i, va := range a {
		vb := b[i]
		if va != vb {
			return false
		}
	}
	return true
}

func in(arr []string, elem string) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}
