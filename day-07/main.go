package main

import (
	"errors"
	"strconv"
	"strings"
)

func rootDirectoryFromInput(input []string) (*directory, error) {
	dir := newDirectory(nil)
	for _, line := range input {
		f := strings.Fields(line)
		switch f[0] {
		case "$": // `$ cd /` OR $ cd <name>` OR `$ ls`
			if f[1] == "ls" {
				continue
			}

			var err error
			dir, err = dir.navigate(f[2])
			if err != nil {
				return nil, err
			}

		case "dir": // `dir <dir name>`
			dir.mkdir(f[1])

		default: // `<size> <file name>`
			size, err := strconv.Atoi(f[0])
			if err != nil {
				return nil, err
			}

			dir.mkfile(f[1], size)
		}
	}

	return dir.navigate("/")
}

type sizer interface {
	size() int
}

type directory struct {
	parent  *directory
	content map[string]sizer
}

var _ sizer = (*directory)(nil)

func newDirectory(parent *directory) *directory {
	return &directory{parent: parent, content: map[string]sizer{}}
}

func (dir *directory) size() int {
	result := 0
	for _, s := range dir.content {
		result += s.size()
	}
	return result
}

func (dir *directory) navigate(p ...string) (*directory, error) {
	if len(p) == 1 && p[0] == "/" {
		for dir.parent != nil {
			dir = dir.parent
		}
		return dir, nil
	}

	result := dir
	for _, d := range p {
		if d == ".." {
			result = result.parent
			continue
		}

		next, ok := result.content[d]
		if !ok {
			return nil, errors.New("next directory not found")
		}

		nextDir, ok := next.(*directory)
		if !ok {
			return nil, errors.New("next directory is not a directory")
		}
		result = nextDir
	}

	return result, nil
}

func (dir *directory) mkdir(name string) *directory {
	result := newDirectory(dir)
	dir.content[name] = result
	return result
}

func (dir *directory) mkfile(name string, size int) {
	dir.content[name] = file(size)
}

type file int

var _ sizer = file(0)

func (f file) size() int {
	return int(f)
}
