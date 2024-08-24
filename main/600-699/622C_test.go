// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/622/C
// https://codeforces.com/problemset/status/622/problem/C
func Test_cf622C(t *testing.T) {
	testCases := [][2]string{
		{
			`6 4
1 2 1 1 3 5
1 4 1
2 6 2
3 4 1
3 4 2`,
			`2
6
-1
4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf622C)
}
