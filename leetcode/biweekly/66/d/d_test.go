// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[[0,1,1,0],[1,1,1,1]]`, 
			`2`,
		},
		{
			`[[1,1,1],[1,1,1]]`, 
			`2`,
		},
		{
			`[[1,0,1],[0,0,0],[1,0,1]]`, 
			`0`,
		},
		{
			`[[1,1,1,1,0],[1,1,1,1,1],[1,1,1,1,1],[0,1,0,0,1]]`, 
			`13`,
		},
		
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, countPyramids, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-66/problems/count-fertile-pyramids-in-a-land/
