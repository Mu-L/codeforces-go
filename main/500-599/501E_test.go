// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/501/E
// https://codeforces.com/problemset/status/501/problem/E?friends=on
func Test_cf501E(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2 2 2`,
			`6`,
		},
		{
			`6
3 6 5 3 3 5`,
			`0`,
		},
		{
			`5
5 5 2 5 2`,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf501E)
}
