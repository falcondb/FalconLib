package exercise

import (
"fmt"
"testing"
)

func TestAreTheyEqual(t *testing.T) {
	testAreTheyEqual( []int {1, 2, 3, 4}, []int {1, 4, 3, 2}, true, t)
	testAreTheyEqual( []int {1, 2, 3, 1}, []int {1, 4, 3, 2}, false, t)

	testAreTheyEqual( []int {1, 2, 3, 1}, []int {1, 2, 3, 1}, true, t)
	testAreTheyEqual( []int {1, 2, 3, 1}, []int {1, 3, 2, 1}, true, t)
	testAreTheyEqual( []int {10, 2, 3, 1}, []int {1, 3, 2, 1}, false, t)
	testAreTheyEqual( []int {10, 2, 3, 1}, []int {1, 3, 2, 5}, false, t)

	testAreTheyEqual( []int {1, 2, 3, 4}, []int {4, 3, 2, 1}, true, t)

	testAreTheyEqual( []int {1, 2, 3, 4}, []int {3, 2, 1, 4}, true, t)

}

func testAreTheyEqual(a, b []int, exp bool, t  *testing.T){
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

func testCountSubarrays(a []int, exp []uint16, t  *testing.T) {
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

func TestGetMilestoneDays (t *testing.T) {
	testGetMilestoneDays([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, []int{100, 200, 500}, []int{4, 6, 10}, t)
}

func testGetMilestoneDays (re, ms, exp []int, t *testing.T) {
	ok, res := getMilestoneDays(re, ms)
	if !ok {
		t.Errorf("Test name is %s, ", t.Name())
	}

	if !arraysEqual(res, exp){
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

	head := listNode{1, &listNode{2, &listNode{3, &listNode {4, nil}}}}

	rev, pre := reverseList(&head)

	for n, pre := rev, uint16(1 << uint16(16)-1); n != nil; n = n.Next {
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

	testReverseEven(&listNode{1, &listNode{2, &listNode{8,nil}}},
		&listNode{1, &listNode{8, &listNode{2, nil}}},
		t)

	testReverseEven(&listNode{1, &listNode{2, &listNode{8,&listNode{0, &listNode{5, nil}}}}},
		&listNode{1, &listNode{0, &listNode{8,&listNode{2, &listNode{5, nil}}}}},
		t)

	testReverseEven(&listNode{1, &listNode{2, &listNode{8, &listNode{9, &listNode{12, &listNode{16, nil}}}}}},
					&listNode{1, &listNode{8, &listNode{2, &listNode{9, &listNode{16, &listNode{12, nil}}}}}},
					t)
}

func testReverseEven(h *listNode, exp *listNode, t *testing.T) {
	res:= reverseEven(h)

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
	testNrDiffTriangles([]triangle {triangle{2, 2, 3}, triangle {3, 2, 2}, triangle{2, 5, 6}}, 2, t)
	testNrDiffTriangles([]triangle {triangle{8, 4, 6}, triangle {100, 101, 102}, triangle{84, 93, 173}}, 3, t)
	testNrDiffTriangles([]triangle {triangle{5, 8, 9}, triangle {5, 9, 8}, triangle{9, 5, 8},
									triangle{9, 8, 5}, triangle{8, 9, 5}, triangle{8, 5, 9}}, 1, t)}

func testNrDiffTriangles(a tas, exp uint, t *testing.T) {
	res := nrDiffTriangles(a)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestNumberOfWays (t *testing.T) {
	TestnumberOfWays([]int{1, 5, 3, 3, 3}, 6, 4, t)
	TestnumberOfWays([]int{1, 2, 3, 4, 3}, 6, 2, t)
}

func TestnumberOfWays (a []int, k int,  exp uint, t *testing.T) {
	res := numberOfWays(a, k)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}



func TestBstSearchRange(t *testing.T){
	// case 1
	r := bstNode{5, nil, nil}
	testBstSearchRange(&r, 6,10, []uint16{}, t)

	// case 2
	r = bstNode{20, &bstNode{8, &bstNode{4, nil, nil}, &bstNode{12, nil, nil}}, &bstNode{22, nil, nil}}
	testBstSearchRange(&r, 10, 22, []uint16{12, 20, 22}, t)
}

func testBstSearchRange(root *bstNode, min, max uint16, exp []uint16, t *testing.T){
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
	n0, n1, n2, n3, n4, n5, n6, n7 :=  graphNode{0, nil}, graphNode{1, nil}, graphNode{2, nil},
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