// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc393/tasks/abc393_f
// 提交：https://atcoder.jp/contests/abc393/submit?taskScreenName=abc393_f
// 对拍：https://atcoder.jp/contests/abc393/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc393_f&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc393/submissions?f.Status=AC&f.Task=abc393_f&orderBy=source_length
func Test_f(t *testing.T) {
	testCases := [][2]string{
		{
			`5 3
2 4 1 3 3
2 5
5 2
5 3`,
			`2
1
2`,
		},
		{
			`10 8
2 5 6 5 2 1 7 9 7 2
7 8
5 2
2 3
2 6
7 3
8 9
9 6
8 7`,
			`4
1
1
2
1
5
3
4`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
