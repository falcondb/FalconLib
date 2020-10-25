package exercise

import (
	"fmt"
	"testing"
)

func TestAreTheyEqual(t *testing.T) {
	testAreTheyEqual([]int{1, 2, 3, 4}, []int{1, 4, 3, 2}, true, t)
	testAreTheyEqual([]int{1, 2, 3, 1}, []int{1, 4, 3, 2}, false, t)

	testAreTheyEqual([]int{1, 2, 3, 1}, []int{1, 2, 3, 1}, true, t)
	testAreTheyEqual([]int{1, 2, 3, 1}, []int{1, 3, 2, 1}, true, t)
	testAreTheyEqual([]int{10, 2, 3, 1}, []int{1, 3, 2, 1}, false, t)
	testAreTheyEqual([]int{10, 2, 3, 1}, []int{1, 3, 2, 5}, false, t)

	testAreTheyEqual([]int{1, 2, 3, 4}, []int{4, 3, 2, 1}, true, t)

	testAreTheyEqual([]int{1, 2, 3, 4}, []int{3, 2, 1, 4}, true, t)

}

func testAreTheyEqual(a, b []int, exp bool, t *testing.T) {
	if exp != areTheyEqual(a, b) {
		fmt.Printf("Input:%v\t %v\t Exp:%v\n", a, b, exp)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestCountSubarrays(t *testing.T) {
	testCountSubarrays([]int{3, 4, 1, 6, 2}, []uint16{1, 3, 1, 5, 1}, t)
	testCountSubarrays([]int{10, 6, 9, 0, 5, 6}, []uint16{6, 1, 5, 1, 2, 3}, t)
}

func testCountSubarrays(a []int, exp []uint16, t *testing.T) {
	res := countSubarrays(a)

	if !assertSlicesEqual(res, exp) {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func assertSlicesEqual(res, exp []uint16) bool {
	for i, v := range res {
		if v != exp[i] {
			return false
		}
	}
	return true
}

func assertIntSlicesEqual(res, exp []int) bool {
	for i, v := range res {
		if v != exp[i] {
			return false
		}
	}
	return true
}

func TestGetMilestoneDays(t *testing.T) {
	testGetMilestoneDays([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{100, 200, 500}, []int{4, 6, 10}, t)
}

func testGetMilestoneDays(re, ms, exp []int, t *testing.T) {
	ok, res := getMilestoneDays(re, ms)
	if !ok {
		t.Errorf("Test name is %s, ", t.Name())
	}

	if !arraysEqual(res, exp) {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func arraysEqual(a, b []int) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestReverseList(t *testing.T) {

	head := listNode{1, &listNode{2, &listNode{3, &listNode{4, nil}}}}

	rev, pre := reverseList(&head)

	for n, pre := rev, uint16(1<<uint16(16)-1); n != nil; n = n.Next {
		if n.V >= pre {
			t.Errorf("Test name is %s, ", t.Name())
		}
		pre = n.V
	}

	if pre == nil || pre.Next != nil || pre.V != 1 {
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestReverseEven(t *testing.T) {
	testReverseEven(&listNode{1, nil},
		&listNode{1, nil},
		t)

	testReverseEven(&listNode{1, &listNode{2, &listNode{8, nil}}},
		&listNode{1, &listNode{8, &listNode{2, nil}}},
		t)

	testReverseEven(&listNode{1, &listNode{2, &listNode{8, &listNode{0, &listNode{5, nil}}}}},
		&listNode{1, &listNode{0, &listNode{8, &listNode{2, &listNode{5, nil}}}}},
		t)

	testReverseEven(&listNode{1, &listNode{2, &listNode{8, &listNode{9, &listNode{12, &listNode{16, nil}}}}}},
		&listNode{1, &listNode{8, &listNode{2, &listNode{9, &listNode{16, &listNode{12, nil}}}}}},
		t)
}

func testReverseEven(h *listNode, exp *listNode, t *testing.T) {
	res := reverseEven(h)

	var rn, en *listNode
	for rn, en = res, exp; rn != nil && en != nil; rn, en = rn.Next, en.Next {
		if en.V != rn.V {
			t.Errorf("Test name is %s, ", t.Name())
		}
	}

	if rn != nil || en != nil {
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestNrDiffTriangles(t *testing.T) {
	testNrDiffTriangles([]triangle{triangle{2, 2, 3}, triangle{3, 2, 2}, triangle{2, 5, 6}}, 2, t)
	testNrDiffTriangles([]triangle{triangle{8, 4, 6}, triangle{100, 101, 102}, triangle{84, 93, 173}}, 3, t)
	testNrDiffTriangles([]triangle{triangle{5, 8, 9}, triangle{5, 9, 8}, triangle{9, 5, 8},
		triangle{9, 8, 5}, triangle{8, 9, 5}, triangle{8, 5, 9}}, 1, t)
}

func testNrDiffTriangles(a tas, exp uint, t *testing.T) {
	res := nrDiffTriangles(a)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestNumberOfWays(t *testing.T) {
	TestnumberOfWays([]int{1, 5, 3, 3, 3}, 6, 4, t)
	TestnumberOfWays([]int{1, 2, 3, 4, 3}, 6, 2, t)
}

func TestnumberOfWays(a []int, k int, exp uint, t *testing.T) {
	res := numberOfWays(a, k)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestBstSearchRange(t *testing.T) {
	// case 1
	r := bstNode{5, nil, nil}
	testBstSearchRange(&r, 6, 10, []uint16{}, t)

	// case 2
	r = bstNode{20, &bstNode{8, &bstNode{4, nil, nil}, &bstNode{12, nil, nil}}, &bstNode{22, nil, nil}}
	testBstSearchRange(&r, 10, 22, []uint16{12, 20, 22}, t)
}

func testBstSearchRange(root *bstNode, min, max uint16, exp []uint16, t *testing.T) {
	res := bstSearchRange(root, min, max)

	if !assertSlicesEqual(res, exp) {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestAllPossibleSubsets(t *testing.T) {

	res := allPossibleSubsets([]uint16{1, 2, 3})

	if res == nil || len(res) == 0 {
		t.Errorf("Test name is %s, ", t.Name())
	}

	fmt.Printf("Res: %v\t\n", res)
}

func TestTopoSorting(t *testing.T) {
	n0, n1, n2, n3, n4, n5, n6, n7 := graphNode{0, nil}, graphNode{1, nil}, graphNode{2, nil},
		graphNode{3, nil}, graphNode{4, nil}, graphNode{5, nil}, graphNode{6, nil}, graphNode{7, nil}
	n0.Links = []*graphNode{&n1, &n2, &n3}
	n1.Links = []*graphNode{&n4, &n6}
	n2.Links = []*graphNode{&n4, &n5, &n7}
	n3.Links = []*graphNode{&n4, &n5, &n6, &n7}
	n4.Links = []*graphNode{&n6, &n7}
	n5.Links = []*graphNode{&n7}
	g := graph{[]*graphNode{&n0, &n1, &n2, &n3, &n4, &n5, &n6, &n7}}

	_, res := g.topoSorting()

	for _, n := range res {
		fmt.Printf("%v, ", n.V)
	}
}

func TestPosNegNumss(t *testing.T) {
	testPosNegNumss([]int{-1, -2, -3, 4, 5, 6}, t)
	testPosNegNumss([]int{-1, 2, -3, -4, 5, 6}, t)
	testPosNegNumss([]int{-1, 2, 3, -4, 5, -6}, t)
	testPosNegNumss([]int{1, 2, 3, -4, -5, -6}, t)
	testPosNegNumss([]int{1, 2, -3, 4, -5, -6}, t)
}

func testPosNegNumss(a []int, t *testing.T) {
	posNegNums(a)

	expSign := a[0] > 0
	for _, v := range a {
		if expSign && v < 0 || !expSign && v > 0 {
			fmt.Printf("Res: %v\t\n", a)
			t.Errorf("Test name is %s, ", t.Name())
			return
		} else {
			expSign = !expSign
		}
	}
}

func TestJumpGame(t *testing.T) {
	testJumpGame([]int{2, 3, 1, 1, 4}, true, t)
	testJumpGame([]int{3, 2, 1, 0, 4}, false, t)

}

func testJumpGame(a []int, exp bool, t *testing.T) {
	res := jumpGame(a)
	if res != exp {
		fmt.Printf("input: %v\t\n", a)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestLongestCS(t *testing.T) {
	testLongestCS([]uint{100, 4, 200, 1, 3, 2}, 4, t)
	testLongestCS([]uint{100, 105, 101, 1, 3, 2, 102, 104, 103}, 6, t)

}

func testLongestCS(a []uint, exp uint, t *testing.T) {
	res := longestCS(a)
	if res != exp {
		fmt.Printf("input: %v\t Res: %v\n", a, res)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestRehashing(t *testing.T) {
	res := rehashing([][]int{nil, []int{21, 9}, []int{14}, nil})
	fmt.Printf("Res: %v\n", res)
}

func TestHeapify(t *testing.T) {
	h := Heap{[]int{3, 2, 1, 4, 5}}
	testHeapify(&h, t)
	h = Heap{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}}
	testHeapify(&h, t)
}

func testHeapify(h *Heap, t *testing.T) {
	h.heapify()
	for i, v := range h.data {
		if h.data[i>>1] > v {
			fmt.Printf("Res: %v\n", h.data)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}

func TestCombinationSum(t *testing.T) {
	testCombinationSum([]int{2, 3}, 3, true, t)
	testCombinationSum([]int{2, 3, 6, 7}, 7, true, t)
	testCombinationSum([]int{2, 3, 6, 7}, 5, true, t)
	testCombinationSum([]int{2, 3, 6, 7}, 4, false, t)
}

func testCombinationSum(a []int, sum int, isOK bool, t *testing.T) {
	res, ok := CombinationSum(a, sum)
	if ok != isOK {
		fmt.Printf("input: %v\t Res: %v\n", a, res)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
	fmt.Printf("Res: %v\n", res)
}

func TestIsPower2(t *testing.T) {
	fmt.Printf("input: %v\t Res: %v\n", 2, isPower2(2))
	fmt.Printf("input: %v\t Res: %v\n", 7, isPower2(7))
}

func TestCWMoreWatter(t *testing.T) {
	testCWMoreWatter([]uint{1, 3, 2}, 2, t)
	testCWMoreWatter([]uint{1, 3, 2, 2}, 4, t)
	testCWMoreWatter([]uint{1, 5, 2, 2, 1, 4}, 16, t)
	testCWMoreWatter([]uint{1, 5, 2, 2, 1, 5}, 20, t)
}

func testCWMoreWatter(a []uint, exp uint, t *testing.T) {
	res := CWMoreWatter(a)
	if res != exp {
		fmt.Printf("input: %v\t Res: %v\n", a, res)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMinSubarrSum(t *testing.T) {
	testMinSubarrSum([]int{2, 3, 1, 2, 4, 3}, 7, 2, true, t)
	testMinSubarrSum([]int{1, 2, 3, 4, 5}, 100, 0, false, t)

	testMinSubarrSum([]int{2, 3, 1, 2, 4, 3}, 2, 1, true, t)
	testMinSubarrSum([]int{2, 30, 10, 20, 4, 3}, 65, 5, true, t)

}

func testMinSubarrSum(a []int, s, exp int, isOK bool, t *testing.T) {
	res, ok := MinSubarrSum(a, s)
	if ok != isOK || res != exp {
		fmt.Printf("input: %v\t Res: %v\n", a, res)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestKClosest(t *testing.T) {
	testKClosest([]int{1, 2, 3}, []int{2, 1, 3}, 2, 3, t)
	testKClosest([]int{1, 4, 6, 8}, []int{4, 1, 6}, 3, 3, t)
	testKClosest([]int{1, 4, 6, 8}, []int{4, 1, 6}, 3, 5, t)

	testKClosest([]int{1, 4, 6, 8}, []int{1, 4, 6, 8}, 0, 5, t)

}

func testKClosest(a, exp []int, r, k int, t *testing.T) {
	res := KClosest(a, r, k)

	for i, v := range res {
		if v != exp[i] {
			fmt.Printf("input: %v\t Res: %v\n", a, res)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}

func TestKFreqWords(t *testing.T) {
	testKFreqWords([]string{"yes", "lint", "code", "yes", "code", "baby", "you", "baby", "chrome",
		"safari", "lint", "code", "body", "lint", "code"}, []string{"code", "lint", "baby"}, 3, t)

	testKFreqWords([]string{"yes", "lint", "code", "yes", "code", "baby", "you", "baby", "chrome",
		"safari", "lint", "code", "body", "lint", "code"}, []string{"code", "lint"}, 2, t)

	testKFreqWords([]string{"yes", "lint", "code", "yes", "code", "baby", "you", "baby", "chrome",
		"safari", "lint", "code", "body", "lint", "code"}, []string{"code", "lint", "baby", "yes"}, 4, t)

	testKFreqWords([]string{"yes", "code", "code", "yes"}, []string{"code", "yes"}, 4, t)

	testKFreqWords([]string{"yes", "code", "code", "yes", "yes"}, []string{"yes", "code"}, 4, t)

}

func testKFreqWords(a, exp []string, k int, t *testing.T) {
	res := KFreqWords(a, k)

	for i, v := range res {
		if exp[i] != v {
			fmt.Printf("input: %v\t Res: %v\n", a, res)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}

func TestCombSum(t *testing.T) {
	testCombSum([]uint{1, 2, 4}, 4, 6, t)
	testCombSum([]uint{1, 2}, 4, 5, t)
}

func testCombSum(a []uint, tg, exp uint, t *testing.T) {
	res := CombSum(a, tg)
	if exp != res {
		fmt.Printf("input: %v\t Res: %v\n", a, res)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMissNumber(t *testing.T) {
	testMissNumber([]byte("19201234567891011121314151618"), 20, 1, 7, t)
	testMissNumber([]byte("2119201734567891011121314151618"), 21, 1, 2, t)
	testMissNumber([]byte("56412"), 6, 3, 7, t)
}

func testMissNumber(a []byte, r uint, exp1, exp2 uint, t *testing.T) {
	m1, m2, _ := MissNumber(a, r)
	if m1 != exp1 && m2 != exp1 || m1 != exp2 && m2 != exp2 {
		fmt.Printf("input: %v\t Res: %v\t%v\n", a, m1, m2)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestPartSubset(t *testing.T) {
	testPartSubset([]uint{1, 5, 11, 5}, true, t)
	testPartSubset([]uint{1, 2, 3, 9}, false, t)
}

func testPartSubset(a []uint, exp bool, t *testing.T) {
	res := PartSubset(a)
	if res != exp {
		fmt.Printf("input: %v\t\n", a)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestBTLCS(t *testing.T) {
	testBTLCS(&bstNode{1, &bstNode{2, &bstNode{3, nil, nil}, nil}, &bstNode{0, nil, nil}}, 4, t)
	testBTLCS(&bstNode{3, &bstNode{2, nil, nil}, &bstNode{2, nil, nil}}, 2, t)
}

func testBTLCS(r *bstNode, exp uint, t *testing.T) {
	res := BTLCS(r)
	if res != exp {
		fmt.Printf("input: %v\t\n", *r)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMaxAvgSubarr(t *testing.T) {
	testMaxAvgSubarr([]int{1, 12, -5, -6, 50, 3}, 3, 47, 3, t)
	testMaxAvgSubarr([]int{5}, 1, 5, 1, t)

	testMaxAvgSubarr([]int{1, 60, 50, 50, 50, 3}, 3, 160, 3, t)
	testMaxAvgSubarr([]int{1, 60, 50, 50, 50, 3}, 4, 210, 4, t)
	testMaxAvgSubarr([]int{1, 60, 50, 50, 50, 3}, 2, 110, 2, t)
	testMaxAvgSubarr([]int{1, 60, -100, 100, -50, 3}, 2, 61, 2, t)
}

func testMaxAvgSubarr(ns []int, l int, ems, eml int, t *testing.T) {
	ms, ml := MaxAvgSubarr(ns, l)
	if ms != ems || ml != eml {
		fmt.Printf("input: %v\t%v\t%v\t\n", ns, ms, ml)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMinSwap2Sorted(t *testing.T) {
	testMinSwap2Sorted([]uint{7, 1, 3, 2, 4, 5, 6}, 5, t)
	testMinSwap2Sorted([]uint{2, 3, 4, 1, 5}, 3, t)
}

func testMinSwap2Sorted(ns []uint, e uint, t *testing.T) {
	r := minSwap2Sorted(ns)
	if r != e {
		fmt.Printf("input: %v\t%v\t\n", ns, r)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMaxOfUniq(t *testing.T) {
	testMaxOfUniq([]int{1, 2, 4, 4}, 4, t)
	testMaxOfUniq([]int{1, 1, 1, 1}, 3, t)
	testMaxOfUniq([]int{1, 1, 1, 4}, 4, t)
	testMaxOfUniq([]int{1, 1, 1, 3}, 4, t)
	testMaxOfUniq([]int{1, 1, 1, 1, 2}, 4, t)
	testMaxOfUniq([]int{1, 1, 1, 1, 1, 1, 1, 2}, 4, t)

	testMaxOfUniq([]int{7, 1, 3, 2, 4, 5, 6}, 7, t)
	testMaxOfUniq([]int{2, 2, 3, 3, 4}, 5, t)
}

func testMaxOfUniq(ns []int, e uint, t *testing.T) {
	r := maxOfUniq(ns)
	if r != e {
		fmt.Printf("input: %v\t%v\t\n", ns, r)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestBTVertOrder(t *testing.T) {
	//[[9],[3,15],[20],[7]]
	//3
	///\
	///  \
	//9  20
	//   /\
	//  /  \
	// 15   7
	testBTVertOrder(&bstNode{3, &bstNode{9, nil, nil}, &bstNode{20, &bstNode{15, nil, nil}, &bstNode{7, nil, nil}}},
		[][]uint16{[]uint16{9}, []uint16{3, 15}, []uint16{20}, []uint16{7}}, t)

	//Output: [[4],[9],[3,0,1],[8],[7]]
	// 	   3
	//	  /\
	//	9   8
	// /  \/  \
	// 4  01   7
	testBTVertOrder(&bstNode{3, &bstNode{9, &bstNode{4, nil, nil}, &bstNode{0, nil, nil}}, &bstNode{8, &bstNode{1, nil, nil}, &bstNode{7, nil, nil}}},
		[][]uint16{[]uint16{4}, []uint16{9}, []uint16{3, 0, 1}, []uint16{8}, []uint16{7}}, t)
}

func testBTVertOrder(rt *bstNode, e [][]uint16, t *testing.T) {
	r := BTVertOrder(rt)

	for i, _ := range r {
		for j, _ := range r[i] {
			if r[i][j] != e[i][j] {
				fmt.Printf("input: %v\t\n", r)
				t.Errorf("Test name is %s, ", t.Name())
				return
			}
		}
	}
}

func TestWallsGates(t *testing.T) {
	testWallsGates([][]int{[]int{SIGNROOM, -1, 0, SIGNROOM}, []int{SIGNROOM, SIGNROOM, SIGNROOM, -1}, []int{SIGNROOM, -1, SIGNROOM, -1}, []int{0, -1, SIGNROOM, SIGNROOM}}, [][]int{[]int{3, -1, 0, 1}, []int{2, 2, 1, -1}, []int{1, -1, 2, -1}, []int{0, -1, 3, 4}}, t)
	testWallsGates([][]int{[]int{0, -1}, []int{SIGNROOM, SIGNROOM}}, [][]int{[]int{0, -1}, []int{1, 2}}, t)
}

func testWallsGates(b [][]int, e [][]int, t *testing.T) {
	wallsGates(b)
	for i, _ := range b {
		for j, _ := range b[i] {
			if b[i][j] != e[i][j] {
				fmt.Printf("input: %v\t\n", b)
				t.Errorf("Test name is %s, ", t.Name())
				return
			}
		}
	}
}

func TestCout1s(t *testing.T) {
	testCout1s(5, []uint{0, 1, 1, 2, 1, 2}, t)
	testCout1s(3, []uint{0, 1, 1, 2}, t)

}

func testCout1s(m uint, e []uint, t *testing.T) {
	r := cout1s(m)
	for i, v := range r {
		if e[i] != v {
			fmt.Printf("input: %v\t%v\t\n", m, r)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}

func TestMaxPali(t *testing.T) {
	testMaxPali("bbbab", 4, t)
	testMaxPali("bbbbb", 5, t)
}

func testMaxPali(s string, e uint, t *testing.T) {
	r := maxPali(s)
	if e != r {
		fmt.Printf("input: %v\t%v\t\n", s, r)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestOnesAndZeors(t *testing.T) {
	testOnesAndZeors([]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4, t)
	testOnesAndZeors([]string{"10", "0001", "111001", "1", "0"}, 7, 7, 5, t)
}

func testOnesAndZeors(ss []string, os, zs int, e uint, t *testing.T) {
	r := onesAndZeors(ss, os, zs)
	if e != r {
		fmt.Printf("input: %v\t%v\t\n", ss, r)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestCoinChange(t *testing.T) {
	testCoinChange([]int{1, 2, 5}, 11, 3, t)
	testCoinChange([]int{2}, 3, -1, t)

	testCoinChange([]int{1, 2, 5, 10}, 18, 4, t)
	testCoinChange([]int{2, 5, 10}, 58, 9, t)
	testCoinChange([]int{1, 2, 5, 10}, 58, 8, t)
}

func testCoinChange(cs []int, tg int, e int, t *testing.T) {
	r := coinChangeBuf(cs, tg)
	if e != r {
		fmt.Printf("input: %v\t%v\t\n", cs, r)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestNumIslands(t *testing.T) {

	g := bigIsl{[][]byte{[]byte{1, 1, 0, 0, 0}, []byte{0, 1, 0, 0, 1}, []byte{0, 0, 0, 1, 1}, []byte{0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1}}, 2, 0}
	//testNumIslands( &g,2, t)

	g = bigIsl{[][]byte{[]byte{1, 1, 0, 0, 0}, []byte{0, 1, 0, 0, 1}, []byte{0, 0, 0, 1, 1}, []byte{0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1}}, 1, 0}
	//testNumIslands( &g,3, t)

	g = bigIsl{[][]byte{[]byte{1, 1, 0, 0, 0}, []byte{0, 1, 0, 0, 1}, []byte{0, 0, 0, 1, 1}, []byte{0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1}}, 4, 0}
	testNumIslands(&g, 0, t)

	g = bigIsl{[][]byte{[]byte{1, 1, 0, 0, 0}, []byte{0, 1, 0, 0, 1}, []byte{0, 1, 1, 1, 1}, []byte{0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 1}}, 4, 0}
	testNumIslands(&g, 1, t)

	g = bigIsl{[][]byte{[]byte{1, 1, 0, 0, 0}, []byte{1, 1, 0, 0, 1}, []byte{0, 0, 0, 1, 1}, []byte{0, 0, 0, 1, 1}, []byte{0, 0, 0, 1, 1}}, 4, 0}
	testNumIslands(&g, 2, t)

	g = bigIsl{[][]byte{[]byte{1, 0}, []byte{0, 1}}, 1, 0}
	testNumIslands(&g, 2, t)

}

func testNumIslands(g *bigIsl, e int, t *testing.T) {
	g.calc()
	if e != g.iss {
		fmt.Printf("input: %v\t%v\t\n", g.iss, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestShoestPalindrome(t *testing.T) {
	testShoestPalindrome([]byte("aacecaaa"), 1, t)
	testShoestPalindrome([]byte("dcbabcd"), 0, t)
	testShoestPalindrome([]byte("abcd"), 3, t)
}

func testShoestPalindrome(s []byte, e int, t *testing.T) {
	res := stestPali(s)
	if e != res {
		fmt.Printf("input: %v\t%v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestSplitStr(t *testing.T) {
	testSplitStr([]byte("123"), t)
	testSplitStr([]byte("12345"), t)
}

func testSplitStr(s []byte, t *testing.T) {
	res := splitStr(s)
	for _, ss := range res {
		for _, s := range ss {
			fmt.Printf("%v\t", s)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-------------\n")
}

func TestWordBrk3(t *testing.T) {
	testWordBrk3([]string{"Cat", "Mat", "Ca", "tM", "at", "C", "Dog", "og", "Do"}, "CatMat", 3, t)
	testWordBrk3([]string{"Cat", "Mat", "Ca", "tM", "at", "C", "Dog", "og", "Do"}, "CatDo", 2, t)
	testWordBrk3([]string{}, "a", 0, t)
}

func testWordBrk3(dict []string, s string, e int, t *testing.T) {
	res := wordBrk3(dict, s)
	if e != res {
		fmt.Printf("input: %v\t%v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestSwues(t *testing.T) {
	testSwues([]int{1, 2, 1, 3, 3}, 3, 5, t)
	testSwues([]int{1, 2, 1, 2, 1}, 3, 3, t)
	testSwues([]int{1, 2, 1, 1, 1}, 3, 2, t)
	testSwues([]int{1, 2, 1, 1, 1}, 2, 4, t)
}

func testSwues(ns []int, ws, e int, t *testing.T) {
	res := swues(ns, ws)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestTrimBST(t *testing.T) {
	testTrimBST(&bstNode{8, &bstNode{3, &bstNode{1, nil, nil},
		&bstNode{6, &bstNode{4, nil, nil}, &bstNode{7, nil, nil}}},
		&bstNode{10, nil, &bstNode{14, &bstNode{13, nil, nil}, nil}}},
		5, 13, t)
}

func testTrimBST(r *bstNode, min, max uint16, t *testing.T) {
	res := trimBST(r, min, max)
	printBST(res)
}

func TestCutRod(t *testing.T) {
	testCutRod([]uint{1, 5, 8, 9, 10, 17, 17, 20}, 8, 22, t)
	testCutRod([]uint{3, 5, 8, 9, 10, 17, 17, 20}, 8, 24, t)

}

func testCutRod(ns []uint, ws, e uint, t *testing.T) {
	res := cutRod(ns, ws)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMinPart(t *testing.T) {
	testMinPart([]int{1, 6, 11, 5}, 1, t)
	testMinPart([]int{1, 2, 3, 4}, 0, t)

}

func testMinPart(ns []int, e int, t *testing.T) {
	res := minPart(ns)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestBstRightView(t *testing.T) {
	testBstRightView(&bstNode{1, &bstNode{2, nil, &bstNode{5, nil, nil}}, &bstNode{3, nil, &bstNode{4, nil, nil}}}, []uint16{1, 3, 4}, t)
	testBstRightView(&bstNode{1, &bstNode{2, nil, nil}, &bstNode{3, nil, nil}}, []uint16{1, 3}, t)
}

func testBstRightView(rt *bstNode, e []uint16, t *testing.T) {
	res := bstRightView(rt)
	for i, v := range res {
		if e[i] != v {
			fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}


func TestMinMeetingRooms(t *testing.T) {
	testMinMeetingRooms([][]int{[]int{0, 30}, []int{5, 10},[]int{15, 20}}, 2,t)
	testMinMeetingRooms([][]int{[]int{7,10},[]int{2,4}}, 1, t)
}

func testMinMeetingRooms(ms [][]int, e int, t *testing.T) {
	res := minMeetingRooms(ms)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestLadderLength(t *testing.T) {
	testLadderLength("hit","cog", []string{"hot","dot","dog","lot","log","cog"}, 5,t)
	testLadderLength("hit","cog", []string{"hot","dot","dog","lot","log"}, 5,t)
}

func testLadderLength(bs, es string, ws []string, e int, t *testing.T) {
	res := ladderLength(bs, es, ws)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestDecodeString(t *testing.T) {
	testDecodeString("3[a]2[bc]", "aaabcbc",t)
	testDecodeString("3[a2[c]]", "accaccacc", t)
	testDecodeString("2[abc]3[cd]ef", "abcabccdcdcdef", t)
	testDecodeString("abc3[cd]xyz", "abccdcdcdxyz", t)
	testDecodeString("3[a]2[b4[F]c]", "aaabFFFFcbFFFFc", t)
}

func testDecodeString(s, e string, t *testing.T) {
	res := decodeStringIt(s)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}

	//bs := decodeString([]byte(s))
	//if e != string(bs) {
	//	fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
	//	t.Errorf("Test name is %s, ", t.Name())
	//	return
	//}

}

func TestMinWindow(t *testing.T) {
	testMinWindow("abcdebdde", "bde", "bcde", t)
	testMinWindow("abcdebdde", "bde", "bcde", t)
}

func testMinWindow(a, b, e string, t *testing.T) {
	res := minWindow(a, b)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestLongestSubarray(t *testing.T) {
	testLongestSubarray([]int{8,2,4,7}, 4, 2, t)
	testLongestSubarray([]int{10,1,2,4,7,2}, 5, 4, t)
	testLongestSubarray([]int{4,2,2,2,4,4,2,2}, 0, 3, t)
}

func testLongestSubarray(ns []int, l, e int, t *testing.T) {
	res := longestSubarray(ns, l)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

//asteroidCollision
func TestAsteroidCollision(t *testing.T) {
	testAsteroidCollision([]int{5, 10, -5}, []int{5,10}, t)
	testAsteroidCollision([]int{-2,-2,1,-2}, []int{-2,-2,-2}, t)
	testAsteroidCollision([]int{10, 2, -5}, []int{10}, t)
}

func testAsteroidCollision(ns, e []int, t *testing.T) {
	res := asteroidCollision(ns)
	if !assertIntSlicesEqual(res, e) {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestTimeMap(t *testing.T) {
	//tm := Constructor()
	//tm.Set("foo","bar",1)
	//e := tm.Get("foo", 1)
	//fmt.Printf("Res:%v\n", e)
	//e = tm.Get("foo", 3)
	//fmt.Printf("Res:%v\n", e)
	//tm.Set("foo","bar2",4)
	//e = tm.Get("foo", 4)
	//fmt.Printf("Res:%v\n", e)
	//e = tm.Get("foo", 5)
	//fmt.Printf("Res:%v\n", e)


	tm := Constructor()
	//["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
	tm.Set("love","high",10)
	tm.Set("love","low",20)
	e := tm.Get("love", 5)
	fmt.Printf("Res:%v\n", e)
	e = tm.Get("love", 10)
	fmt.Printf("Res:%v\n", e)
	e = tm.Get("love", 15)
	fmt.Printf("Res:%v\n", e)
	e = tm.Get("love", 20)
	fmt.Printf("Res:%v\n", e)
	e = tm.Get("love", 25)
	fmt.Printf("Res:%v\n", e)
}

func TestNumMatrix(t *testing.T) {
	nm := NumMatrixConstructor([][]int{[]int{3,0,1,4,2}, []int{5,6,3,2,1}, []int{1,2,0,1,5}, []int{4,1,0,1,7},
		[]int{1,0,3,0,5}})
	fmt.Printf("Res:%v\n", nm.SumRegion(2,1,4,3))
	fmt.Printf("Res:%v\n", nm.SumRegion(1,1,2,2))
	fmt.Printf("Res:%v\n", nm.SumRegion(1,2,2,4))
}



func TestRandomPick(t *testing.T) {
	rp := RPIConstructor([]int{1,2,3,3,3})

	for i:=0; i < 10; i++ {
		fmt.Printf("Res:%v\n", rp.Pick(3))
	}

	fmt.Printf("Res:%v\n", rp.Pick(1))
	fmt.Printf("Res:%v\n", rp.Pick(2))
	fmt.Printf("Res:%v\n", rp.Pick(4))
}

//[[],[2,4,6],[1,4,8,9],[7,8],[1,2,8,9],[6,9],[1,5,7,8,9],[3,6,9],[2,3,4,6,9],[2,4,5,6,7,8]]
func TestIsBipartite(t *testing.T) {
	testIsBipartite([][]int{[]int{},[]int{2,4,6},[]int{1,4,8,9},[]int{7,8},[]int{1,2,8,9},[]int{6,9},[]int{1,5,7,8,9},[]int{3,6,9},[]int{2,3,4,6,9},[]int{2,4,5,6,7,8}}, false, t)

}

func testIsBipartite(ns [][]int, e bool, t *testing.T) {
	res := isBipartite(ns)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

//sortList(head *ListNode)
func TestSortList(t *testing.T) {
	//4->2->1->3
	h := sortList(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}})
	for n := h; n != nil; n=n.Next{fmt.Printf("%v\n", n.Val)}

	//-1->5->3->4->0
	h = sortList(&ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0,nil}}}}})
	for n := h; n != nil; n=n.Next{fmt.Printf("%v\n", n.Val)}
}


//groupAnagrams
func TestGroupAnagrams(t *testing.T) {
	res := groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})

	for _, ss := range res {
		for _, s := range ss {
			fmt.Printf("%s, ", s)
		}
		fmt.Printf("\n")
	}
}


func TestPartitionLabels(t *testing.T) {
	testPartitionLabels("ababcbacadefegdehijhklij", []int{9,7,8}, t)
}

func testPartitionLabels(s string, e []int, t *testing.T) {
	res := partitionLabels(s)
	for i, v := range res {
		if e[i] != v {
			fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}


func TestMostCommonWord(t *testing.T) {
	testMostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball", t)
}

func testMostCommonWord(s string, b []string, e string, t *testing.T) {
	res := mostCommonWord(s, b)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestValidateStackSequences(t *testing.T) {
	testValidateStackSequences([]int{1,2,3,4,5}, []int{4,5,3,2,1}, true, t)
}

func testValidateStackSequences(ps, pp []int, e bool, t *testing.T) {
	res := validateStackSequences(ps, pp)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestExpressiveWords(t *testing.T) {
	testExpressiveWords("heeellooo", []string{"hello", "hi", "helo"}, 1, t)

	testExpressiveWords("zzzzzyyyyy", []string{"zzyy","zy","zyy"}, 3, t)

	testExpressiveWords("heeellooo", []string{"heeelloooworld"}, 0, t)
}

func testExpressiveWords(S string, ws []string, e int, t *testing.T) {
	res := expressiveWords(S, ws)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestIsNStraightHand(t *testing.T) {
	testIsNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true, t)
	testIsNStraightHand([]int{1, 1, 2, 2, 3, 3}, 2, false, t)
}

func testIsNStraightHand(h []int, w int, e bool, t *testing.T) {
	res := isNStraightHand(h, w)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v\n", h)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestMaxScore(t *testing.T) {
	testMaxScore([]int{1,2,3,4,5,6,1}, 3, 12, t)
	testMaxScore([]int{2,2,2}, 2, 4, t)
	testMaxScore([]int{9,7,7,9,7,7,9}, 7, 55, t)
}

func testMaxScore(h []int, w int, e int, t *testing.T) {
	res := maxScore(h, w)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v\n", h)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestLengthOfLongestSubstring(t *testing.T) {
	testLengthOfLongestSubstring("abcabcbb", 3, t)
	testLengthOfLongestSubstring("bbbbb",1, t)
	testLengthOfLongestSubstring("pwwkew", 3, t)
}

func testLengthOfLongestSubstring(s string, e int, t *testing.T) {
	res := lengthOfLongestSubstring(s)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v\n", s)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestDistanceK(t *testing.T) {
	testDistanceK(&TreeNode{0,
		&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, nil},
		&TreeNode{3, nil,nil}},
		1, 0, t)

	testDistanceK(
		&TreeNode{3,
		&TreeNode{5,
			&TreeNode{6, nil, nil},
			&TreeNode{2,
				&TreeNode{7, nil, nil},
				&TreeNode{4, nil, nil},
			},
		},
		&TreeNode{1,
			&TreeNode{0, nil, nil},
			&TreeNode{8, nil, nil},
			},
		}, 5, 2, t)

	testDistanceK(&TreeNode{0,nil,&TreeNode{1, nil,
		&TreeNode{2,nil,&TreeNode{3, nil,
			&TreeNode{4, nil, nil}}}}},2, 2, t)

	testDistanceK(&TreeNode{0,
		&TreeNode{2, nil, nil},
		&TreeNode{1, &TreeNode{3, nil,nil}, nil}},
		3, 3, t)
}

func testDistanceK(r *TreeNode, tg, k int, t *testing.T) {
	res := distanceK(r, &TreeNode{tg, nil, nil}, k)
	fmt.Printf("Res: %v\t", res)
}



func TestZigZagConvert(t *testing.T) {
	testZigZagConvert("PAYPALISHIRING","PAHNAPLSIIGYIR", 3,  t)
	testZigZagConvert("PAYPALISHIRING", "PINALSIGYAHRPI", 4, t)

}

func testZigZagConvert(s, e string, l int, t *testing.T) {
	res := convert(s, l)
	if e != res {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v\n", s)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestMinAvailableDuration(t *testing.T) {
	testMinAvailableDuration([][]int{[]int{0,2}}, [][]int{[]int{1, 3}}, 1, []int{1, 2}, t)
	testMinAvailableDuration([][]int{[]int{0,1},[]int{100,1000100}}, [][]int{[]int{90,1000100},[]int{0,2}}, 1000000, []int{100,1000100}, t)
	testMinAvailableDuration([][]int{[]int{216397070,363167701}, []int{98730764,158208909}, []int{441003187,466254040}, []int{558239978,678368334}, []int{683942980,717766451}},
		[][]int{[]int{50490609,222653186}, []int{512711631,670791418}, []int{730229023,802410205}, []int{812553104,891266775}, []int{230032010,399152578}}, 456085, []int{98730764,99186849}, t)
}

func testMinAvailableDuration(a, b [][]int, d int, e []int, t *testing.T) {
	res := minAvailableDuration(a, b, d)
	if e[0] != res[0] || e[1] != res[1] {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v%v\n", a, b)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMultiply(t *testing.T) {
	testMultiply([][]int{[]int{1,0,0}, []int{-1,0,3}},[][]int{[]int{7,0,0}, []int{0,0,0}, []int{0,0,1}}, [][]int{[]int{7,0,0}, []int{-7,0,3}},  t)

}

func testMultiply(a, b [][]int, e [][]int, t *testing.T) {
	res := multiply(a, b)
	if len(res) != len(e) {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v\n", a)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
	for i, _ := range res {
		if !assertIntSlicesEqual(res[i], e[i]) {
			fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
			fmt.Printf("Input: %v\n", a)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}


func TestNumOfMinutes(t *testing.T) {
	testNumOfMinutes(10,3, []int{8,9,8,-1,7,1,2,0,3,0}, []int{224,943,160,909,0,0,0,643,867,722}, 3665, t)
}

func testNumOfMinutes(n, h int, m, it []int, e int, t *testing.T) {
	res := numOfMinutes(n,h, m, it)
	if res != e {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v%v\n", m, it)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}



func TestFindReplaceString (t *testing.T) {
	testFindReplaceString("vmokgggqzp", []int{3,5,1}, []string{"kg","ggq","mo"}, []string{"s","so","bfr"}, "vbfrssozp", t)
	testFindReplaceString("vmokgggqzp", []int{1, 3,5}, []string{"mo","kg","ggq","mo"}, []string{"bfr", "s","so"}, "vbfrssozp", t)
}

func testFindReplaceString (s string, ind []int, src, tag []string, e string, t *testing.T) {

	res := findReplaceString(s, ind, src, tag)
	if res != e {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v%v\n", s, ind)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}


func TestMinSumOfLengths (t *testing.T) {
	//testMinSumOfLengths([]int{2,2,4,4,4,4,4,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}, 20, 23, t)
	testMinSumOfLengths([]int{3,2,2,4,3}, 3, 2, t)
}

func testMinSumOfLengths (d []int, s int, e int, t *testing.T) {

	res := minSumOfLengths(d, s)
	if res != e {
		fmt.Printf("Res: %v\tExp: %v\t\n", res, e)
		fmt.Printf("Input: %v%v\n", s, d)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}