package exercise

import (
	"container/heap"
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
	for h = head; h.Next != nil && h.Next.V&0x1 == 1; h = h.Next {
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

func nrDiffTriangles(a tas) uint {
	switch {
	case a == nil || len(a) == 0:
		return 0
	case len(a) == 1:
		return 1
	}

	for i := 0; i < len(a); i++ {
		a[i].edgeSort()
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
				bd[c.y][c.x - 1] = ct
			}
			if c.y > 0 && bd[c.y-1][c.x] == SIGNROOM {
				ne = append(ne, &coor{c.x, c.y - 1})
				bd[c.y - 1][c.x] = ct
			}
			if c.x+1 < len(bd[0]) && bd[c.y][c.x+1] == SIGNROOM {
				ne = append(ne, &coor{c.x + 1, c.y})
				bd[c.y][c.x + 1] = ct
			}
			if c.y+1 < len(bd) && bd[c.y+1][c.x] == SIGNROOM {
				ne = append(ne, &coor{c.x, c.y + 1})
				bd[c.y + 1][c.x] = ct
			}
		}
		ce = ne
	}
}


//Counting Bits
func cout1s(m uint) []uint {
	res := make([]uint,0, m+1)

	for i := uint(0); i <= m; i++ {
		n := i
		ct := uint(0)
		for n != 0 {
			ct++
			n = n & (n-1)
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

	if s[0] == s[len(s) - 1] {
		max = 2 + maxPali(s[1:len(s)-1])
	} else {
		sr := maxPali(s[1:])
		if max < sr {
			max = sr
		}

		sr = maxPali(s[:len(s) - 1])
		if max < sr {
			max = sr
		}
		sr = maxPali(s[1:len(s) - 1])
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
	for _, c := range ss[len(ss) - 1]{
		switch c {
		case '0':
			lzs++
		case '1':
			los++
		}
	}

	tl := 1 + onesAndZeors(ss[:len(ss)-1], os - los, zs - lzs)
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
		if mc != -1 && min > mc + 1 {
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

	for prog := true ; prog; prog = false {
		for _, c := range cs {
			for p := 1; p < t+1 && buf[p] != 0; p++ {
				if p + c < t + 1 && buf[p+c] > buf[p] {
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