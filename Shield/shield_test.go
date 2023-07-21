package Shield

import (
	"fmt"
	"sync"
	"testing"
)

var a = map[int]byte{}
var s = NewShield()

func test(protect bool, wg *sync.WaitGroup) {
	if protect {
		_ = s.Protect(func() error {
			for i := 0; i < 100; i++ {
				a[i] = 'b'
			}
			return nil
		})
	} else {
		for i := 0; i < 100; i++ {
			a[i] = 'b'
		}
	}

	wg.Done()
}

func TestNewShield(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test(true, wg)
	}
	wg.Wait()
	fmt.Println(a)
}
