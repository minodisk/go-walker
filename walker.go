package walker

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Walk(name string, cb func(string, os.FileInfo) (bool, error)) (err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return err
	}
	isContinuous, err := cb(name, fi)
	if err != nil {
		return err
	}
	if !isContinuous {
		return nil
	}
	if fi.IsDir() {
		return WalkUnder(name, cb)
	}
	return nil
}

func WalkUnder(dirname string, cb func(string, os.FileInfo) (bool, error)) (err error) {
	fis, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		name := filepath.Join(dirname, fi.Name())
		isContinuous, err := cb(name, fi)
		if err != nil {
			return err
		}
		if !isContinuous {
			return nil
		}
		if fi.IsDir() {
			err := WalkUnder(name, cb)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func FindDirs(dirname string) (dirnames []string) {
	Walk(dirname, func(name string, fi os.FileInfo) (bool, error) {
		if fi.IsDir() {
			dirnames = append(dirnames, name)
		}
		return true, nil
	})
	return dirnames
}

func FindFiles(dirname string) (filenames []string) {
	Walk(dirname, func(name string, fi os.FileInfo) (bool, error) {
		if !fi.IsDir() {
			filenames = append(filenames, name)
		}
		return true, nil
	})
	return filenames
}
