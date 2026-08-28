package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2244C(in io.Reader, out io.Writer) {
	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		g := gcd44(x, y)
		for i := range g {
			b := []int{}
			for j := i; j < n; j += g {
				b = append(b, a[j])
			}
			slices.Sort(b)
			for j := i; j < n; j += g {
				a[j] = b[0]
				b = b[1:]
			}
		}

		if slices.IsSorted(a) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2244C(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd44(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
