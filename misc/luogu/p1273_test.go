// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P1273
func Test_p1273(t *testing.T) {
	testCases := [][2]string{
		{
			`5 3
2 2 2 5 3
2 3 2 4 3
3 4 2`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p1273)
}
