package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
