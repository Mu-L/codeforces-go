// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P2014
func Test_p2014(t *testing.T) {
	testCases := [][2]string{
		{
			`7  4
2  2
0  1
0  4
2  1
7  1
7  6
2  2`,
			`13`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p2014)
}
