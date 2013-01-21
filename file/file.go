// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO. 
//
// Package file provides wrappers around os/io calls for convenience

package file

import (
	"errors"
	"os"
	"io"
	"strings"
	"../find"
	"regexp"
)

// file.IsDir returns whether or not the given string is a directory
func IsDir(s string) bool {
	info, err := os.Stat(s)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// destination can be either a directory or a file; if it is a directory, src
// is copied into that directory with the same file name. otherwise, it is
// copied into the specified file.
func Copy(s, d string) error {
	if s == "" {
		return errors.New("empty src provided to file.Copy")
	}
	if d == "" {
		return errors.New("empty dst provided to file.Copy")
	}

	src, err := os.Open(s)
	if err != nil {
		return err
	}

	// if the destination is a directory, copy src into it. otherwise, copy src
	// to a file with the destination name. return an error if we can't create
	// that file.
	if IsDir(d) {
		paths := strings.Split(s, "/")
		file := paths[len(paths)-1]
		if file == "" {
			file = paths[len(paths)-2]
		}
		if file == "" {
			return errors.New("invalid src provided to file.Copy")
		}
		dst, err := os.Create(d +"/"+ file)
		if err != nil {
			return err
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
		return nil
	} else {
		dst, err := os.Create(d)
		if err != nil {
			return err
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}

// copy src directory, s, to dst directory, d; d will be created if it does not
// already exist.
func DirCopy(s, d string) error {
	if s == "" {
		return errors.New("empty src provided to file.DirCopy")
	}
	if d == "" {
		return errors.New("empty dst provided to file.DirCopy")
	}

	// create dst if it does not exist
	info, err := os.Stat(d)
	if err != nil {
		info, err = os.Stat(s)
		if err != nil {
			return err
		}
		err = os.MkdirAll(d, info.Mode())
		if err != nil {
			return err
		}
	}

	// recurse over the src tree, creating any required directories and copying
	// any files into those directories
	err = find.Find(s, func(p string, i os.FileInfo) {
		regex, _ := regexp.Compile(s)
		dst := regex.ReplaceAllLiteralString(p, d)
		if i.IsDir() {
			info, _ := os.Stat(p)
			os.MkdirAll(dst, info.Mode())
		} else {
			Copy(p, dst)
		}
	})

	if err != nil {
		return err
	}
	return nil
}
