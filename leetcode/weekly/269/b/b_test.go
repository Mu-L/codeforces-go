// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[7,4,3,9,1,8,5,2,6]`, `3`, 
			`[-1,-1,-1,5,4,4,-1,-1,-1]`,
		},
		{
			`[100000]`, `0`, 
			`[100000]`,
		},
		{
			`[8]`, `100000`, 
			`[-1]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, getAverages, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-269/problems/k-radius-subarray-averages/
