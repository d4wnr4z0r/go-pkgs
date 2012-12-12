// copyright 2012 Jacob Pipkin                                                      
// DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE                                  
// TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION              
//                                                                              
// 0. You just DO WHAT THE FUCK YOU WANT TO.                                    
//                                                                              
// Package find provides (a) mechanism(s) to search for files/folders under
// a directory tree.
package find

import (
	"os"
	"io"
	"regexp"
)

var results = make([]string, 0)

// recursive function that does the actual finding and processing
func Find(path string, fn func(s string, i os.FileInfo)) (err error) {
	dirh, err := os.Open(path)
	if err != nil {
		return err
	}

	infoSlice := make([]os.FileInfo, 1)
	infoSlice, err = dirh.Readdir(1)

	for err != io.EOF {
		info := infoSlice[0]
		fn(path +"/"+ info.Name(), info)

		if info.Mode().IsDir() {
			Find(path +"/"+ info.Name(), fn)
		}
		infoSlice, err = dirh.Readdir(1)
	}
	return nil
}

// generic find for names matching a given regex
func Regex(path string, r string) ([]string, error) {
	// zero out the results var
	results = []string{}

	regex, err := regexp.Compile(r)
	if err != nil {
		return nil, err
	}

	err = Find(path, func(s string, i os.FileInfo) {
		if regex.FindString(i.Name()) != "" {
			results = append(results, s)
		}
	})

	if err != nil {
		return nil, err
	}
	return results, nil
}

func DirRegex(path string, r string) ([]string, error) {
	// zero out the results var
	results = []string{}

	regex, err := regexp.Compile(r)
	if err != nil {
		return nil, err
	}

	err = Find(path, func(s string, i os.FileInfo) {
		if i.Mode().IsDir() {
			if regex.FindString(i.Name()) != "" {
				results = append(results, s)
			}
		}
	})

	if err != nil {
		return nil, err
	}
	return results, nil
}
