package exercise

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

//FBrecruiting/portal/interview_prep_hub

func areTheyEqual(arr_a []int, arr_b []int) bool {

	if len(arr_a) != len(arr_b) {
		return false
	}

	h, t := -1, len(arr_a)-1

	for i, _ := range arr_a {
		if arr_a[i] != arr_b[i] {
			h = i
			break
		}
	}

	if h == -1 {
		return true
	}

	for ; t > h && arr_a[t] == arr_b[t]; t-- {
	}

	return reverseMatch(arr_a[h:t+1], arr_b[h:t+1])
}

func reverseMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[len(b)-1-i] != v {
			return false
		}
	}

	return true
}

func countSubarrays(a []int) []uint16 {
	if a == nil {
		return nil
	}

	res := make([]uint16, len(a))

	for i, v := range a {

		res[i] = 1

		for be := i - 1; be >= 0 && a[be] <= v; be-- {
			res[i]++
		}

		for af := i + 1; af < len(a) && a[af] <= v; af++ {
			res[i]++
		}
	}

	return res
}

func getMilestoneDays(revenues, milestones []int) (bool, []int) {
	if revenues == nil || milestones == nil {
		return false, nil
	}

	res := make([]int, len(milestones))

	mi, tr := 0, 0

	for d, r := range revenues {
		tr += r
		if tr >= milestones[mi] {
			res[mi] = d + 1
			mi++
		}
	}

	return true, res
}

func reverseList(head *listNode) (*listNode, *listNode) {
	if head == nil {
		return nil, nil
	}

	fh := listNode{0, nil}
	tail := head
	for n := head; n != nil; {
		tmp := fh.Next
		fh.Next = n
		n = n.Next
		fh.Next.Next = tmp
	}

	return fh.Next, tail
}

func reverseEven(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	return __reverseEven(head)
}

func __reverseEven(head *listNode) *listNode {
	if head == nil {
		return nil
	}

	var h, t *listNode
	for h = head; h.Next != nil && h.Next.V&0x01 == 1; h = h.Next {
	}

	if h.Next == nil {
		return head
	}

	for t = h.Next; t.Next != nil && t.Next.V&0x01 == 0; t = t.Next {
	}

	if t.Next != nil {
		tmp := t.Next
		t.Next = nil

		nh, nt := reverseList(h.Next)
		h.Next, nt.Next = nh, tmp
		__reverseEven(nt.Next)
	} else {
		h.Next, _ = reverseList(h.Next)
	}

	return head
}

type triangle struct {
	a, b, c uint
}

type tas []triangle

// Len is the number of elements in the collection.
func (a tas) Len() int { return len(a) }

// Less reports whether the element with
// index i should sort before the element with index j.
func (a tas) Less(i, j int) bool {
	switch {
	case a[i].a > a[j].a:
		return true
	case a[i].a == a[j].a:
		switch {
		case a[i].b > a[j].b:
			return true
		case a[i].b == a[j].b:
			switch {
			case a[i].c >= a[j].c:
				return true
			default:
				return false
			}
		default:
			return false
		}
	default:
		return false
	}
}

// Swap swaps the elements with indexes i and j.
func (a tas) Swap(i, j int) {
	a[i].a, a[j].a = a[j].a, a[i].a
	a[i].b, a[j].b = a[j].b, a[i].b
	a[i].c, a[j].c = a[j].c, a[i].c
}

func (a *triangle) sameAs(b *triangle) bool {
	return a.a == b.a && a.b == b.b && a.c == b.c
}

func (a *triangle) edgeSort() {
	edges := []int{int(a.a), int(a.b), int(a.c)}
	sort.Ints(edges)
	a.a, a.b, a.c = uint(edges[0]), uint(edges[1]), uint(edges[2])
}

func (a *triangle) edgeReorder() {
	if a.a < a.b {
		a.a, a.b = a.b, a.a
	}
	if a.a < a.c {
		a.a, a.c = a.c, a.a
	}
	if a.b < a.c {
		a.b, a.c = a.c, a.b
	}
}

func nrDiffTriangles(a tas) uint {
	switch {
	case a == nil || len(a) == 0:
		return 0
	case len(a) == 1:
		return 1
	}

	for i := 0; i < len(a); i++ {
		a[i].edgeReorder()
	}

	sort.Sort(a)

	var c uint = 0
	for i := 0; i < len(a)-1; i++ {
		if !a[i].sameAs(&a[i+1]) {
			c++
		}
	}
	c++
	return c
}

func decodeString(inp []byte) []byte {

	if inp == nil || len(inp) == 0 {
		return []byte{}
	}

	pre, reps, h, t := dsSubproblem(inp)

	if reps != 0 {
		pre = append(pre, generateReps(reps, inp[h+1:t])...)
		pre = append(pre, decodeString(inp[t+1:])...)
	}

	return pre
}

func dsSubproblem(p []byte) ([]byte, int, int, int) {
	i, h, t := 0, -1, len(p)
	reps := 0
	pre := make([]byte, 0)

	for ; i < len(p) && p[i] >= 'a' && p[i] <= 'z'; i++ {
		pre = append(pre, p[i])
	}

	for ; i < len(p) && p[i] >= '0' && p[i] <= '9'; i++ {
		reps = reps*10 + int(p[i]) - int('0')
	}

	for h = i; h < len(p) && p[h] != '['; h++ {
	}
	for t = len(p) - 1; t >= 0 && p[t] != ']'; t-- {
	}

	return pre, reps, h, t
}

func generateReps(reps int, substr []byte) []byte {
	res := make([]byte, 0)
	for ; reps > 0; reps-- {
		res = append(res, substr...)
	}
	return res
}

func numberOfWays(a []int, k int) uint {

	wc := make(map[int]uint)
	var ct uint = 0

	for _, c := range a {
		wc[c]++
	}

	for _, c := range a {
		if wc[c] != 0 {
			switch {
			case k-c == c:
				ct += wc[c] * (wc[c] - 1) >> 1
				wc[c] = 0
			default:
				ct += wc[c] * wc[k-c]
				wc[c], wc[k-c] = 0, 0
			}
		}
	}

	return ct
}

type bstNode struct {
	v           uint16
	left, right *bstNode
}

func bstSearchRange(root *bstNode, min, max uint16) []uint16 {
	if root == nil {
		return nil
	}

	res := make([]uint16, 0)

	switch {
	case root.v < min:
		return bstSearchRange(root.right, root.v, max)
	case root.v > max:
		return bstSearchRange(root.left, min, root.v)
	default:
		res = append(res, bstSearchRange(root.left, min, root.v)...)
		res = append(res, root.v)
		res = append(res, bstSearchRange(root.right, root.v, max)...)
	}

	return res
}

func allPossibleSubsets(a []uint16) [][]uint16 {
	if a == nil {
		return nil
	}

	res := make([][]uint16, 1<<uint(len(a)))
	var i, j uint16
	for i = 0; i < 1<<uint(len(a)); i++ {
		cur := make([]uint16, 0)
		for j = 0; j < uint16(len(a)); j++ {
			if (i & (1 << uint(j))) != 0 {
				cur = append(cur, a[j])
			}
		}
		res[i] = cur
	}

	return res
}

type graphNode struct {
	V     uint16
	Links []*graphNode
}

type graph struct {
	nodes []*graphNode
}

func (g *graph) topoSorting() (bool, []*graphNode) {
	if g.nodes == nil {
		return false, nil
	}

	res := make([]*graphNode, 0, len(g.nodes))

	roots := make([]*graphNode, 0)
	d2n := make(map[*graphNode]uint)

	for _, n := range g.nodes {
		if n.Links != nil {
			for _, on := range n.Links {
				d2n[on]++
			}
		}
	}

	for _, v := range g.nodes {
		_, ok := d2n[v]
		if !ok {
			roots = append(roots, v)
		}
	}

	for len(roots) > 0 {
		cn := roots[len(roots)-1]
		roots = roots[:len(roots)-1]
		res = append(res, cn)

		for _, n := range cn.Links {
			d2n[n]--
			if d2n[n] == 0 {
				roots = append(roots, n)
			}
		}
	}

	return len(g.nodes) == len(res), res
}

// this code has bugs.
func posNegNums(a []int) {
	if a == nil || len(a) < 3 {
		return
	}

	pc := 0
	for _, v := range a {
		if v > 0 {
			pc++
		}
	}

	np := pc >= len(a)&01

	for c, pb := 0, 0; pb < len(a); {
		switch np {
		case true:
			switch {
			case a[pb] > 0:
				a[c], a[pb] = a[pb], a[c]
				c += 2
				pb = c
			default:
				pb++
			}
		default:
			switch {
			case a[pb] < 0:
				a[c], a[pb] = a[pb], a[c]
				c += 2
				pb = c
			default:
				pb++
			}
		}
	}
}

func jumpGame(js []int) bool {

	rc := make([]bool, len(js))
	rc[0] = true

	for i := 0; i < len(js) && rc[i]; i++ {
		for j := 1; j <= js[i] && i+j < len(js); rc[i+j], j = true, j+1 {
		}
	}

	return rc[len(js)-1]
}

// Longest Consecutive Sequence
// assume the range is from 0 to 255
func longestCS(nums []uint) uint {
	bm := []uint32{0, 0, 0, 0, 0, 0, 0, 0}

	for _, v := range nums {
		bm[v>>5] |= 1 << (v & 0x1f)
	}

	lcs, cl := uint(0), uint(0)

	for i := 0; i < 8; i++ {
		for j := uint(0); j < 32 && bm[i]>>j != 0; j++ {
			switch bm[i] & (1 << j) {
			case 0:
				if cl > lcs {
					lcs = cl
				}
				cl = 0
			default:
				cl++
			}
		}
	}

	if cl > lcs {
		lcs = cl
	}

	return lcs
}

func CombinationSum(nums []int, s int) ([][]int, bool) {
	if len(nums) == 0 {
		switch s {
		case 0:
			return [][]int{[]int{}}, true
		default:
			return nil, false
		}
	}

	var res [][]int
	found := false

	cans, ok := CombinationSum(nums[1:], s-nums[0])
	if ok {
		for i := 0; i < len(cans); i++ {
			cans[i] = append(cans[i], nums[0])
		}
		res = cans
		found = true
	}

	cans, ok = CombinationSum(nums[1:], s)
	if ok {
		res = append(res, cans...)
		found = true
	}

	return res, found
}

func rehashing(hm [][]int) [][]int {
	cap := len(hm) << 1
	res := make([][]int, cap)

	for _, b := range hm {
		for _, v := range b {
			res[v%cap] = append(res[v%cap], v)
		}
	}

	return res
}

type Heap struct {
	data []int
}

func (h *Heap) heapifyAt(ind uint) {
	for pp := ind >> 1; ind != 0 && h.data[pp] > h.data[ind]; {
		h.data[pp], h.data[ind] = h.data[ind], h.data[pp]
		ind = ind >> 1
		pp = ind >> 1
	}
}

func (h *Heap) heapify() {
	for i, _ := range h.data {
		h.heapifyAt(uint(i))
	}
}

func isPower2(n uint) bool {
	return n&(n-1) == 0
}

//Container With Most Water
func CWMoreWatter(ls []uint) uint {
	if ls == nil || len(ls) == 0 {
		return 0
	}

	rm := uint(0)
	for s, e := 0, len(ls)-1; s < e; {
		switch {
		case ls[s] > ls[e]:
			if rm < ls[e]*uint(e-s) {
				rm = ls[e] * uint(e-s)
			}
			e--
		default:
			if rm < ls[s]*uint(e-s) {
				rm = ls[s] * uint(e-s)
			}
			s++
		}
	}
	return rm
}

// Minimum Size Subarray Sum
func MinSubarrSum(nums []int, t int) (int, bool) {
	if nums == nil || len(nums) == 0 {
		return 0, false
	}

	s, e, cs := 0, 0, 0
	for ; e < len(nums); e++ {
		cs += nums[e]
		if cs >= t {
			break
		}
	}

	if e == len(nums) {
		return 0, false
	}

	md := 1 << 30
	for true {
		if cs >= t {
			if e-s < md {
				md = e - s
			}
			cs -= nums[s]
			s++
		} else {
			e++
			if e == len(nums) {
				break
			}
			cs += nums[e]
		}
	}

	return md + 1, true
}

//Find K Closest Elements
func KClosest(e []int, t, k int) []int {
	res := make([]int, 0, k)
	p, ok := binarySearch(e, t)

	be, af := p, p+1
	if ok {
		res = append(res, e[p])
		k--
		be, af = p-1, p+1
	}

	for k > 0 && be >= 0 && af < len(e) {
		switch {
		case be >= 0 && af < len(e):
			switch {
			case t-e[be] <= e[af]-t:
				res = append(res, e[be])
				be--
				k--
			default:
				res = append(res, e[af])
				af++
				k--
			}
		case be == -1 && af < len(e):
			res = append(res, e[af])
			af++
			k--
		case be >= 0 && af == len(e):
			res = append(res, e[be])
			be--
			k--
		}
	}

	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

type Wdc struct {
	s *string
	c uint
}

type WCHeap []*Wdc

func (h WCHeap) Len() int      { return len(h) }
func (h WCHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h WCHeap) Less(i, j int) bool {
	if h[i].c < h[j].c {
		return true
	}
	if h[i].c == h[j].c {
		return *h[i].s > *h[j].s
	}
	return false
}
func (h *WCHeap) Push(x interface{}) {
	*h = append(*h, x.(*Wdc))
}

func (h *WCHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// K Frequent Words
func KFreqWords(wds []string, k int) []string {
	if wds == nil || len(wds) == 0 {
		return wds
	}
	if k == 0 {
		return make([]string, 0, 0)
	}

	ci := make(map[string]*Wdc, len(wds))

	for i, _ := range wds {
		wdc, ok := ci[wds[i]]
		switch ok {
		case true:
			wdc.c++
		default:
			ci[wds[i]] = &Wdc{&wds[i], 1}
		}
	}

	h := &WCHeap{}
	sc := 0
	for _, wc := range ci {
		sc++
		heap.Push(h, wc)
		if sc > k {
			heap.Pop(h)
		}
	}

	var res []string
	if k > h.Len() {
		res = make([]string, h.Len())
	} else {
		res = make([]string, k)
	}
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = *heap.Pop(h).(*Wdc).s
	}
	return res
}

// Combination Sum
func CombSum(n []uint, t uint) uint {
	cs := make([]uint, t+1)
	cs[0] = 1
	for ct := uint(1); ct <= t; ct++ {
		for _, v := range n {
			if ct >= v {
				cs[ct] += cs[ct-v]
			}
		}
	}

	return cs[t]
}

// Find the Missing Number, r < 30
func MissNumber(s []byte, r uint) (uint, uint, bool) {
	if s == nil || len(s) == 0 {
		return r + 1, r + 1, false
	}

	cc := make([]uint, 10)
	for i := uint(1); i <= r; i++ {
		cc[i%10]++
		if i > 9 {
			cc[i/10]++
		}
	}

	wc := make([]uint, 10)
	for _, v := range s {
		wc[v-'0']++
	}

	mns := make([]uint, 0, 2)

	for i := uint(0); i < 10; i++ {
		if cc[i] != wc[i] {
			mns = append(mns, i)
		}
	}

	switch len(mns) {
	case 1:
		return mns[0], r + 1, true
	case 2:
		return mns[0], mns[1], true
	default:
		return r + 1, r + 1, false
	}
}

// Partition Equal Subset Sum
func PartSubset(ns []uint) bool {
	t := uint(0)
	for _, v := range ns {
		t += v
	}

	if t&0x1 == 1 {
		return false
	}

	t = t >> 1

	return subsetSum(ns, t)
}

func subsetSum(ns []uint, t uint) bool {
	if len(ns) == 0 {
		return t == 0
	}

	if t >= ns[len(ns)-1] && subsetSum(ns[:len(ns)-1], t-ns[len(ns)-1]) {
		return true
	}

	return subsetSum(ns[:len(ns)-1], t)
}

// Binary Tree Longest Consecutive Sequence II
var btLCSMax = uint(0)

func BTLCS(r *bstNode) uint {
	btLCSMax = 0
	btLCS(r)
	return btLCSMax
}

func btLCS(r *bstNode) (uint, uint) {
	if r == nil {
		return 0, 0
	}

	lam, ldm, ram, rdm := uint(1), uint(1), uint(1), uint(1)

	if r.left != nil {
		if r.left.v == r.v+1 || r.left.v+1 == r.v {
			lam, ldm = btLCS(r.left)
			if r.left.v == r.v+1 {
				lam++
			}
			if r.left.v+1 == r.v {
				ldm++
			}
		}
	}






	if r.right != nil {
		if r.right.v == r.v+1 || r.right.v+1 == r.v {
			ram, rdm = btLCS(r.right)
			if r.right.v == r.v+1 {
				ram++
			}
			if r.right.v+1 == r.v {
				rdm++
			}
		}
	}

	if lam+rdm-1 > btLCSMax {
		btLCSMax = lam + rdm - 1
	}
	if ldm+ram-1 > btLCSMax {
		btLCSMax = ldm + ram - 1
	}

	if lam < ram {
		lam = ram
	}
	if ldm < rdm {
		ldm = rdm
	}

	return lam, ldm
}

// Maximum Average Subarray II
func MaxAvgSubarr(ns []int, l int) (int, int) {
	if ns == nil || len(ns) < l {
		return 0, 1
	}

	sumbuf := make([]int, len(ns)+1)

	for i, v := range ns {
		sumbuf[i+1] = sumbuf[i] + v
	}

	msum, ml := sumbuf[len(ns)]-sumbuf[len(ns)-l], l

	for e := len(ns); e >= l; e-- {
		for h := e - l; h >= 0; h-- {
			cs := sumbuf[e] - sumbuf[h]
			if msum*(e-h) < cs*ml {
				msum = cs
				ml = e - h
			}
		}
	}

	return msum, ml
}

type K9board struct {
	bd             [][]byte
	sx, sy, tx, ty int
	bh, bw         int
}

type coor struct {
	x, y int
}

//Knight Shortest Path
func (b *K9board) K9Path() int {
	// sanity check to be added
	cands := make([]coor, 0)
	cands = append(cands, coor{b.sx, b.sy})
	st := -1

	for len(cands) != 0 {
		cp := cands[0]
		st++
		if b.wins(&cp) {
			return st
		}
		cands = cands[1:]

		b.bd[cp.x][cp.y] = 1
		switch {
		case cp.x > 0 && cp.y > 1 && b.bd[cp.x-1][cp.y-2] != 1:
			cands = append(cands, coor{cp.x - 1, cp.y - 2})
		case cp.x > 1 && cp.y > 0 && b.bd[cp.x-2][cp.y-1] != 1:
			cands = append(cands, coor{cp.x - 2, cp.y - 1})
		case cp.x > 0 && cp.y+1 < b.bh && b.bd[cp.x-1][cp.y+2] != 1:
			cands = append(cands, coor{cp.x - 1, cp.y + 2})
		case cp.x > 1 && cp.y < b.bh && b.bd[cp.x-2][cp.y+1] != 1:
			cands = append(cands, coor{cp.x - 2, cp.y + 1})

		case cp.x < b.bw && cp.y > 1 && b.bd[cp.x+1][cp.y-2] != 1:
			cands = append(cands, coor{cp.x + 1, cp.y - 2})
		case cp.x+1 > b.bw && cp.y > 0 && b.bd[cp.x+2][cp.y-1] != 1:
			cands = append(cands, coor{cp.x + 2, cp.y - 1})
		case cp.x < b.bw && cp.y+1 < b.bh && b.bd[cp.x+1][cp.y+2] != 1:
			cands = append(cands, coor{cp.x + 1, cp.y + 2})
		case cp.x+1 > b.bw && cp.y > b.bh && b.bd[cp.x+2][cp.y+1] != 1:
			cands = append(cands, coor{cp.x + 2, cp.y + 1})
		}
	}
	return -1
}

func (b *K9board) wins(p *coor) bool {
	if p.x == b.tx && p.y == b.ty {
		return true
	}
	return false
}

//Minimum number of swaps required to sort an array of first N number
func minSwap2Sorted(ns []uint) uint {
	if ns == nil {
		return 0
	}

	ct := uint(0)
	for i, _ := range ns {
		for ns[i] != uint(i)+1 {
			ct++
			ns[ns[i]-1], ns[i] = ns[i], ns[ns[i]-1]
		}
	}
	return ct
}

//Maximum number of unique values
func maxOfUniq(ns []int) uint {
	if ns == nil {
		return 0
	}

	sort.Ints(ns)

	l := len(ns)

	bm := make([]int, l>>5+1)

	for _, v := range ns {
		switch {
		case bm[(v-1)>>5]&(1<<uint((v-1)&0x1f)) == 0:
			bm[(v-1)>>5] = (1 << uint((v-1)&0x1f)) | bm[(v-1)>>5]
		case bm[(v)>>5]&(1<<uint((v)&0x1f)) == 0:
			bm[(v)>>5] = (1 << uint((v)&0x1f)) | bm[(v)>>5]
		case bm[(v+1)>>5]&(1<<uint((v+1)&0x1f)) == 0:
			bm[(v+1)>>5] = (1 << uint((v+1)&0x1f)) | bm[(v+1)>>5]
		}
	}

	ct := uint(0)

	for _, v := range bm {
		for v != 0 {
			ct++
			v = v & (v - 1)
		}
	}

	return ct
}

//Binary Tree Vertical Order Traversal
type dlink struct {
	pre, after *dlink
	ns         []uint16
}

func BTVertOrder(r *bstNode) [][]uint16 {
	if r == nil {
		return nil
	}

	rtln := dlink{nil, nil, nil}
	btVertOrder(r, &rtln)

	res := make([][]uint16, 0)
	for cn := &rtln; cn != nil; cn = cn.pre {
		res = append(res, cn.ns)
	}

	for i := 0; i < len(res)>>1; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	for cn := rtln.after; cn != nil; cn = cn.after {
		res = append(res, cn.ns)
	}

	return res
}

func btVertOrder(r *bstNode, ln *dlink) {
	if r == nil {
		return
	}

	if ln.ns == nil {
		ln.ns = []uint16{r.v}
	} else {
		ln.ns = append(ln.ns, r.v)
	}

	if r.left != nil {
		if ln.pre == nil {
			ln.pre = &dlink{nil, ln, nil}
		}
		btVertOrder(r.left, ln.pre)
	}
	if r.right != nil {
		if ln.after == nil {
			ln.after = &dlink{ln, nil, nil}
		}
		btVertOrder(r.right, ln.after)
	}
}

//Factorization
func factorization(num uint) [][]uint {
	res := make([][]uint, 0)
	if num == 1 {
		return res
	}

	for i := uint(2); i*i <= num; i++ {
		if num%i == 0 {
			subres := factorization(num / i)
			for _, v := range subres {
				res = append(res, append([]uint{i}, v...))
			}
		}
	}
	res = append(res, []uint{num})

	return res
}

const (
	SIGNROOM = 0x7fffffff
	SIGNGATE = 0
	SIGNWALL = 0xffffffff
)

//Walls and Gates
func wallsGates(bd [][]int) {
	if bd == nil || len(bd) == 0 {
		return
	}

	ce := make([]*coor, 0)

	for y, _ := range bd {
		for x, v := range bd[y] {
			if v == SIGNGATE {
				ce = append(ce, &coor{x, y})
			}
		}
	}

	ct := 0
	for len(ce) != 0 {
		ne := make([]*coor, 0)
		ct++
		for _, c := range ce {

			if c.x > 0 && bd[c.y][c.x-1] == SIGNROOM {
				ne = append(ne, &coor{c.x - 1, c.y})
				bd[c.y][c.x-1] = ct
			}
			if c.y > 0 && bd[c.y-1][c.x] == SIGNROOM {
				ne = append(ne, &coor{c.x, c.y - 1})
				bd[c.y-1][c.x] = ct
			}
			if c.x+1 < len(bd[0]) && bd[c.y][c.x+1] == SIGNROOM {
				ne = append(ne, &coor{c.x + 1, c.y})
				bd[c.y][c.x+1] = ct
			}
			if c.y+1 < len(bd) && bd[c.y+1][c.x] == SIGNROOM {
				ne = append(ne, &coor{c.x, c.y + 1})
				bd[c.y+1][c.x] = ct
			}
		}
		ce = ne
	}
}

//Counting Bits
func cout1s(m uint) []uint {
	res := make([]uint, 0, m+1)

	for i := uint(0); i <= m; i++ {
		n := i
		ct := uint(0)
		for n != 0 {
			ct++
			n = n & (n - 1)
		}
		res = append(res, ct)
	}

	return res
}

// Longest Palindromic Subsequence
func maxPali(s string) uint {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	max := uint(0)

	if s[0] == s[len(s)-1] {
		max = 2 + maxPali(s[1:len(s)-1])
	} else {
		sr := maxPali(s[1:])
		if max < sr {
			max = sr
		}

		sr = maxPali(s[:len(s)-1])
		if max < sr {
			max = sr
		}
		sr = maxPali(s[1 : len(s)-1])
		if max < sr {
			max = sr
		}
	}

	return max
}

// Ones and Zeroes
func onesAndZeors(ss []string, os, zs int) uint {
	if len(ss) == 0 || os < 0 || zs < 0 {
		return 0
	}

	los, lzs := 0, 0
	for _, c := range ss[len(ss)-1] {
		switch c {
		case '0':
			lzs++
		case '1':
			los++
		}
	}

	tl := 1 + onesAndZeors(ss[:len(ss)-1], os-los, zs-lzs)
	ntl := onesAndZeors(ss[:len(ss)-1], os, zs)

	if tl > ntl {
		return tl
	} else {
		return ntl
	}
}

//Coin Change
func coinChange(cs []int, t int) int {
	if t < 0 {
		return -1
	}

	if t == 0 {
		return 0
	}

	min := 0x7fffffff
	for _, c := range cs {
		mc := coinChange(cs, t-c)
		if mc != -1 && min > mc+1 {
			min = mc + 1
		}
	}

	if min == 0x7fffffff {
		return -1
	}

	return min
}

func coinChangeBuf(cs []int, t int) int {
	buf := make([]int, t+1)

	for i, _ := range buf {
		buf[i] = 0x7fffffff
	}

	for _, c := range cs {
		buf[c] = 1
	}

	prog := true
	for prog {
		prog = false
		for _, c := range cs {
			for p := 1; p < t+1 && buf[p] != 0; p++ {
				if p+c < t+1 && buf[p+c] > buf[p] {
					buf[p+c] = buf[p] + 1
					prog = true
				}
			}
		}
	}

	if buf[t] == 0x7fffffff {
		return -1
	}

	return buf[t]
}

// Number of Big Islands
type bigIsl struct {
	bd         [][]byte
	isize, iss int
}

func (g *bigIsl) calc() {

	for i := range g.bd {
		for j := range g.bd[0] {
			if g.bd[i][j] == 1 {
				g.explore(i, j)
			}
		}
	}
}

func (g *bigIsl) explore(x, y int) {
	backlog := make([]coor, 0)
	backlog = append(backlog, coor{x, y})

	cs := 1
	if cs == g.isize {
		g.iss++
	}
	g.bd[x][y] = 2

	for len(backlog) != 0 {
		cur := backlog[0]
		backlog = backlog[1:]

		if cur.x > 0 && g.bd[cur.x-1][cur.y] == 1 {
			cs++
			if cs == g.isize {
				g.iss++
			}
			g.bd[cur.x-1][cur.y] = 2
			backlog = append(backlog, coor{cur.x - 1, cur.y})
		}

		if cur.y > 0 && g.bd[cur.x][cur.y-1] == 1 {
			cs++
			if cs == g.isize {
				g.iss++
			}
			g.bd[cur.x][cur.y-1] = 2
			backlog = append(backlog, coor{cur.x, cur.y - 1})
		}

		if cur.y+1 < len(g.bd[0]) && g.bd[cur.x][cur.y+1] == 1 {
			cs++
			if cs == g.isize {
				g.iss++
			}
			g.bd[cur.x][cur.y+1] = 2
			backlog = append(backlog, coor{cur.x, cur.y + 1})
		}

		if cur.x+1 < len(g.bd) && g.bd[cur.x+1][cur.y] == 1 {
			cs++
			if cs == g.isize {
				g.iss++
			}
			g.bd[cur.x+1][cur.y] = 2
			backlog = append(backlog, coor{cur.x + 1, cur.y})
		}
	}
}

// Shortest Palindrome
func stestPali(s []byte) int {

	for i := len(s); i > 0; i-- {
		if isPali(s[:i]) {
			return len(s) - i
		}
	}

	return len(s) - 1
}

func isPali(s []byte) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i <= len(s)>>1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Split String
func splitStr(s []byte) [][]string {
	res := make([][]string, 0)

	if len(s) == 0 {
		return res
	}
	if len(s) == 1 {
		return append(res, []string{string(s)})
	}
	if len(s) == 2 {
		res = append(res, []string{string(s)})
		return append(res, []string{string(s[0]), string(s[1])})
	}

	s1 := splitStr(s[1:])
	for i, _ := range s1 {
		s1[i] = append(s1[i], string(s[0]))
	}
	res = append(res, s1...)

	s2 := splitStr(s[2:])
	for i, _ := range s2 {
		s2[i] = append(s2[i], string(s[:2]))
	}
	res = append(res, s2...)

	return res
}

// Word Break III
func wordBrk3(dict []string, st string) int {
	if len(st) == 0 {
		return 0
	}

	di := make(map[string]byte)
	for _, w := range dict {
		di[w] = 1
	}

	return _wb(di, st)
}

func _wb(di map[string]byte, st string) int {

	if len(st) == 0 {
		return 1
	}

	res := 0

	for i, _ := range st {
		_, ok := di[string(st[:i+1])]
		if ok {
			res += _wb(di, st[i+1:])
		}
	}

	return res
}

// Sliding Window Unique Elements Sum
func swues(ns []int, ws int) int {
	if len(ns) < ws {
		return 0
	}

	wc := make(map[int]int)
	for i := 0; i < ws; i++ {
		_, ok := wc[ns[i]]
		switch ok {
		case true:
			wc[ns[i]]++
		case false:
			wc[ns[i]] = 1
		}
	}

	nu := 0
	for _, c := range wc {
		if c == 1 {
			nu++
		}
	}

	res := nu

	for i := ws; i < len(ns); i++ {
		if ns[i] != ns[i-ws] {
			wc[ns[i-ws]]--
			if wc[ns[i-ws]] == 0 {
				nu--
			} else if wc[ns[i-ws]] == 1 {
				nu++
			}

			wc[ns[i]]++
			if wc[ns[i]] == 1 {
				nu++
			} else if wc[ns[i]] == 2 {
				nu--
			}
		}
		res += nu
	}

	return res
}

// Trim a Binary Search Tree
func trimBST(rt *bstNode, min, max uint16) *bstNode {
	if rt == nil {
		return nil
	}

	switch {
	case rt.v < min:
		return trimBST(rt.right, min, max)
	case rt.v > max:
		return trimBST(rt.left, min, max)
	default:
		rt.left = trimBST(rt.left, min, max)
		rt.right = trimBST(rt.right, min, max)
		return rt
	}
}

func printBST(rt *bstNode) {
	if rt == nil {
		return
	}

	buff := []*bstNode{rt}

	for len(buff) != 0 {
		nbuf := make([]*bstNode, 0)
		for _, v := range buff {
			fmt.Printf("%v\t", v.v)
			if v.left != nil {
				nbuf = append(nbuf, v.left)
			}
			if v.right != nil {
				nbuf = append(nbuf, v.right)
			}
		}
		buff = nbuf
		fmt.Println()
	}
}

//Cutting a Rod
func cutRod(vs []uint, l uint) uint {
	maxv := make([]uint, l+1)

	for cur := uint(1); cur <= l; cur++ {
		max := uint(0)
		for i := uint(0); i < cur; i++ {
			if vs[i]+maxv[cur-i-1] > max {
				max = vs[i] + maxv[cur-i-1]
			}
		}
		maxv[cur] = max
	}
	return maxv[l]
}

// Minimum Partition
func minPart(ns []int) int {
	sum, min := 0, 0x7fffffff
	for _, v := range ns {
		sum += v
	}

	nbs := uint(len(ns))

	for i := 1; i < (1 << nbs); i++ {
		cs := 0
		for j := uint(0); j < nbs; j++ {
			if i&(1<<j) != 0 {
				cs += ns[j]
			}
		}

		if min > absolute((cs<<1)-sum) {
			min = absolute((cs << 1) - sum)
		}
	}

	return min
}

func absolute(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Binary Tree Right Side View
func bstRightView(rt *bstNode) []uint16 {
	res := make([]uint16, 0)

	if rt == nil {
		return res
	}

	buf := make([]*bstNode, 0)
	buf = append(buf, rt)

	for len(buf) != 0 {
		res = append(res, buf[0].v)
		nbuf := make([]*bstNode, 0)
		for _, n := range buf {
			if n.right != nil {
				nbuf = append(nbuf, n.right)
			}
			if n.left != nil {
				nbuf = append(nbuf, n.left)
			}
		}
		buf = nbuf
	}

	return res
}

//Give change
var cs = []int{1, 4, 16, 64}

func minChanges(mt int) (int, bool) {
	if mt <= 0 {
		return 0, false
	}
	if mt == 0 {
		return 0, true
	}

	res, ok := 0x1fffffff, false

	for i:=0; i< len(cs); i++ {
		cm, cok := minChanges(mt - cs[i])
		if cok { ok = true }
		if cok && (cm < res) {
			res = cm
		}
	}

	return res, ok
}

func minChanges2(mt int) (int, bool) {
	buf := make([]int, mt + 1)

	cmax := cs[len(cs)-1]

	for i := 0; i <= cmax; i++ {
		if buf[i] != 0 {
			for ci := 0; ci < len(cs); ci++ {
				if buf[i+cs[ci]] != 0 {
					buf[i+cs[ci]] = i
				}
			}
		}
		cmax += cs[len(cs)-1]
	}

	return buf[mt],  buf[mt] != 0
}


/**
253. Meeting Rooms II
Share
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

Example 1:

Input: [[0, 30],[5, 10],[15, 20]]
Output: 2
 */

func minMeetingRooms(ints [][]int) int {

	if ints == nil || len(ints) == 0 {
		return 0
	}

	mts := intervals{}
	for _, v := range ints {
		mts = append(mts, interval{v[0], v[1]})
	}

	sort.Sort(mts)

	ol := &intervalheap{}
	latest, max := -1, 0

	for _, v := range mts {
		if v.st < latest {
			for len(*ol) != 0 {
				top := heap.Pop(ol).(interval)
				if top.end > v.st {
					heap.Push(ol, top)

					break
				}
			}
		} else {
			ol = &intervalheap{}
		}

		heap.Push(ol, v)
		if max < len(*ol) {
			max = len(*ol)
		}
		if latest < v.end {
			latest = v.end
		}
	}

	return max
}

func canAttendMeetings(ints [][]int) bool {
	if ints == nil || len(ints) == 0 {
		return true
	}

	mts := intervals{}
	for _, v := range ints {
		mts = append(mts, interval{v[0], v[1]})
	}

	sort.Sort(mts)
	lst := -1
	for _, v := range mts {
		if v.st < lst {
			return false
		}

		if lst < v.end {
			lst = v.end
		}
	}

	return true
}

/**
127. Word Ladder
Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5
 */

func ladderLength(bw, ew string, wl []string) int {

	cands := make([]string, 0)
	cands = append(cands, bw)
	ss := 1
	wl = append(wl, ew)

	for len(cands) != 0 {
		ncands := make([]string, 0)
		for _, c := range cands {
			if c == ew {
				return ss
			}
			for i, _ := range wl {
				if oneMismatch(wl[i], c) {
					ncands = append(ncands, wl[i])
					wl[i] = string(append([]byte(wl[i]), ' '))
				}
			}
		}
		ss++
		cands = ncands
	}

	return 0
}

func oneMismatch(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	mm := 0

	for i := 0; i < len(a) && mm < 2; i++ {
		if a[i] != b[i] {
			mm++
		}
	}

	return mm == 1
}

/**
394. Decode String
Input: s = "3[a]2[bc]"
Output: "aaabcbc"
 */

func decodeStringIt(s string) string {

	rmd := make([]token, 0)

	tks := tokenize(s)

	for _, t := range tks {
		switch t.tp {
		case 1:
			if len(rmd) == 0 {
				rmd = append(rmd, t)
			} else {
				rmd[len(rmd)-1].cs = append(rmd[len(rmd)-1].cs, t.cs...)
			}

		case 2:
			rmd = append(rmd, t)

		case 3:
			// NOP
		case 4:
			st := genStr(rmd[len(rmd)-1])
			rmd = rmd[:len(rmd)-1]
			if len(rmd) != 0 {
				rmd[len(rmd)-1].cs = append(rmd[len(rmd)-1].cs, st...)
			} else {
				rmd = append(rmd, token{st, 1, 1})
			}
		}
	}

	return string(rmd[0].cs)
}

func genStr(t token) []byte {
	res := make([]byte,0)
	for i:= 0; i < t.v; i++ {
		res = append(res, t.cs...)
	}
	return res
}

func tokenize(st string) []token {

	tks := make([]token, 0)
	for _, c := range []byte(st) {
		switch {
		case c  == '[' :
			tks = append(tks, token{nil, 0, 3})
		case c  == ']' :
			tks = append(tks, token{nil, 0, 4})
		case c >= 'a' && c <= 'z'  || c >= 'A' && c <= 'Z':
			if len(tks) != 0  && tks[len(tks)-1].tp == 1 {
				tks[len(tks)-1].cs = append(tks[len(tks)-1].cs, c)
			} else {
				tks = append(tks, token{[]byte{c}, 1, 1})
			}
		case c >= '0' && c <= '9':
			if len(tks) != 0 && tks[len(tks)-1].tp == 2 {
				tks[len(tks)-1].v = tks[len(tks)-1].v * 10 + int(c - '0')
			} else {
				tks = append(tks, token{nil, int(c - '0'), 2})
			}
		}
	}

	return tks
}

type token struct {
	cs []byte
	v int
	tp int
}

/**
727. Minimum Window Subsequence
Input:
S = "abcdebdde", T = "bde"
Output: "bcde"
 */

func minWindow(S string, T string) string {

	s, e := _mw(S, T)
	if s == -1 {
		return ""
	} else {
		return string(S[s:e])
	}
}

func _mw(s, t string) (int, int) {
	if len(t) == 0 {
		return 0,0
	}
	if len(s) == 0 {
		return -1, -1
	}

	rs, re := -1, -1
	mlen := 0x7fffffff

	if s[0] == t[0] {
		st, end := _mw(s[1:], t[1:])
		if st != -1 {
			mlen = end + 1
			rs, re = 0, 1 + end
		}
		st, end = _mw(s[1:], t)
		if st != -1 && end - st + 1 < mlen{
			mlen = end - st + 1
			rs, re = st+1, end+1
		}
	} else {
		st, end := _mw(s[1:], t)
		if st != -1 && end - st + 1 < mlen{
			mlen = end - st + 1
			rs, re = st+1, end+1
		}
	}

	return rs, re
}


/**
1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
Given an array of integers nums and an integer limit, return the size of the longest non-empty subarray such that the absolute difference between any two elements of this subarray is less than or equal to limit.
Input: nums = [8,2,4,7], limit = 4
Output: 2
*/

func longestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {return 0}

	ci, ca := nums[0], nums[0]
	l, st, s := 0, 0, 0

	for i, v := range nums {
		switch {
		case v > ca:
			if v > ci + limit {
				if i - st > l {
					l = i - st
				}
				s, ci, ca = minMax(nums[st:i+1], v - limit, v)
				st += s
			}
			ca = v
		case v < ci:
			if v < ca - limit {
				if i - st > l {l = i -st}
				s, ci, ca = minMax(nums[st:i+1], v, v+limit)
				st += s
			}
			ci = v
		default:
		}
	}
	if len(nums) - st > l {l = len(nums) - st}
	return l
}

func minMax(ns []int, min, max int) (int , int , int) {
	if len(ns) == 0 {
		return 0, 0, 0
	}
	e := len(ns) - 1
	ci, ca := ns[e], ns[e]
	for ; e >= 0 && ns[e] <= max && ns[e] >= min; e-- {
		if ns[e] > ca { ca = ns[e]}
		if ns[e] < ci { ci = ns[e]}
	}

	return e + 1, ci, ca
}


func asteroidCollision(asteroids []int) []int {

	res := make([]int, 0)
	rr := make([]int, 0)

	for _, v := range asteroids {
		switch {
		case v < 0:
			switch {
			case len(rr) == 0:
				res = append(res, v)
			default:
				rri := len(rr) -1
				for ; rri >= 0; rri-- {
					if rr[rri] > -v { rr = rr[:rri+1]; break}
					if rr[rri] == -v { rr = rr[:rri]; break}
				}
				if rri < 0 {res = append(res, v); rr = rr[:0]}
			}
		default:
			rr = append(rr, v)
		}
	}

	for _, v := range rr {res = append(res, v)}

	return res
}


type TimeMap struct {
	kvm map[string]tvs
	updated bool
}

type timedValue struct {
	v string
	ts int
}

type tvs []timedValue

func (s tvs) Len() int {return len(s)}
func (s tvs) Less(i, j int) bool {return s[i].ts < s[j].ts }
func (s tvs) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func (s tvs) searchES(k int) int {
	if len(s) == 0 {return -1}
	l, r, m := 0, len(s)-1, 0
	for l <= r {
		m = l + (r-l)/2

		if s[m].ts == k {
			return m
		}
		if s[m].ts < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if l == len(s) { if s[r].ts < k {return r} else {return -1}}
	if r == -1 { if s[l].ts < k {return l} else { return -1}}
	if s[l].ts > k {return l-1}

	return l
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	tm := TimeMap{ make(map[string]tvs), false}
	return tm
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	tvss, ok := this.kvm[key]
	if !ok {
		this.kvm[key] = tvs{timedValue{value, timestamp}}
	} else {
		tvss = append(tvss, timedValue{value, timestamp})
		sort.Sort(tvss)
		this.kvm[key] = tvss
	}
}


func (this *TimeMap) Get(key string, timestamp int) string {
	tvss, ok := this.kvm[key]
	if !ok {
		return ""
	} else {
		m := tvss.searchES(timestamp)
		if  m>=0 {return tvss[m].v} else {return ""}
	}
}

/**
304. Range Sum Query 2D - Immutable
Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
 */

type NumMatrix struct {
	m [][]int
	us [][]int
}


func NumMatrixConstructor(matrix [][]int) NumMatrix {

	nm := NumMatrix{matrix, make([][]int, len(matrix))}
	for i:= 0; i < len(matrix); i++ {nm.us[i] = make([]int, len(matrix[0]))}
	nm.us[0][0] = matrix[0][0]
	for i:=1; i < len(matrix[0]); i++ {
		nm.us[0][i] += matrix[0][i] + nm.us[0][i-1]
	}

	for i:=1; i < len(matrix); i++ {nm.us[i][0] = nm.us[i-1][0] + matrix[i][0]}

	for i:=1; i < len(matrix); i++ {
		for j:=1; j < len(matrix[0]); j++ {
			nm.us[i][j] = nm.us[i-1][j] + nm.us[i][j-1] - nm.us[i-1][j-1] + matrix[i][j]
		}
	}
	return nm
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {


	switch {
	case row1 == 0 && col1 == 0:
		return this.us[row2][col2]
	case row1==0:
		return this.us[row2][col2] - this.us[row2][col1]
	case col1==0:
		return this.us[row2][col2] - this.us[row1][col2]
	default:
		return this.us[row2][col2] - this.us[row1-1][col2] - this.us[row2][col1-1] + this.us[row1-1][col1-1]
	}
}

/**
398. Random Pick Index
Given an array of integers with possible duplicates, randomly output the index of a given target number.
You can assume that the given target number must exist in the array.
 */

type Solution struct {
	idx map[int][]int
}


func RPIConstructor(nums []int) Solution {
	s := Solution{make(map[int][]int)}

	for i,v := range nums{
		ps, ok := s.idx[v]
		if ok {
			s.idx[v] = append(ps, i)
		} else {
			s.idx[v] = []int{i}
		}
	}

	return s
}


func (this *Solution) Pick(target int) int {
	ps, ok := this.idx[target]
	if ok {
		rd := rand.Intn(len(ps))
		return ps[rd]
	}
	return -1
}


type MaxStack struct {
	st []int
	ms [][]int
}


/** initialize your data structure here. */
func maxStackConstructor() MaxStack {
	return MaxStack{make([]int, 0), make([][]int, 0)}
}


func (this *MaxStack) Push(x int)  {
	this.st = append(this.st, x)
	if len(this.ms) == 0 {
		this.ms = append(this.ms, []int{x,0})
	} else if x >= this.ms[len(this.ms)-1][0] {
		this.ms = append(this.ms, []int{x,len(this.st) -1})
	}
}


func (this *MaxStack) Pop() int {
	if len(this.st) == 0 {return -1}
	v := this.st[len(this.st) - 1]
	if v == this.ms[len(this.ms)-1][0] {this.ms = this.ms[:len(this.ms)-1]}
	this.st = this.st[:len(this.st) - 1]
	return v
}


func (this *MaxStack) Top() int {
	return this.st[len(this.st) - 1]
}


func (this *MaxStack) PeekMax() int {
	return this.ms[len(this.ms) -  1][0]
}


func (this *MaxStack) PopMax() int {
	if len(this.ms) == 0 {return -1}

	idx := this.ms[len(this.ms)-1][1]
	res := this.ms[len(this.ms)-1][0]
	this.ms = this.ms[:len(this.ms)-1]

	if idx == len(this.st) - 1 {
		this.st = this.st[:len(this.st)-1]
	} else {
		this.st = append(this.st[:idx], this.st[idx+1:]...)

		for i:= idx; i < len(this.st); i++ {
			if len(this.ms) == 0 || this.st[i] >= this.ms[len(this.ms) - 1][0] {
				this.ms = append(this.ms, []int{this.st[i], i})
			}
		}
	}
	return  res
}


/**
785. Is Graph Bipartite?
 */
func isBipartite(graph [][]int) bool {
	if graph == nil || len(graph) == 0 {return true}

	ws := make([]byte, len(graph))
	cns := make([]int, 0)

	cns = append(cns, 0)
	ws[0] = 1
	for len(cns) != 0 {
		ncs := make([]int,0)
		for _, s := range cns {
			for _, e := range graph[s] {
				switch ws[e] {
				case 0:
					if ws[s] == 1 {ws[e] = 2} else {ws[e] = 1}
					ncs = append(ncs, e)
				case 1:
					if ws[s] != 2 {return false}
				case 2:
					if ws[s] != 1 {return false}
				}
			}
		}
		cns = ncs
	}

	return true
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	h, _ :=_sortList(head)
	return h
}

func _sortList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}

	nh, nt := &ListNode{0, nil}, &ListNode{0, nil}
	for cn := head.Next; cn != nil; cn = cn.Next {
		if cn.Val >= head.Val {
			t:= &ListNode{cn.Val, nt.Next}
			nt.Next= t
		} else {
			t:= &ListNode{cn.Val, nh.Next}
			nh.Next= t
		}
	}

	var nhh, nht, nth, ntt *ListNode
	if nh.Next == nil {
		nhh = head
	} else {
		nhh, nht = _sortList(nh.Next)
		nht.Next = head
	}
	if nt.Next == nil {
		ntt = head
	} else {
		nth, ntt = _sortList(nt.Next)
		head.Next = nth
	}
	ntt.Next = nil
	return nhh, ntt
}


/**
49. Group Anagrams
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 */
func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs)==0 { return [][]string{}}

	ana := make(map[string]*[]string)

	for _, s := range strs {
		bs := sts(s)
		sort.Sort(bs)
		ss, ok := ana[string(bs)]
		if ok {
			*ss = append(*ss, s)
		} else {
			ana[string(bs)] = &[]string{s}
		}
	}

	res := make([][]string, 0)
	for _, v := range ana { res = append(res, *v) }
	return res
}

type sts []byte
func (s sts) Len() int {return len(s)}
func (s sts) Swap(i, j int) {s[i], s[j] = s[j], s[i]}
func (s sts) Less(i, j int) bool {return s[i] < s[j]}


/**
763. Partition Labels
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
 */

func partitionLabels(S string) []int {
	if len(S) == 0 {return []int{}}

	res := make([]int, 0)

	li := make([]int, 26)

	for i := len(S)-1; i >=0; i-- {
		if li[S[i] - 'a'] == 0 {li[S[i] - 'a'] = i}
	}

	s, e := 0, 0
	for s < len(S) {
		e = li[S[s] - 'a']
		for c := s; c <= e; c++ {
			if li[S[c]- 'a'] > e { e = li[S[c] - 'a'] }
		}
		res = append(res, e - s + 1)
		s = e+1
	}

	return res
}


/**
819. Most Common Word
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
 */

func mostCommonWord(p string, banned []string) string {

	wc := make(map[string]uint)
	bw := make(map[string]bool)

	for _, w := range banned {
		bw[w] = true
	}

	i := 0
	for w := nextToken(p, &i); i < len(p); w = nextToken(p, &i) {
		_, b := bw[w]
		if len(w) != 0 && !b {
			c, ok := wc[w]
			if ok {
				wc[w] = c+1
			} else {
				wc[w] = 1
			}
		}
	}

	mwc, res := uint(0), ""

	for w,c := range wc {
		if c > mwc { mwc = c; res = w}
	}

	return res
}

func nextToken(st string, i *int) string {
	res := make([]byte, 0)
	for  ; *i < len(st) && ((st[*i] >= 'a' && st[*i] <= 'z') || (st[*i] >= 'A' && st[*i] <= 'Z')); *i++ {
		if st[*i] > 'Z' {
			res = append(res, st[*i])
		} else {
			res = append(res, st[*i] - 'A' + 'a')
		}
	}

	for ; *i < len(st) && !((st[*i] >= 'a' && st[*i] <= 'z') || (st[*i] >= 'A' && st[*i] <= 'Z')); *i++ {}

	return string(res)
}

/**
946. Validate Stack Sequences
Given two sequences pushed and popped with distinct values, return true if and only if this could have
been the result of a sequence of push and pop operations on an initially empty stack.
Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
 */

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped){
		return false
	}

	st := make([]int,0)
	ppi, psi := 0, 0

	for ; ppi < len(popped) ; {
		// try pop
		if len(st) > 0 && st[len(st) - 1] == popped[ppi] {
			ppi++
			st = st[:len(st) - 1]
		} else {
			// try push
			if psi == len(pushed) {
				return false
			}
			st = append(st, pushed[psi])
			psi++
		}
	}

	return ppi == len(popped)
}

/**
809. Expressive Words
Sometimes people repeat letters to represent extra feeling, such as "hello" -> "heeellooo", "hi" -> "hiiii".
In these strings like "heeellooo", we have groups of adjacent letters that are all the same:  "h", "eee", "ll", "ooo".

For some given string S, a query word is stretchy if it can be made to be equal to S by any number of
applications of the following extension operation: choose a group consisting of characters c,
and add some number of characters c to the group so that the size of the group is 3 or more.

For example, starting with "hello", we could do an extension on the group "o" to get "hellooo",
but we cannot get "helloo" since the group "oo" has size less than 3.  Also, we could do another extension like
"ll" -> "lllll" to get "helllllooo".  If S = "helllllooo", then the query word "hello" would be stretchy
because of these two extension operations: query = "hello" -> "hellooo" -> "helllllooo" = S.

Given a list of query words, return the number of words that are stretchy.
Example:
Input:
S = "heeellooo"
words = ["hello", "hi", "helo"]
Output: 1
Explanation:
We can extend "e" and "o" in the word "hello" to get "heeellooo".
We can't extend "helo" to get "heeellooo" because the group "ll" is not size 3 or more.
 */

func expressiveWords(S string, words []string) int {

	if len(S) == 0 || len(words) == 0 {return 0}

	ss := 0

	for ss < len(S) {
		c, f := getWordFreq(S, &ss)
		nw := make([]string, 0)
		for _, w := range words {
			ok, ni := canReplaced(c, f, w)
			if ok {
				nw = append(nw, w[ni:])
			}
		}
		words = nw
	}

	wc := 0
	for _, w := range words {
		if len(w) == 0 {wc++}
	}
	return wc
}

func getWordFreq(s string, st *int) (byte, int) {
	if len(s) == 0 {return '1', 0}
	c, f := s[*st], 1
	*st++
	for ; *st < len(s) && s[*st] == c; *st++ { f++ }
	return c,f
}

func canReplaced(c byte, f int, w string) (bool, int) {
	wi := 0
	wc, wf := getWordFreq(w, &wi)
	if wf == 0 {return false, 0}

	switch {
		case wc == c:
			switch {
			case f == wf || (wf < f && f > 2):
				return true, wf
			default:
				return false, wf
			}
		default:
		return false, 0
	}
}


/**
846. Hand of Straights
Alice has a hand of cards, given as an array of integers.
Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.

Return true if and only if she can.
Example 1:
Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
Output: true
 */

func isNStraightHand0(hand []int, W int) bool {
	if len(hand) == 0 { return false }

	if len(hand) % W != 0 { return false }

	sort.Ints(hand)

	st, cc := 0, 0
	nc := hand[st]

	for st < len(hand) {
		nc = hand[st]
		for i := st; i < len(hand) && cc < W; i++ {
			switch hand[i] {
			case nc:
				nc++
				hand[i] = 0x10000000
				cc++
			case nc - 1:
			case 0x10000000:
			default:
				return false
			}
		}

		if cc == W {
			cc = 0
		} else {
			return false
		}

		p := st
		for ; p < len(hand) && hand[p] == 0x10000000; p++ {}
		st = p
	}

	if cc != 0 {
		return false
	}

	for _, h := range hand {
		if h != 0x10000000 {
			return false
		}
	}

	return true
}

func isNStraightHand(hand []int, W int) bool {
	if len(hand) == 0 || len(hand) % W != 0 { return false }

	sort.Ints(hand)

	fk := &listNode{0, nil}
	for _, v := range hand {
		n := &listNode{uint16(v), fk.Next}
		fk.Next = n
	}

	cc := W
	for fk.Next != nil {
		fk.V, cc = fk.Next.V, W
		for cur := fk; cur.Next != nil && cc > 0;  {
			if cur.Next.V == fk.V {
				cc--
				fk.V--
				cur.Next = cur.Next.Next
			} else if cur.Next.V == fk.V + 1 {
				cur = cur.Next
			} else  {
				return false
			}
		}

		if cc != 0 {
			return false
		}
	}

	return cc == 0
}

/**
1423. Maximum Points You Can Obtain from Cards
There are several cards arranged in a row, and each card has an associated number of points.
The points are given in the integer array cardPoints.
In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
Your score is the sum of the points of the cards you have taken.

Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
Example 1:
Input: cardPoints = [1,2,3,4,5,6,1], k = 3
Output: 12
 */

func maxScore0(cardPoints []int, k int) int {
	if k == 0 || len(cardPoints) == 0 { return 0 }
	if k == 1 {
		switch {
		case cardPoints[0] > cardPoints[len(cardPoints) - 1]:
			return cardPoints[0]
		default:
			return cardPoints[len(cardPoints) - 1]
		}
	}

	if k >= len(cardPoints) {
		res := 0
		for _, v := range cardPoints {
			res += v
		}
		return res
	}

	rm := maxScore(cardPoints[1:], k - 1)
	lm := maxScore(cardPoints[:len(cardPoints) - 1], k - 1)

	switch {
	case cardPoints[0] + rm > cardPoints[len(cardPoints) - 1] + lm:
		return cardPoints[0] + rm
	default:
		return cardPoints[len(cardPoints) - 1] + lm
	}
}

func maxScore(cp []int, k int) int {
	if k == 0 || len(cp) == 0 { return 0 }
	if k >= len(cp) {
		res := 0
		for _, v := range cp {
			res += v
		}
		return res
	}

	ls := 0
	for c, i := 0, 0; c < k && i < len(cp); i++ {
		ls += cp[i]
		c++
	}

	max, rs := ls, 0

	for c, i := 0, len(cp) - 1; c < k && i >= 0; {
		rs += cp[i];  ls -= cp[k - c - 1];
		c++; i--

		if rs + ls > max {
			max = rs + ls
		}
	}

	return max
}
