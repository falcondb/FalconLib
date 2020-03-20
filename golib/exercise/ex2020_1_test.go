package exercise

import (
	"fmt"
	"strings"
	"testing"
)

func TestMergeTwo(t *testing.T) {

	a1 := []int{3, 7, 5}
	a2 := []int{6, 4, 2}

	testMergeTwo(a1, a2, t)

	var a0 []int
	testMergeTwo(a0, a2, t)
	testMergeTwo(a1, a0, t)

	a3 := []int {1, 10, 20, 100, 9}
	testMergeTwo(a1, a3, t)
	testMergeTwo(a3, a2, t)
}


func testMergeTwo(a1 []int , a2 []int, t *testing.T) {
	m := MergeTwo(a1, a2)

	if len(m) != len(a1)+len(a2) {
		fmt.Println(m)
		t.Errorf("Failed in the test of MergeTwo")
	}

	for i := 0; i < len(m)-1; i++ {
		if m[i] > m [i+1]{
			fmt.Println(m)
			t.Errorf("Failed in the test of MergeTwo")
		}
	}
}


func TestReverseWords(t *testing.T) {
	testReverseWords("the sky is blue", t)
	testReverseWords("hello world", t)
}

func testReverseWords(w string, t *testing.T) {
	r := ReverseWords(w)
	ws := strings.Fields(w)
	pp := 1 << 32 - 1

	for _, s := range ws {
		if strings.Index(r, s) > pp {
			fmt.Println(w)
			fmt.Println(r)
			t.Errorf("Failed in the test of MergeTwo")
		}
	}
}

func TestLastNList (t *testing.T) {
	testLastNList(nil, 2, false, 0, t)

	fh :=  listNode {0, nil}
	var cur *listNode
	var i uint16
	for i=1 ; i < 11; i++ {
		cur = &listNode{i, fh.Next}
		fh.Next = cur
	}

	testLastNList(fh.Next, 100, false, 0, t)
	testLastNList(fh.Next, 11, false, 0, t)
	testLastNList(fh.Next, 10, true, 10, t)
	testLastNList(fh.Next, 9, true, 9, t)
	testLastNList(fh.Next, 1, true, 1, t)
	testLastNList(fh.Next, 3, true, 3, t)

}

func testLastNList(head * listNode, p int, eOK bool, exp uint16, t *testing.T) {
	res, ok := lastNList(head, p)

	if ok != eOK || res != exp {
		t.Fail()
	}
}

func TestBS (t *testing.T) {
	arr := make([]int, 0)
	for i := 10; i > 0; i-- {
		arr = append(arr, i*2)
	}

	testBS(nil, 0, false, -1, t)

	testBS(arr, 2, true, 9, t)
	testBS(arr, 4, true, 8, t)
	testBS(arr, 20, true, 0, t)
	testBS(arr, 12, true, 4, t)

	testBS(arr, 11, false, 5, t)
	testBS(arr, 17, false, 2, t)
}

func testBS(data []int, k int , eOK bool, exp int, t *testing.T) {
	res, ok := binarySearch(data, k)
	if ok != eOK || res != exp {
		t.Fail()
	}
}

func TestClosest (t *testing.T) {
	ds1, ds2 := []int {3, 6, 7, 4}, []int{2, 8, 9, 3}

	testClosest(ds1, ds2, 0, t)

	ds1, ds2 = []int {1, 2, 3, 4}, []int {7, 6, 5}
	testClosest(ds1, ds2, 1, t)
}

func testClosest(ds1 []int, ds2 []int,  exp int, t *testing.T) {
	md, ok := closest(ds1, ds2)
	if ok != nil || md != exp {
		t.Fail()
	}
}

func TestCalcString(t *testing.T) {

	testCalcString([]byte{2, '+', 3}, 5, t)

	testCalcString([]byte{2, '+', 3, '+', 2}, 7, t)
	testCalcString([]byte{2, '+', 3, '*', 2}, 8, t)
	testCalcString([]byte{2, '*', 3, '+', 2}, 8, t)
	testCalcString([]byte{2, '*', 3, '*', 2}, 12, t)

	testCalcString([]byte{2, '*', 3, '*', 2, '+', 5}, 17, t)
	testCalcString([]byte{2, '+', 3, '*', 2, '+', 5}, 13, t)
	testCalcString([]byte{2, '+', 3, '*', 2, '+', 5, '*', 2}, 18, t)
}

func testCalcString(fm []byte, exp int, t *testing.T) {
	if calcString(fm) != exp {
		fmt.Println(fm)
		t.Fail()
	}
}


func TestMWordSearch(t *testing.T) {
	w := []byte("ABCCED")
	m := [][]byte {[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	rows, cols := 3, 4

	sd := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		sd[i] = make([]bool, cols)
	}
	wm := wordMatrixSearch{m, sd, w}

	_testMWordSearch(&wm, true, t)

	/************************************/
	w = []byte("ABCCEF")
	m = [][]byte {[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	rows, cols = 3, 4

	sd = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		sd[i] = make([]bool, cols)
	}
	wm = wordMatrixSearch{m, sd, w}
	_testMWordSearch(&wm, false, t)

	/************************************/
	w = []byte("ASFCEE")
	m = [][]byte {[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	rows, cols = 3, 4

	sd = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		sd[i] = make([]bool, cols)
	}
	wm = wordMatrixSearch{m, sd, w}

	_testMWordSearch(&wm, true, t)

	/************************************/
	w = []byte("ACBC")
	m = [][]byte {[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	rows, cols = 3, 4

	sd = make([][]bool, rows)
	for i := 0; i < rows; i++ {
		sd[i] = make([]bool, cols)
	}
	wm = wordMatrixSearch{m, sd, w}

	_testMWordSearch(&wm, false, t)
}

func _testMWordSearch(wm *wordMatrixSearch, exp bool, t *testing.T) {
	if exp != wm.mWordSearch(0, 0, 0) {
		t.Fail()
	}
}

func TestCanWordBreak(t *testing.T) {
	d:= map[string]bool {"lint":true, "code":true, "a":true, "bb":true}
	wb := WordBreak{d}
	_testCanWordBreak(&wb, "lintcode", true, t)
	_testCanWordBreak(&wb, "lintode", false, t)
	_testCanWordBreak(&wb, "codea", true, t)
	_testCanWordBreak(&wb, "bcodea", false, t)

	_testCanWordBreak(&wb, "abb", true, t)
	_testCanWordBreak(&wb, "a", true, t)
	_testCanWordBreak(&wb, "aaaaa", true, t)
	_testCanWordBreak(&wb, "bbbb", true, t)
	_testCanWordBreak(&wb, "bbbbb", false, t)

}

func _testCanWordBreak(wb *WordBreak, s string, exp bool, t *testing.T){
	if exp != wb.CanWordBreak(s) {
		fmt.Printf("%v\n",s)
		t.Errorf("Test name is %s", t.Name())
	}
}

// TODO: debug
func testTopKFreWord( t *testing.T){
	wds := []string{"yes", "lint", "code",
		"yes", "code", "baby",
		"you", "baby", "chrome",
		"safari", "lint", "code",
		"body", "lint", "code"}
	exp := map[string]bool {"code":true, "lint":true, "baby":true}
	_testTopKFreWords(wds, 3, exp, t)
}

func _testTopKFreWords(wds []string, k int, exp map[string]bool, t *testing.T){
	res, err := TopKFreWords(wds, k)

	if err != nil {
		fmt.Printf("%v\n", k)
		t.Errorf("Test name is %s", t.Name())
	} else {
		for _, w := range res {
			_, found := exp[w]
			if !found {
				fmt.Printf("%v\n", k)
				t.Errorf("Test name is %s", t.Name())
				return
			}
		}
	}
}

func TestRemoveArbitrarySpace(t *testing.T) {
	if removeArbitrarySpace("The  sky   is blue") != "The sky is blue" {
		t.Errorf("Test name is %s, ", t.Name())
	}

	if removeArbitrarySpace("  low               ercase  ") != "low ercase" {
		t.Errorf("Test name is %s, ", t.Name())
	}

}

func TestMinPartition(t *testing.T) {
	_testMinPartition([]int {1}, 1, t)
	_testMinPartition([]int {1, 6}, 5, t)
	_testMinPartition([]int {1, 6, 3}, 2, t)
	_testMinPartition([]int {1, 6, 11, 5}, 1, t)
	_testMinPartition([]int {1, 2, 3, 4}, 0, t)
	_testMinPartition([]int {1, 2, 3, 4, 5}, 1, t)

	_testMinPartition([]int {10, 2, 100, 500, 5}, 383, t)
	_testMinPartition([]int {10, 2, 100, 50, 5}, 33, t)
	_testMinPartition([]int {10, 2, 20, 30, 5}, 3, t)
}

func _testMinPartition(nums []int, exp int, t *testing.T) {
	res := minPartition(nums)
	if res != exp {
		fmt.Printf("Nums:%v, result:%v\n", nums, res)
		t.Errorf("Test name is %s, ", t.Name())
	}
}

func TestPpids(t *testing.T) {
	res := ppids([]int {1, 3, 10, 5}, []int {3, 0, 5, 3}, 5)
	fmt.Printf("%v\n", res)

	fmt.Printf("%v\n",ppids([]int {1, 2, 3}, []int {0, 1, 1}, 2))
}

func TestInterleavingString(t *testing.T) {
	testInterleavingString("","", "", true, t)
	testInterleavingString("","dbbca", "dbbca", true, t)
	testInterleavingString("aabcc","dbbca", "aadbbcbcac", true, t)
	testInterleavingString("","", "1", false, t)
	testInterleavingString("aabcc","dbbca", "aadbbbaccc", false, t)
}

func testInterleavingString(sa, sb, sm string, exp bool, t *testing.T) {
	res := interleavingString(sa, sb, sm)
	if res != exp {
		fmt.Printf("SA:%v, SB:%v, SM:%v, result:%v\n", sa, sb, sm, res)
		t.Errorf("Test name is %s, ", t.Name())
	}
}


func TestFlatListOrInt( t *testing.T) {

	l := listOrInt{  //[[1,1],2,[1,1]]
		[]listOrInt{ //[1,1],2,[1,1]
			listOrInt{ // [1,1]
				[] listOrInt{
					listOrInt{nil, 1},
					listOrInt{nil, 1},
				},
				-1,
			},
			listOrInt{nil, 2}, // 2
			listOrInt{
				[] listOrInt{
					listOrInt{nil, 1},
					listOrInt{nil, 1},
				},
				-1,
			},
		},
		-1,
	}

	testFlatListOrInt(&l, []int{1,1,2,1,1}, t)
	// [4,[3,[2,[1]]]]
	l = listOrInt {
		[]listOrInt{
			listOrInt{nil, 4},
			listOrInt {
				[]listOrInt {
					listOrInt{nil, 3},
					listOrInt {
						[]listOrInt {
							listOrInt{nil, 2},
							listOrInt {
								[]listOrInt {
									listOrInt{nil, 1},
								},
								-1,
							},
						},
						-1,
					},
				},
				-1,
			},
		},
		-1,
	}
	testFlatListOrInt(&l, []int{4,3,2,1}, t)

}

func testFlatListOrInt(li *listOrInt, exp []int, t *testing.T) {
	res := flatListOrInt(li)

	if len(res) != len(exp) {
		fmt.Printf("Res:%v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}

	for i, v := range res {
		if v != exp[i] {
			fmt.Printf("Res:%v\t Exp:%v\n", res, exp)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}