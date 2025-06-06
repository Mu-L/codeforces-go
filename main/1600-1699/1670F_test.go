// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1670/F
// https://codeforces.com/problemset/status/1670/problem/F?friends=on
func Test_cf1670F(t *testing.T) {
	testCases := [][2]string{
		{
			`3 1 5 1`,
			`13`,
		},
		{
			`4 1 3 2`,
			`4`,
		},
		{
			`2 1 100000 15629`,
			`49152`,
		},
		{
			`100 56 89 66`,
			`981727503`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1670F)
}
