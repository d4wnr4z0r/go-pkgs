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

	if IsDir(d) {
		paths := strings.Split(s, "/")
		file := paths[len(paths)-1]
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
		} else {
			return nil
		}
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
