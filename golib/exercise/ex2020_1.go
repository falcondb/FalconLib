package exercise

import (
	"container/heap"
	"errors"
	"sort"
	"strings"
)

func MergeTwo(arr1 []int, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	if arr1 == nil {
		return arr2
	}
	if arr2 == nil {
		return arr1
	}

	m := make([]int, 0)
	i1, i2 := 0, 0
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] < arr2[i2] {
			m = append(m, arr1[i1])
			i1++
		} else {
			m = append(m, arr2[i2])
			i2++
		}
	}

	if i1 < len(arr1) {

		m = append(m, arr1[i1:len(arr1)]...)
	} else {
		m = append(m, arr2[i2:len(arr2)]...)
	}
	return m
}

func ReverseWords(ws string) string {
	st, cur := 0, 0
	var res []byte
	for cur < len(ws) {
		if ws[cur] != ' ' {
			cur++
		} else if st != cur {
			res = append(res, ws[st:cur]...)
			res = append(res, ' ')
			cur++
			st = cur
		} else {
			cur++
			st = cur
		}
	}

	if cur != st {
		res = append(res, ws[st:cur]...)
	}
	return string(res)
}

func MagicHash(key string, size int) int {
	mb := 33
	sum := 100
	for _, c := range key {
		sum += int(c) * mb
		mb *= 33
	}
	return int(sum / size)
}

type listNode struct {
	V    uint16
	Next *listNode
}

func lastNList(head *listNode, p int) (uint16, bool) {
	if head == nil || p == 0 {
		return 0, false
	}

	cur, fol := head, head

	var i int
	for i = 0; i < p && cur != nil; i, cur = i+1, cur.Next {
	}

	if i < p {
		return 0, false
	}

	for ; cur != nil; cur, fol = cur.Next, fol.Next {
	}

	return fol.V, true
}

/*
assumption: data is ordered in descending order
*/
func binarySearch(data []int, k int) (int, bool) {
	if data == nil || len(data) == 0 {
		return -1, false
	}

	l, r, m := 0, len(data)-1, 0
	for l <= r {
		m = l + (r-l)/2

		if data[m] == k {
			return m, true
		}
		if data[m] < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return m, false
}

func closest(ds1 []int, ds2 []int) (int, error) {
	if ds1 == nil || ds2 == nil || len(ds1) == 0 || len(ds2) == 0 {
		return 0, errors.New("invalid input")
	}

	var md int = 1<<31 - 1

	if len(ds1) < len(ds2) {
		ds1, ds2 = ds2, ds1
	}

	sort.Ints(ds2)

	for _, v := range ds1 {

		switch {
		case v < ds2[0] && md > ds2[0]-v:
			md = ds2[0] - v
		case v > ds2[len(ds2)-1] && md > v-ds2[len(ds2)-1]:
			md = ds2[len(ds2)-1]
		default:
			t, f := binarySearch(ds2, v)
			if f {
				return 0, nil
			}
			switch {
			case v-ds2[t-1] < md:
				md = v - ds2[t-1]
			case ds2[t]-v < md:
				md = ds2[t] - v
			}
		}
	}
	return md, nil
}

func calcString2(fm []byte) int {
	if fm == nil || len(fm) == 0 {
		return 0
	}

	if len(fm) == 1 {
		return int(fm[0])
	}

	var adds []int
	mt := 1
	m := false
	for p := 1; p < len(fm); p += 2 {
		switch {
		case fm[p] == '+':
			if m {
				adds = append(adds, int(fm[p-1])*mt)
			} else {
				adds = append(adds, int(fm[p-1]))
			}
			mt = 1
		case fm[p] == '*':
			mt *= int(fm[p-1])
			m = true
		}
	}

	if fm[len(fm)-2] == '+' {
		adds = append(adds, int(fm[len(fm)-1]))
	} else {
		adds = append(adds, mt*int(fm[len(fm)-1]))
	}

	var res int
	for _, v := range adds {
		res += v
	}
	return res
}


func calcString(fm []byte) int {
	if fm == nil || len(fm) == 0 {
		return 0
	}

	if len(fm) == 1 {
		return int(fm[0])
	}

	stk := 0

	v, l, f := getNum(fm)

	if f {
		stk = v
		fm = fm[l:]
	}

	res := 0
	for add, f := getSign(fm); f; add, f = getSign(fm) {
		fm = fm[1:]
		v, l , f = getNum(fm)
		if add && f {
				res += stk
				stk = v
				fm = fm[l:]
		} else if f {
			stk = stk * v
			fm = fm[l:]
		} else {
			break
		}

	}

	return res + stk
}

func getSign(st []byte) (bool, bool) {
	if st == nil || len(st) == 0 {
		return true, false
	}

	switch  st[0] {
	case '+':
		return true, true
	case '*':
		return false, true
	default:
		return true, false
	}

}

func getNum(st []byte) (int, int, bool) {
	if st == nil || len(st) == 0 || st[0] < '0' || st[0] > '9' {
		return 0xffff, 0, false
	}

	value := 0
	l := 0
	for _, v := range st {
		if v < '0' || v > '9' {
			break
		} else {
			l++
			value = value * 10 + int(v - '0')
		}
	}

	return value, l, true
}


type dLinkedNode struct {
	K, V       int
	Prev, Next *dLinkedNode
}

/*
LRU Cache
*/
type LRUCache struct {
	Cap   int
	Head  *dLinkedNode
	Tail  *dLinkedNode
	index map[int]*dLinkedNode
}

func (cache *LRUCache) Init(c int) {
	cache.Cap = c
	cache.Head = &dLinkedNode{-1, -1, nil, nil}
	cache.Tail = &dLinkedNode{-1, -1, nil, nil}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	cache.index = make(map[int]*dLinkedNode)
}

func (cache *LRUCache) Get(k int) (int, bool) {
	v, ok := cache.index[k]
	if !ok {
		return -1, false
	}

	v.Prev.Next = v.Next
	v.Next.Prev = v.Prev

	cache.moveToTail(v)

	return v.V, true
}

func (cache *LRUCache) moveToTail(v *dLinkedNode) {
	v.Prev = cache.Tail.Prev
	cache.Tail.Prev.Next = v
	cache.Tail.Prev = v
	v.Next = cache.Tail
}

func (cache *LRUCache) set(k int, v int) {
	n, ok := cache.index[k]
	if ok {
		n.V = v
	}

	if len(cache.index) == cache.Cap {
		delete(cache.index, cache.Head.Next.K)
		cache.Head = cache.Head.Next
	}

	nn := &dLinkedNode{k, v, nil, nil}
	cache.index[k] = nn
	cache.moveToTail(nn)
}

type wordMatrixSearch struct {
	m  [][]byte
	sd [][]bool
	w  []byte
}

func (wm *wordMatrixSearch) mWordSearch(p, x, y int) bool {
	if len(wm.w) == p {
		return true
	}

	if !wm.sd[y][x] && wm.m[y][x] == wm.w[p] {
		wm.sd[y][x] = true
		if x > 0 && wm.mWordSearch(p+1, x-1, y) {
			return true
		}
		if y > 0 && wm.mWordSearch(p+1, x, y-1) {
			return true
		}
		if x < len(wm.m[0])-1 && wm.mWordSearch(p+1, x+1, y) {
			return true
		}
		if y < len(wm.m)-1 && wm.mWordSearch(p+1, x, y+1) {
			return true
		}
		wm.sd[y][x] = false
		return false
	} else {
		return false
	}
}

type WordBreak struct {
	dict map[string]bool
}

func (wb *WordBreak) CanWordBreak(t string) bool {
	if len(t) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		_, ok := wb.dict[t[:i+1]]
		if ok {
			if wb.CanWordBreak(t[i+1:]) {
				return true
			}
		}
	}
	return false
}

type wordCount struct {
	w string
	c int
}

type wcHeap []*wordCount

func (wc wcHeap) Len() int      { return len(wc) }
func (wc wcHeap) Swap(i, j int) { wc[i], wc[j] = wc[j], wc[i] }
func (wc wcHeap) Less(i, j int) bool {
	return wc[i].c > wc[j].c || wc[i].c == wc[j].c && wc[i].w < wc[j].w
}

func (wc *wcHeap) Push(x interface{}) {
	item := x.(*wordCount)
	*wc = append(*wc, item)
}

func (wc *wcHeap) Pop() interface{} {
	old := *wc
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*wc = old[0 : n-1]
	return item
}

func TopKFreWords(wds []string, k int) ([]string, error) {
	if wds == nil || len(wds) < k {
		return nil, errors.New("invalid input")
	}

	wfc := make(map[string]int)

	for _, w := range wds {
		c, ok := wfc[w]
		if !ok {
			wfc[w] = 1
		} else {
			wfc[w] = c + 1
		}
	}

	var res []string
	if len(wfc) <= k {
		for v, _ := range wfc {
			res = append(res, v)
		}
	} else {
		h := make(wcHeap, 0)
		heap.Init(&h)
		cur := 0
		for w, c := range wfc {
			heap.Push(&h, wordCount{w, c})
			cur++
			if cur > k {
				heap.Pop(&h)
			}
		}

		for i := 0; i < k; i++ {
			res = append(res, heap.Pop(&h).(*wordCount).w)
		}
	}
	return res, nil
}

func removeArbitrarySpace(s string) string {
	space := false
	var res []byte
	for _, c := range []byte(s) {
		switch c {
		case ' ':
			if !space {
				res = append(res, c)
				space = true
			}
		default:
			res = append(res, c)
			space = false
		}
	}
	return strings.TrimSpace(string(res))
}

func minPartition(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	min := 1<<31 - 1
	ca := make([]int, 1<<uint16(len(nums)))
	ca[0], ca[1] = 0, nums[0]
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 1<<uint16(i); j++ {
			ca[j+1<<uint16(i)] = nums[i] + ca[j]
		}
	}

	for i := 0; i < len(ca); i++ {
		switch {
		case sum-ca[i] == ca[i]:
			return 0
		case ca[i] > sum-ca[i] && min > (ca[i]+ca[i]-sum):
			min = ca[i] + ca[i] - sum
		case ca[i] < sum-ca[i] && min > (sum-ca[i]-ca[i]):
			min = sum - ca[i] - ca[i]
		default:
			// no op to update min
		}
	}
	return min
}

type interval struct {
	st, end int
}

type intervals []interval

func (is intervals) Len() int {
	return len(is)
}

func (is intervals) Less(i, j int) bool {
	switch {
	case is[i].st < is[j].st:
		return true
	case is[i].st == is[j].st:
		return is[i].end < is[j].end
	default:
		return false
	}
}

func (is intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type intervalheap []interval

func (h intervalheap) Len() int {
	return len(h)
}

func (h intervalheap) Less(i, j int) bool {
	return h[i].end < h[j].end
}

func (h intervalheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intervalheap) Push(x interface{}) {
	*h = append(*h, x.(interval))
}

func (h *intervalheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeIntervalArrays(a1 []interval, a2 []interval) []interval {
	if a1 == nil || len(a1) == 0 {
		return a2
	}
	if a2 == nil || len(a2) == 0 {
		return a1
	}

	var m []interval
	cur := interval{0, 0}

	var i, j int

	for i < len(a1) && j < len(a2) {
		switch {
		//TODO
		}
	}

	m = append(m, cur)
	return m
}

func mergeInterval(l *interval, r *interval) (*interval, *interval) {
	switch {
	case l == nil:
		return &interval{r.st, r.end}, nil
	case r == nil:
		return &interval{l.st, l.end}, nil
	case l.st == r.st && l.end > r.end:
		return &interval{l.st, l.end}, nil
	case l.st == r.st:
		return &interval{r.st, r.end}, nil
	case l.st < r.st && l.end < r.st:
		return &interval{l.st, l.end}, &interval{r.st, r.end}
	case r.st < l.st && r.end < l.st:
		return &interval{r.st, r.end}, &interval{l.st, l.end}
	case l.end == r.st:
		return &interval{l.st, r.end}, nil
	case r.end == r.st:
		return &interval{r.end, l.end}, nil
		// TODO: handle all the possible cases
	default:
		return nil, nil
	}
}


func ppids(pid []int, ppid []int, k int) []int {
	if pid == nil || ppid == nil || len(pid) != len(ppid) {
		return nil
	}
	if len(pid) == 1 {
		return pid
	}

 	r := make(map[int]int, len(pid))

	for i,p := range ppid {
		r[p] = pid[i]
	}

	p := k
	var res []int

	for p != 0 {
		res = append(res, p)
		pp, ok := r[p]
		if ok {
			p = pp

		} else {
			p = 0
		}
	}

	return res
}

func interleavingString(sa, sb, sm string) bool {
	switch {
	case len(sa) == 0 :
		return sb == sm
	case len(sb) == 0:
		return sa == sm
	case len(sm) == 0:
		return false
	case len(sa) + len(sb) != len(sm):
		return false
	}

	if sb[len(sb) - 1] == sm[len(sm) - 1]  && interleavingString(sa, sb[:len(sb) - 1], sm[:len(sm) - 1]) {
		return true
	}

	if sa[len(sa) - 1] == sm[len(sm) - 1]  && interleavingString(sa[:len(sa) - 1], sb, sm[:len(sm) - 1]) {
		return true
	}

	return false
}

type listOrInt struct {
	l []listOrInt
	v int
}


func flatListOrInt(li *listOrInt) []int {
	res :=  make([]int, 0)

	if li == nil {
		return res
	}

	switch {
		case li.l == nil :
			res = append(res, li.v)
	default:
		for _, v := range li.l {
			res = append(res, flatListOrInt(&v)...)
		}
	}

	return res
}