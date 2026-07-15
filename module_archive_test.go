package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func TestTrackedFilesHaveNoCaseInsensitiveCollisions(t *testing.T) {
	output, err := exec.Command("git", "ls-files", "-z").Output()
	if errors.Is(err, exec.ErrNotFound) {
		t.Skip("git is unavailable")
	}
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && bytes.Contains(exitErr.Stderr, []byte("not a git repository")) {
			t.Skip("source is not a Git checkout")
		}
		t.Fatalf("list tracked files: %v", err)
	}

	var paths []string
	for _, path := range bytes.Split(output, []byte{0}) {
		if len(path) > 0 {
			paths = append(paths, string(path))
		}
	}

	for i, left := range paths {
		for _, right := range paths[i+1:] {
			if left != right && strings.EqualFold(left, right) {
				t.Errorf("tracked paths collide on case-insensitive filesystems: %q and %q", left, right)
			}
		}
	}
}
