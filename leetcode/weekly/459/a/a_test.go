// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, checkDivisibility, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-459/problems/check-divisibility-by-digit-sum-and-product/
// https://leetcode.cn/problems/check-divisibility-by-digit-sum-and-product/
