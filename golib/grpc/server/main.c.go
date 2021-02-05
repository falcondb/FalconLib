package main
import (
	"errors"
	"fmt"
)



func Max(data []int) (int, error) {

if data == nil || len(data) == 0 {
return 0, errors.New("Invalid input")
}

res := 0x8000000

for _, v := range data {
if v > res {
res = v
}
}

return res, nil
}

func main()  {

testdata := []int{3, 2, -1, 0, -1, -100000}
max, err := Max(testdata)

if err != nil {
fmt.Printf("Test Data: %v\tMax: %d", testdata, max)
} else {
fmt.Printf("Errpr: %v", err)
}

testdata = []int{}
max, err = Max(testdata)

if err != nil {
fmt.Printf("Test Data: %v\tMax: %d", testdata, max)
} else {
fmt.Printf("Errpr: %v", err)
}

testdata = []int{0, 0, 0, 0x70000000, 0x9fffffff}
max, err = Max(testdata)

if err != nil {
fmt.Printf("Test Data: %v\tMax: %d", testdata, max)
} else {
fmt.Printf("Errpr: %v", err)
}

}