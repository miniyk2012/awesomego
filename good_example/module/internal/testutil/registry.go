package testutil

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

var tests = make(map[string][][]func(*testing.T))

func Register(ft ...func(*testing.T)) int {
	// Use the directory of the Caller's file
	// to map the tests. Why this can be useful
	// will be shown later.
	fmt.Printf("register %v", ft)
	_, f, _, _ := runtime.Caller(1)
	dir := filepath.Dir(f)

	tests[dir] = append(tests[dir], ft)

	// This is not necessary, but a function with a return
	// can be used in a top-level variable declaration which
	// can be used to avoid unnecessary init() functions.
	return 0
}
