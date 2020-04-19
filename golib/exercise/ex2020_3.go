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
