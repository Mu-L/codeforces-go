// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/keyence2020/tasks/keyence2020_c
// 提交：https://atcoder.jp/contests/keyence2020/submit?taskScreenName=keyence2020_c
// 对拍：https://atcoder.jp/contests/keyence2020/submissions?f.LanguageName=Go&f.Status=AC&f.Task=keyence2020_c&orderBy=source_length
// 最短：https://atcoder.jp/contests/keyence2020/submissions?f.Status=AC&f.Task=keyence2020_c&orderBy=source_length
func Test_c(t *testing.T) {
	testCases := [][2]string{
		{
			`4 2 3`,
			`1 2 3 4`,
		},
		{
			`5 3 100`,
			`50 50 50 30 70`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
