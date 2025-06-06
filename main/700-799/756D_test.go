// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/756/D
// https://codeforces.com/problemset/status/756/problem/D?friends=on
func Test_cf756D(t *testing.T) {
	testCases := [][2]string{
		{
			`3
aaa`,
			`1`,
		},
		{
			`2
ab`,
			`3`,
		},
		{
			`4
babb`,
			`11`,
		},
		{
			`7
abacaba`,
			`589`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf756D)
}
