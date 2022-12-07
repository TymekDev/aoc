package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootDirectoryFromInput(t *testing.T) {
	root, err := rootDirectoryFromInput(input())
	require.NoError(t, err)

	tests := []struct {
		path   []string
		result int
	}{
		{[]string{"a", "e"}, 584},
		{[]string{"a"}, 94853},
		{[]string{"d"}, 24933642},
		{nil, 48381165},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			d, err := root.navigate(tt.path...)
			require.NoError(t, err)
			assert.Equal(t, tt.result, d.size())
		})
	}
}

func TestDirectoryTraverse(t *testing.T) {
	root, err := rootDirectoryFromInput(input())
	require.NoError(t, err)

	result := 0
	root.traverse(func(d *directory) {
		if s := d.size(); s <= 100000 {
			result += s
		}
	})

	assert.Equal(t, 95437, result)
}

func TestDirectoryNavigate(t *testing.T) {
	root := newDirectory(nil)
	c := root.mkdir("c")
	e := c.mkdir("e")

	tests := []struct {
		d      *directory
		path   []string
		result *directory
	}{
		{root, []string{}, root},
		{root, []string{"c", ".."}, root},
		{root, []string{"c"}, c},
		{root, []string{"c", "e"}, e},
		{c, []string{".."}, root},
		{c, []string{"/"}, root},
		{e, []string{".."}, c},
		{e, []string{"/"}, root},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			d, err := tt.d.navigate(tt.path...)
			require.NoError(t, err)
			assert.Equal(t, tt.result, d)
		})
	}
}

func TestMkdir(t *testing.T) {
	root := newDirectory(nil)
	c := root.mkdir("c")
	e := c.mkdir("e")

	require.Equal(t, c.parent, root)
	require.Equal(t, e.parent, c)
	require.Equal(t, e.parent.parent, root)
}

func TestDirectorySize(t *testing.T) {
	tests := []struct {
		d      *directory
		result int
	}{
		{
			&directory{
				content: map[string]sizer{
					"a": file(1),
				},
			},
			1,
		},
		{
			&directory{
				content: map[string]sizer{
					"a": file(1),
					"b": file(2),
				},
			},
			3,
		},
		{
			&directory{
				content: map[string]sizer{
					"a": file(1),
					"b": file(2),
					"c": &directory{
						content: map[string]sizer{
							"d": file(5),
						},
					},
				},
			},
			8,
		},
		{
			&directory{
				content: map[string]sizer{
					"a": file(1),
					"b": file(2),
					"c": &directory{
						content: map[string]sizer{
							"d": file(5),
							"e": &directory{
								content: map[string]sizer{
									"f": file(12),
								},
							},
						},
					},
				},
			},
			20,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tt.result, tt.d.size())
		})
	}
}

func input() []string {
	return strings.Split(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`, "\n")
}
