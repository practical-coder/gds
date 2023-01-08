package counter

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	cnt := Counter{}
	for i := 0; i < 100; i++ {
		cnt.Inc()
	}
	fmt.Println(cnt)
	for i := 0; i < 50; i++ {
		cnt.Dec()
	}
	fmt.Println(cnt)
	cnt.Reset()
	fmt.Println(cnt)
	cnt.Dec()
	cnt.Inc()
	if cnt.Get() != 1 {
		t.Errorf("expected: 1, but get: %d", cnt.Get())
	}
}
