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

func testCountSubarrays(a []int, exp []uint16, t  *testing.T){
	res := countSubarrays(a)

	for i, v := range res {
		if v != exp[i] {
			fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
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