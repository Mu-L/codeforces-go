package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf659G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := 0
	f := 0
	for i, v := range a {
		ans += v - 1
		if i > 0 {
			ans += f * (min(a[i-1], v) - 1) % mod
		}
		if i == n-1 {
			continue
		}
		r := a[i+1]
		if v <= r {
			if i == 0 || a[i-1] >= v {
				f = (f + 1) * (v - 1)
			} else {
				f = f*(a[i-1]-1) + v - 1
			}
		} else {
			if i == 0 || a[i-1] >= v || a[i-1] >= r {
				f = (f + 1) * (r - 1)
			} else {
				f = f*(a[i-1]-1) + r - 1
			}
		}
		f %= mod
	}
	Fprint(out, ans%mod)
}

//func main() { cf659G(bufio.NewReader(os.Stdin), os.Stdout) }
