package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize(t *testing.T) {
	type tc struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		expected  string
	}

	cases := []tc{
		{
			name:      "single file",
			path:      filepath.Join("testdata", "test.txt"),
			recursive: false,
			human:     false,
			all:       false,
			expected:  "6",
		},
		{
			name:      "directory (no hidden)",
			path:      filepath.Join("testdata", "dir2"),
			recursive: false,
			human:     false,
			all:       false,
			expected:  "0",
		},
		{
			name:      "directory include hidden",
			path:      filepath.Join("testdata", "dir2"),
			recursive: false,
			human:     false,
			all:       true,
			expected:  "7",
		},
		{
			name:      "directory recursive human all",
			path:      filepath.Join("testdata", "dir1"),
			recursive: true,
			human:     true,
			all:       true,
			expected:  "22B",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			size, err := GetPathSize(c.path, c.recursive, c.human, c.all)

			require.NoError(t, err)
			require.Equal(t, c.expected, size)
		})
	}
}
