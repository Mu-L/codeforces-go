// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P2375
func Test_p2375(t *testing.T) {
	testCases := [][2]string{
		{
			`3
aaaaa
ab
abcababc`,
			`36
1
32`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p2375)
}
