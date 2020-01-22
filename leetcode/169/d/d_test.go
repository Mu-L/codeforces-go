// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`["SEND","MORE"]`, `"MONEY"`}, {`["SIX","SEVEN","SEVEN"]`, `"TWENTY"`}, {`["THIS","IS","TOO"]`, `"FUNNY"`}, {`["LEET","CODE"]`, `"POINT"`}}
	exampleOuts := [][]string{{`true`}, {`true`}, {`true`}, {`false`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, isSolvable, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
