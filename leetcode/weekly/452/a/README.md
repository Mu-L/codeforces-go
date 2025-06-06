## 错误做法一

判断所有元素的乘积是否等于 $\textit{target}^2$。

这会错在 $\textit{nums}=[1,2,8]$，$\textit{target}=4$ 的情况。正确答案是 $\texttt{false}$。

## 错误做法二

如果有元素不是 $\textit{target}$ 的因子，则返回 $\texttt{false}$。否则，判断所有元素的乘积是否等于 $\textit{target}^2$。

这会错在 $\textit{nums}=[6,8,12]$，$\textit{target}=24$ 的情况。正确答案是 $\texttt{false}$。

## 方法一：递归

**前置题目**：[78. 子集](https://leetcode.cn/problems/subsets/)。

由于数组长度 $n$ 很小，可以枚举每个 $\textit{nums}[i]$ **分给第一个子集还是分给第二个子集**。

**细节**：为防止乘积过大导致溢出，可以在乘积大于 $\textit{target}$ 时返回 $\texttt{false}$。（Python 用户可以忽略）

如果两个子集的乘积都等于 $\textit{target}$，返回 $\texttt{true}$。

**注**：由于题目保证 $n\ge 3$ 且 $\textit{nums}$ 的所有元素互不相同，所以当一个子集是空集时，乘积为 $1$，另一个乘积一定大于 $1$，所以这种情况一定不符合要求。所以无需判断子集是空集的情况。

```py [sol-Python3]
class Solution:
    def checkEqualPartitions(self, nums: List[int], target: int) -> bool:
        def dfs(i: int, mul1: int, mul2: int) -> bool:
            if i == len(nums):
                return mul1 == mul2 == target
            return dfs(i + 1, mul1 * nums[i], mul2) or dfs(i + 1, mul1, mul2 * nums[i])
        return dfs(0, 1, 1)
```

```java [sol-Java]
class Solution {
    public boolean checkEqualPartitions(int[] nums, long target) {
        return dfs(0, 1, 1, nums, target);
    }

    private boolean dfs(int i, long mul1, long mul2, int[] nums, long target) {
        if (mul1 > target || mul2 > target) {
            return false;
        }
        if (i == nums.length) {
            return mul1 == target && mul2 == target;
        }
        return dfs(i + 1, mul1 * nums[i], mul2, nums, target) ||
               dfs(i + 1, mul1, mul2 * nums[i], nums, target);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool checkEqualPartitions(vector<int>& nums, long long target) {
        auto dfs = [&](this auto&& dfs, int i, long long mul1, long long mul2) -> bool {
            if (mul1 > target || mul2 > target) {
                return false;
            }
            if (i == nums.size()) {
                return mul1 == target && mul2 == target;
            }
            return dfs(i + 1, mul1 * nums[i], mul2) || dfs(i + 1, mul1, mul2 * nums[i]);
        };
        return dfs(0, 1, 1);
    }
};
```

```go [sol-Go]
func checkEqualPartitions(nums []int, target int64) bool {
	tar := int(target)
	var dfs func(int, int, int) bool
	dfs = func(i, mul1, mul2 int) bool {
		if mul1 > tar || mul2 > tar {
			return false
		}
		if i == len(nums) {
			return mul1 == tar && mul2 == tar
		}
		return dfs(i+1, mul1*nums[i], mul2) || dfs(i+1, mul1, mul2*nums[i])
	}
	return dfs(0, 1, 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n)$，其中 $n$ 是 $\textit{nums}$ 的长度。搜索树是一棵高为 $\mathcal{O}(n)$ 的二叉树，有 $\mathcal{O}(2^n)$ 个节点，所以遍历搜索树需要 $\mathcal{O}(2^n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。递归需要 $\mathcal{O}(n)$ 的栈空间。

## 方法二：二进制枚举

枚举下标全集 $U=\{0,1,2,\ldots, n-1\}$ 的**非空真子集** $S$，计算子集 $S$ 的 $\textit{nums}[i]$ 的乘积以及补集 $\complement_US$ 的 $\textit{nums}[i]$ 的乘积。

原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

**优化**：根据对称性，无需枚举 $n-1$ 在 $S$ 中的情况，也就是说，二进制的最高位一定是 $0$。

[本题视频讲解](https://www.bilibili.com/video/BV1Dz76zfEdi/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def checkEqualPartitions(self, nums: List[int], target: int) -> bool:
        n = len(nums)
        for s in range(1, 1 << (n - 1)):
            mul1 = mul2 = 1
            for i, x in enumerate(nums):
                if s >> i & 1:  # i 在集合 s 中
                    mul1 *= x
                else:  # i 在 s 的补集中
                    mul2 *= x
            if mul1 == target and mul2 == target:
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean checkEqualPartitions(int[] nums, long target) {
        int n = nums.length;
        int u = 1 << (n - 1);
        for (int s = 1; s < u; s++) {
            long mul1 = 1, mul2 = 1;
            for (int i = 0; i < n && mul1 <= target && mul2 <= target; i++) {
                if ((s >> i & 1) > 0) { // i 在集合 s 中
                    mul1 *= nums[i];
                } else { // i 在 s 的补集中
                    mul2 *= nums[i];
                }
            }
            if (mul1 == target && mul2 == target) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool checkEqualPartitions(vector<int>& nums, long long target) {
        int n = nums.size();
        int u = 1 << (n - 1);
        for (int s = 1; s < u; s++) {
            long long mul1 = 1, mul2 = 1;
            for (int i = 0; i < n && mul1 <= target && mul2 <= target; i++) {
                if (s >> i & 1) { // i 在集合 s 中
                    mul1 *= nums[i];
                } else { // i 在 s 的补集中
                    mul2 *= nums[i];
                }
            }
            if (mul1 == target && mul2 == target) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func checkEqualPartitions(nums []int, target int64) bool {
	tar := int(target)
	for s := 1; s < 1<<(len(nums)-1); s++ {
		mul1, mul2 := 1, 1
		for i, x := range nums {
			if s>>i&1 > 0 { // i 在集合 s 中
				mul1 = min(mul1*x, tar+1) // 与 tar+1 取 min，防止溢出
			} else { // i 在 s 的补集中
				mul2 = min(mul2*x, tar+1)
			}
		}
		if mul1 == tar && mul2 == tar {
			return true
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法三：折半枚举

首先，判断所有元素的乘积是否等于 $\textit{target}^2$，如果不等于，返回 $\texttt{false}$。

把 $\textit{nums}$ 均分成左右两个数组 $A$ 和 $B$。

从数组 $A$ 中，选择一些数，放入第一个集合，乘积为 $a_1$；其余元素放入第二个集合，乘积为 $b_1$。

从数组 $B$ 中，选择一些数，放入第一个集合，乘积为 $b_2$；其余元素放入第二个集合，乘积为 $a_2$。

题目要求 $a_1\cdot b_2 = b_1\cdot a_2$，变形得

$$
\dfrac{a_1}{b_1} = \dfrac{a_2}{b_2}
$$

我们可以用一个哈希集合维护所有 $\dfrac{a_1}{b_1}$ 的最简分数，另一个哈希集合维护所有 $\dfrac{a_2}{b_2}$ 的最简分数。

遍历第二个哈希集合，判断元素是否在第一个哈希集合中。或者说，两个集合的交集不为空。

**注**：不建议用浮点数计算，有精度误差。

```py [sol-Python3]
class Solution:
    def calc(self, nums: List[int], target: int) -> Set[Tuple[int, int]]:
        st = set()

        def dfs(i: int, a: int, b: int) -> None:
            if a > target or b > target:
                return
            if i == len(nums):
                g = gcd(a, b)
                st.add((a // g, b // g))  # 最简分数
                return
            dfs(i + 1, a * nums[i], b)
            dfs(i + 1, a, b * nums[i])

        dfs(0, 1, 1)
        return st

    def checkEqualPartitions(self, nums: List[int], target: int) -> bool:
        if prod(nums) != target ** 2:
            return False

        m = len(nums) // 2
        set1 = self.calc(nums[:m], target)
        set2 = self.calc(nums[m:], target)
        return len(set1 & set2) > 0  # 交集不为空
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    private record Pair(long a, long b) {
    }

    public boolean checkEqualPartitions(int[] nums, long target) {
        BigInteger prodAll = Arrays.stream(nums)
                .mapToObj(BigInteger::valueOf)
                .reduce(BigInteger.ONE, BigInteger::multiply);
        if (!prodAll.equals(BigInteger.valueOf(target).pow(2))) {
            return false;
        }

        int m = nums.length / 2;
        int[] left = Arrays.copyOfRange(nums, 0, m);
        int[] right = Arrays.copyOfRange(nums, m, nums.length);
        Set<Pair> set1 = calc(left, target);
        Set<Pair> set2 = calc(right, target);

        for (Pair p : set1) {
            if (set2.contains(p)) {
                return true;
            }
        }
        return false;
    }

    private Set<Pair> calc(int[] nums, long target) {
        Set<Pair> st = new HashSet<>();
        dfs(0, 1, 1, nums, target, st);
        return st;
    }

    private void dfs(int i, long a, long b, int[] nums, long target, Set<Pair> st) {
        if (a > target || b > target) {
            return;
        }
        if (i == nums.length) {
            long g = gcd(a, b);
            st.add(new Pair(a / g, b / g)); // 最简分数
            return;
        }
        dfs(i + 1, a * nums[i], b, nums, target, st);
        dfs(i + 1, a, b * nums[i], nums, target, st);
    }

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
    set<pair<long long, long long>> calc(const vector<int>& nums, long long target) {
        set<pair<long long, long long>> st; // 可以用 unordered_set + 自定义哈希
        auto dfs = [&](this auto&& dfs, int i, long long a, long long b) -> void {
            if (a > target || b > target) {
                return;
            }
            if (i == nums.size()) {
                long long g = gcd(a, b);
                st.emplace(a / g, b / g); // 最简分数
                return;
            }
            dfs(i + 1, a * nums[i], b);
            dfs(i + 1, a, b * nums[i]);
        };
        dfs(0, 1, 1);
        return st;
    }

public:
    bool checkEqualPartitions(vector<int>& nums, long long target) {
        __int128 prod_all = 1;
        for (int x : nums) {
            prod_all *= x;
        }
        if (prod_all != (__int128) target * target) {
            return false;
        }

        int m = nums.size() / 2;
        auto set1 = calc(vector<int>(nums.begin(), nums.begin() + m), target);
        auto set2 = calc(vector<int>(nums.begin() + m, nums.end()), target);

        for (auto& p : set1) {
            if (set2.contains(p)) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func calc(nums []int, target int) map[[2]int]struct{} {
	set := map[[2]int]struct{}{}
	var dfs func(int, int, int)
	dfs = func(i, a, b int) {
		if a > target || b > target {
			return
		}
		if i == len(nums) {
			g := gcd(a, b)
			set[[2]int{a / g, b / g}] = struct{}{} // 最简分数
			return
		}
		dfs(i+1, a*nums[i], b)
		dfs(i+1, a, b*nums[i])
	}
	dfs(0, 1, 1)
	return set
}

func checkEqualPartitions(nums []int, target int64) bool {
	prodAll := big.NewInt(1)
	for _, x := range nums {
		prodAll.Mul(prodAll, big.NewInt(int64(x)))
	}
	square := big.NewInt(target)
	square.Mul(square, square)
	if prodAll.Cmp(square) != 0 {
		return false
	}

	m := len(nums) / 2
	set1 := calc(nums[:m], int(target))
	set2 := calc(nums[m:], int(target))

	for p := range set1 {
		if _, ok := set2[p]; ok {
			return true
		}
	}
	return false
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^{n/2}\log \textit{target})$，其中 $n$ 是 $\textit{nums}$ 的长度。$\log \textit{target}$ 是计算 GCD 的复杂度。
- 空间复杂度：$\mathcal{O}(2^{n/2})$。

## 思考题

如果把题目改成分成非空前后缀呢？

欢迎在评论区分享你的思路/代码。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
