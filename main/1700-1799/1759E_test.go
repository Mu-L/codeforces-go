// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1759/problem/E
// https://codeforces.com/problemset/status/1759/problem/E
func Test_cf1759E(t *testing.T) {
	testCases := [][2]string{
		{
			`8
4 1
2 1 8 9
3 3
6 2 60
4 5
5 1 100 5
3 2
38 6 3
1 1
12
4 6
12 12 36 100
4 1
2 1 1 15
3 5
15 1 13`,
			`4
3
3
3
0
4
4
3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1759E)
}
