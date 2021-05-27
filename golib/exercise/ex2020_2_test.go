package exercise

import (
	"fmt"
	"testing"
)

func TestLowToUpCase(t *testing.T) {
	testLowToUpCase("aAbbb", "ABBB", t)
	testLowToUpCase("aAzdB$DA+", "AZD", t)
}

func testLowToUpCase(in, exp string, t  *testing.T){
	if exp != lowToUpCase(in) {
		fmt.Printf("Res:%v\t Exp:%v\n", lowToUpCase(in), exp)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}

func TestMergeIntevals(t *testing.T) {
	//[(1,3)]
	inp := []interval{interval{1,3}}
	testMergeIntevals(inp, inp, t)

	//Input:  [(1,3),(2,6),(8,10),(15,18)]
	//Output: [(1,6),(8,10),(15,18)]

	in2 := []interval{interval{1,3}, interval{2,6}, interval{8,10}, interval{15,18}}
	exp2 := []interval{interval{1,6}, interval{8,10}, interval{15,18}}
	testMergeIntevals(in2, exp2, t)


	in3 := []interval{interval{1,2}, interval{2,6}, interval{3,4}, interval{5,18}}
	exp3 := []interval{interval{1,18}}
	testMergeIntevals(in3, exp3, t)

}

func testMergeIntevals(in, exp []interval, t *testing.T) {
	res := mergeIntevals(in)

	if len(res) != len(exp) {
		fmt.Printf("Res:%v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}

	for i,v := range res {
		if v != exp[i] {
			fmt.Printf("Res:%v\t Exp:%v\n", res, exp)
			t.Errorf("Test name is %s, ", t.Name())
			return
		}
	}
}

func TestRotateImage(t *testing.T) {

	in := [][]uint {
		[]uint{1,2,3},
		[]uint{4,5,6},
		[]uint{7,8,9},
	}

	exp := [][]uint {
		[]uint{7,4,1},
		[]uint{8,5,2},
		[]uint{9,6,3},
	}


	testRotateImage(in, exp, t)



	in = [][]uint {
		[]uint{ 5, 1, 9,11},
		[]uint{ 2, 4, 8,10},
		[]uint{13, 3, 6, 7},
		[]uint{15,14,12,16},
	}

	exp = [][]uint {
		[]uint{15,13, 2, 5},
		[]uint{14, 3, 4, 1},
		[]uint{12, 6, 8, 9},
		[]uint{16, 7,10,11},
	}

	testRotateImage(in, exp, t)
}

func testRotateImage(in, exp [][]uint, t *testing.T) {
	rotateImage(in)

	for y,_ := range in {
		for x, _ := range in[y] {
			if in[y][x] != exp[y][x] {
				fmt.Printf("Res:%v\t Exp:%v\n", in, exp)
				t.Errorf("Test name is %s, ", t.Name())
				return
			}
		}
	}
}


func TestNumIslands2(t *testing.T) {
	testNumIslands2([][]byte{[]byte{1, 1, 1, 1, 0}, []byte{1, 1, 0, 1, 0}, []byte{1, 1, 0, 0, 0}, []byte{0,0,0,0,0}}, 1, t)
}

func testNumIslands2(in [][]byte, exp int, t  *testing.T){
	res := numIslands(in)
	if exp != res {
		fmt.Printf("Res:%v\t Exp:%v\n", res, exp)
		t.Errorf("Test name is %s, ", t.Name())
		return
	}
}