// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P2152
func Test_p2152(t *testing.T) {
	testCases := [][2]string{
		{
			`12
54`,
			`6`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p2152)
}
