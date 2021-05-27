package exercise

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
)

/*
953. Verifying an Alien Dictionary

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.

 */
func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}
	co := make([]byte, 26)

	for i, c := range order {co[c-'a']=(byte)(i)}

	for i:= 1; i < len(words); i++ {
		j:= 0
		for ; j < len(words[i-1]) && j < len(words[i]) && co[words[i-1][j] - 'a'] == co[words[i][j] - 'a']; j++ {}
		if j < len(words[i-1]) && j < len(words[i]) {
			if co[words[i-1][j] - 'a']  > co[words[i][j] - 'a'] {
				return false
			}
		} else if j == len(words[i]) && j != len(words[i-1]) {
			return  false
		}
	}

	return true
}

/**
973. K Closest Points to Origin
 */

type EDpoints [][]int

func (h EDpoints) Len() int      {return len(h)}
func (h EDpoints) Less(i, j int) bool {
	return (h[i][0] * h[i][0] + h[i][1] * h[i][1]) < (h[j][0] * h[j][0] + h[j][1] * h[j][1])
}
func (h EDpoints) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func (h *EDpoints) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *EDpoints) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h EDpoints) Peek() interface{} {
	return h[len(h)-1]
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) < k {
		return points
	}

	edps := EDpoints(make([][]int,0))
	heap.Init(&edps)

	for i:=0; i < k; i++ {heap.Push(&edps, points[i])}
	for i:=k; i < len(points); i++ {
		peek := edps.Peek().([]int)
		if peek[0] * peek[0] + peek[1]*peek[1] > points[i][0] * points[i][0] + points[i][1] * points[i][1] {
			edps.Pop()
			edps.Push(points[i])
		}
	}

	res := make([][]int, 0)

	for edps.Len() > 0 {
		res = append(res, edps.Pop().([]int))
	}

	return res
}

/*
680. Valid Palindrome II

 */

func validPalindrome(wd string) bool {
	if palindrome(wd) {
		return true
	}

	s, e := 0, len(wd) - 1
	for wd[s] == wd[e] {
		s++
		e--
	}

	if palindrome(wd[s+1:e+1]) {
		return true
	}
	if palindrome(wd[s:e]) {
		return true
	}
	return false
}

func palindrome(wd string) bool {
	for s, e := 0, len(wd) - 1; s <= e; {
		if wd[s] != wd[e] {
			return  false
		}
		s++
		e--
	}
	return true
}

/*
238. Product of Array Except Self
 */
func productExceptSelf(nums []int) []int {
	// assume overflow won't happen

	tp := int64 (1)
	zi, zc := 0, 0
	res := make([]int, len(nums))

	for i, n := range nums {
		if n == 0 {
			zc++; zi = i
			if zc > 1 {
				break
			}
		} else {
			tp *= int64(n)
		}
	}

	if zc > 1 {
		return res
	}
	if zc == 1 {
		res[zi] = int(tp)
		return res
	}

	for i, n := range nums {
		res[i] = int(tp / int64(n))
	}
	return res
}

/**
415. Add Strings
 */
func addStrings(num1 string, num2 string) string {
	res := make([]byte, 0)
	c := uint8(0)

	if len(num1) < len(num2) { num2, num1 = num1, num2}

	for i:=0; i < len(num2); i++ {
		s := num1[len(num1) - 1 - i] - '0' + num2[len(num2) - 1 - i] - '0' + c
		c = s / 10
		s = s % 10
		res = append(res, s + '0')
	}

	for i := len(num2); i < len(num1); i++ {
		s := num1[len(num1) - 1 - i] - '0' + c
		c = s / 10
		s = s % 10
		res = append(res, s + '0')
	}
	if c == 1 {
		res = append(res, '1')
	}

	for i := 0; i < len(res) >> 1; i++ {res[i], res[len(res) - 1 - i ] = res[len(res) - 1 - i ], res[i]}
	return string(res)
}

/**
124. Binary Tree Maximum Path Sum
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var currentMaxPathSum int = -0xfffffff
func maxPathSum(root *TreeNode) int {
	currentMaxPathSum = root.Val
	maxSingPath(root)
	return currentMaxPathSum
}

func maxSingPath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lm := maxSingPath(root.Left)
	rm := maxSingPath(root.Right)
	rootMax := root.Val
	if lm > 0 {
		rootMax += lm
	}
	if rm > 0 {
		rootMax += rm
	}

	if rootMax > currentMaxPathSum {
		currentMaxPathSum = rootMax
	}

	if lm > 0 && lm >= rm {
		return lm + root.Val
	}

	if rm > 0 && rm >= lm {
		return rm + root.Val
	}

	return root.Val
}

/**
1762. Buildings With an Ocean View
 */
func findBuildings(heights []int) []int {
	res := make([]int, 0)
	if heights == nil || len(heights) == 0 {
		return res
	}

	hest := -1

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > hest {
			hest = heights[i]
			res = append(res, i)
		}
	}

	for i := 0; i < len(res) >> 1; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}

	return res
}

/**
56. Merge Intervals
 */

type sortItvs [][]int
func (itv sortItvs) Len() int {return len(itv)}
func (itv sortItvs) Swap(i,j int) {itv[i], itv[j] = itv[j], itv[i]}
func (itv sortItvs) Less(i, j int) bool {
	if itv[i][0] < itv[j][0] {
		return true
	} else if itv[i][0] == itv[j][0] {
		if itv[i][1] > itv[j][1] {
			return true
		}
	}
	return false
}

func merge(intervals [][]int) [][]int {

	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		return res
	}

	sort.Sort(sortItvs(intervals))
	cur := intervals[0]

	for i:=1; i < len(intervals); i++ {
		if cur[1] < intervals[i][0] {
			res = append(res, cur)
			cur = intervals[i]
		} else {
			cur[1] = intervals[i][1]
		}
	}

	res = append(res, cur)
	return res
}

/**
523. Continuous Subarray Sum
 */
func checkSubarraySum(nums []int, k int) bool {

	smap := make(map[int]int)
	csum := 0

	for i, c := range nums {
		csum += c

		if smap[csum % k] != 0 && i - smap[csum % k] > 0 || (csum % k == 0 && i > 0) {
			return true
		} else if smap[csum % k] == 0 {
			smap[csum % k] = i+1
		}
	}

	return false
}

type Node struct {
	Val int
	Left *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {

	if root == nil {
		return nil
	}

	head, tail := treeToDoublyListHelper(root)
	tail.Right = head
	head.Left = tail
	return head
}

func treeToDoublyListHelper(root *Node) (*Node, *Node) {
	var head, tail *Node
	if root.Left != nil {
		head, tail = treeToDoublyListHelper(root.Left)
		tail.Right = root
		root.Left = tail
		tail = root
	} else {
		head, tail = root, root
	}

	if root.Right != nil {
		rhead, rtail := treeToDoublyListHelper(root.Right)
		tail.Right = rhead
		rhead.Left = tail
		tail = rtail
	} else {
		tail.Right = nil
	}

	return head, tail
}

/**
199. Binary Tree Right Side View
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	res = rightSideViewHelper(root, res, 1)

	return  res
}

func rightSideViewHelper(root *TreeNode, view []int, depth int) []int {
	if root == nil {
		return view
	}

	if len(view) < depth {
		view = append(view, root.Val)
	}

	view = rightSideViewHelper(root.Right, view, depth + 1)
	view = rightSideViewHelper(root.Left, view, depth + 1)

	return view
}

/**
1650. Lowest Common Ancestor of a Binary Tree III
 */

//type Node struct {
//    Val int
//    Left *Node
//    Right *Node
//    Parent *Node
//}
//
//func lowestCommonAncestor(p *Node, q *Node) *Node {
//
//	if lcaHelp(p, q.Val) {
//		return p
//	}
//	if lcaHelp(q, p.Val) {
//		return q
//	}
//
//	for cur := p; cur.Parent != nil; cur = cur.Parent {
//		if cur.Parent.Left == cur && lcaHelp(cur.Parent.Right, q.Val) || cur.Parent.Right == cur && lcaHelp(cur.Parent.Left, q.Val) {
//			return cur.Parent
//		}
//	}
//
//	return nil
//}

func lcaHelp(r *Node, t int) bool {
	if r == nil {
		return false
	}

	if r.Val == t {
		return true
	}

	return lcaHelp(r.Left, t) || lcaHelp(r.Right, t)
}

type intHeap []int

func (h intHeap) Len()int {return len(h)}

func (h intHeap) Less(i, j int)bool {return h[i]<h[j]}

func (h intHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	hp := intHeap(make([]int, 0))

	for _, n := range nums {
		if len(hp) < k {
			heap.Push(&hp, n)
		} else {
			if hp[0] < n {
				heap.Push(&hp, n)
				if len(hp) > k {
					heap.Pop(&hp)
				}
			}
		}
	}

	return hp[0]
}

/**
125. Valid Palindrome
 */
func isPalindrome(wd string) bool {

	res := false
	for s, e := 0, len(wd) - 1; s < e; {
		 if !isAlpha(wd[s]) {
		 	s++
		 	continue
		 }
		 if !isAlpha(wd[e]) {
			 e--
			 continue
		 }
		 if !isSameAlpha(wd[s], wd[e]) {
		 	return false
		 } else {
		 	res = true
		 }
		 s++
		 e--
	}

	return res
}

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c<= 'Z'
}

func isSameAlpha(a, b byte) bool {
	if a >= 'A' && a<= 'Z' {
		a = a - 'A' + 'a'
	}
	if b >= 'A' && b<= 'Z' {
		b = b - 'A' + 'a'
	}

	return a == b
}

/**
270. Closest Binary Search Tree Value
 */
var cvRes int
func closestValue(root *TreeNode, target float64) int {
	cvRes = root.Val
	closestValueHelper(root, target)
	return cvRes
}

func closestValueHelper(root *TreeNode, t float64) {
	if root == nil {
		return
	}
	if math.Abs(float64(root.Val) - t) < math.Abs(float64(cvRes) - t) {
		cvRes = root.Val
	}

	if float64(root.Val) > t {
		closestValueHelper(root.Left, t)
	} else {
		closestValueHelper(root.Right, t)
	}
}

/**
249. Group Shifted Strings
 */
func groupStrings(sts []string) [][]string {
	gm := make(map[string][]string)

	for _, s := range sts {
		if len(s) == 0 {
			continue
		}

		d := s[0] - 'a'
		shed := make([]byte,0)
		for _, c := range s {
			if byte(c) - d < 'a' {
				shed = append(shed, byte(c)-d + 26)
			} else {
				shed = append(shed, byte(c)-d)
			}
		}

		gm[string(shed)] = append(gm[string(shed)], s)
	}

	res := make([][]string, 0)

	for _, v := range gm {
		res = append(res, v)
	}

	return res
}

/**
62. Unique Paths
 */

func uniquePaths(m int, n int) int {
	if m > n { m,n = n, m}

	pre := make([]int, m)
	cur := make([]int, m)
	for i:=0; i < m; i++ {pre[i] = 1}
	for j:=1; j < n; j++ {
		cur[0] = 1
		for i:=1; i < m; i++ { cur[i] = cur[i-1] + pre[i]}
		pre = cur
	}

	return cur[m-1]
}

/**
1047. Remove All Adjacent Duplicates In String
 */
func removeDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}

	res := []byte{s[0]}
	for t := 1; t < len(s); t++ {
		if len(res) >0 && s[t] == res[len(res) - 1] {
			res = res[:len(res) - 1]
		} else {
			res = append(res, s[t])
		}
	}

	return string(res)
}

/**
189. Rotate Array
 */

func rotate(nums []int, k int)  {
	if len(nums) <= k {
		return
	}

	sh := append([]int{}, nums[len(nums) - k:]...)

	for i:= len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i - k]
	}

	for i:= 0; i < k; i++ {
		nums[i] = sh[i]
	}
}


/**
42. Trapping Rain Water
 */
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	l, r := 0, len(height) - 1
	lm, rm := height[0], height[len(height) - 1]
	res := 0

	for l < r {
		if height[l] < height[r] {
			if lm < height[l] {lm = height[l]}
			res += lm - height[l]
			l++
		} else {
			if rm < height[r] {rm = height[r]}
			res += rm - height[r]
			r--
		}
	}
	return res
}

/**
1570. Dot Product of Two Sparse Vectors
 */


/**
78. Subsets
 */
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int,0)

	for i := 0; i < 1 << uint(len(nums)); i++ {
		ns := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i & (1 << uint(j)) != 0x0 {
				ns = append(ns, nums[j])
			}
		}
		res = append(res, ns)
	}

	return res
}

/**
17. Letter Combinations of a Phone Number
 */
var d2c = [][]byte {[]byte{}, []byte{'a', 'b', 'c'}, []byte{'e', 'f', 'g'}, []byte{'g', 'h', 'i'},
[]byte{'j', 'k', 'l'}, []byte{'m', 'n', 'o'}, []byte{'p', 'q', 'r', 's'}, []byte{'t', 'u', 'v'}, []byte{'w', 'x', 'y', 'z'}}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	res := make([]string, 0)
	for _, c := range d2c[digits[len(digits) - 1] - '1'] {
			for _, w := range letterCombinations(digits[:len(digits) - 1]) {
				res = append(res, string(append([]byte(w),c)))
			}
	}

	return res
}

/**
670. Maximum Swap
 */

func maximumSwap(num int) int {
	inbytes := []byte(strconv.Itoa(num))
	h, t := len(inbytes), len(inbytes)
	bg := len(inbytes) - 1

	for i:= len(inbytes) - 1; i >=0; i-- {
		if inbytes[i] > inbytes[bg] {
			bg = i
		}
		if inbytes[i] < inbytes[bg] {
			h, t = i, bg
		}
	}
	inbytes[h], inbytes[t] = inbytes[t], inbytes[h]
	res, _ := strconv.Atoi(string(inbytes))
	return res
}


/**
140. Word Break II
 */

func wordBreak(s string, wordDict []string) []string {
	wd := make(map[string]bool)
	for _, w := range wordDict {wd[w] = true}

	res := make([][]string, len(s))

	for i := 0; i < len(s); i++ {
		for j := 0 ; j < i ; j++ {
			_, ok := wd[s[i - j : i+1]]
			if ok {
				for _, ws := range res[i-j-1] {
					res[i] = append(res[i], fmt.Sprintf("%v %v",ws, s[ i - j : i + 1]))
				}
			}
		}

		_, ok := wd[s[: i+1]]
		if ok {
			res[i] = append(res[i], s[:i+1])
		}
	}

	return res[len(s) - 1]
}

/**
621. Task Scheduler
 */
func leastInterval(tasks []byte, n int) int {
	if tasks == nil || len(tasks) == 0 {
		return 0
	}

	tfs := make([]int, 26)

	for _, t := range tasks {
		tfs[t-'A']++
	}

	mf, nm := 0, 0
	for i:=0; i<26; i++ {
		if tfs[i] > mf {
			mf = tfs[i]
			nm = 1
		} else if tfs[i] == mf {
			nm++
		}
	}
	af := (mf - 1) * (n+1) + nm
	if len(tasks) > af {
		return len(tasks)
	} else {
		return af
	}
}

/**
1570. Dot Product of Two Sparse Vectors
 */
/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
type SparseVector struct {
	nz []int
	vs []int
}

func SVConstructor(nums []int) SparseVector {
	if nums == nil || len(nums) == 0 {
		return SparseVector{nil,nil}
	}
	nz, vs := make([]int, 0), make([]int, 0)
	for i, v := range nums {
		if v != 0 {
			nz = append(nz, i)
			vs = append(vs, v)
		}
	}

	return SparseVector{nz, vs}
}

// Return the dotProduct of two sparse vectors
func (l *SparseVector) dotProduct(r SparseVector) int {
	if l == nil || len(l.nz ) == 0 || len(r.nz) == 0 {
		return 0
	}

	li, ri := 0, 0
	res := 0

	for li < len(l.nz) && ri < len(r.nz) {
		if l.nz[li] == r.nz[ri] {
			res += l.vs[li] * r.vs[ri]
			li++
			ri++
		} else if l.nz[li] < r.nz[ri] {
			li++
		} else {
				ri++
		}
	}

	return res
}

/**
given a word A and string N, find if N contains words that have the same characters and frequency of characters as A.
 */

func wdContains(t, src string) []int {
	if len(t) == 0 || len(src) == 0 || len(src) < len(t) {
		return nil
	}

	cf := make([]int, 26)
	res := make([]int, 0)
	for _, c := range t {cf[c-'a']++}

	ccf := make([]int, 26)
	for i:=0; i < len(t); i++ {ccf[src[i]-'a']++}

	if wdfMatch(cf, ccf) {
		res = append(res, 0)
	}

	for i:= len(t); i < len(src); i++ {
		ccf[src[i-len(t)] - 'a']--
		ccf[src[i] - 'a']++
		if wdfMatch(cf, ccf){
			res = append(res, i - len(t) + 1)
		}
	}
	return res
}

func wdfMatch(cf, ccf []int) bool {
	m := true
	for ab := 0; ab < 26; ab++ {
		if cf[ab] != ccf[ab] {
			m = false
			break
		}
	}
	return m
}


/**
1541. Minimum Insertions to Balance a Parentheses String
 */


func minInsertions2(s string) int {
	lp, res := 0, 0
	pr := false

	for _, c := range s {
		if c == '(' {
			if pr {
				res++
				if lp == 0 {
					res++
				} else {
					lp--
				}
			}
			lp++
			pr = false
		} else {
			if pr {
				if lp > 0 {
					lp--
				} else {
					res++
				}
				pr = false
			} else {
				pr = true
			}
		}
	}

	if pr {
		res++
		if lp == 0 {
			res++
		} else {
			lp--
		}
	}
	if lp > 0 {
		res += lp << 1
	}

	return res
}


func minInsertions(s string) int {
	res := 0
	lc := 0
	for i:=0; i < len(s); {
		if s[i] == '(' {
			lc++
			i++
		} else {
			if i < len(s) - 1 && s[i+1] == ')' {
				i+=2
			} else {
				res++
				i++
			}
			if lc == 0 {
				res++
			} else {
				lc--
			}
		}
	}

	res += lc << 1
	return res
}

/**
314. Binary Tree Vertical Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var voResP [][]int
var voResN [][]int
func verticalOrder(root *TreeNode) [][]int {
	voResP = [][]int{}
	voResN = [][]int{[]int{}}
	res := make([][]int, 0)
	voHelper(root, 0)

	for i:= len(voResN) - 1; i >0; i--{
		res = append(res, voResN[i])
	}
	for i:= 0; i < len(voResP); i++ {
		res = append(res, voResP[i])
	}
	return res
}

func voHelper(r *TreeNode, ni int)  {
	if r == nil {
		return
	}

	if ni < 0 {
		if len(voResN) <= -ni {
			voResN = append(voResN, []int{})
		}
		voResN[-ni] = append(voResN[-ni], r.Val)
	} else {
		if len(voResP) <= ni {
			voResP = append(voResP, []int{})
		}
		voResP[ni] = append(voResP[ni], r.Val)
	}

	voHelper(r.Left, ni-1)
	voHelper(r.Right, ni+1)
}


/**
252. Meeting Rooms
 */
func canAttendMeetings2(intervals [][]int) bool {

	if len(intervals) < 2 {
		return true
	}
	sort.Sort(ints(intervals))

	for i:=0; i < len(intervals) -1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}

type ints [][]int

func (is ints) Len() int {return len(is)}
func (is ints) Less(i, j int) bool {
	if is[i][0] == is[j][0] {
		return is[i][1] < is[j][1]
	}
	return is[i][0] < is[j][0]
}
func (is ints) Swap(i, j int) {is[i], is[j] = is[j], is[i]}

/**
1361. Validate Binary Tree Nodes
 */
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	if n == 0 {
		return true
	}

	vm := make([]bool, n)

	for _, v := range leftChild {
		if v == -1 {
			continue
		}
		if vm[v] {
			return false
		}
		vm[v] = true
	}

	for _, v := range rightChild {
		if v == -1 {
			continue
		}
		if vm[v] {
			return false
		}
		vm[v] = true
	}

	ct := false
	rt := -1
	for i:=0; i < len(vm); i++ {
		if !vm[i] {
			if !ct {
				ct = true
				rt = i
			} else {
				return false
			}
		}
	}
	
	if !ct {
		return false
	}
	
	cands := []int{rt}
	vct := 0
	for len(cands) != 0 {
		vct++
		if leftChild[cands[0]] != -1 {
			cands = append(cands, leftChild[cands[0]])
		}
		if rightChild[cands[0]] != -1 {
			cands = append(cands, rightChild[cands[0]])
		}
		cands = cands[1:]
	}
	return vct == n
}

/**
389. Find the Difference
 */
func findTheDifference(s string, t string) byte {
	wc := make([]int, 26)
	
	for _, c :=  range s {
		wc[c-'a']++
	}
	
	for _, c := range t {
		if wc[c-'a'] == 0 {
			return byte(c)
		}
		wc[c-'a']--
	}

	return 0
}


/**
848. Shifting Letters
 */

func shiftingLetters(s string, shifts []int) string {
	for i:= len(shifts) - 2; i >= 0; i-- {shifts[i] += shifts[i+1]}

	bs := []byte(s)

	for i:=0; i<len(bs)-1; i++ {
		bs[i] = byte((int(bs[i] - 'a') + shifts[i])%26) +'a'
	}

	return string(bs)
}


/**
24. Swap Nodes in Pairs
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	fh := ListNode{0,head}
	h := &fh
	for ; h.Next != nil && h.Next.Next != nil ; h = h.Next.Next {
		t := h.Next.Next
		// swap
		f := h.Next
		h.Next = t
		f.Next = t.Next
		t.Next = f
	}

	return fh.Next
}

/**
246. Strobogrammatic Number
 */

func isStrobogrammatic(num string) bool {
	mir := []int {0, 1, -1, -1, -1, -1, 9, -1, 8, 6}


	for l, r:= 0, len(num) - 1; l <= r; {
		lv, rv := num[l] - '0', num[r] - '0'

		if mir[lv] == -1 || mir[rv] == -1 || mir[lv] != int(rv) {
			return false
		}

		l++
		r--
	}

	return true
}

/**
1676. Lowest Common Ancestor of a Binary Tree IV
 */

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

//var lcaNS map[int]bool
//var lcaNode *TreeNode
//var lcaLen int
//func lowestCommonAncestor(root *TreeNode, ns []*TreeNode) *TreeNode {
//	for _, n := range ns {lcaNS[n.Val] = true}
//	lcaNode = nil
//	lcaLen = len(ns)
//	return lcaNode
//}
//
//func lcaHelper(root *TreeNode) int {
//	if root == nil || lcaNode != nil {
//		return 0
//	}
//	res := 0
//	_, ok := lcaNS[root.Val]
//	if ok {
//		res++
//	}
//	res = res + lcaHelper(root.Left) + lcaHelper(root.Right)
//	if res == lcaLen {
//		lcaNode = root
//	}
//
//	return res
//}

/**
605. Can Place Flowers
 */

func canPlaceFlowers(fb []int, n int) bool {
	lz := -2

	for i, v := range fb {
		if v != 0 {
			nz := i - lz - 1
			lz = i
			if nz > 2 {
				n -= (nz - 1) >> 1
			}
		}
	}

	nz := len(fb) - 1 - lz - 1
	if nz > 2 {
		n -= (nz - 1) >> 1
	}

	return n <= 0
}

/**
763. Partition Labels
 */
func partitionLabels(s string) []int {
	cl := make([]int, 26)
	for i, v := range s {cl[v-'a'] = i}

	res := make([]int, 0)
	lst := -1
	st := 0
	for i, v := range s {
		if cl[v-'a'] > lst {
			lst = cl[v-'a']
		}
		if i == lst {
			res = append(res, i - st + 1)
			st = i+1
		}
	}

	return res
}


/**
340. Longest Substring with At Most K Distinct Characters
 */
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	wc := make([]int, 26)
	dcs := 0
	res := 0
	h, e := 0, 0

	for e < len(s) {
		if wc[s[e]-'a'] == 0 {
			dcs++
		}
		wc[s[e]-'a']++
		if dcs <= k {
			if e - h + 1 > res {
				res = e - h
			}
		} else {
			for ;h <= e && dcs != k; h++ {
				if wc[s[h]-'a'] == 1 {
					dcs--
				}
				wc[s[e]-'a']--
			}
		}
		e++
	}

	return res
}

/**
921. Minimum Add to Make Parentheses Valid
 */
func minAddToMakeValid(s string) int {
	res, ls := 0, 0

	for _, c := range s {
		if c == '(' {
			ls++
		} else {
				if ls > 0 {
					ls--
				} else {
						res++
				}
		}
	}

	return res + ls
}

/**
658. Find K Closest Elements
 */
func findClosestElements(a []int, k int, x int) []int {
	if k == 0 {
		return nil
	}
	i, _ := binarySearch(a, x)

	sm, lg := []int{a[i]}, make([]int, 0)
	k--
	s, e := i - 1, i + 1

	for ; k > 0 && (s >= 0 || e < len(a)); k-- {
		if s < 0 {
			lg = append(lg, a[e])
			e++
			continue
		}
		if e == len(a) {
			sm = append([]int{a[s]}, sm...)
			s--
			continue
		}

		if math.Abs(float64(a[s]-x)) <= math.Abs(float64(a[e]-x)) {
			sm = append([]int{a[s]}, sm...)
			s--
		} else {
			lg = append(lg, a[e])
			e++
		}
	}

	return append(sm, lg...)
}

/**
163. Missing Ranges
 */
func findMissingRanges(nums []int, lower int, upper int) []string {
	lower--
	upper++
	res := make([]string, 0)
	nums = append([]int{lower}, nums...)
	nums = append(nums, upper)

	for i := 1; i < len(nums); i++ {
		v := nums[i]
		if v != lower + 1 {
			if v == lower+2 {
				res = append(res, strconv.Itoa(v - 1))
			} else {
				res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(lower + 1), strconv.Itoa(v-1)))
			}
		}
		lower = v
	}

	return res
}

/**
126. Word Ladder II
 */
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	vm := make(map[string]bool)

	cands := make([]string, 0)
	cands = append(cands, beginWord)
	path := make(map[string][]string)
	res := make([][]string, 0)

	for len(cands) != 0 {
		ncs := make([]string, 0)
		for _, w := range cands {
			if vm[w] {
				continue
			}
			if oneDiffChar(w, endWord) {
				// found a result
				pp, _ := path[w]
				res = append(res, append(pp, endWord))
			} else {
				for i := 0; i < len(wordList) && !vm[wordList[i]]; i++ {
					if oneDiffChar(w, wordList[i]) {
						ncs = append(ncs, wordList[i])
						path[wordList[i]] = append(path[w], wordList[i])
					}
				}
			}
			cands = ncs
		}
	}

	return res
}

func oneDiffChar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := false

	for i:=0; i < len(a); i++ {
		if a[i] != b[i] {
			if diff {
				return false
			} else {
				diff = true
			}
		}
	}

	return diff
}

/**
162. Find Peak Element
 */
//func findPeakElement(nums []int) int {
//	if len(nums) < 3 {
//		return 0
//	}
//
//	s, e := 0, len(nums) - 1
//
//	for s <= e {
//		m := s + (e - s) >> 1
//		if m == 0 {
//			return 0
//		} else if m == len(nums) - 1 {
//			return len(nums) - 1
//		} else {
//			if nums[m] > nums[m -1] && nums[m] > nums[m + 1] {
//				return m
//			} else if nums[m] < nums[m - 1] {
//				e = m - 1
//			} else {
//				s = m + 1
//			}
//		}
//	}
//
//	return -1
//}

func findPeakElement(nums []int) int {
	nums = append(nums, -1 << 31)
	nums = append([]int{-1 << 31}, nums...)

	s, e := 0, len(nums) - 1

	for s <= e {
		m := s + (e - s) >> 1
		if nums[m] > nums[m -1] && nums[m] > nums[m + 1] {
			return m - 1
		} else if nums[m] < nums[m - 1] {
			e = m - 1
		} else {
			s = m + 1
		}
	}

return -1
}

/**
515. Find Largest Value in Each Tree Row
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	cands := []*TreeNode{root}

	for len(cands) != 0 {
		ll := -1 << 31
		ncs := make([]*TreeNode, 0)

		for _, n := range cands {
			if n.Val > ll {
				ll = n.Val
			}
			if n.Left != nil {
				ncs = append(ncs, n.Left)
			}
			if n.Right != nil {
				ncs = append(ncs, n.Right)
			}
		}

		res = append(res, ll)
		cands = ncs
	}

	return res
}

/**
938. Range Sum of BST
 */
var res938 int
func rangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		return res938
	}

	cands := []*TreeNode{root}

	for len(cands) != 0 {
		ncs := make([]*TreeNode, 0)
		for _, n := range cands {
			if n.Val >= low && n.Val <= high {
				res938 += n.Val
			}
			if n.Left != nil && n.Val >= low {
				ncs = append(ncs, n.Left)
			}
			if n.Right != nil && n.Val <= high {
				ncs = append(ncs, n.Right)
			}
		}
		cands = ncs
	}
	return res938
}

/**
240. Search a 2D Matrix II
 */
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	r, c := len(matrix) - 1, 0
	for r >= 0 && c < len(matrix[0]) {
		if matrix[r][c] == target {
			return true
		} else if  matrix[r][c] > target {
			c++
		} else {
			r--
		}
	}
	return false
}

/**
746. Min Cost Climbing Stairs
 */
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	cost = append(cost, 0)
	mc := make([]int, len(cost))

	for i := len(cost) - 3; i >= 0; i-- {

		if mc[i+1] < mc[i+2] {
			mc[i] = cost[i] + mc[i+1]
		} else {
			mc[i] = cost[i] + mc[i+2]
		}
	}

	if mc[0] < mc[1] {
		return mc[0]
	} else {
		return mc[1]
	}
}

/**
543. Diameter of Binary Tree
 */
var dbtres int
func diameterOfBinaryTree(root *TreeNode) int {
	dbthelper(root)
	return dbtres - 1
}

func dbthelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := dbthelper(root.Left)
	rd := dbthelper(root.Right)

	if ld + rd + 1 > dbtres {
		dbtres = ld + rd + 1
	}

	if ld > rd {
		return 1 + ld
	} else {
		return 1 + rd
	}
}

/**
349. Intersection of Two Arrays
 */
func intersection(nums1 []int, nums2 []int) []int {
	nc1, nc2 := make(map[int]bool), make(map[int]bool)

	for _, n := range nums1 {
		nc1[n] = true
	}

	for _, n := range nums2 {
		nc2[n] = true
	}

	res := make([]int, 0)
	for k, _ := range nc1 {
		_, ok := nc2[k]

		if ok {
			res = append(res, k)
		}
	}

	return res
}

/**
228. Summary Ranges
 */
func summaryRanges(nums []int) []string {
	if nums == nil || len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	s, e := 0, 1
	for ; e < len(nums); e++ {
		if nums[e] != nums[e - 1] + 1 {
			if s + 1 == e {
				res = append(res, strconv.Itoa(nums[s]))
			} else {
				res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(nums[s]), strconv.Itoa(nums[e-1])))
			}
			s = e
		}
	}

	if s + 1 == e {
		res = append(res, strconv.Itoa(nums[s]))
	} else {
		res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(nums[s]), strconv.Itoa(nums[e-1])))
	}

	return res
}

/**
226. Invert Binary Tree
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

/**
203. Remove Linked List Elements
 */
func removeElements(head *ListNode, val int) *ListNode {
	fh := ListNode{0, head}

	for c:= &fh; c.Next != nil; {
		if c.Next.Val == val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return fh.Next
}

/**
278. First Bad Version
*/
//func firstBadVersion(n int) int {
//	s, e, m := 1, n, 0
//
//	for s < e {
//		m = s + (e-s)>> 1
//		if isBadVersion(m) {
//			e = m
//		} else {
//			s = m + 1
//		}
//	}
//
//	return s
//}

/**
567. Permutation in String
 */

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cc1, cc2 := make([]int, 26), make([]int, 26)
	mc := 0

	for _, c := range s1 {
		cc1[c-'a']++
	}

	for i:=0; i < len(s1); i++ {
		cc2[s2[i]-'a']++
	}

	for i:=0; i<26; i++ {
		if cc1[i] == cc2[i] {
			mc++
		}
	}

	if mc == 26 {
		return true
	}

	for i:= len(s1); i < len(s2); i++ {
		if cc2[s2[i-len(s1)] - 'a'] == cc1[s2[i-len(s1)] - 'a'] {
			mc--
		}
		cc2[s2[i-len(s1)] - 'a']--

		if cc2[s2[i-len(s1)] - 'a'] == cc1[s2[i-len(s1)] - 'a'] {
			mc++
		}


		if cc2[s2[i] - 'a'] == cc1[s2[i] - 'a'] {
			mc--
		}
		cc2[s2[i] - 'a']++
		if cc2[s2[i] - 'a'] == cc1[s2[i] - 'a'] {
			mc++
		}


		if mc == 26 {
			return true
		}
	}

	return false
}

/**
160. Intersection of Two Linked Lists
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la := linkLen(headA)
	lb := linkLen(headB)

	if la < lb {
		lb, la = la, lb
		headA, headB = headB, headA
	}

	for i := la - lb; i > 0; i-- {
		headA = headA.Next
	}

	for ; headA != nil && headB != nil; {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func linkLen(h *ListNode) int {
	l := 0
	for ; h != nil; h=h.Next {l++}
	return l
}

/**
777. Swap Adjacent in LR String
 */
func canTransform(start string, end string) bool {

	nrs := 0
	si, ei:=0, 0

	for si != len(start) && ei != len(end) {
		for ; si < len(start) && start[si] != 'L'; si++ {
			if start[si] == 'R' {
				nrs++
			}
		}
		si++
		for ; si < len(start) && start[si] == 'X'; si++ {}

		for ; ei < len(end) && end[ei] != 'L'; ei++ {
			if end[ei] == 'R' {
				nrs--
			}
		}
		ei++
		for ; ei < len(end) && end[ei] == 'X'; ei++ {}

		if  nrs != 0 || si == len(start) && ei != len(end) || si != len(start) && ei == len(end) {
			return false
		}
	}
	return true
}

/**
1213. Intersection of Three Sorted Arrays
 */
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {

	res := make([]int,0)

	i1, i2, i3 := 0, 0, 0

	for i1 < len(arr1) && i2 < len(arr2) && i3 < len(arr3) {

		if arr1[i1] == arr2[i2] && arr3[i3] == arr2[i2] {
			res = append(res, arr1[i1])
			i1++; i2++; i3++
		} else {
			if 	arr1[i1] <= arr2[i2] && arr1[i1] <= arr3[i3] {
				i1++
			} else if 	arr2[i2] <= arr1[i1] && arr2[i2] <= arr3[i3] {
					i2++
			} else {
						i3++
			}
		}
	}

	return res
}

/**
987. Vertical Order Traversal of a Binary Tree
 */
var res987 [][]int
func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	lm, rm := calWidth(root, 0)
	res987 = make([][]int, rm - lm + 1)
	for i:=0; i<len(res987); i++ {res987[i] = make([]int,0)}
	addVTnodes(root, -lm)
	return res987
}

func calWidth(r *TreeNode, p int) (int, int) {

	lm, rm := p, p
	if r.Left != nil {
		lm, rm = calWidth(r.Left, p-1)
	}
	if r.Right != nil {
		rl, rr := calWidth(r.Right, p + 1)
		if rl < lm {
			lm = rl
		}
		if rr > rm {
			rm = rr
		}
	}

	return lm, rm
}

func addVTnodes(r *TreeNode, p int) {
	cands := []*TreeNode{r}
	cp := []int{p}

	for len(cands) != 0 {
		ncs := make([]*TreeNode, 0)
		ncp := make([]int, 0)

		for i, n := range cands {
			res987[cp[i]] = append(res987[cp[i]], n.Val)
			if n.Left != nil {
				ncs = append(ncs, n.Left)
				ncp = append(ncp, cp[i] - 1)
			}
			if n.Right != nil {
				ncs = append(ncs, n.Right)
				ncp = append(ncp, cp[i] + 1)
			}
		}

		cands = ncs
		cp = ncp
	}

}

/**
57. Insert Interval
 */
func insert(intervals [][]int, ni []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{ni}
	}

	res := make([][]int, 0)
	cur := ni

	for i, it := range intervals {
		ni, ord := mergeInterval57(it, cur)
		if ord == -1 {
			res = append(res, cur)
			res = append(res, intervals[i:]...)
			return res
		} else if ord == 1 {
			res = append(res, it)
		} else {
			cur = ni
		}
	}

	res = append(res, cur)
	return res
}

func mergeInterval57(i1, i2 []int) ([]int, int) {
	if i2[1] < i1[0] {
		return nil, -1
	}
	if i2[0] > i1[1] {
		return nil, 1
	}
	l, r := i1[0], i1[1]
	if i1[0] > i2[0] {
		l = i2[0]
	}
	if i1[1] < i2[1] {
		r = i2[1]
	}
	return []int{l, r}, 0
}