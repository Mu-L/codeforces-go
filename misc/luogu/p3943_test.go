// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P3943
func Test_p3943(t *testing.T) {
	testCases := [][2]string{
		{
			`5 2 2 
1 5 
3 4`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p3943)
}
