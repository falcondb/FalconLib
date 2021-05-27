package exercise

import (
	"math/rand"
	"sort"
	"strings"
)

func lowToUpCase(in string) string {
	res := make([]byte, 0)

	for _, c := range in {
		if c >= 'a' && c <= 'z' {
			c -= 32
			res = append(res, byte(c))
		}
	}

	return string(res)
}

func mergeIntevals (ints []interval) []interval {
	if ints == nil {
		return nil
	}

	res := make([]interval, 0)

	if len(ints) == 0 {
		return res
	}

	cur := ints[0]

	for i, v := range ints {
		if i == 0 {
			continue
		}
		switch {
		case cur.end < v.st:
			res = append(res, cur)
			cur = v
		case cur.end >= v.st && cur.end < v.end:
			cur.end = v.end
		}
	}

	res = append(res, cur)
	return res
}

func rotateImage(im [][]uint) {
	w := len(im)

	for l := 0; l <= w >> 1; l++ {

		for p := l; p < w - l - 1; p++ {
			// top (l, p) ==> (p, len
			im[p][w-l-1], im[l][p] = im[l][p], im[p][w-l-1]
			// right
			im[w-l-1][w-p-1], im[l][p] = im[l][p], im[w-p-1][w-l-1]

			im [w-p-1][l], im[l][p] = im[l][p], im [w-p-1][l]
		}
	}
}

/**
91. Decode Ways
Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
 */

func numDecodings(s string) int {

	if len(s) == 0{
		return 0
	}

	buf := make([]int, len(s) + 1)
	buf[len(buf) - 1] = 1

	if s[len(s) - 1] != '0' {
		buf[len(s) - 1] = 1
	} else {
		buf[len(s) - 1] = -1
	}

	for i := len(s) - 2; i >= 0; i-- {
		switch {
		case s[i] == '0':
			if buf[i+1] <= 0 {
				return 0
			} else {
				buf[i] = -buf[i+1]
				buf[i+1] = 0
			}
		case s[i] > '2':
			if buf[i+1] <= 0 {
				return 0
			} else {
				buf[i] = buf[i+1]
			}
		case s[i] == '2':
			if buf[i+1] <= 0 {
				buf[i] = -buf[i+1]
				buf[i+1] = 0
			} else if s[i + 1] > '6' {
				buf[i] = buf[i+1]
			} else {
				buf[i] = buf[i+1] + buf[i + 2]
			}
		case s[i] == '1':
			if buf[i+1] <= 0 {
				buf[i] = -buf[i+1]
				buf[i+1] = 0
			} else {
				buf[i] = buf[i+1] + buf[i+2]
			}
		}
	}

	if buf[0] > 0 {
		return buf[0]
	} else {
		return 0
	}
}

/**
380. Insert Delete GetRandom O(1)
 */

type RandomizedSet struct {
	l []int
	m map[int]int
}


/** Initialize your data structure here. */
func RSConstructor() RandomizedSet {
	return RandomizedSet{make([]int,0), make(map[int]int)}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.m[val]

	if !ok {
		rs.l = append(rs.l, val)
		rs.m[val] = len(rs.l) - 1
		return true
	} else {
		return false
	}
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.m[val]

	if !ok {
		return false
	} else {
		delete(rs.m, val)
		if idx == len(rs.l) - 1 {
			rs.l = rs.l[:len(rs.l) - 1]
		} else {
			rs.l[idx] = rs.l[len(rs.l) - 1]
			rs.m[rs.l[idx]] = idx
			rs.l = rs.l[:len(rs.l) - 1]
		}
		return true
	}
}


/** Get a random element from the set. */
func (rs *RandomizedSet) GetRandom() int {
	return rs.l[rand.Intn(len(rs.l))]
}

/**
1249. Minimum Remove to Make Valid Parentheses
 */

func minRemoveToMakeValid(s string) string {
	moved := make([]int, 0)
	left := make([]int, 0)
	for i, c := range s {
		switch c {
		case '(':
			left = append(left, i)
		case ')':
			if len(left) == 0 {
				moved = append(moved, i)
			} else {
				left = left[:len(left) -1]
			}
		}
	}

	for _, i := range left {
		moved = append(moved, i)
	}

	sort.Ints(moved)

	res := make([]byte,0)

	for i, m := 0, 0; i < len(s); i++ {
		if m == len(moved) {
			res = append(res, s[i:]...)
			break
		}
		if i != moved[m] {
			res = append(res, s[i])
		} else {

			m++
		}
	}

	return string(res)
}

/**
560. Subarray Sum Equals K
 */

func subarraySum(ns []int, k int) int {
	if ns == nil || len(ns) == 0 {
		return 0
	}
	for i := 1; i < len(ns); i++ {
		ns[i] = ns[i-1] + ns[i]
	}

	sc := make(map[int]int)
	res := 0

	sc[0] = 1
	for _, s := range ns {
		c, ok := sc[s-k]
		if ok {
			res += c
		}
		_, ok = sc[s]
		if ok {
			sc[s]++
		} else {
			sc[s] = 1
		}
	}

	return res
}


/**
939. Minimum Area Rectangle
 */
func minAreaRect(points [][]int) int {
	pm := make(map[int]bool)
	res := 0xffffff
	for _, p := range points {
		pm[p[0] * 40001 + p[1]] = true
	}

	for i:= 0; i < len(points); i++ {
		for j:=i; j < len(points); j++ {
			if points[i][0] == points[j][0] || points[i][1] == points[j][1] {
				continue
			}

			_, ok1 := pm[points[i][0] * 40001 + points[j][1]]
			_, ok2 := pm[points[j][0] * 40001 + points[i][1]]

			if ok1 && ok2 {
				area := (points[i][0] - points[j][0]) * (points[i][1] - points[j][1])
				if area < 0 {
					area = -area
				}
				if area < res {
					res = area
				}

			}
		}
	}
	if res == 0xffffff {
		return 0
	}
	return res
}

/**
1711. Count Good Meals
 */

func countPairs(deliciousness []int) int {
	dc := make(map[int]int)
	dd := make([]int, 0)

	for _, d := range deliciousness {
		_, ok := dc[d]
		if !ok {
			dc[d] = 1
			dd = append(dd, d)
		} else {
			dc[d]++
		}
	}

	sct, dct := 0, 0

	for k, v := range dc {
		for i := uint(0); i <= 21; i++ {
			t := 1 << i - k

			c, ok := dc[t]
			if ok {
				if t == k {
					sct += v * (v-1) >> 1
				} else {
					dct += v * c
				}
			}
		}
	}

	return sct + (dct >> 1)
}


/**
1010. Pairs of Songs With Total Durations Divisible by 60
 */

func numPairsDivisibleBy60(time []int) int {
	wc := make([]int, 60)
	res := 0

	for _, t := range time {
		if t % 60 == 0 {
			res += wc[0]
		} else {
			res += wc[60 - t % 60]
		}
		wc[t % 60]++
	}

	return  res
}

/**
200. Number of Islands
 */

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	res := 0
	for i:=0; i < len(grid); i++ {
		for j:=0; j < len(grid[0]); j++ {
			if grid[i][j] == byte(1) {
				res++
				explore(grid, i, j)
			}
		}
	}

	 return res
}

func explore(g [][]byte, i, j int) {
	g[i][j] = byte(2)
	if i > 0 && g[i-1][j] == byte(1) {
		explore(g, i-1, j)
	}
	if i < len(g) && g[i+1][j] == byte(1) {
		explore(g, i+1, j)
	}
	if j < len(g) && g[i][j+1] == byte(1) {
		explore(g, i, j+1)
	}

	if j > 0 && g[i][j-1] == byte(1) {
		explore(g, i, j-1)
	}
}

/**
937. Reorder Data in Log Files
 */
type dwstring []string

func (ws dwstring) Less(i, j int) bool {
	si := strings.IndexByte(ws[i], ' ')
	sj := strings.IndexByte(ws[i], ' ')
	idi := ws[i][:si]
	idj := ws[j][:sj]


	isdi, isdj := false, false
	if ws[i][si+1] >= '0' && ws[i][si+1] <= '9' {
		isdi = true
	}
	if ws[j][sj+1] >= '0' && ws[j][sj+1] <= '9' {
		isdj = true
	}

	if isdi && !isdj {
		return false
	}
	if !isdi && isdj {
		return true
	}
	if isdi && isdj {
		return idi < idj
	} else {
		if idi == idj {
			return ws[i][:si] < ws[j][:sj]
		}
		return idi < idj
	}


}
func (ws dwstring) Len() int { return len(ws) }
func (ws dwstring) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func reorderLogFiles(logs []string) []string {
	sort.Sort((dwstring)(logs))
	return logs
}

/**
572. Subtree of Another Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	cands := make([]*TreeNode, 0)
	cands = append(cands, s)
	for len(cands) != 0 {
		if isSameTree(cands[0], t) {
			return true
		}
		cands = cands[1:]
		cands = append(cands, s.Left, s.Right)
	}

	return false
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if t == nil {
		return false
	}

	return s.Val == s.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, s.Right)
}

/**
994. Rotting Oranges
 */
func orangesRotting(grid [][]int) int {
	rots := make([][]int, 0)
	fsct := 0
	days := -1

	for i := 0; i < len(grid); i++ {
		for j:= 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				rots = append(rots, []int{i, j})
			} else if grid[i][j] == 1 {
				fsct++
			}
		}
	}

	if len(rots) == 0 && fsct == 0 {
		return 0
	}

	for len(rots) != 0 {
		nrots := make([][]int, 0)
		days++
		for _, coor := range rots {
			if coor[0] > 0 && grid[coor[0] - 1][coor[1]] == 1 {
				fsct--
				grid[coor[0] - 1][coor[1]] = 2
				nrots = append(nrots, []int{coor[0] - 1,coor[1]})
			}
			if coor[0] < len(grid) - 1 && grid[coor[0] + 1][coor[1]] == 1 {
				fsct--
				grid[coor[0] + 1][coor[1]] = 2
				nrots = append(nrots, []int{coor[0] + 1,coor[1]})
			}
			if coor[1] > 0 && grid[coor[0]][coor[1] - 1] == 1 {
				fsct--
				grid[coor[0]][coor[1]-1] = 2
				nrots = append(nrots, []int{coor[0], coor[1] - 1})
			}
			if coor[1] < len(grid[0]) - 1 && grid[coor[0]][coor[1] + 1] == 1 {
				fsct--
				grid[coor[0]][coor[1] + 1] = 2
				nrots = append(nrots, []int{coor[0],coor[1] + 1})
			}
		}
		rots = nrots
	}

	if fsct > 0 {
		return -1
	}
	return days
}