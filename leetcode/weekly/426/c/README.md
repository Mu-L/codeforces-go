**核心思路**：对于 $\textit{answer}[i]$ 来说，新添加的边的一个端点必然是 $i$。为什么？因为用其他节点当作端点，只会让第二棵树的节点到 $i$ 距离变得更远，答案更小。

新添加的边，连到第二棵树的哪个节点上呢？

**暴力枚举**第二棵树的节点 $j$，用 DFS 计算距离 $j$ 不超过 $k-1$ 的节点个数 $\textit{cnt}_j$。这里 $k-1$ 是因为新添加的边也算在距离中。所有 $\textit{cnt}_j$ 取最大值，记作 $\textit{max}_2$。新添加的边就连到 $\textit{max}_2$ 对应的节点上。

同样地，暴力枚举第一棵树的节点 $i$，用 DFS 计算距离 $i$ 不超过 $k$ 的节点个数 $\textit{cnt}_i$。那么 $\textit{answer}[i] = \textit{cnt}_i + \textit{max}_2$。

**小结**：贪心思考后，把问题拆分成第一棵树与第二棵树**独立计算**，从而可以暴力通过。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1tAzoY1EUN/?t=16m03s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def build_tree(self, edges: List[List[int]], k: int) -> Callable[[int, int, int], int]:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        def dfs(x: int, fa: int, d: int) -> int:
            if d > k:
                return 0
            cnt = 1
            for y in g[x]:
                if y != fa:
                    cnt += dfs(y, x, d + 1)
            return cnt
        return dfs

    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]], k: int) -> List[int]:
        max2 = 0
        if k:
            dfs = self.build_tree(edges2, k - 1)  # 注意这里传的是 k-1
            max2 = max(dfs(i, -1, 0) for i in range(len(edges2) + 1))

        dfs = self.build_tree(edges1, k)
        return [dfs(i, -1, 0) + max2 for i in range(len(edges1) + 1)]
```

```java [sol-Java]
class Solution {
    public int[] maxTargetNodes(int[][] edges1, int[][] edges2, int k) {
        int max2 = 0;
        if (k > 0) {
            List<Integer>[] g = buildTree(edges2);
            for (int i = 0; i < edges2.length + 1; i++) {
                max2 = Math.max(max2, dfs(i, -1, 0, g, k - 1)); // 注意这里传的是 k-1
            }
        }

        List<Integer>[] g = buildTree(edges1);
        int[] ans = new int[edges1.length + 1];
        for (int i = 0; i < ans.length; i++) {
            ans[i] = dfs(i, -1, 0, g, k) + max2;
        }
        return ans;
    }

    private List<Integer>[] buildTree(int[][] edges) {
        List<Integer>[] g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return g;
    }

    private int dfs(int x, int fa, int d, List<Integer>[] g, int k) {
        if (d > k) {
            return 0;
        }
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, d + 1, g, k);
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<int>> buildTree(vector<vector<int>>& edges) {
        vector<vector<int>> g(edges.size() + 1);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }
        return g;
    }

    int dfs(int x, int fa, int d, vector<vector<int>>& g, int k) {
        if (d > k) {
            return 0;
        }
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, d + 1, g, k);
            }
        }
        return cnt;
    }

public:
    vector<int> maxTargetNodes(vector<vector<int>>& edges1, vector<vector<int>>& edges2, int k) {
        int max2 = 0;
        if (k > 0) {
            auto g = buildTree(edges2);
            for (int i = 0; i < edges2.size() + 1; i++) {
                max2 = max(max2, dfs(i, -1, 0, g, k - 1)); // 注意这里传的是 k-1
            }
        }

        auto g = buildTree(edges1);
        vector<int> ans(edges1.size() + 1);
        for (int i = 0; i < ans.size(); i++) {
            ans[i] = dfs(i, -1, 0, g, k) + max2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func buildTree(edges [][]int, k int) func(int, int, int) int {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int, int) int
	dfs = func(x, fa, d int) int {
		if d > k {
			return 0
		}
		cnt := 1
		for _, y := range g[x] {
			if y != fa {
				cnt += dfs(y, x, d+1)
			}
		}
		return cnt
	}
	return dfs
}

func maxTargetNodes(edges1, edges2 [][]int, k int) []int {
	max2 := 0
	if k > 0 {
		dfs := buildTree(edges2, k-1) // 注意这里传的是 k-1
		for i := range len(edges2) + 1 {
			max2 = max(max2, dfs(i, -1, 0))
		}
	}

	dfs := buildTree(edges1, k)
	ans := make([]int, len(edges1)+1)
	for i := range ans {
		ans[i] = dfs(i, -1, 0) + max2
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+m^2)$，其中 $n$ 是 $\textit{edges}_1$ 的长度，$m$ 是 $\textit{edges}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 直径优化

对于第一棵树来说，如果 [树的直径【基础算法精讲 23】](https://www.bilibili.com/video/BV17o4y187h1/)小于等于 $k$，那么就不需要暴力算了，第一棵树的每个节点都是目标节点。

同理，对于对于第二棵树来说，如果其直径小于 $k$，那么也不需要暴力算了，第二棵树的每个节点都是目标节点。

```py [sol-Python3]
class Solution:
    def calc_tree(self, edges: List[List[int]], k: int) -> Tuple[int, Callable[[int, int, int], int]]:
        g = [[] for _ in range(len(edges) + 1)]
        for x, y in edges:
            g[x].append(y)
            g[y].append(x)

        diameter = 0
        def dfs_diameter(x: int, fa: int) -> int:
            nonlocal diameter
            max_len = 0
            for y in g[x]:
                if y != fa:
                    sub_len = dfs_diameter(y, x) + 1
                    diameter = max(diameter, max_len + sub_len)
                    max_len = max(max_len, sub_len)
            return max_len
        dfs_diameter(0, -1)

        def dfs(x: int, fa: int, d: int) -> int:
            if d > k:
                return 0
            cnt = 1
            for y in g[x]:
                if y != fa:
                    cnt += dfs(y, x, d + 1)
            return cnt

        return diameter, dfs

    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]], k: int) -> List[int]:
        n, m = len(edges1) + 1, len(edges2) + 1

        max2 = 0
        if k:
            diameter, dfs = self.calc_tree(edges2, k - 1)
            if diameter < k:
                max2 = m  # 第二棵树的每个节点都是目标节点
            else:
                max2 = max(dfs(i, -1, 0) for i in range(m))

        diameter, dfs = self.calc_tree(edges1, k)
        if diameter <= k:
            return [n + max2] * n  # 第一棵树的每个节点都是目标节点
        return [dfs(i, -1, 0) + max2 for i in range(n)]
```

```java [sol-Java]
class Solution {
    public int[] maxTargetNodes(int[][] edges1, int[][] edges2, int k) {
        int n = edges1.length + 1;
        int m = edges2.length + 1;
    
        int max2 = 0;
        if (k > 0) {
            List<Integer>[] g = buildTree(edges2);
            diameter = 0;
            dfs(0, -1, g);

            if (diameter < k) {
                max2 = m; // 第二棵树的每个节点都是目标节点
            } else {
                for (int i = 0; i < m; i++) {
                    max2 = Math.max(max2, dfs(i, -1, 0, g, k - 1));
                }
            }
        }

        List<Integer>[] g = buildTree(edges1);
        diameter = 0;
        dfs(0, -1, g);     

        int[] ans = new int[n];
        if (diameter <= k) {
            Arrays.fill(ans, n + max2); // 第一棵树的每个节点都是目标节点
        } else {
            for (int i = 0; i < ans.length; i++) {
                ans[i] = dfs(i, -1, 0, g, k) + max2;
            }
        }
        return ans;
    }

    private List<Integer>[] buildTree(int[][] edges) {
        List<Integer>[] g = new ArrayList[edges.length + 1];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int[] e : edges) {
            int x = e[0];
            int y = e[1];
            g[x].add(y);
            g[y].add(x);
        }
        return g;
    }

    private int diameter;

    // 求树的直径
    private int dfs(int x, int fa, List<Integer>[] g) {
        int maxLen = 0;
        for (int y : g[x]) {
            if (y != fa) {
                int subLen = dfs(y, x, g) + 1;
                diameter = Math.max(diameter, maxLen + subLen);
                maxLen = Math.max(maxLen, subLen);
            }
        }
        return maxLen;
    }

    private int dfs(int x, int fa, int d, List<Integer>[] g, int k) {
        if (d > k) {
            return 0;
        }
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, d + 1, g, k);
            }
        }
        return cnt;
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<vector<int>> buildTree(vector<vector<int>>& edges) {
        vector<vector<int>> g(edges.size() + 1);
        for (auto& e : edges) {
            int x = e[0], y = e[1];
            g[x].push_back(y);
            g[y].push_back(x);
        }
        return g;
    }

    int dfs_diameter(int x, int fa, const vector<vector<int>>& g, int& diameter) {
        int max_len = 0;
        for (int y : g[x]) {
            if (y != fa) {
                int sub_len = dfs_diameter(y, x, g, diameter) + 1;
                diameter = max(diameter, max_len + sub_len);
                max_len = max(max_len, sub_len);
            }
        }
        return max_len;
    }

    int dfs(int x, int fa, int d, vector<vector<int>>& g, int k) {
        if (d > k) {
            return 0;
        }
        int cnt = 1;
        for (int y : g[x]) {
            if (y != fa) {
                cnt += dfs(y, x, d + 1, g, k);
            }
        }
        return cnt;
    }

public:
    vector<int> maxTargetNodes(vector<vector<int>>& edges1, vector<vector<int>>& edges2, int k) {
        int n = edges1.size() + 1;
        int m = edges2.size() + 1;
    
        int max2 = 0;
        if (k > 0) {
            auto g = buildTree(edges2);
            int diameter = 0;
            dfs_diameter(0, -1, g, diameter);

            if (diameter < k) {
                max2 = m; // 第二棵树的每个节点都是目标节点
            } else {
                for (int i = 0; i < m; i++) {
                    max2 = max(max2, dfs(i, -1, 0, g, k - 1));
                }
            }
        }

        auto g = buildTree(edges1);
        int diameter = 0;
        dfs_diameter(0, -1, g, diameter);

        vector<int> ans(n);
        if (diameter <= k) {
            ranges::fill(ans, n + max2); // 第一棵树的每个节点都是目标节点
        } else {
            for (int i = 0; i < n; i++) {
                ans[i] = dfs(i, -1, 0, g, k) + max2;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func calcTree(edges [][]int, k int) (diameter int, dfs func(int, int, int) int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfsDiameter func(int, int) int
	dfsDiameter = func(x, fa int) (maxLen int) {
		for _, y := range g[x] {
			if y != fa {
				subLen := dfsDiameter(y, x) + 1
				diameter = max(diameter, maxLen+subLen)
				maxLen = max(maxLen, subLen)
			}
		}
		return
	}
	dfsDiameter(0, -1)

	dfs = func(x, fa, d int) int {
		if d > k {
			return 0
		}
		cnt := 1
		for _, y := range g[x] {
			if y != fa {
				cnt += dfs(y, x, d+1)
			}
		}
		return cnt
	}

	return
}

func maxTargetNodes(edges1, edges2 [][]int, k int) []int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	max2 := 0
	if k > 0 {
		diameter, dfs := calcTree(edges2, k-1)
		if diameter < k {
			max2 = m // 第二棵树的每个节点都是目标节点
		} else {
			for i := range m {
				max2 = max(max2, dfs(i, -1, 0))
			}
		}
	}

	diameter, dfs := calcTree(edges1, k)
	ans := make([]int, n)
	if diameter <= k {
		for i := range ans {
			ans[i] = n + max2 // 第一棵树的每个节点都是目标节点
		}
	} else {
		for i := range ans {
			ans[i] = dfs(i, -1, 0) + max2
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+m^2)$，其中 $n$ 是 $\textit{edges}_1$ 的长度，$m$ 是 $\textit{edges}_2$ 的长度。**注**：在随机情况下，大小为 $n$ 的树的期望高度只有 $h=\Theta(\sqrt n)$，小于 $k$ 的期望值，大概率可以跑到 $\mathcal{O}(n+m)$。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
