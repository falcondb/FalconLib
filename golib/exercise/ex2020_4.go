package exercise

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
)


/**
329. Longest Increasing Path in a Matrix
Given an integer matrix, find the length of the longest increasing path.
From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary

Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
*/

type lip struct {
	d,w int
	lp uint
	m [][]int
	db [][]uint
	cands []int
}

func createLIP(m [][]int) *lip {
	l := lip{len(m), len(m[0]), 0, m, make([][]uint, len(m)),  make([]int, 0)}
	for i, _ := range l.db { l.db[i] = make([]uint, l.w)}
	return &l
}

func longestIncreasingPath(m [][]int) int {

	if m == nil || len(m) == 0 {
		return 0
	}

	l :=  createLIP(m)

	for i, _ := range m {
		for j, _ := range m[i] {
			if l.read4update(i, j) == 0 {
				l.cands = append(l.cands, i * l.w + j)
			}
		}
	}

	for len(l.cands) != 0 {
		c := l.cands[0]
		l.cands = l.cands[1:]
		l.updateNbs(c/l.w, c%l.w)
	}

	return int(l.lp) + 1
}

func  (s *lip) read4update( i, j int) uint {
	nbs := uint(0)
	if i > 0 && s.m[i][j] > s.m[i-1][j] {
		nbs++
	}
	if i < s.d -1 && s.m[i][j] > s.m[i + 1][j] {
		nbs++
	}
	if j > 0 && s.m[i][j] > s.m[i][j-1] {
		nbs++
	}
	if j < s.w - 1 && s.m[i][j] > s.m[i][j+1] {
		nbs++
	}
	s.db[i][j] = nbs << 29
	return nbs
}

func (s *lip) updateNbs( i, j int)  {

	if (s.db[i][j] & 0x1fffffff) > s.lp {
		s.lp = s.db[i][j] & 0x1fffffff
	}
	// update pathlongest
	if i > 0 && s.m[i][j] < s.m[i-1][j] {
		if (s.db[i-1][j] & 0x1fffffff) < (s.db[i][j] & 0x1fffffff) + 1 {
			s.db[i-1][j] = (s.db[i-1][j] & 0xe0000000) + (s.db[i][j] & 0x1fffffff) + 1
		}
		s.db[i-1][j] = s.db[i-1][j] - 0x20000000
		if s.db[i-1][j] & 0xe0000000 == 0 {
			s.cands = append(s.cands, i * s.w - s.w + j)
		}
	}
	if i < s.d - 1 && s.m[i][j] < s.m[i + 1][j] {
		if (s.db[i+1][j] & 0x1fffffff) < (s.db[i][j] & 0x1fffffff) + 1 {
			s.db[i+1][j] = (s.db[i+1][j] & 0xe0000000) + (s.db[i][j] & 0x1fffffff) + 1
		}
		s.db[i+1][j] = s.db[i+1][j] - 0x20000000
		if s.db[i+1][j] & 0xe0000000 == 0 {
			s.cands = append(s.cands, i * s.w + s.w + j)
		}
	}
	if j > 0 && s.m[i][j] < s.m[i][j-1] {
		if (s.db[i][j-1] & 0x1fffffff) < (s.db[i][j] & 0x1fffffff) + 1 {
			s.db[i][j-1] = (s.db[i][j-1] & 0xe0000000) + (s.db[i][j] & 0x1fffffff) + 1
		}
		s.db[i][j-1] = s.db[i][j-1] - 0x20000000
		if s.db[i][j-1] & 0xe0000000 == 0 {
			s.cands = append(s.cands, i * s.w + j-1)
		}
	}
	if j < s.w - 1 && s.m[i][j] < s.m[i][j+1] {
		if (s.db[i][j+1] & 0x1fffffff) < (s.db[i][j] & 0x1fffffff) + 1 {
			s.db[i][j+1] = (s.db[i][j+1] & 0xe0000000) + (s.db[i][j] & 0x1fffffff) + 1
		}
		s.db[i][j+1] = s.db[i][j+1] - 0x20000000
		if s.db[i][j+1] & 0xe0000000 == 0 {
			s.cands = append(s.cands, i * s.w + j+1)
		}
	}
}

/**
1060. Missing Element in Sorted Array
Given a sorted array A of unique numbers, find the K-th missing number starting from the leftmost number of the array.
Input: A = [4,7,9,10], K = 1
Output: 5
Explanation:
The first missing number is 5.

 */
func missingElement(ns []int, k int) int {
	if k == 0 {
		return ns[0]
	}

	for i := 1; i < len(ns); i++ {
		d := ns[i] - ns[i-1] - 1
		if k <= d {
			return ns[i-1] + k
		} else {
			k -= d
		}
	}

	return ns[len(ns) - 1] + k

}

/*
729. My Calendar I

 */

type ctn struct {
	s, e int
	l, r *ctn
}

type MyCalendar struct {
	rt *ctn
}

func MyCalendarConstructor() MyCalendar {
	return MyCalendar{nil}
}


func (c *MyCalendar) Book(s, e int) bool {
	if c.rt == nil {
		c.rt = &ctn{s, e, nil, nil}
		return true
	}

	cur, par := c.rt, c.rt
	for cur != nil {
		par = cur
		switch {
		case cur.e <= s:
			cur = cur.r
		case cur.s >= e:
			cur = cur.l
		default:
			return false
		}
	}

	if par.e <= s {
		par.r = &ctn{s, e, nil, nil}
	} else {
		par.l = &ctn{s, e, nil, nil}
	}

	return true
}

/**
1138. Alphabet Board Path
 */

func alphabetBoardPath(target string) string {
	t := []byte(target)
	if len(t) == 0 {
		return ""
	}

	res := make([]byte, 0)
	pre := byte(0)

	for p:=0; p < len(t); p++ {
		vd := int((t[p]-'a')/5) - int(pre/5)
		hd := int((t[p]-'a')%5) - int(pre%5)

		if hd < 0 {
			for i:= hd; i != 0; i++ {res = append(res, 'L')}
		}
		if vd > 0 {
			for i:= 0; i != vd; i++ {res = append(res, 'D')}
		}
		if vd < 0 {
			for i:= vd; i != 0; i++ {res = append(res, 'U')}
		}
		if hd > 0 {
			for i:= 0; i != hd; i++ {res = append(res, 'R')}
		}

		pre = t[p] - 'a'
		res = append(res, '!')
	}

	return string(res)
}

/*
900. RLE Iterator
Write an iterator that iterates through a run-length encoded sequence.
The iterator is initialized by RLEIterator(int[] A), where A is a run-length encoding of some sequence.  More specifically, for all even i, A[i] tells us the number of times that the non-negative integer value A[i+1] is repeated in the sequence.
The iterator supports one function: next(int n), which exhausts the next n elements (n >= 1) and returns the last element exhausted in this way.  If there is no element left to exhaust, next returns -1 instead.
For example, we start with A = [3,8,0,9,2,5], which is a run-length encoding of the sequence [8,8,8,5,5].  This is because the sequence can be read as "three eights, zero nines, two fives".
 */

type RLEIterator struct {
	ci, cv, ct int
	d []int
}


func RLEConstructor(A []int) RLEIterator {
	if len(A) > 1 {
		return RLEIterator{0, A[1], A[0], A}
	}
	return RLEIterator{0, 0, 0, A}
}


func (it *RLEIterator) Next(n int) int {
	switch {
	case n <= it.ct:
		it.ct -= n
		return it.cv
	default:
		n -= it.ct
		for it.ci+=2; it.ci < len(it.d) - 1; it.ci +=2 {
			if n <= it.d[it.ci] {
				it.cv = it.d[it.ci+1]
				it.ct = it.d[it.ci] - n
				return it.cv
			} else {
				n -= it.d[it.ci]
			}
		}
		it.ct = 0
		return -1
	}
}


/*
299. Bulls and Cows
Input: secret = "1807", guess = "7810"
Output: "1A3B"
Explanation: Bulls are connected with a '|' and cows are underlined:
"1807"
  |
"7810"
 */

func getHint(secret string, guess string) string {
	s := []byte(secret)
	g := []byte(guess)
	bs :=  getBulls(&s, &g)
	cs := getCows(s, g)
	return fmt.Sprintf("%dA%dB", bs, cs)
}

func getBulls(s, g *[]byte) int {
	ct := 0
	for i:=0; i < len(*s) && i < len(*g); i++ {
		if (*s)[i] == (*g)[i] {
			ct++
			(*s)[i], (*g)[i] = 'x', 'x'
		}
	}
	return ct
}

func getCows(s, g []byte) int {
	swc, gwc := make([]int, 10), make([]int, 10)
	for _, c := range s {
		if c != 'x' {
			swc[c-'0'] ++
		}
	}
	for _, c := range g {
		if c != 'x' {
			gwc[c-'0'] ++
		}
	}
	ct := 0
	for i:=0; i<10; i++ {
		if swc[i] >= gwc[i] {
			ct += gwc[i]
		} else {
			ct += swc[i]
		}
	}

	return ct
}

/**
1146. Snapshot Array
Implement a SnapshotArray that supports the following interface:
SnapshotArray(int length) initializes an array-like data structure with the given length.  Initially, each element equals 0.
void set(index, val) sets the element at the given index to be equal to val.
int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id

Input: ["SnapshotArray","set","snap","set","get"]
[[3],[0,5],[],[0,6],[0,0]]
Output: [null,null,0,null,5]
Explanation:
SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
snapshotArr.set(0,5);  // Set array[0] = 5
snapshotArr.snap();  // Take a snapshot, return snap_id = 0
snapshotArr.set(0,6);
snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
 */

type SnapshotArray struct {
	sva [][]*snapValue
	nsid int
}

type snapValue struct {
	sid, v int
}

func SAConstructor(l int) SnapshotArray {
	sa := SnapshotArray{make([][]*snapValue, l), 0}
	return sa
}


func (sa *SnapshotArray) Set(index int, val int)  {
	svs := sa.sva[index]
	if len(svs) != 0 && svs[len(svs)-1].sid == -1 {
		svs[len(svs)-1].v = val
	} else {
		sa.sva[index] = append(svs, &snapValue{-1, val})
	}
}


func (sa *SnapshotArray) Snap() int {
	sid := sa.nsid
	sa.nsid++
	for _, v := range sa.sva {
		if len(v) != 0 {
			v[len(v) - 1].sid = sid
		}
	}
	return sid
}


func (sa *SnapshotArray) Get(index int, id int) int {
	svs := sa.sva[index]
	if svs == nil {
		return 0
	}

	if svs[len(svs) - 1].sid == -1 {
		if len(svs) == 1 {
			if sa.nsid == id {
				return svs[0].v
			} else {
				return 0
			}
		}
		if svs[len(svs) - 2].sid < id {
			return svs[len(svs) - 1].v
		}
	}

	s, e := 0, len(svs) - 1
	for s < e {
		m := (s + e) >> 1
		if svs[m].sid == id {
			return svs[m].v
		} else if svs[m].sid < id {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	if e >= 0 {
		return svs[e].v
	}

	return 0
}


// 53. Maximum Subarray

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max, cur, b, e  := 0, 0, 0, 0

	for ; e < len(nums); e++ {
		cur += nums[e]
		if cur > max {
			max = cur
		}
		if cur < 0 {
			for ; b <= e && cur < 0; b++ {
				cur -= nums[b]
			}
		}
	}

	return max
}

// 283. Move Zeroes
func moveZeroes(n []int)  {
	for zi, i:=0, 0; i < len(n) && i != 0; i++ {
		n[zi], n[i] = n[i], n[zi]
		zi++
	}
}

// 20. Valid Parentheses
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	st := make([]byte,1)
	st[0] = s[0]

	for i := 1; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			st = append(st, s[i])
		case ')':
			if len(st) == 0 || st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st) - 1]
		case ']':
			if len(st) == 0 || st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st) - 1]
		case '}':
			if len(st) == 0 || st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st) - 1]
		default:
			return false
		}
	}

	return len(st) == 0
}


// 230. Kth Smallest Element in a BST
func kthSmallest(root *TreeNode, k int) int {
	sm, v := _kthSmallest(root, k)

	if sm == k {
		return v
	}
	return 0
}

func _kthSmallest(root *TreeNode, k int) (int, int) {
	if root == nil {
		return 0, 0
	}
	lsmc, rsmc, tv := 0, 0, 0
	if root.Left != nil {
		lsmc, tv = _kthSmallest(root.Left, k)
		if lsmc == k {
			return k, tv
		}
		if lsmc == k - 1 {
			return k, root.Val
		}
	}

	if root.Right != nil {
		rsmc, tv = _kthSmallest(root.Left, k - lsmc - 1)
		if rsmc == k - lsmc - 1 {
			return k, tv
		}
	}

	return lsmc + rsmc + 1, 0
}

type Concordance struct {
	wStats map[string]*wordStats
}

type wordStats struct {
	word        *string
	occurrences uint
	sentenceIDs []uint
}

type wss []*wordStats

func (s wss) Len() int {return len(s)}
func (s wss) Less(i, j int) bool { return strings.Compare(*s[i].word, *s[j].word) < 0}
func (s wss) Swap(i, j int) { s[i], s[j] = s[j], s[i]}

/*
Convert the wordStats instance to a string with the expected format
 */
func (ws * wordStats) toFormattedString() string {
	// if the given wordStats instance is invalid, just return an empty string
	if ws == nil || ws.word == nil {
		return ""
	}

	// convert the sentence IDs to a comma separated string
	sentenceIDs := make([]byte, 0)
	for _, sid := range ws.sentenceIDs {
		sentenceIDs = append(sentenceIDs,  strconv.FormatUint(uint64(sid), 10)...)
		sentenceIDs = append(sentenceIDs, ',')
	}
	// remove the last ',' generated above to meet the expected format
	if len(sentenceIDs) > 0 {
		sentenceIDs = sentenceIDs[:len(sentenceIDs)-1]
	}

	// The output format section and the output sample section in the problem description didn't explicitly specify
	// the delimiter between the word and the its occurrence is one whitespace or a tab character, here I assume it is
	// a white space between them.
	return fmt.Sprintf("%s: {%d:%s}", *ws.word, ws.occurrences, sentenceIDs)
}

/*
 process one line from the input, update the word statistics for every valid word in the line.
 Here we assume a word doesn't cross multiple lines in the input
 */
func (cdc *Concordance) parseOneLine(line string, csid *uint)  {
	cur := make([]byte, 0)
	for _, c := range []byte(line) {
		switch typeOfChar(c) {
		case 0:
			cur = append(cur, c)
			continue
		case 1:
			*csid++
		}
		if len(cur) != 0 {
			cdc.wordStats(cur, csid)
			cur = make([]byte, 0)
		}

	}
}

/*
Check the type of the character:
	- alphabet: return 0
	- line delimiters: return 1
	- other charaters as word delimiters: 2
Here we don't consider the hyphen, '-' to link words
 */
func typeOfChar(c byte) uint8 {
	if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
		return 0
	}

	switch c {
	case '.', '!', '?':
		return 1
	default:
		return 2
	}
}

/*
Update the words' statistics in a hash map
 */
func (cdc *Concordance) wordStats(cs []byte, id *uint) {
	w := strings.ToLower(string(cs))
	ws, ok := cdc.wStats[w]
	if !ok {
		cdc.wStats[string(w)] = &wordStats{&w, 1, []uint{*id}}
	} else {
		ws.sentenceIDs = append(ws.sentenceIDs, *id)
		ws.occurrences++
	}
}


func generateAndPrintConcordance(lines []string) {
	// sanity check and return fast
	if len(lines) == 0 {
		return
	}

	// init
	cdc := Concordance{make(map[string]*wordStats)}
	sid := uint(1)

	// parse the words line by line and update the words' statistics
	for _, l := range lines {
		cdc.parseOneLine(l, &sid)
	}

	// sort the words by alphabet order
	ss := make([]*wordStats, 0)
	for _, v := range cdc.wStats {
		ss = append(ss, v)
	}
	sort.Sort(wss(ss))

	// output the words' statistics
	for _, s := range ss {
		fmt.Println(s.toFormattedString())
	}
}


/* 1041. Robot Bounded In Circle
On an infinite plane, a robot initially stands at (0, 0) and faces north. The robot can receive one of three instructions:
"G": go straight 1 unit;
"L": turn 90 degrees to the left;
"R": turn 90 degrees to the right.
The robot performs the instructions given in order, and repeats them forever.
Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
Example 1:
Input: instructions = "GGLLGG"
Output: true
Explanation: The robot moves from (0,0) to (0,2), turns 180 degrees, and then returns to (0,0).
When repeating these instructions, the robot remains in the circle of radius 2 centered at the origin.
 */

type plane struct {
	x, y int
	dir int
	update func(p *plane, i int32)
}

func brmove (p *plane, i int32) {
	switch i {
	case 'G':
		if p.dir&0x1 == 0 {
			p.y += int(p.dir - 1)
		} else {
			p.x += int(p.dir - 2)
		}
	case 'L':
		p.dir = (p.dir + 1) & 0x3
	case 'R':
		p.dir = (p.dir - 1) & 0x3
	}
}

func isRobotBounded(ins string) bool {
	p := plane{0, 0, 0, brmove}

	for _, i := range ins {
		p.update(&p, i)
	}

	return p.x == 0 && p.y == 0 || p.dir != 0
}


/*
739. Daily Temperatures
Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.
For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].
Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].
 */


func dailyTemperatures(T []int) []int {

	res := make([]int, len(T))
	wt := make([]int, 0)
	wt = append(wt, len(T)-1)

	for i := len(T) - 2; i >= 0; i-- {
		for j := len(wt) - 1; j >= 0 && T[i] >= T[wt[j]] ; j-- {
			wt = wt[:len(wt)-1]
		}

		if len(wt) != 0 {
			res[i] = wt[len(wt)-1] - i
		}

		wt = append(wt, i)
	}

	return res
}

/*
387. First Unique Character in a String
 */

func firstUniqChar(s string) int {
	cc, cd := uint(0), uint(0)


	for _, c := range s {
		if cd & (1 << uint(c - 'a')) == 0 {
			if cc & (1 << uint(c - 'a')) == 0 {
				cc += uint(1 << uint(c - 'a'))
			} else {
				cd += uint(1 << uint(c - 'a'))
			}
		}
	}

	for i, c := range s {
		if (cd & (1 << uint(c - 'a')) == 0) && (cc & (1 << uint(c - 'a')) != 0) {
			return i
		}
	}

	return -1
}

/*
1048. Longest String Chain
Given a list of words, each word consists of English lowercase letters.
Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".
A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.
Return the longest possible length of a word chain with words chosen from the given list of words.

Example 1:
Input: words = ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: One of the longest word chain is "a","ba","bda","bdca".
Example 2:
Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
 */



func longestStrChain(words []string) int {
	sort.Sort(wa(words))

	cands := make([]ws, 0)
	res := uint(0)

	for _, s := range words {
		if len(cands) == 0 || len(s) == len(cands[0].s) {
			cands = append(cands, ws{s, 1})
		} else if len(s) != len(cands[0].s)+1 {
			cands = nil
		} else {
			ncd := make([]ws, 0)
			added := false
			for _, c := range cands {
				if ispredecessor(c.s, s) {
					if added {
						if ncd[len(ncd) - 1 ].lps < 1 + c.lps {
							ncd[len(ncd) - 1 ].lps = 1 + c.lps
						}
					} else {
						ncd = append(ncd, ws{s, 1 + c.lps})
					}
				} else {
					if res < c.lps {
						res = c.lps
					}
				}
			}
			cands = ncd
		}
	}

	for _, c := range cands {
		if res < c.lps {
			res = c.lps
		}
	}

	return int(res)
}

type wa []string

func (a wa) Len() int {return len(a)}
func (a wa) Less(i, j int) bool {return len(a[i]) < len(a[j])}
func (a wa) Swap(i, j int) {a[i], a[j] = a[j], a[i]}

type ws struct {
	s string
	lps uint
}


func ispredecessor(s, pc string) bool {
	if len(s) + 1 != len(pc) {
		return false
	}

	pci := 0

	for i, _ := range s {
		if s[i] != pc[i] {
			pci++
			if s[i] != pc[i] {
				return false
			}
		}

		pci++
	}

	return true
}


/*
1429. First Unique Number
You have a queue of integers, you need to retrieve the first unique integer in the queue.
Implement the FirstUnique class:
FirstUnique(int[] nums) Initializes the object with the numbers in the queue.
int showFirstUnique() returns the value of the first unique integer of the queue, and returns -1 if there is no such integer.
void add(int value) insert value to the queue.
 */

type FirstUnique struct {
	cands []*numStats
	qa map[int]*numStats
}


func FUConstructor(nums []int) FirstUnique {
	fu := FirstUnique{make([]*numStats, 0), make(map[int]*numStats)}
	for _, n := range nums {
		fu.Add(n)
	}
	return fu
}


func (fu *FirstUnique) ShowFirstUnique() int {
	if len(fu.cands) != 0 {
		return fu.cands[0].v
	}
	return -1
}


func (fu *FirstUnique) Add(n int)  {
	v, ok := fu.qa[n]
	if !ok {
		cd := &numStats{n, 1}
		fu.qa[n] = cd
		fu.cands = append(fu.cands, cd)
	} else if v.c == 1{
		v.c++
	}

	i := 0
	for ; i < len(fu.cands); i++ {
		if fu.cands[i].c == 1 {
			fu.cands = fu.cands[i:]
			break
		}
	}

	if i == len(fu.cands) {
		fu.cands = nil
	}
}

type numStats struct {
	v, c int
}



/**
359. Logger Rate Limiter
Design a logger system that receives a stream of messages along with their timestamps. Each unique message should only be printed at most every 10 seconds (i.e. a message printed at timestamp t will prevent other identical messages from being printed until timestamp t + 10).
All messages will come in chronological order. Several messages may arrive at the same timestamp.
Implement the Logger class:
Logger() Initializes the logger object.
bool shouldPrintMessage(int timestamp, string message) Returns true if the message should be printed in the given timestamp, otherwise returns false.
 */

type Logger struct {
	buf map[string]int
}


/** Initialize your data structure here. */
func LConstructor() Logger {
	return Logger{make(map[string]int)}
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (l *Logger) ShouldPrintMessage(ts int, m string) bool {
	tss, ok := l.buf[m]
	if !ok {
		l.buf[m] = ts
		return true
	}

	if ts >= tss -10 {
		return false
	}

	l.buf[m] = ts
	return true
}

/*
525. Contiguous Array
Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
Example 1:
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
Example 2:
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
 */

func findMaxLength(nums []int) int {
	sbuf := make(map[int]int)
	res, sum := 0, 0
	sbuf[0] = -1

	for i, v := range nums {
		if v == 0 {
			sum--
		} else {
			sum++
		}

		ei, ok := sbuf[sum]
		if !ok {
			sbuf[sum] = i
		} else {
			if i - ei  > res {
				res = i - ei
			}
		}
	}

	return res
}

/**
811. Subdomain Visit Count
Example 2:
Input:
["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
Output:
["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
Explanation:
We will visit "google.mail.com" 900 times, "yahoo.com" 50 times, "intel.mail.com" once and "wiki.org" 5 times. For the subdomains, we will visit "mail.com" 900 + 1 = 901 times, "com" 900 + 50 + 1 = 951 times, and "org" 5 times.

 */

func subdomainVisits(cpdomains []string) []string {
	dc := make(map[string]int)

	for _, s := range cpdomains {
		ts := strings.Split(s, " ")
		if len(ts) != 2 {
			continue
		}
		vs, err := strconv.Atoi(ts[0])
		if err != nil {
			continue
		}
		for i, c := range ts[1] {
			if c == '.' {
				d := ts[1][i+1:]
				_, ok := dc[d]
				if !ok {
					dc[d] = vs
				} else {
					dc[d] += vs
				}
			}
		}
		d := ts[1]
		_, ok := dc[d]
		if !ok {
			dc[d] = vs
		} else {
			dc[d] += vs
		}
	}

	res := make([]string, 0)
	for d, v := range dc {
		res = append(res, fmt.Sprintf("%v %v", v, d))
	}

	return  res
}


/**
636. Exclusive Time of Functions

 */

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	fh := make([]fch, 0)
	pts := 0

	for _, e := range logs {
		tks := strings.Split(e, ":")
		if len(tks) != 3 {
			continue
		}

		id, _ := strconv.Atoi(tks[0])
		ts, _ := strconv.Atoi(tks[2])


		switch tks[1] {
		case "start":
			if len(fh) != 0 {
				res[fh[len(fh) - 1].id] += ts - pts
			}
			fh = append(fh, fch{id, ts})
		case "end":
			if fh[len(fh) - 1].id == id && pts == fh[len(fh) - 1].st && fh[len(fh) - 1].st != ts {
				res[id] += ts - pts + 1
			} else {
				res[id] += ts - pts
			}
			fh = fh[:len(fh)-1]
		}
		pts = ts
	}

	return res
}

type fch struct {
	id, st int
}



const (
	URL = "https://do.mcquay.us/ssh/short.txt"
	FAILUREKW = "Failed password"
)


func appInt() {
	resp, err := http.Get("https://do.mcquay.us/ssh/short.txt")
	defer resp.Body.Close()

	if err != nil {
		// some error handling to return an error to the upstream
		return
	}

	ipc := make(map[string]uint)

	s := bufio.NewScanner(resp.Body)

	parseLog(s, ipc)

	processIPStats(ipc)
}

func processIPStats (ipc map[string]uint) {
	fs := fsa{}

	for k, v := range ipc {
		fs = append(fs, &failureStat{k, v})
	}

	sort.Sort(fs)

	for _, s := range fs {
		fmt.Printf("%v\t%v\n", s.ct, s.IP)
	}
}

func parseLog(s *bufio.Scanner,  ipc map[string]uint) {
	for s.Scan() {
		line := s.Text()

		if !strings.Contains(line, FAILUREKW){
			continue
		}

		ip, err := getIP(line)

		if err == nil {
			_, ok := ipc[ip]
			if !ok {
				ipc[ip] = 1
			} else {
				ipc[ip]++
			}
		}
	}
}

func getIP(line string) (string, error) {
	tks := strings.Split(line, " ")

	for _, tk := range tks {
		ip := net.ParseIP(tk)
		if ip != nil {
			return ip.String(), nil
		}
	}

	return "", errors.New("no IP addr")
}


type failureStat struct {
	IP string
	ct uint
}

type fsa []*failureStat

func (f fsa) Len() int {return len(f)}
func (f fsa) Less (i, j int) bool {return f[i].ct > f[j].ct}
func (f fsa) Swap (i, j int) {f[i], f[j] = f[j], f[i]}



/**
253. Meeting Rooms II
Given an array of meeting time intervals intervals where intervals[i] = [starti, endi], return the minimum number of conference rooms required.
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2
Input: intervals = [[7,10],[2,4]]
Output: 1
 */

func minMeetingRooms(intervals [][]int) int {
	sort.Sort(mts(intervals))

	res := 0

	h := &mth{}
	for _, m := range intervals {
		for ; h.Len() != 0 && h.Peek().([]int)[1] <= m[0];  {
			fmt.Println(h.Peek().([]int))
			heap.Pop(h)
		}
		heap.Push(h, m)
		if h.Len() > res {
			res = h.Len()
		}
	}

	return res
}

type mts [][]int
func (f mts) Len() int {return len(f)}
func (f mts) Less (i, j int) bool {return f[i][0] < f[j][0] || (f[i][0] == f[j][0] && f[i][1] < f[j][1])}
func (f mts) Swap (i, j int) {f[i][0], f[j][0] = f[j][0], f[i][0]; f[i][1], f[j][1] = f[j][1], f[i][1]}

type mth [][]int

func (h *mth) Len() int { return len(*h) }
func (h *mth) Swap (i, j int) { (*h)[i][0], (*h)[j][0] = (*h)[j][0], (*h)[i][0]; (*h)[i][1], (*h)[j][1] = (*h)[j][1], (*h)[i][1]}
func (h *mth) Less(i, j int) bool {return (*h)[i][1] < (*h)[j][1] }
func (h *mth) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *mth) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *mth) Peek() interface{}  { return (*h)[0] }


type BinaryMatrix struct {
	Get func(int, int) int
	Dimensions func() []int
}

func leftMostColumnWithOne(bm BinaryMatrix) int {

	ds := bm.Dimensions()
	r, c := ds[1], 0
	for ; r < ds[0] && c >= 0; {
		if bm.Get(r, c) == 0 {
			c--
		} else {
			r++
		}
	}
	return c-1
}

/**
362. Design Hit Counter
*/

type HitCounter struct {
	bm []int
	lh int
}


/** Initialize your data structure here. */
func HCConstructor() HitCounter {
	return HitCounter{nil, 0}
}


/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (hc *HitCounter) Hit(ts int)  {
	hc.bm = append(hc.bm, ts)
}


/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (hc *HitCounter) GetHits(ts int) int {
	if len(hc.bm) != 0 && hc.bm[len(hc.bm) - 1] <= ts - 300 {
		hc.bm = nil
		return 0
	}

	for i, v := range hc.bm {
		if v > ts - 300 {
			hc.bm = hc.bm[i:]
			break
		}
	}

	return len(hc.bm)
}



/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */


