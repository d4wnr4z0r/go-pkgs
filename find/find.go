// copyright 2012 Jacob Pipkin
//
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
// 0. You just DO WHAT THE FUCK YOU WANT TO.
//
// Package find provides mechanisms to search for files/folders under a
// directory tree.
package find

import (
	"os"
	"io"
	"regexp"
)

var results = make([]string, 0)

// this might be a bit of a nasty hack. I don't like it. but it works.
// making this OO might be a solution. not that I think OO makes sense here.
var first = true

func clearVars() {
	results = []string{}
	first = true
}

// function that does the actual finding and processing
func Find(path string, fn func(s string, i os.FileInfo)) (err error) {
	dirh, err := os.Open(path)
	if err != nil {
		return err
	}

	infoSlice := make([]os.FileInfo, 1)
	infoSlice, err = dirh.Readdir(1)

	if err != nil {
		return err
	}

	for err != io.EOF {
		info := infoSlice[0]
		if first {
			fn(path, info)
		} else {
			fn(path +"/"+ info.Name(), info)
		}
		first = false

		if info.Mode().IsDir() {
			Find(path +"/"+ info.Name(), fn)
		}
		infoSlice, err = dirh.Readdir(1)
	}
	return nil
}

// find for anything matching a given regex
func Regex(path string, r string) ([]string, error) {
	clearVars()

	regex, err := regexp.Compile(r)
	if err != nil {
		return nil, err
	}

	err = Find(path, func(s string, i os.FileInfo) {
		if regex.FindString(regexp.QuoteMeta(i.Name())) != "" {
			results = append(results, s)
		}
	})

	if err != nil {
		return nil, err
	}
	return results, nil
}

// find for directories matching a given regex
func DirRegex(path string, r string) ([]string, error) {
	clearVars()

	regex, err := regexp.Compile(r)
	if err != nil {
		return nil, err
	}

	err = Find(path, func(s string, i os.FileInfo) {
		if i.Mode().IsDir() {
			if regex.FindString(regexp.QuoteMeta(i.Name())) != "" {
				results = append(results, s)
			}
		}
	})

	if err != nil {
		return nil, err
	}
	return results, nil
}
