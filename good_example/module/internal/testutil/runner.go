package testutil

import "testing"

func TestAll(t *testing.T) {
	// TODO setup ...

	defer func() {
		// TODO teardown ...
	}()

	// run
	for _, dir := range tests {
		for _, file := range dir {
			for _, test := range file {
				test(t)
			}
		}
	}
}
