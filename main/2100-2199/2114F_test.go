// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2114/F
// https://codeforces.com/problemset/status/2114/problem/F?friends=on
func Test_cf2114F(t *testing.T) {
	testCases := [][2]string{
		{
			`8
4 6 3
4 5 3
4 6 2
10 45 3
780 23 42
11 270 23
1 982800 13
1 6 2`,
			`2
-1
-1
3
3
3
6
-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2114F)
}
