// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1974/D
// https://codeforces.com/problemset/status/1974/problem/D?friends=on
func Test_cf1974D(t *testing.T) {
	testCases := [][2]string{
		{
			`10
6
NENSNE
3
WWW
6
NESSWS
2
SN
2
WE
4
SSNN
4
WESN
2
SS
4
EWNN
4
WEWE`,
			`RRHRRH
NO
HRRHRH
NO
NO
RHRH
RRHH
RH
RRRH
RRHH`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1974D)
}
