// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	exampleIns := [][]string{{`[4,2,3,0,3,1,2]`, `5`}, {`[4,2,3,0,3,1,2]`, `0`}, {`[3,0,2,1,2]`, `2`}}
	exampleOuts := [][]string{{`true`}, {`true`}, {`false`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, canReach, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
