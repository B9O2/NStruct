package ScrollArray

import (
	"fmt"
	"testing"
)

func TestNewRollArray(t *testing.T) {
	ra := NewScrollArray(2)
	for i := 0; i < 106; i++ {
		ra.Append(i)
	}
	fmt.Println(ra.array)
	if v, ok := ra.LoadWithEid(104); ok {
		fmt.Println(v)
	} else {
		fmt.Println("DELETED")
	}
	fmt.Println(ra.Load(0))

}
