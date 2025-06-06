package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1815D(in io.Reader, out io.Writer) {
	const M = 998244353
	const inv2 = (M + 1) / 2
	var T, s, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &k)
		if k == 1 {
			Fprintln(out, s%M)
			continue
		}
		if k > 2 {
			Fprintln(out, (s%2+s)%M*(s/2%M+1)%M*inv2%M)
			continue
		}

		n := bits.Len(uint(s))
		f := make([][2]struct{ ways, f int }, n+1)
		f[n][0].ways = 1
		for i := n - 1; i >= 0; i-- {
			for carry := range 2 {
				res := &f[i][carry]
				for si := range 3 {
					if (si+carry)&1 == s>>i&1 {
						q := f[i+1][(si+carry)>>1]
						res.ways += q.ways
						res.f += q.f + si&1<<i%M*q.ways
					}
				}
				res.ways %= M
				res.f %= M
			}
		}
		Fprintln(out, f[0][0].f)
	}
}

//func main() { cf1815D(bufio.NewReader(os.Stdin), os.Stdout) }
