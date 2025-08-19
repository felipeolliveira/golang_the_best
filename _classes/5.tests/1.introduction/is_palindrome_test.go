/*
* Testing in Go need some patterns
*
* - `[file]_test.go`: this sufix enable test to file
* - `func Test[name]`: this prefix run test for specific functions
*   - import "testing": default package to test
*   - `(t *testing.T)`: this args add all testing functionality
*   - testing has multiple functions to control a test
*     - https://pkg.go.dev/testing#pkg-types
*
* Comparing, Matchers, Assertions, Mocks
* - gotest.tools/assert: assertions, matchers
* - gomock: default mock package
* - testify: assert, require, mock, suite
* - reflect: deep comparing
* - go-cmp: deep comparing more than reflect
* - gomega/ginkgo: descritive tests.
*
* Cache is automatic
* - go clean --testcache
*
* */
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	str := IsPalindrome("ana")
	assert.True(t, str)
}
