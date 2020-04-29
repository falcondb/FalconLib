package exercise

import "sort"

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
    return []byte {}
  }

  pre, reps, h, t := dsSubproblem(inp)

  if reps != 0 {
    pre = append(pre, generateReps(reps, inp[h+1:t])...)
    pre = append(pre, decodeString(inp[t+1:])...)
  }

  return pre
}


func dsSubproblem(p []byte) ([]byte, int, int, int) {
  i,h,t := 0, -1, len(p)
  reps := 0
  pre := make([]byte, 0)

  for ; i< len(p) && p[i] >= 'a' && p[i] <= 'z'; i++ {
    pre = append(pre, p[i])
  }

  for ; i < len(p) && p[i] >= '0' && p[i] <= '9'; i++ {
    reps = reps * 10 + int(p[i]) - int('0')
  }

  for h = i; h < len(p) && p[h] != '['; h++ { }
  for t = len(p) - 1; t >=0 && p[t] != ']'; t-- { }

  return pre, reps, h, t
}

func generateReps(reps int, substr []byte) []byte {
  res := make([]byte, 0)
  for ; reps >0; reps-- {res = append(res, substr...)}
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
			case k - c == c:
				ct += wc[c] * (wc[c] - 1 ) >> 1
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
	v uint16
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

func allPossibleSubsets(a []uint16)[][]uint16 {
	if a == nil {
		return nil
	}

	res:= make([][]uint16, 1<<uint(len(a)))
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
	V uint16
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
		cn := roots[len(roots) - 1]
		roots = roots[:len(roots) - 1]
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
	if a == nil  || len(a) < 3 {
		return
	}

	pc := 0
	for _, v := range a {
		if v > 0 {
			pc++
		}
	}

	np := pc >= len(a) & 01

	for c, pb := 0, 0; pb < len(a); {
		switch np {
		case true:
			switch {
			case a[pb] > 0:
				a[c], a[pb] = a[pb], a[c]
				c+=2
				pb=c
			default:
				pb++
			}
		default:
			switch {
			case a[pb] < 0:
				a[c], a[pb] = a[pb], a[c]
				c+=2
				pb=c
			default:
				pb++
			}
		}
	}
}