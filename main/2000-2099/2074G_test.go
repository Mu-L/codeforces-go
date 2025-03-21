// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2074/G
// https://codeforces.com/problemset/status/2074/problem/G?friends=on
func Test_cf2074G(t *testing.T) {
	testCases := [][2]string{
		{
			`6
3
1 2 3
4
2 1 3 4
6
2 1 2 1 1 1
6
1 2 1 3 1 5
9
9 9 8 2 4 4 3 5 3
9
9 9 3 2 4 4 8 5 3`,
			`6
24
5
30
732
696`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2074G)
}
