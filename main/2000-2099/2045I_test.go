// Generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/2045/I
// https://codeforces.com/problemset/status/2045/problem/I?friends=on
func Test_cf2045I(t *testing.T) {
	testCases := [][2]string{
		{
			`5 4
3 2 1 3 2`,
			`13`,
		},
		{
			`3 3
1 1 1`,
			`2`,
		},
		{
			`3 3
2 1 2`,
			`5`,
		},
		{
			`4 4
4 3 1 2`,
			`12`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2045I)
}

func TestCompare_cf2045I(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		n := rg.Int(1, 4)
		m := rg.Int(1,4)
		rg.NewLine()
		rg.IntSlice(n, 1, m)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, cf2045I, cf2045IWA)
}

func cf2045IWA(in io.Reader, out io.Writer) {
	var n, m, v, ans, c int
	Fscan(in, &n, &m)
	pre := make([]int, m)
	f := make(fenwick45, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		v--
		if pre[v] > 0 {
			f.update(pre[v], -1)
		}
		f.update(i, 1)
		ans += f.query(pre[v]+1, i-1) * 2
		pre[v] = i
	}
	for _, i := range pre {
		if i > 0 {
			ans += f.query(i+1, n)
			c++
		}
	}
	Fprint(out, (ans+c)/2+c*(m-c))
}
