package main

// https://space.bilibili.com/206214
func countValidSplit(nums []int, skip int) (cnt int) {
	n := len(nums)
	// suf[i] 表示后缀 [i,n-1]（除去 skip）的 GCD
	suf := make([]int, n+1)
	for j := n - 1; j >= 0; j-- {
		if j != skip {
			suf[j] = gcd(suf[j+1], nums[j])
		} else {
			suf[j] = suf[j+1]
		}
	}

	pre := 0
	for j, x := range nums {
		if j != skip {
			pre = gcd(pre, x)
			// 现在 pre 是前缀 [0,j]（除去 skip）的 GCD
			if pre == suf[j+1] {
				cnt++
			}
		}
	}
	return
}

func maxValidSplits(nums []int) int {
	ans := countValidSplit(nums, -1) // 不删除元素

	// countValidSplit 只会调用 O(log max(nums)) 次
	g := 0
	for i, x := range nums {
		if g > 0 && x%g == 0 { // x 不改变前缀 GCD
			continue
		}
		g = gcd(g, x)
		ans = max(ans, countValidSplit(nums, i))
	}

	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
