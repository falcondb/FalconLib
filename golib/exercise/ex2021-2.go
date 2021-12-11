package exercise

import (

	"sort"
)

func minDifference2(nums []int) int {
	if len(nums) < 5 {
		return 0
	}

	sort.Ints(nums)
	res := 0xffffffff
	for i:=0; i<3; i++ {
		dif := nums[len(nums) - 4 + i] - nums[i]
		if dif < res {
			res = dif
		}
	}

	return res
}

var min4, max4 = make([]int, 4), make([]int, 4)

func updateMinMax4(v int){
	n := v
	for i:=0; i < 4; i++ {
		if n > max4[i] {
			n, max4[i] = max4[i], n
		}
	}
	n = v
	for i:=3; i >= 0; i-- {
		if n < min4[i] {
			n, min4[i] = min4[i], n
		}
	}
}

func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	for i:=0; i < 4; i++ {
		min4[i] = 0x7fffffff
		max4[i] = -0xffffffff
	}

	for _, n := range nums {
		updateMinMax4(n)
	}

	res := 0xffffffff
	for i:=0; i<=3; i++ {
		dif := max4[i] - min4[i]
		if dif >=0 && dif < res {
			res = dif
		}
	}

	return res
}


/**
1525. Number of Good Ways to Split a String
 */
func numSplits(s string) int {
	twc := make([]int,26)
	tds := 0
	for _, c := range s {
		if twc[c-'a'] == 0 {tds++}
		twc[c-'a']++
	}

	res:=0
	cwc:= make([]bool, 26)
	cds:=0

	for i:=0; i < len(s) - 1; i++ {
		if !cwc[s[i]-'a'] {
			cds++
			cwc[s[i]-'a'] = true
		}
		twc[s[i]-'a']--
		if twc[s[i]-'a'] == 0 {
			tds--
		}

		if cds == tds {
			res++
		}
	}

	return res
}


type wls []string
func (s wls) Len() int {return len(s)}
func (s wls) Less(i, j int) bool {return len(s[i]) < len(s[j])}
func (s wls) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func longestStrChain(words []string) int {
	sort.Sort(wls(words))

	nl := make([]int, len(words))
	res := 0
	lidx := make(map[int][]int)

	for i, w := range words {
		lidx[len(w)] = append(lidx[len(w)], i)
	}

	for l := 1; l <= len(words[len(words) - 1]); l++ {
		for _, v := range lidx[l] {
			for _, sv := range lidx[l-1] {
				if isPed(words[sv], words[v]) {
					if nl[v] < nl[sv] + 1 {
						nl[v] = nl[sv] + 1
					}
					if res < nl[v] {
						res = nl[v]
					}
				}
			}
		}
	}

	return res + 1
}

func isPed(s, t string) bool {
	if len(s) + 1 != len(t) {
		return false
	}
	del := false
	for si, ti := 0, 0; si < len(s) && ti < len(t) ; {
		if s[si] == t[ti] {
			si++
			ti++
		} else if !del {
			ti++
			del = true
		} else {
			return false
		}
	}
	return true
}

/**
735. Asteroid Collision
 */

func asteroidCollision2(asteroids []int) []int {
	res := make([]int, 0)
	cr := make([]int, 0)

	for _, a := range asteroids {
		if a > 0 {
			cr = append(cr, a)
		} else {
			gone := false
			for i:= len(cr) - 1; i>=0; i-- {
				if cr[i] == -a {
					cr = cr[:len(cr) - 1];
					gone = true
				} else if cr[i] < -a {
					cr = cr[:len(cr) - 1]
				} else {
					gone = true
				}
				if gone {
					break
				}
			}
			if !gone {
				res = append(res, a)
			}
		}
	}

	res = append(res, cr...)
	return res
}

/**
853. Car Fleet
 */
func carFleet(target int, position []int, speed []int) int {
	cs := make(cars, len(position))
	for i:=0; i < len(position); i++ {
		cs[i] = car{position[i], speed[i]}
	}

	sort.Sort(cs)

	res := 0
	cl := float32(0)

	for _, c := range cs {
		iat :=float32(target-c.pos)/float32(c.sp)
		if iat > cl {
			res++
			cl = iat
		}
	}
	return res
}

type car struct {
	pos, sp int
}

type cars []car
func (a cars) Len() int {return len(a)}
func (a cars) Less(i, j int) bool {return a[i].pos > a[j].pos}
func (a cars) Swap(i, j int) {a[i], a[j] = a[j], a[i]}


/**
690. Employee Importance
 */
type Employee struct {
    Id int
    Importance int
    Subordinates []int
}

func getImportance(es []*Employee, id int) int {
	ti := 0
	id2ix := make(map[int]int)
	cs := make([]int, 0)
	for i, e := range es {
		if e.Id == id {
			cs = append(cs, e.Id)
		}
		id2ix[e.Id] = i
	}

	for len(cs) != 0 {
		ncs := make([]int, 0)
		for _, e := range cs {
			ncs = append(ncs, es[id2ix[e]].Subordinates...)
			ti += es[id2ix[e]].Importance
		}
		cs = ncs
	}

	return ti
}

/**
792. Number of Matching Subsequences
 */

func numMatchingSubseq(s string, words []string) int {
	res := 0
	for _, t := range words {
		si, ti := 0, 0
		for ; si < len(s) && ti < len(t); {
			if s[si] == t[ti] {
				ti++
			}
			si++
		}
		if ti == len(t) {
			res ++
		}
	}
	return res
}

/**
1110. Delete Nodes And Return Forest
 */
var res1110 = make([]*TreeNode, 0)
var dk1110 = make(map[int]bool)
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	for _, d := range to_delete {
		dk1110[d] = true
	}

	if !dnt(root) {
		res1110 = append(res1110, root)
	}

	return res1110
}

func dnt(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if dnt(root.Left) {
		root.Left = nil
	}
	if dnt(root.Right) {
		root.Right = nil
	}

	if dk1110[root.Val] {
		if root.Left != nil {
			res1110 = append(res1110,root.Left)
		}
		if root.Right != nil {
			res1110 = append(res1110,root.Right)
		}

		return true
	} else {
		return false
	}
}


/**
1277. Count Square Submatrices with All Ones
 */
func countSquares(matrix [][]int) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	h, w := len(matrix), len(matrix[0])
	cs := 0

	for j:=0; j<h; j++ {
		for i:=0; i<w; i++ {
			if i > 0 && j > 0 && matrix[j][i] == 1 {
				matrix[j][i] = MinI3(matrix[j-1][i], matrix[j][i-1], matrix[j-1][i-1]) + 1
			}
			cs += matrix[j][i]
		}
	}
	return cs
}

func MinI3(a, b, c int) int {
	if a <= b {
		if a <=c {
			return a
		} else {
			return c
		}
	} else {
		if b <=c {
			return b
		} else {
			return c
		}
	}
}

func MaxI2(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func MaxI3(a, b, c int) int {
	if a <= b {
		if b <=c {
			return c
		} else {
			return b
		}
	} else {
		if a <= c {
			return c
		} else {
			return a
		}
	}
}

/**
562. Longest Line of Consecutive One in Matrix
 */
func longestLine(mat [][]int) int {
	w, h := len(mat[0]), len(mat)
	prh, prv, prd, prrd := make([]int, w), make([]int, w), make([]int, w), make([]int, w)
	res := 0
	pd := 0
	for j:=0; j<h; j++ {
		for i:=0; i< w; i++ {
			if mat[j][i] == 1 {
				if i == 0 {
					prh[i] = 1
					pd = 1
				} else {
					pd = pd + 1
					prd[i-1] = pd -1
					prh[i] = prh[i-1] + 1
				}
				prv[i] = prv[i] + 1
				if i <= w - 1 {
					prrd[i] = 1
				} else {
					prrd[i] = prrd[i+1]
				}
			}
			res = MaxI3(res, MaxI3(prh[i], prv[i], prd[i]), prrd[i])
		}
	}

	return res
}

/**
1417. Reformat The String
 */

func reformat(s string) string {
	res:=make([]byte,0)
	cs,ds:= make([]byte, 0), make([]byte,0)
	for _, c := range []byte(s) {
		if c >= '0' && c <= '9' {
			ds = append(ds, c)
		} else {
			cs = append(cs, c)
		}
	}

	if len(cs) == len(ds) + 1 {
		for i:=0; i < len(ds); i++ {
			res = append(res, cs[i])
			res = append(res, ds[i])
		}
		res = append(res, cs[len(cs) - 1])
	} else if len(cs) == len(ds) - 1 {
		for i:=0; i < len(cs); i++ {
			res = append(res, ds[i])
			res = append(res, cs[i])
		}
		res = append(res, cs[len(ds) - 1])
	} else if len(cs) == len(ds) {
		for i := 0; i < len(cs); i++ {
			res = append(res, ds[i])
			res = append(res, cs[i])
		}
	}
	return string(res)
}


func flipEquiv(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	}

	if r1.Val == r2.Val {
		return flipEquiv(r1.Left, r2.Left) && flipEquiv(r1.Right, r2.Right) || flipEquiv(r1.Left, r2.Right) && flipEquiv(r1.Right, r2.Left)
	}
	return false
}

/**
67. Add Binary
 */
func addBinary(a string, b string) string {
	res := make([]byte, 0)
	cr := byte(0)
	// a is longer
	if len(a) < len(b) {a, b = b, a}

	for ai, bi := len(a) - 1, len(b) -1; bi >= 0;  {
		if a[ai] == '1' {
			cr++
		}
		if b[bi] == '1' {
			cr++
		}

		if cr & 0x1 == 0x1 {
			res = append(res, '1')
		} else {
			res = append(res, '0')
		}

		cr = cr >> 1
		ai--; bi--
	}

	for ai := len(a) - 1 - len(b); ai >= 0; ai-- {
		if a[ai] == '1' {
			cr++
		}

		if cr & 0x1 == 0x1 {
			res = append(res, '1')
		} else {
			res = append(res, '0')
		}
		cr = cr >> 1
	}

	if cr > 0 {
		res = append(res, '1')
	}

	for i:=0; i < len(res)>>1; i++ {
		res[i], res[len(res) - 1 -i] = res[len(res) - 1 -i], res[i]
	}
	return string(res)
}


/**
708. Insert into a Sorted Circular Linked List
 */

/**
when looping the nodes, special care of the first node.
handle the case all nodes have the same value
 */

 type SCNode struct {
     Val int
    Next *SCNode
 }
func scinsert(aNode *SCNode, x int) *SCNode {
	nn := &SCNode{x, nil}
	if aNode == nil {
		nn.Next = nn
		return nn
	}

	if aNode.Next == nil {
		aNode.Next = nn
		nn.Next = aNode
	}

	added := false
	lp := false
	for cn := aNode; lp && !added; {
		if cn.Val < cn.Next.Val && cn.Val < x && cn.Next.Val > x || cn.Val == x || cn.Val > cn.Next.Val &&( x > cn.Val || x < cn.Next.Val){
			nn.Next = cn.Next
			cn.Next = nn
			added = true
		}

		cn = cn.Next
		if cn == aNode {
			lp = true
		}
	}

	if !added {
		nn.Next = aNode.Next
		aNode.Next = nn
	}

	return aNode
}


 /**
 173. Binary Search Tree Iterator
  */


/**
211. Design Add and Search Words Data Structure
 */

/**
Need a flag in the tree node of if the word ends here
 */
type WordDictionary struct {
	fh *PreNode
}

type PreNode struct {
	ns []*PreNode
	ep bool
}


func WDConstructor() WordDictionary {
	return WordDictionary{&PreNode{make([]*PreNode, 26), false}}
}

func (t *WordDictionary) AddWord(word string)  {
	cn := t.fh
	for _, c := range word {
		ci := c - 'a'
		if cn.ns[ci] == nil {
			cn.ns[ci] = &PreNode{make([]*PreNode, 26), false}
		}
		cn = cn.ns[ci]
	}
	cn.ep = true
}

func (t *WordDictionary) Search(word string) bool {
	return wdsh(t.fh, word)
}

func wdsh(r *PreNode, word string) bool {
	if len(word) == 0 {
		return r.ep
	}

	c := word[0] - 'a'
	if word[0] != '.' {
		return r.ns[c] != nil && wdsh(r.ns[c], word[1:])
	} else {
		for i:=0; i<26; i++{
			if r.ns[i] != nil && wdsh(r.ns[i], word[1:]) {
				return true
			}
		}
	}

	return false
}

/**
986. Interval List Intersections
 */
func intervalIntersection(fl [][]int, sl [][]int) [][]int {
	res := make([][]int, 0)

	for f, s := 0, 0; f < len(fl) && s < len(sl); {
		up := fl[f][0]
		if fl[f][0] < sl[s][0] {up = sl[s][0]}
		low := fl[f][1]
		if fl[f][1] > sl[s][1] {low = sl[s][1]}

		if up <= low {
			res = append(res, []int{up, low})
		}

		if fl[f][1] > sl[s][1] {
			s++
		} else {
			f++
		}
	}

	return res
}

/**
536. Construct Binary Tree from String
 */
//func str2tree(s string) *TreeNode {
//	st536 = make([]*TreeNode, 0)
//}
//
//var st536 []*TreeNode
//
//func s2th(s string) *TreeNode{
//	if len(s) == 0 {
//		return nil
//	}
//	var cn *TreeNode
//	pos := false
//	for i:=0; i < len(s); i++ {
//		if s[i] == '(' {
//			st536 = append(st536, cn)
//		} else if s[i] == ')' {
//			r := st536[len(st536)-1]
//			if r.Left == nil {
//				r.Left = cn
//			} else {
//				r.Right = cn
//			}
//			cn = st536[len(st536)-1]
//			st536 = st536[:len(st536)-1]
//		} else if s[i] == '-' {
//			pos = false
//		} else {
//			cv := 0
//			for p:=i; p < len(s) && s[p] >= '0' && s[p] <= '9'; p++ {
//				cv = cv * 10 + int(s[i] - '0')
//			}
//			if !pos {cv = -cv}
//			cn = &TreeNode{cv, nil, nil}
//		}
//	}
//
//	return cn
//}

/**
71. Simplify Path
 */

func simplifyPath(path string) string {
	ps := make([][]byte,0)
	np := make([]byte, 0)

	if len(path)==0 {
		return "/"
	}
	if path[len(path) - 1] != '/' {path += "/"}

	for i:=0; i<len(path); i++ {
		switch path[i] {
		case '/':
			if len(np) == 2 && np[0] == '.' && np[1] == '.' {
				if len(ps) > 0 {
					ps = ps[:len(ps)-1]
				}
			} else if len(np) > 0 && (len(np) != 1 || np[0] != '.') {
				ps = append(ps, np)
			}
			np = make([]byte, 0)
		default:
			np = append(np, path[i])
		}
	}
	if len(ps) == 0 {
		return "/"
	}
	res := make([]byte, 0)
	for _, p := range ps {
		res = append(res, '/')
		res = append(res, p...)
	}

	return string(res)
}



func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	t, be := head, head
	for ; t != nil && k>0; t=t.Next{k--}

	for ; t.Next != nil; {t = t.Next; be = be.Next}

	fh := ListNode{0, head}

	fh.Next = head.Next
	head.Next = be.Next.Next
	tmp := be.Next

	tmp.Next = fh.Next
	be.Next = head

	return tmp
}

/**
317. Shortest Distance from All Buildings
 */


/**
277. Find the Celebrity
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		if n <= 1 {
			return n
		}
		cd := 0
		for i:=1; i < n; i++ {
			if knows(cd, i) {
				cd = i
			}
		}
		for i:=0; i<n; i++ {
			if knows(cd, i) || !knows(i, cd) {
				return -1
			}
		}
		return cd
	}

}

/**
759. Employee Free Time
 */
type Interval struct {
     Start,     End   int
}

type invs []*Interval

func (iv invs) Len() int {return len(iv)}
func (iv invs) Less(i, j int) bool {return iv[i].Start < iv[j].Start}
func (iv invs) Swap(i, j int) {iv[i], iv[j] = iv[j], iv[i]}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	if schedule == nil || len(schedule) == 0 {
		return nil
	}

	ivs := make([]*Interval, 0)

	for _, p := range schedule {
		for _, t := range p {
			ivs = append(ivs, t)
		}
	}

	sort.Sort(invs(ivs))

	ce := -1
	res := make([]*Interval, 0)
	for _, iv := range ivs {
		if iv.Start > ce {
			res = append(res, &Interval{ce, iv.Start})
			ce = iv.End
		} else {
			if iv.End > ce {
				ce = iv.End
			}
		}
	}

	return res[1:]
}

/**
1673. Find the Most Competitive Subsequence
 */
func mostCompetitive(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}

	res := make([]int,0)
	for i:= 0; i < len(nums); i++ {
		for len(res) != 0 && res[len(res) - 1] > nums[i] && len(res) + len(nums) - i > k {
			res = res[:len(res)-1]
		}
		res = append(res, nums[i])
	}
	return res[:k]
}

/**
26. Remove Duplicates from Sorted Array
 */

func removeDuplicatesInt(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	fp, p := 1, 1

	for ; p < len(nums); p++{
		if nums[p] != nums[p-1] {
			nums[fp] = nums[p]
			fp++
		}
	}

	return fp-1
}

/**
418. Sentence Screen Fitting
 */
func wordsTypingBF(sentence []string, rows int, cols int) int {

	res:=0
	wi:=0
	for r:=1; r<=rows; r++ {
		rms:=cols
		for ; len(sentence[wi]) <= rms; {
			if wi == len(sentence) - 1 {
				res++
			}
			rms -= len(sentence[wi]) + 1
			wi = (wi+1)%len(sentence)
		}
	}

	return res
}

func wordsTyping(sentence []string, rows int, cols int) int {
	nw := make([]int, len(sentence))
	rps := make([]int, len(sentence))
	for i, _ := range sentence {
		sps := cols
		p:=i
		for ; len(sentence[p]) <= sps; {
			sps -= len(sentence[p]) + 1
			if p == len(sentence) - 1 {
				rps[i]++
			}
			p = (p+1)%len(sentence)
		}
		nw[i] = p
	}

	res := 0
	for cp := 0; rows >=1; rows-- {
		res += rps[cp]
		cp = nw[cp]
	}

	return res
}

/**
1937. Maximum Number of Points with Cost
 */

//func maxPoints(points [][]int) int64 {
//	if len(points) == 0 {
//		return 0
//	}
//	pr := make([]int, len(points[0]))
//	w := len(points[0])
//	for c:=0; c<len(points[0]); c++ {pr[c]=points[0][c]}
//	for r:=1; r < len(points); r++ {
//		cr := make([]int, w)
//		for c:=0; c<w; {
//			cr[c] = pr[c]
//			for p:=c+1; p < w; p++ {
//				if pr[p] >= pr[c] + p - c {
//					c=p
//					break
//				} else {
//					if cr[p] < pr[c] + p - c {
//						cr[p] = pr[c] + p - c
//					}
//				}
//			}
//		}
//		for c:=w-1; c>=0; {
//
//		}
//	}
//}

/*
5. Longest Palindromic Substring
 */
var lp5r string
var lp5s []byte
func longestPalindrome(s string) string {
	max := -1
	if len(s) <= 1 {
		return s
	}
	lp5r = s
	for i:=0; i < len(s); i++ {
		l1:= lpextend(i, i+1)
		l2:= lpextend(i-1, i+1) + 1
		if l1 > l2 && l1 > max {
			lp5r = s[i-l1+1: i+1+l1-1]
			max = l1
		} else if l2 > l1 && l2 > max {
			lp5r = s[i-l2-1: i+l2+1]
			max = l2
		}
	}

	return lp5r
}

func lpextend(l, r int) int {
	d:= 0
	for ; l - d >= 0 && r+d < len(lp5s) && lp5s[l - d]  == lp5s[r+d]; d++ {}

	return d
}

/**
370. Range Addition
 */


func getModifiedArray(length int, updates [][]int) []int {
	res :=make([]int, length)
	if len(updates) == 0 {
		return res
	}

	for _, u := range updates {
		res[u[0]] += u[2]
		if u[1] + 1 < length {
			res[u[1] + 1] -= u[2]
		}
	}

	for i:= 1; i < len(res); i++ {res[i] += res[i-1]}
	return res
}

/**
259. 3Sum Smaller
 */
var data259 []int
func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	data259 = nums
	res := 0
	for i:=0; i < len(data259) - 2; i++ {
		res += twoSumSmaller(i+1, target - nums[i])
	}
	return res
}

func twoSumSmaller(s int, t int) int {
	e := len(data259) -1
	res := 0
	for s < e {
		if data259[s] + data259[e] < t {
			res += e - s
			s++
		} else {
			e--
		}
	}
	return res
}

/**
518. Coin Change 2
 */

func change(amount int, coins []int) int {
	cgs := make([]int, amount + 1)
	cgs[0] = 1
	for _, c := range coins {
		for v:=c; v < amount; v++ {
			cgs[v] += cgs[v - c]
		}
	}

	return cgs[amount]
}


func minSol(k int) int {

	for i:= 1; i > 0; i++ {
		gys := make([]bool, 2*k)

		st := 0
		rmgs := 0

		found := false
		for !found {
			for c := 0; c < i; {
				if !gys[st] {
					c++
				}
				if c == i {
					break
				}
				st = (st+1) % (2*k)
			}

			if st < k {
				found = true

			} else {
				gys[st] = true
				rmgs++

				if rmgs == k {
					return i
				}
			}
		}
	}

	return 0
}

/**
766. Toeplitz Matrix
 */
func isToeplitzMatrix(matrix [][]int) bool {
	w := len(matrix[0])
	h := len(matrix)

	for i:=0; i < w ; i++ {
		for j:= 0; j +1 < h && i +j +1 < w; j++ {
			if matrix[j][i+j] != matrix[j+1][i+j+1] {
				return false
			}
		}
	}

	for j := 1; j < h; j++ {
		for i:=0; i+1 < w && j+i+1 < h; i++ {
			if matrix[j+i][i] != matrix[j+i+1][i+1] {
				return false
			}
		}
	}

	return true
}

/**
536. Construct Binary Tree from String
 */
func str2tree(s string) *TreeNode {
	st := make([]*TreeNode, 0)
	n := &TreeNode{0, nil, nil}
	ng := false
	for _, c := range s {
		switch c {
		case '-':
			ng = true
		case '(':
			if ng {
				n.Val = -n.Val
			}

			st = append(st, n)
			n = &TreeNode{0, nil, nil}
			ng = false
		case ')':
			if len(st) != 0 {
				pn := st[len(st) - 1]
				if ng {
					n.Val = -n.Val
				}
				if pn.Left == nil {
					pn.Left = n
				} else if pn.Right == nil {
					pn.Right = n
				}
				n = pn
				st = st[:len(st)-1]
			} else {
				return n
			}
		default:
			n.Val = n.Val*10 + int(c - '0')
		}
	}

	return n
}

/**
1011. Capacity To Ship Packages Within D Days
 */
func shipWithinDays(weights []int, days int) int {
	max, min := 0, 0
	res := 0
	for _, w := range weights {
		if w > min { min = w }
		max += w
	}

	for min <= max {
		cc := (max + min)>>1
		ds := 1
		rc := cc
		for _, w := range weights {
			if w <= rc {
				rc -= w
			} else {
				rc = cc - w
				ds++
			}
		}
		if ds == days {
			res = ds
		}
		if ds <= days { max = cc - 1
		} else { min = cc + 1 }
	}
	return res
}

/**
1382. Balance a Binary Search Tree
 */
var ns1382 = make([]*TreeNode, 0)
func balanceBST(root *TreeNode) *TreeNode {
	return nil

}

func getNodes(r *TreeNode) {
	if r != nil {
		getNodes(r.Left)

		ns1382 = append(ns1382, r)
		getNodes(r.Right)
	}
}

func buildBBST(s, e int) *TreeNode {
	if s > e {
		return nil
	}
	nr := ns1382[(s+e)>>1]
	nr.Left = buildBBST(s, (s+e)>>1 - 1)
	nr.Right = buildBBST((s+e)>>1 + 1, e)
	return nr
}

/**
1868. Product of Two Run-Length Encoded Arrays
 */
func findRLEArray(encoded1 [][]int, encoded2 [][]int) [][]int {
	if len(encoded1) == 0 || len(encoded2) == 0 {
		return nil
	}

	fc, fl, sc, sl := 0, encoded1[0][1], 0, encoded2[0][1]
	res := make([][]int, 0)
	for  {
		a := encoded1[fc][0]
		b := encoded2[sc][0]
		cl := sl
		if fl < sl {
			cl = fl
		}
		pd := a * b
		if len(res) > 0 && pd == res[len(res)-1][0] {
			res[len(res)-1][1] += cl
		} else {
			res = append(res, []int{pd, cl})
		}

		fl-=cl
		sl-=cl

		if fl == 0 {
			fc++
			if fc >= len(encoded1) {
				break
			}
			fl = encoded1[fc][1]
		}


		if sl == 0 {
			sc++
			if sc >= len(encoded2) {
				break
			}
			sl = encoded2[sc][1]
		}
	}

	return res
}

/**
133. Clone Graph
 */

type Node struct {
Val int
Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nm := make(map[*Node]*Node)
	cds := []*Node{node}
	vis := make(map[*Node]bool)
	nr := &Node{node.Val, nil}
	nm[node]=nr

	for len(cds) != 0 {
		for _, on := range cds {
			cds = cds[1:]
			if vis[on] { continue}
			vis[on] = true
			for _, nei := range on.Neighbors {
				if cn, ok := nm[nei]; ok {
					nm[on].Neighbors = append(nm[on].Neighbors, cn)
				} else {
					newn := &Node{nei.Val, nil}
					nm[nei] = newn
					nm[on].Neighbors = append(nm[on].Neighbors, newn)
				}
				if !vis[nei] {
					cds = append(cds, nei)
				}
			}
		}
	}

	return nr
}

/**
1891. Cutting Ribbons
 */
func maxLength(ribbons []int, k int) int {
	max := 0
	for _, r := range ribbons {
		if r > max {max = r}
	}
	min := 1
	res := 0
	for min <= max {
		nl := (min + max)>>1
		nc := numCuts(ribbons, nl)
		if nc >= k { res = nl}
		if nc >= k {min = nl + 1} else {max = nl - 1}

	}
	return res
}

func numCuts(rs []int, len int) int {
	nc := 0
	for _, r := range rs {
		nc+=r/len
	}
	return nc
}


//type Solution struct {
//	nms []int
//}
//
//
//func Constructor(nums []int) Solution {
//	return Solution{nums}
//}

//
//func (this *Solution) RSPick(target int) int {
//	ct, res := 0, -1
//
//	for i, n := range this.nms {
//		if n == target {ct++}
//		if rand.Intn(ct) == 0 {res = i}
//	}
//
//	return res
//}

/**
1344. Angle Between Hands of a Clock
 */
func angleClock(hour int, minutes int) float64 {

	hp := 5 *  (float64(hour)  + float64(minutes / 60 ))
	mp := float64(minutes)
	if hp < mp {
		hp, mp = mp, hp
	}
	return (hp - mp) * float64(60) / float64(265)
}


/**
1004. Max Consecutive Ones III
 */
func longestOnes1(nums []int, k int) int {
	zi := make([]int, k)
	ci := 0
	res := 0
	for i, n := range nums {
		if n == 0 {
			if res < i - zi[ci] {res = i - zi[ci] + 1}
			zi[ci] = i
			ci = (ci + 1) % k
		}
	}

	if len(nums) - zi[ci] > res {res = len(nums) - zi[ci]}

	return res
}

func longestOnes(nums []int, k int) int {
	l, r := 0, 0

	for ; r < len(nums); r++ {
		if nums[r] == 0 {k--}
		if k < 0 {
			k += 1 - nums[r]
			l++
		}
	}
	return r - l
}

/**
266. Palindrome Permutation
 */
func canPermutePalindrome(s string) bool {

	wc := make(map[int32]int)
	ct := 0
	for _, c := range s {
		wc[c]++
		if wc[c] & 0x1 != 0 {ct++} else {ct--}
	}

	return (ct & 0xfffffffe) == 0
}

/**
1539. Kth Missing Positive Number
 */
func findKthPositive(arr []int, k int) int {
	c := 1

	for _, v := range arr {
		if v == c { c++; continue }
		k--
		if k == 0 { return c }
		c++
	}

	return -1
}


func isMonotonic(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	i:=1
	inc := false
	for ; i<len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc = true
			break
		} else if nums[i] < nums[i-1] {
			break
		}
	}
	i++
	for ; i < len(nums); i++ {
		if nums[i] > nums[i-1]  {
			if !inc {
			return false
			}
		} else if nums[i] < nums[i-1] {
			if inc {
				return false
			}
		}
	}

	return true
}

/**
724. Find Pivot Index
 */
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	si := make([]int, len(nums))
	cs := 0

	for i, n := range nums {
		cs += n
		si[i] = cs
	}

	for i, n := range nums {
		if si[i] - n == cs - si[i] {
			return i
		}
	}

	return -1
}

/**
637. Average of Levels in Binary Tree
 */

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	cds := make([]*TreeNode, 0)
	cds = append(cds, root)
	res := make([]float64, 0)

	for len(cds) != 0 {
		ncd := make([]*TreeNode, 0)
		sum := 0
		for _, n := range cds {
			sum += n.Val
			if n.Left != nil {
				ncd = append(ncd, n.Left)
			}
			if n.Right != nil {
				ncd = append(ncd, n.Right)
			}
		}
		res = append(res, float64(sum) / float64(len(cds)))
		cds = ncd
	}

	return res
}

/**
461. Hamming Distance
 */
func hammingDistance(x int, y int) int {
	x = x ^ y
	y = 0
	for x != 0 {
		y++
		x = x & (x-1)
	}
	return y
}

/**
408. Valid Word Abbreviation
 */

func validWordAbbreviation(word string, abbr string) bool {

	if len(abbr) == 0 && len(word) == 0 || len(word) == 0 && len(abbr) != 0 {
		return false
	}

	if len(abbr) == 0 {
		return true
	}

	i:=0
	for ; i<len(abbr); i++ {
		if abbr[i] > '9' || abbr[i] < '0' {
			if len(word) < i || word[i] != abbr[i] {
				return false
			}
		} else {
			break
		}
	}
	wb := i

	l := int(abbr[i]) - '0'
	for ; i < len(abbr); i++ {
		if abbr[i] <= '9' && abbr[i] >= '0' {
			l = l*10 + int(abbr[i]) - '0'
		}
	}

	if wb-1 + l >= len(word) {
		return false
	}

	return validWordAbbreviation(word[wb-1 + l:], abbr[i:])
}

/**
366. Find Leaves of Binary Tree
 */
var res366 [][]int
func findLeaves(root *TreeNode) [][]int {
	res366 = make([][]int, 1)

	findLeavesHelper(root)

	return res366
}

func findLeavesHelper(r *TreeNode) int {
	if r == nil {
		return -1
	}

	ld := findLeavesHelper(r.Left)
	rd := findLeavesHelper(r.Right)

	if ld < rd { ld = rd }
	ld++

	if len(res366) <= ld {
		for i:= len(res366)-1; i < ld; i++ {res366 = append(res366, nil)}
	}

	res366[ld] = append(res366[ld], r.Val)

	return ld
}

/**
28. Implement strStr()
 */

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}


	for i := 0; i < len(haystack) -  len(needle) + 1; i++ {
		j := 0
		for ; j < len(needle) && haystack[i+j] == needle[j]; j++ {		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}


/**
14. Longest Common Prefix
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 { return "" }

	res := make([]byte, 0)

	ml := 1<<7

	for _, s := range strs {
		if len(s) < ml {ml = len(s)}
	}

	for si := 0; si < ml; si++ {
		cc := strs[0][si]
		for j := 0; j < len(strs); j++ {
			if strs[j][si] != cc {
				return string(res)
			}
		}
		res = append(res, cc)
	}

	return string(res)
}

/**
169. Majority Element
 */

func majorityElement2(nums []int) int {
	c, cc := 0, 0

	for _, n := range nums {
		if n == c {cc++} else {
			cc--
			if cc <0 {
				c = n
				cc = 0
			}
		}

	}

	return c
}

/**
190. Reverse Bits
 */

func reverseBits(num uint32) uint32 {
	res := uint32(0)

	for i := 32; i > 0; i-- {
		res = res << 1
		if num & 0x1 == 1 {
			res++
		}
		num = num >> 1
	}

	return res
}


/**
404. Sum of Left Leaves
 */

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil { return 0 }

	res := 0
	cands := []*TreeNode{root}

	for len(cands) != 0 {
		ncs := make([]*TreeNode , 0)
		for _, n := range cands {
			if n.Left != nil && n.Left.Left == nil && n.Left.Right == nil {
				res += n.Left.Val
			}
			if n.Left != nil {
				ncs = append(ncs, n.Left)
			}
			if n.Right != nil {
				ncs = append(ncs, n.Right)
			}

		}
		cands = ncs
	}

	return res
}

/**
509. Fibonacci Number
 */

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	n -= 1
	a, b := 0, 1

	for n > 0 {
		a = a + b
		b, a = a, b
	}

	return a
}