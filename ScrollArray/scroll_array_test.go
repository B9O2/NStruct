package ScrollArray

import (
	"fmt"
	"testing"
)

func TestNewRollArray(t *testing.T) {
	ra := NewScrollArray(20)
	for i := 0; i < 10; i++ {
		ra.Append(i)
	}
	fmt.Println(ra.array)
	if v, ok := ra.LoadWithEid(104); ok {
		fmt.Println(v)
	} else {
		fmt.Println("DELETED")
	}
	fmt.Println(ra.Load(0))
	index := 0
	ra.Range(func(i interface{}) bool {
		fmt.Println(index, i)
		index += 1
		return true
	})

}
