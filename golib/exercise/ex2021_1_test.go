package exercise

import (
	"fmt"
	"testing"
)

func TestIsAlienSorted(t *testing.T) {
	testIsAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true, t)
	testIsAlienSorted([]string{"word","world","row"},	"worldabcefghijkmnpqstuvxyz", false, t)
}

func testIsAlienSorted(ws []string, order string, exp bool, t *testing.T) {
	res := isAlienSorted(ws, order)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}


func TestKkClosest(t *testing.T) {

	testKkClosest([][]int{[]int{6,10}, []int{-3,3}, []int{-2, 5}, []int{0,2}}, [][]int{[]int{0,2}, []int{-3,3}, []int{-2,5}}, 3, t)
	testKkClosest([][]int{[]int{1,3},[]int{-2,2}}, [][]int{[]int{-2, 2}} ,1, t)
	testKkClosest([][]int{[]int{3,3}, []int{5, -1}, []int{-2, 4}}, [][]int{[]int{3,3}, []int{-2,4}}, 2, t)

}

func testKkClosest(a, exp [][]int,  k int, t *testing.T) {
	res := kClosest(a, k)

	fmt.Printf("Res: %v\t Exp:%v\n", res, exp)

}



func TestProductExceptSelf(t *testing.T) {

	testProductExceptSelf([]int{1,2,3,4}, []int{24,12,8,6}, t)

}

func testProductExceptSelf(a, exp []int,  t *testing.T) {
	res := productExceptSelf(a)

	if !assertIntSlicesEqual(res, exp) {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
	}
}



func TestMaxPathSum(t *testing.T) {

	//testMaxPathSum(&TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, &TreeNode{5, nil,nil}, nil}, nil}, nil}, nil}, 15, t)

	testMaxPathSum(&TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil,&TreeNode{5, nil,nil}}}}}, 15, t)
}

func testMaxPathSum(r *TreeNode, exp int, t *testing.T) {
	res := maxPathSum(r)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
	}
}


func TestCheckSubarraySum(t *testing.T) {

	//testCheckSubarraySum([]int{1, 0}, 2, false, t)
	// testCheckSubarraySum([]int{5, 0, 0, 0}, 3, true, t)
	testCheckSubarraySum([]int{1000000000}, 1000000000, false, t)
}

func testCheckSubarraySum(r []int, k int, exp bool, t *testing.T) {
	res := checkSubarraySum(r, k)

	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
	}
}


//func TestTreeToDoublyList(t *testing.T) {
//
//	testTreeToDoublyList(&Node{4, &Node{2, &Node{1, nil,nil}, &Node{3, nil,nil}},
//		&Node{5, nil, nil}}, t)
//}

//func testTreeToDoublyList(r *Node, t *testing.T) {
//	res := treeToDoublyList(r)
//	fmt.Printf(" %v", res.Val)
//	for cur := res.Right; cur != res; cur = cur.Right {
//		fmt.Printf(" %v", cur.Val)
//	}
//
//}


func TestRightSideView(t *testing.T) {

	testRightSideView(&TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4,nil,nil}}}, []int{1, 3, 4}, t)
}

func testRightSideView(r *TreeNode, exp []int, t *testing.T) {
	res := rightSideView(r)
	if !assertIntSlicesEqual(res, exp) {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
	}
}


func TestFindKthLargest(t *testing.T) {

	testFindKthLargest([]int{3,2,1,5,6,4}, 2, 5, t)
}

func testFindKthLargest(r []int, k, exp int, t *testing.T) {
	res := findKthLargest(r, k)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
	}
}



func TestIsPalindrome(t *testing.T) {

	//testIsPalindrome("A man, a plan, a canal: Panama", true, t)
	testIsPalindrome("0p", false, t)
}

func testIsPalindrome(r string, exp bool, t *testing.T) {
	res := isPalindrome(r)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
	}
}


func TestRotate(t *testing.T) {
	n := []int{1,2,3,4,5,6,7}
	rotate(n, 3)

	fmt.Printf("%v", n)
}

func TestSubsets(t *testing.T) {
	fmt.Printf("%v",subsets([]int{1,2, 3}))
}


func TestMaximumSwap(t *testing.T) {
	testMaximumSwap(98368, 98863, t)
	testMaximumSwap(115, 511, t)
	testMaximumSwap(2736, 7236, t)
}

func testMaximumSwap(r int, exp int, t *testing.T) {
	res := maximumSwap(r)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}


func TestWordBreak(t *testing.T) {
	res := wordBreak("catsanddog", []string{"cat","cats","and","sand","dog"})
	fmt.Printf("%v",res)
}

func TestWdContains(t *testing.T) {
	res := wdContains("baca", "bacaercd")
	fmt.Printf("%v",res)

	fmt.Printf("%v",wdContains("race", "bacaercd"))
}




func TestMinInsertions(t *testing.T) {
	testMinInsertions(		"))())(", 3, t)
	testMinInsertions(	"()())))()", 3, t)
}

func testMinInsertions(s string, exp int, t *testing.T) {
	res := minInsertions(s)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}

}



func TestVerticalOrder(t *testing.T) {


	res := verticalTraversal(&TreeNode{3, &TreeNode{9, &TreeNode{4, nil, nil}, &TreeNode{0, nil, nil}},
		&TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{7, nil, nil}}})

	//res := verticalOrder(&TreeNode{3, &TreeNode{9, &TreeNode{4, nil, nil}, &TreeNode{0, nil, nil}},
	//	&TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{7, nil, nil}}})

	//fmt.Printf("%v", res)

	res = verticalTraversal(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}})

	//res := verticalOrder(&TreeNode{3, &TreeNode{9, &TreeNode{4, nil, nil}, &TreeNode{0, nil, nil}},
	//	&TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{7, nil, nil}}})

	fmt.Printf("%v", res)
}

func TestValidateBinaryTreeNodes(t *testing.T) {
	testValidateBinaryTreeNodes(	[]int{1,0}, []int{-1,-1}, false, t)
	testValidateBinaryTreeNodes(	[]int{1,0,3,-1}, []int{-1,-1,-1,-1}, false, t)

}

func testValidateBinaryTreeNodes(l, r []int, exp bool, t *testing.T) {
	res := validateBinaryTreeNodes( len(l), l, r)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}

}

func TestCanPlaceFlowers(t *testing.T) {
	testCanPlaceFlowers(	[]int{0,0,1,0,1}, 1, true, t)

}

func testCanPlaceFlowers(l []int, n int, exp bool, t *testing.T) {
	res := canPlaceFlowers( l, n)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}

}


func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	testLengthOfLongestSubstringKDistinct(	"aba", 1, 1, t)

}

func testLengthOfLongestSubstringKDistinct(l string, n int, exp int, t *testing.T) {
	res := lengthOfLongestSubstringKDistinct( l, n)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}

}



func TestFindClosestElements(t *testing.T) {
	testFindClosestElements(	[]int{1,2,3,4,5}, 4, 3,  t)

}

func testFindClosestElements(l []int, k, v int, t *testing.T) {
	res := findClosestElements( l, k, v)

		fmt.Printf("Res: %v\t Exp:%v\n", res, true)
		t.Errorf("Test name is %s, ", t.Name())

}


func TestFindMissingRanges(t *testing.T) {
	testFindMissingRanges(	[]int{0,1,3,50,75}, 0,99,  t)

}

func testFindMissingRanges(l []int, k, v int, t *testing.T) {
	res := findMissingRanges( l, k, v)

	fmt.Printf("Res: %v\t Exp:%v\n", res, true)
	t.Errorf("Test name is %s, ", t.Name())

}



func TestRangeSumBST(t *testing.T) {
	testRangeSumBST(&TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{1, nil, nil}, nil},&TreeNode{7, &TreeNode{6, nil, nil }, nil}},
		&TreeNode{15, &TreeNode{13, nil,nil}, &TreeNode{18, nil, nil}}}, 6, 10,23,  t)

}

func testRangeSumBST(l *TreeNode, k, v, exp int, t *testing.T) {
	res := rangeSumBST( l, k, v)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}



func TestDiameterOfBinaryTree(t *testing.T) {
	testDiameterOfBinaryTree(&TreeNode{1, &TreeNode{2, nil, nil}, nil}, 1,  t)

}

func testDiameterOfBinaryTree(l *TreeNode, exp int, t *testing.T) {
	res := diameterOfBinaryTree( l)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}


func TestCheckInclusion(t *testing.T) {
	testCheckInclusion("ab", "eidbaooo", true, t)

}

func testCheckInclusion(l,r string, exp bool, t *testing.T) {
	res := checkInclusion( l, r)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}



func TestGetIntersectionNode(t *testing.T) {
	cn := &ListNode{1, &ListNode{8, &ListNode{4, &ListNode{5, nil}}}}
	testGetIntersectionNode(&ListNode{4, cn},
	&ListNode{5, &ListNode{6, cn}}, 1, t)

}

func testGetIntersectionNode(l,r *ListNode, exp int, t *testing.T) {
	res := getIntersectionNode( l, r)
	if res.Val != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}



func TestCanTransform(t *testing.T) {
	testCanTransform("RXXLRXRXL",	"XRLXXRRLX", true, t)

}

func testCanTransform(l,r string, exp bool, t *testing.T) {
	res := canTransform( l, r)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}



func TestScoreOfParentheses(t *testing.T) {
	testScoreOfParentheses("(())()",	3, t)
	testScoreOfParentheses("(())",	2, t)
	testScoreOfParentheses("()",	1, t)

}

func testScoreOfParentheses(l string, exp int, t *testing.T) {
	res := scoreOfParentheses( l)
	if res != exp {
		fmt.Printf("Res: %v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
	}
}


func TestOoddEvenList(t *testing.T) {
	testOoddEvenList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5,nil}}}}}, t)

}

func testOoddEvenList(l *ListNode, t *testing.T) {
	res := oddEvenList(l)
	for ; res != nil; res = res.Next {
		fmt.Println(res.Val)
	}
}


func TestCopyRandomList(t *testing.T) {
	n4 := &RDNode{1, nil, nil}
	n3 := &RDNode{10, n4, nil}
	n2 := &RDNode{11, n3, nil}
	n1 := &RDNode{13, n2, nil}
	n0 := &RDNode{7, n1, nil}

	n1.Random = n0
	n2.Random = n4
	n3.Random = n2
	n4.Random = n0

	testCopyRandomList(n0, t)

}

func testCopyRandomList(l *RDNode, t *testing.T) {
	res := copyRandomList(l)
	for ; res != nil; res = res.Next {
		fmt.Printf("\n%v ", res.Val)
		if  res.Random != nil {
			fmt.Printf("%v ", res.Random.Val)
		}
	}
}


func TestMinSteps(t *testing.T) {
	b := [][]byte{ []byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	testExist(b, "ABCCED", true, t)

}

func testExist(b [][]byte, s string, exp bool, t *testing.T) {
	res := exist(b, s)
	if  res != exp {
		fmt.Printf("%v ", res)
	}
}


func TestSpiralOrder(t *testing.T) {
	m := [][]int{[]int{1,2,3,4}, []int{5,6,7,8}, []int{9,10,11,12}}
	testSpiralOrder(m, t)

}

func testSpiralOrder(m [][]int, t *testing.T) {
	res := spiralOrder(m)
		fmt.Printf("%v ", res)
}


func TestLastStoneWeight(t *testing.T) {
	fmt.Printf("%v ", lastStoneWeight([]int{2,7,4,1,8,1}))
}



func TestLowestCommonAncestor(t *testing.T) {
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, &TreeNode{6, nil,nil}, &TreeNode{2, &TreeNode{7, nil, nil}, n4}}
	n0 := &TreeNode{0, nil, nil}
	n8 := &TreeNode{8, nil, nil}
	r := &TreeNode{3, n5,
		&TreeNode{1, n0, n8}}
	fmt.Printf("%v ", lowestCommonAncestor(r, n0, n8).Val)
}


func TestMinDifference(t *testing.T) {
	fmt.Printf("%v ", minDifference([]int{20,66,68,57,45,18,42,34,37,58}))
}

func TestNumSplits(t *testing.T) {
	//fmt.Printf("%v ",numSplits("abcd"))

	fmt.Printf("%v ",numSplits("aaaaa"))
	//fmt.Printf("%v ",numSplits("acbadbaada"))
}


func TestLongestStrChain(t *testing.T) {
	fmt.Printf("%v ", longestStrChain([]string{"abcd","dbqca"}))
}


func TestAsteroidCollision2(t *testing.T) {
	fmt.Printf("%v ", asteroidCollision2([]int{5,10,-5}))
}

func TestCarFleet(t *testing.T) {
	fmt.Printf("%v ", carFleet(10, []int{6,8}, []int{3,2}))
}

func TestCarFTestCountSquaresleet(t *testing.T) {
	countSquares([][]int{[]int{0,1,1,1}, []int{1,1,1,1}, []int{0,1,1,1}})
}


func TestScinsert(t *testing.T) {
	last := &SCNode{1, nil}
	st := &SCNode{3, &SCNode{4, last}}
	last.Next = st
	scinsert(st, 2)
}

func TestSimplifyPath(t *testing.T) {
	fmt.Printf("%v ", simplifyPath("/../"))
}

func TestSwapNodes(t *testing.T) {
	h := &ListNode{1,  &ListNode{2,  &ListNode{3,  &ListNode{4, &ListNode{5, nil}}}}}
	swapNodes(h, 3)
}

func TestMostCompetitive(t *testing.T) {
	//mostCompetitive([]int{2,4,3,3,5,4,9,6}, 4)
	mostCompetitive([]int{71,18,52,29,55,73,24,42,66,8,80,2}, 3)
}


func TestWordsTyping(t *testing.T) {
	wordsTyping([]string{"a","abc"}, 1, 2000)
}

func TestLongestPalindrome(t *testing.T) {
	longestPalindrome("babad")
}
func TestChange2(t *testing.T) {
	change(5, []int{1,2,5})
}

func TestChange3(t *testing.T) {
	minSol(10)
}


func TestToeplitzMatrix(t *testing.T) {
	isToeplitzMatrix([][]int{[]int{44,35,39}, []int{15,44,35}, []int{17,15,44}, []int{80,17,15}, []int{43,80,17}, []int{77,43,80}})
}

func TestStr2tree(t *testing.T) {
	str2tree("4(2(3)(1))(6(5))")
}

func TestRangeSumBST2(t *testing.T) {
	rangeSumBST(&TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{1, nil, nil}, nil}, &TreeNode{7, &TreeNode{6, nil, nil}, nil}},
	&TreeNode{15, &TreeNode{13,nil, nil}, &TreeNode{18, nil, nil}}}, 6, 10)
}