//go:build e2e
// +build e2e

package e2e_test

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		fmt.Println(222)
	})
}
