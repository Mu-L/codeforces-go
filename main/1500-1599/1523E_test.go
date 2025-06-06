// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1523/E
// https://codeforces.com/problemset/status/1523/problem/E?friends=on
func Test_cf1523E(t *testing.T) {
	testCases := [][2]string{
		{
			`3
3 2
15 2
40 15`,
			`333333338
141946947
329622137`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1523E)
}
