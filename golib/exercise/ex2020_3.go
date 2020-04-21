// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

func areTheyEqual(arr_a []int, arr_b []int) bool {
      // Write your code here
  if len(arr_a) != len(arr_b) {
    return false
  }
  
  h, t := -1, len(arr_a) - 1
  
  for i, _ := range arr_a {
    if arr_a[i] != arr_b[i] {
      h = i
    }
  }
  
  if h == -1 {
    return true
  }
  
  
  for ; t > h && arr_a[t] == arr_b[t]; t-- {}

  return reverseMatch(arr_a[h:t+1], arr_b[h:t+1])
}

func main() {
  // Call areTheyEqual() with test cases here
}


func reverseMatch(a, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  
  for i, v := range a {
    if b[len(b) - 1 - i] != v {
      return false
    }
  }
  
  return true
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
  pre := []byte {}

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
  res := []byte {}
  for ; reps >0; reps-- {res = append(res, substr...)}
  return res
}